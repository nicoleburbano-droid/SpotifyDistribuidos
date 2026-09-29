package entity

type MetadataPodcast struct{
	nombrePodcast string
	tituloEpisodio string
	anfitrion string
	numeroTemporada int
	notasShow string
	clasificacionContenido string
}

// getters

func (m *MetadataPodcast) NombrePodcast() string {
	return m.nombrePodcast
}

func (m *MetadataPodcast) TituloEpisodio() string {
	return m.tituloEpisodio
}

func (m *MetadataPodcast) Anfitrion() string {
	return m.anfitrion
}

func (m *MetadataPodcast) NumeroTemporada() int {
	return m.numeroTemporada
}

func (m *MetadataPodcast) NotasShow() string {
	return m.notasShow
}

func (m *MetadataPodcast) ClasificacionContenido() string {
	return m.clasificacionContenido
}

// setters

func (m *MetadataPodcast) SetNombrePodcast(nombre string) {
	m.nombrePodcast = nombre
}

func (m *MetadataPodcast) SetTituloEpisodio(titulo string) {
	m.tituloEpisodio = titulo
}

func (m *MetadataPodcast) SetAnfitrion(anfitrion string) {
	m.anfitrion = anfitrion
}

func (m *MetadataPodcast) SetNumeroTemporada(numero int) {
	m.numeroTemporada = numero
}

func (m *MetadataPodcast) SetNotasShow(notas string) {
	m.notasShow = notas
}

func (m *MetadataPodcast) SetClasificacionContenido(clasificacion string) {
	m.clasificacionContenido = clasificacion
}