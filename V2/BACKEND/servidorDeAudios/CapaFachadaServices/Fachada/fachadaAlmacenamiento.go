package fachada

import (
	capaaccesodatos "almacenamiento/CapaAccesoADatos/repository"
	clientemetadatos "almacenamiento/CapaFachadaServices/ComponenteClienteMetadatos"
	dtos "almacenamiento/CapaFachadaServices/DTOs"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ErrDatosInvalidos indica que el administrador envió datos incorrectos.
var ErrDatosInvalidos = errors.New("datos del audio inválidos")

// ErrAudioExistente indica que ya hay un audio con ese título.
var ErrAudioExistente = errors.New("ya existe un audio con ese título")

// FachadaAlmacenamiento coordina el guardado del mp3 y el registro de sus
// metadatos en el servidor de metadatos.
type FachadaAlmacenamiento struct {
	repo             *capaaccesodatos.RepositorioAudios
	clienteMetadatos *clientemetadatos.ClienteMetadatos
	serviceMetadatos *clientemetadatos.MetadataAudioService
}

// NuevaFachadaAlmacenamiento crea la fachada.
func NuevaFachadaAlmacenamiento(carpetaAudios string, serviceMetadatos *clientemetadatos.MetadataAudioService) *FachadaAlmacenamiento {
	fmt.Println("Inicializando fachada de almacenamiento...")
	return &FachadaAlmacenamiento{
		repo:             capaaccesodatos.GetRepositorioAudios(carpetaAudios),
		clienteMetadatos: clientemetadatos.NuevoClienteMetadatos(serviceMetadatos),
	}
}

// AlmacenarAudio convierte los metadatos al DTO del tipo, guarda el mp3 con
// el nombre que usará el streaming y registra los metadatos. Si el registro
// falla, borra el archivo.
func (thisF *FachadaAlmacenamiento) AlmacenarAudio(tipo string, metadatosJSON []byte, data []byte) (dtos.RespuestaAlmacenamientoDTO, error) {
	var respuesta dtos.RespuestaAlmacenamientoDTO

	metadatosDTO, nombreArchivo, err := convertirMetadatos(tipo, metadatosJSON)
	if err != nil {
		return respuesta, err
	}

	if thisF.repo.ExisteAudio(nombreArchivo) {
		return respuesta, fmt.Errorf("%w: \"%s\"", ErrAudioExistente, nombreArchivo)
	}

	if err := thisF.repo.GuardarAudio(nombreArchivo, data); err != nil {
		return respuesta, err
	}

	if err := thisF.clienteMetadatos.RegistrarMetadatos(tipo, metadatosDTO); err != nil {
		thisF.repo.EliminarAudio(nombreArchivo)
		return respuesta, err
	}

	respuesta.Codigo = 201
	respuesta.Mensaje = "Audio \"" + nombreArchivo + "\" almacenado correctamente"
	respuesta.NombreArchivo = nombreArchivo
	return respuesta, nil
}

// convertirMetadatos lee el JSON con el DTO del tipo indicado y retorna:
// el DTO y el nombre del archivo (el mismo título que el cliente usa para
// pedir el streaming). Cada tipo exige sus metadatos propios.
func convertirMetadatos(tipo string, metadatosJSON []byte) (interface{}, string, error) {
	switch tipo {
	case "musica":
		var musica dtos.MetadataMusicaDTO
		err := json.Unmarshal(metadatosJSON, &musica)
		return musica, musica.TituloCancion, validarConversion(err, musica.TituloCancion,
			campo{"artista principal", musica.ArtistaPrincipal}, campo{"álbum", musica.Album}, campo{"género musical", musica.Genero})
	case "audiolibro", "audioLibros":
		var libro dtos.MetadataAudioLibrosDTO
		err := json.Unmarshal(metadatosJSON, &libro)
		return libro, libro.TituloLibro, validarConversion(err, libro.TituloLibro,
			campo{"autor", libro.Autor}, campo{"narrador", libro.Narrador})
	case "podcast":
		var podcast dtos.MetadataPodcastDTO
		err := json.Unmarshal(metadatosJSON, &podcast)
		// el archivo se nombra con el episodio: un podcast tiene muchos episodios
		return podcast, podcast.TituloEpisodio, validarConversion(err, podcast.TituloEpisodio,
			campo{"nombre del podcast", podcast.NombrePodcast}, campo{"anfitrión", podcast.Anfitrion})
	case "ruidoBlanco":
		var ruido dtos.MetadataRuidoBlancoDTO
		err := json.Unmarshal(metadatosJSON, &ruido)
		return ruido, ruido.Titulo, validarConversion(err, ruido.Titulo,
			campo{"tipo de sonido", ruido.TipoSonido}, campo{"fuente del audio", ruido.FuenteAudio}, campo{"uso sugerido", ruido.UsoSugerido})
	default:
		return nil, "", fmt.Errorf("%w: el tipo \"%s\" no existe", ErrDatosInvalidos, tipo)
	}
}

// campo es un metadato obligatorio del tipo de audio.
type campo struct {
	nombre string
	valor  string
}

// validarConversion verifica que el JSON sea válido y que el título sirva
// como nombre de archivo.
func validarConversion(errJSON error, titulo string, obligatorios ...campo) error {
	if errJSON != nil {
		return fmt.Errorf("%w: el campo metadatos no es un JSON válido", ErrDatosInvalidos)
	}
	if strings.TrimSpace(titulo) == "" {
		return fmt.Errorf("%w: falta el título del audio", ErrDatosInvalidos)
	}
	for _, c := range obligatorios {
		if strings.TrimSpace(c.valor) == "" {
			return fmt.Errorf("%w: falta el metadato \"%s\"", ErrDatosInvalidos, c.nombre)
		}
	}
	if strings.ContainsAny(titulo, `/\`) || titulo == "." || titulo == ".." {
		return fmt.Errorf("%w: el título no puede contener / ni \\", ErrDatosInvalidos)
	}
	return nil
}
