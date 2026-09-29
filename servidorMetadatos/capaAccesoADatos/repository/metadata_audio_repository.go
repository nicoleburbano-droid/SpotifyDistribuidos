package repository

import "microservicio/capaAccesoADatos/entity"

// MetadataAudioRepository es un repositorio que mantiene en memoria
// el slice de audios y expone las operaciones de búsqueda y registro.

type MetadataAudioRepository struct {
	vectorMetadataAudios []entity.MetadataAudio
	vectorTiposDeAudio []entity.TipoAudio
	vectorMetadataMusica []entity.MetadataMusica
	vectorMetadataAudioLibros []entity.MetadataAudioLibros
	vectorMetadataPodcast []entity.MetadataPodcast
	vectorMetadataRuidoBlanco []entity.MetadataRuidoBlanco
}

// NewMetadataAudioRepository crea el repositorio y lo precarga con la
// metadata de audios de ejemplo.
func NewMetadataAudioRepository() *MetadataAudioRepository {
	this := &MetadataAudioRepository{}
	this.CargarMetadataAudios()
	this.CargarTiposDeAudio()
	this.CargarMetadataMusica()
	this.CargarMetadataAudioLibros()
	this.CargarMetadataPodcast()
	this.CargarMetadataRuidoBlanco()
	
	return this
}

// CargarMetadataAudios inicializa el vector con 5 audios de ejemplo.
func (this *MetadataAudioRepository) CargarMetadataAudios() {
	var objAudio1, objAudio2, objAudio3, objAudio4, objAudio5 entity.MetadataAudio

	objAudio1.SetTitulo("Cancion 1")
	objAudio1.SetDuracion(10)
	objAudio1.SetTipo("Musica")
	objAudio1.SetDisponible(true)

	objAudio2.SetTitulo("Podcast 2")
	objAudio2.SetDuracion(20)
	objAudio2.SetTipo("Podcasts")
	objAudio2.SetDisponible(false)

	objAudio3.SetTitulo("Ruido Blanco 3")
	objAudio3.SetDuracion(30)
	objAudio3.SetTipo("Ruido Blanco")
	objAudio3.SetDisponible(true)

	objAudio4.SetTitulo("Audiolibro 4")
	objAudio4.SetDuracion(40)
	objAudio4.SetTipo("Audiolibros")
	objAudio4.SetDisponible(true)

	objAudio5.SetTitulo("Meditacion 5")
	objAudio5.SetDuracion(50)
	objAudio5.SetTipo("Meditaciones guiadas")
	objAudio5.SetDisponible(false)

	this.vectorMetadataAudios = []entity.MetadataAudio{
		objAudio1, objAudio2, objAudio3, objAudio4, objAudio5,
	}
}

// Inicializa el vector de tipos de audio con los 4 tipos de audio de la aplicacion
func (this *MetadataAudioRepository) CargarTiposDeAudio() {
	var objTipo1, objTipo2, objTipo3, objTipo4 entity.TipoAudio

	objTipo1.SetId(1)
	objTipo1.SetTipo("musica")

	objTipo2.SetId(2)
	objTipo2.SetTipo("audioLibros")

	objTipo3.SetId(3)
	objTipo3.SetTipo("podcast")

	objTipo4.SetId(4)
	objTipo4.SetTipo("ruidoBlanco")

	this.vectorTiposDeAudio = []entity.TipoAudio{
		objTipo1, objTipo2, objTipo3, objTipo4,
	}
}

// Inicializa el vector de musica
func (this *MetadataAudioRepository) CargarMetadataMusica() {
	var objMusica1, objMusica2, objMusica3 entity.MetadataMusica

	objMusica1.SetArtistaPrincipal("Daft Punk")
	objMusica1.SetAlbum("Random Access Memories")
	objMusica1.SetGenero("Electrónica")
	objMusica1.SetTituloCancion("Get Lucky")
	objMusica1.SetSelloDiscografico("Columbia Records")
	objMusica1.SetAnioLanzamiento("2013")

	objMusica2.SetArtistaPrincipal("Michael Jackson")
	objMusica2.SetAlbum("Thriller")
	objMusica2.SetGenero("Pop")
	objMusica2.SetTituloCancion("Billie Jean")
	objMusica2.SetSelloDiscografico("Epic Records")
	objMusica2.SetAnioLanzamiento("1982")

	objMusica3.SetArtistaPrincipal("Queen")
	objMusica3.SetAlbum("A Night at the Opera")
	objMusica3.SetGenero("Rock")
	objMusica3.SetTituloCancion("Bohemian Rhapsody")
	objMusica3.SetSelloDiscografico("EMI Records")
	objMusica3.SetAnioLanzamiento("1975")

	this.vectorMetadataMusica = []entity.MetadataMusica{
		objMusica1, objMusica2, objMusica3,
	}
}

// Inicializa el vector de audiolibros
func (this *MetadataAudioRepository) CargarMetadataAudioLibros() {
	var objLibro1, objLibro2, objLibro3 entity.MetadataAudioLibros

	objLibro1.SetTituloLibro("Cien años de soledad")
	objLibro1.SetAutor("Gabriel García Márquez")
	objLibro1.SetNarrador("Gustavo Bonfigli")
	objLibro1.SetEditorial("Audible Studios")
	objLibro1.SetIsbn(978030747)
	objLibro1.SetCapitulo(1)

	objLibro2.SetTituloLibro("El código Da Vinci")
	objLibro2.SetAutor("Dan Brown")
	objLibro2.SetNarrador("Raúl Amundaray")
	objLibro2.SetEditorial("Planeta Audio")
	objLibro2.SetIsbn(978840805)
	objLibro2.SetCapitulo(5)

	objLibro3.SetTituloLibro("Don Quijote de la Mancha")
	objLibro3.SetAutor("Miguel de Cervantes")
	objLibro3.SetNarrador("Juan Magno")
	objLibro3.SetEditorial("Penguin Audio")
	objLibro3.SetIsbn(978849105)
	objLibro3.SetCapitulo(12)

	this.vectorMetadataAudioLibros = []entity.MetadataAudioLibros{
		objLibro1, objLibro2, objLibro3,
	}
}

// Inicializa el vector de podcast
func (this *MetadataAudioRepository) CargarMetadataPodcast() {
	var objPodcast1, objPodcast2, objPodcast3 entity.MetadataPodcast

	objPodcast1.SetNombrePodcast("The Daily")
	objPodcast1.SetTituloEpisodio("The Global Economy Today")
	objPodcast1.SetAnfitrion("Michael Barbaro")
	objPodcast1.SetNumeroTemporada(2026)
	objPodcast1.SetNotasShow("An analysis of economic shifts across continents.")
	objPodcast1.SetClasificacionContenido("Todo público")

	objPodcast2.SetNombrePodcast("Radio Ambulante")
	objPodcast2.SetTituloEpisodio("Los herederos")
	objPodcast2.SetAnfitrion("Daniel Alarcón")
	objPodcast2.SetNumeroTemporada(12)
	objPodcast2.SetNotasShow("Una historia sobre la memoria y las herencias familiares.")
	objPodcast2.SetClasificacionContenido("Mayores de 14")

	objPodcast3.SetNombrePodcast("Lex Fridman Podcast")
	objPodcast3.SetTituloEpisodio("Artificial Intelligence and the Future")
	objPodcast3.SetAnfitrion("Lex Fridman")
	objPodcast3.SetNumeroTemporada(1)
	objPodcast3.SetNotasShow("Deep dive into technical and philosophical aspects of AI.")
	objPodcast3.SetClasificacionContenido("Todo público")

	this.vectorMetadataPodcast = []entity.MetadataPodcast{
		objPodcast1, objPodcast2, objPodcast3,
	}
}

// Inicializa el vector de ruido blanco
func (this *MetadataAudioRepository) CargarMetadataRuidoBlanco() {
	var objRuido1, objRuido2, objRuido3 entity.MetadataRuidoBlanco

	objRuido1.SetTipoSonido("Ruido Blanco Clásico")
	objRuido1.SetFuenteAudio("Generador Estático Analógico")
	objRuido1.SetUsoSugerido("Concentración y Estudio profunda")
	objRuido1.SetProveedorContenido("CalmSounds Ltd")
	objRuido1.SetDuracionBucle(60)
	objRuido1.SetFrecuenciaDominante("Plana (20Hz - 20kHz)")

	objRuido2.SetTipoSonido("Lluvia Pesada")
	objRuido2.SetFuenteAudio("Grabación Estéreo de Campo")
	objRuido2.SetUsoSugerido("Conciliación del Sueño")
	objRuido2.SetProveedorContenido("Nature Tracks")
	objRuido2.SetDuracionBucle(120)
	objRuido2.SetFrecuenciaDominante("Baja-Media (Pink Noise)")

	objRuido3.SetTipoSonido("Estática de Televisor")
	objRuido3.SetFuenteAudio("Sintetizador Digital")
	objRuido3.SetUsoSugerido("Bloqueo de Tinnitus / Ruido de Fondo")
	objRuido3.SetProveedorContenido("Focus Audio")
	objRuido3.SetDuracionBucle(30)
	objRuido3.SetFrecuenciaDominante("Alta (White Noise Puro)")

	this.vectorMetadataRuidoBlanco = []entity.MetadataRuidoBlanco{
		objRuido1, objRuido2, objRuido3,
	}
}



// BuscarAudio recorre el vector buscando un audio por su título. Retorna el
// audio encontrado y un booleano que indica si la búsqueda tuvo éxito.

func (this *MetadataAudioRepository) BuscarAudio(titulo string) (entity.MetadataAudio, bool) {
	for _, audio := range this.vectorMetadataAudios {
		if audio.GetTitulo() == titulo {
			return audio, true
		}
	}
	return entity.MetadataAudio{}, false
}

// BuscarMusica
func (this *MetadataAudioRepository) BuscarMusica(titulo string) (entity.MetadataMusica, bool) {
	for _, musica := range this.vectorMetadataMusica {
		if musica.GetTituloCancion() == titulo {
			return musica, true
		}
	}
	return entity.MetadataMusica{}, false
}

// BuscarAudioLibro
func (this *MetadataAudioRepository) BuscarAudioLibro(titulo string) (entity.MetadataAudioLibros, bool) {
	for _, libro := range this.vectorMetadataAudioLibros {
		if libro.TituloLibro() == titulo {
			return libro, true
		}
	}
	return entity.MetadataAudioLibros{}, false
}

// BuscarPodcast
func (this *MetadataAudioRepository) BuscarPodcast(titulo string) (entity.MetadataPodcast, bool) {
	for _, podcast := range this.vectorMetadataPodcast {
		if podcast.TituloEpisodio() == titulo {
			return podcast, true
		}
	}
	return entity.MetadataPodcast{}, false
}

// BuscarRuidoBlanco
func (this *MetadataAudioRepository) BuscarRuidoBlanco(tipoSonido string) (entity.MetadataRuidoBlanco, bool) {
	for _, ruido := range this.vectorMetadataRuidoBlanco {
		if ruido.TipoSonido() == tipoSonido {
			return ruido, true
		}
	}
	return entity.MetadataRuidoBlanco{}, false
}

// RegistrarAudio agrega un nuevo audio al vector.

func (this *MetadataAudioRepository) RegistrarAudio(audio entity.MetadataAudio) {
	this.vectorMetadataAudios = append(this.vectorMetadataAudios, audio)
}

// RegistrarMusica
func (this *MetadataAudioRepository) RegistrarMusica(musica entity.MetadataMusica) {
	this.vectorMetadataMusica = append(this.vectorMetadataMusica, musica)
}
// RegistrarAudioLibro
func (this *MetadataAudioRepository) RegistrarAudioLibro(libro entity.MetadataAudioLibros) {
	this.vectorMetadataAudioLibros = append(this.vectorMetadataAudioLibros, libro)
}
// RegistrarPodcast
func (this *MetadataAudioRepository) RegistrarPodcast(podcast entity.MetadataPodcast) {
	this.vectorMetadataPodcast = append(this.vectorMetadataPodcast, podcast)
}
// RegistrarRuidoBlanco
func (this *MetadataAudioRepository) RegistrarRuidoBlanco(ruido entity.MetadataRuidoBlanco) {
	this.vectorMetadataRuidoBlanco = append(this.vectorMetadataRuidoBlanco, ruido)
}

//ObtenerAudiosPorTipo obtiene todo el vector de audios segun un determinado tipo

func (this *MetadataAudioRepository) ObtenerAudiosPorTipo(tipo string) []entity.MetadataAudio {
	var vectorAudiosTipo []entity.MetadataAudio
	for audio := range this.vectorMetadataAudios{
		if this.vectorMetadataAudios[audio].GetTipo() == tipo {
			vectorAudiosTipo=append(vectorAudiosTipo, this.vectorMetadataAudios[audio])
		}
	}
	return vectorAudiosTipo
}

// ObtenerVectorMusica
func (this *MetadataAudioRepository) ObtenerVectorMusica() []entity.MetadataMusica {
	return this.vectorMetadataMusica
}
// ObtenerVectorAudioLibros
func (this *MetadataAudioRepository) ObtenerVectorAudioLibros() []entity.MetadataAudioLibros {
	return this.vectorMetadataAudioLibros
}
// ObtenerVectorPodcast
func (this *MetadataAudioRepository) ObtenerVectorPodcast() []entity.MetadataPodcast {
	return this.vectorMetadataPodcast
}
// ObtenerVectorRuidoBlanco
func (this *MetadataAudioRepository) ObtenerVectorRuidoBlanco() []entity.MetadataRuidoBlanco {
	return this.vectorMetadataRuidoBlanco
}
// ObtenerVectorTiposDeAudio
func (this *MetadataAudioRepository) ObtenerVectorTiposDeAudio() []entity.TipoAudio {
	return this.vectorTiposDeAudio
}