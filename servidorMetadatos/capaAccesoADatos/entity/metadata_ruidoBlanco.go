
package entity

type MetadataRuidoBlanco struct{
	tipoSonido string
	fuenteAudio string
	usoSugerido string
	proveedorContenido string
	duracionBucle int
	frecuenciaDominante string
}

// getters

func (r *MetadataRuidoBlanco) TipoSonido() string {
	return r.tipoSonido
}

func (r *MetadataRuidoBlanco) FuenteAudio() string {
	return r.fuenteAudio
}

func (r *MetadataRuidoBlanco) UsoSugerido() string {
	return r.usoSugerido
}

func (r *MetadataRuidoBlanco) ProveedorContenido() string {
	return r.proveedorContenido
}

func (r *MetadataRuidoBlanco) DuracionBucle() int {
	return r.duracionBucle
}

func (r *MetadataRuidoBlanco) FrecuenciaDominante() string {
	return r.frecuenciaDominante
}

// setters

func (r *MetadataRuidoBlanco) SetTipoSonido(tipo string) {
	r.tipoSonido = tipo
}

func (r *MetadataRuidoBlanco) SetFuenteAudio(fuente string) {
	r.fuenteAudio = fuente
}

func (r *MetadataRuidoBlanco) SetUsoSugerido(uso string) {
	r.usoSugerido = uso
}

func (r *MetadataRuidoBlanco) SetProveedorContenido(proveedor string) {
	r.proveedorContenido = proveedor
}

func (r *MetadataRuidoBlanco) SetDuracionBucle(duracion int) {
	r.duracionBucle = duracion
}

func (r *MetadataRuidoBlanco) SetFrecuenciaDominante(frecuencia string) {
	r.frecuenciaDominante = frecuencia
}