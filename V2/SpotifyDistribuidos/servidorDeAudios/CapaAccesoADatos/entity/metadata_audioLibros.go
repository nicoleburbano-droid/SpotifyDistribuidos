package entity

type MetadataAudioLibros struct{
	tituloLibro string
	autor string
	narrador string
	editorial string
	isbn int
	capitulo int
}

// getters

func (a *MetadataAudioLibros) TituloLibro() string {
	return a.tituloLibro
}

func (a *MetadataAudioLibros) Autor() string {
	return a.autor
}

func (a *MetadataAudioLibros) Narrador() string {
	return a.narrador
}

func (a *MetadataAudioLibros) Editorial() string {
	return a.editorial
}

func (a *MetadataAudioLibros) Isbn() int {
	return a.isbn
}

func (a *MetadataAudioLibros) Capitulo() int {
	return a.capitulo
}

// setters

func (a *MetadataAudioLibros) SetTituloLibro(titulo string) {
	a.tituloLibro = titulo
}

func (a *MetadataAudioLibros) SetAutor(autor string) {
	a.autor = autor
}

func (a *MetadataAudioLibros) SetNarrador(narrador string) {
	a.narrador = narrador
}

func (a *MetadataAudioLibros) SetEditorial(editorial string) {
	a.editorial = editorial
}

func (a *MetadataAudioLibros) SetIsbn(isbn int) {
	a.isbn = isbn
}

func (a *MetadataAudioLibros) SetCapitulo(capitulo int) {
	a.capitulo = capitulo
}
