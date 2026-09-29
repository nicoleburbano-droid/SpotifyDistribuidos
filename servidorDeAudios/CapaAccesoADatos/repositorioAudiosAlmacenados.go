package capaaccesoadatos

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// RepositorioAudios guarda los archivos mp3 en disco. Es un singleton y usa
// un mutex para que dos administradores no escriban el mismo archivo a la vez.
type RepositorioAudios struct {
	mu            sync.Mutex
	carpetaAudios string
}

var (
	instancia *RepositorioAudios
	once      sync.Once
)

// GetRepositorioAudios retorna la única instancia del repositorio.
// La carpeta es donde el servidor de streaming busca los audios.
func GetRepositorioAudios(carpetaAudios string) *RepositorioAudios {
	once.Do(func() {
		instancia = &RepositorioAudios{carpetaAudios: carpetaAudios}
	})
	return instancia
}

// construirRuta arma la ruta del archivo. El servidor de streaming abre el
// archivo con el mismo nombre que el cliente le envía (el título del audio),
// por eso el archivo se guarda con ese nombre exacto.
func (r *RepositorioAudios) construirRuta(nombreArchivo string) string {
	return filepath.Join(r.carpetaAudios, nombreArchivo)
}

// ExisteAudio indica si ya hay un audio guardado con ese nombre.
func (r *RepositorioAudios) ExisteAudio(nombreArchivo string) bool {
	_, err := os.Stat(r.construirRuta(nombreArchivo))
	return err == nil
}

// GuardarAudio escribe los bytes del mp3 en la carpeta de audios.
func (r *RepositorioAudios) GuardarAudio(nombreArchivo string, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := os.MkdirAll(r.carpetaAudios, os.ModePerm); err != nil {
		return fmt.Errorf("error creando la carpeta de audios: %v", err)
	}

	ruta := r.construirRuta(nombreArchivo)
	if err := os.WriteFile(ruta, data, 0644); err != nil {
		return fmt.Errorf("error al guardar archivo: %v", err)
	}
	fmt.Printf("Archivo guardado en %s (%d bytes)\n", ruta, len(data))
	return nil
}

// EliminarAudio borra un archivo; se usa si falla el registro de metadatos
// para no dejar audios sin metadatos.
func (r *RepositorioAudios) EliminarAudio(nombreArchivo string) {
	os.Remove(r.construirRuta(nombreArchivo))
}
