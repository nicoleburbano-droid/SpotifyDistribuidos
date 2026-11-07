
package entity

type MetadataMusica struct {
	artistaPrincipal    string
	album   			string
	genero       		string
	tituloCancion 		string
	selloDiscografico	string
	anioLanzamiento		string
}

// getters

func (m *MetadataMusica) GetArtistaPrincipal() string {
	return m.artistaPrincipal
}

func (m *MetadataMusica) GetAlbum() string {
	return m.album
}

func (m *MetadataMusica) GetGenero() string {
	return m.genero
}

func (m *MetadataMusica) GetTituloCancion() string {
	return m.tituloCancion
}

func (m *MetadataMusica) GetSelloDiscografico() string {
	return m.selloDiscografico
}

func (m *MetadataMusica) GetAnioLanzamiento() string {
	return m.anioLanzamiento
}

// setters

func (m *MetadataMusica) SetArtistaPrincipal(artista string) {
	m.artistaPrincipal = artista
}

func (m *MetadataMusica) SetAlbum(album string) {
	m.album = album
}

func (m *MetadataMusica) SetGenero(genero string) {
	m.genero = genero
}

func (m *MetadataMusica) SetTituloCancion(titulo string) {
	m.tituloCancion = titulo
}

func (m *MetadataMusica) SetSelloDiscografico(sello string) {
	m.selloDiscografico = sello
}

func (m *MetadataMusica) SetAnioLanzamiento(anio string) {
	m.anioLanzamiento = anio
}