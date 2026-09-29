package service

import (
	"microservicio/capaAccesoADatos/entity"
	"microservicio/capaAccesoADatos/repository"
	"microservicio/capaFachadaServices/dto"
)

// MetadataAudioService es la fachada (facade) que expone al controlador las
// operaciones de negocio, ocultando el acceso al repositorio y la
// conversión entre Entity y DTO.

type MetadataAudioService struct {
	repository *repository.MetadataAudioRepository
}

func NewMetadataAudioService(repository *repository.MetadataAudioRepository) *MetadataAudioService {
	return &MetadataAudioService{repository: repository}
}

// RegistrarAudio recibe un DTO, lo convierte a Entity y lo registra en el
// repositorio.

func (this *MetadataAudioService) RegistrarAudio(audioDTO dto.MetadataAudioDTO) {
	var audio entity.MetadataAudio

	audio.SetTitulo(audioDTO.Titulo)
	audio.SetDuracion(audioDTO.Duracion)
	audio.SetTipo(audioDTO.Tipo)
	audio.SetDisponible(audioDTO.Disponible)

	this.repository.RegistrarAudio(audio)
}

// TODO registrarMusica
func (this *MetadataAudioService) RegistrarMusica(musicaDTO dto.MetadataMusicaDTO) {
	var musica entity.MetadataMusica
	musica.SetArtistaPrincipal(musicaDTO.ArtistaPrincipal)
	musica.SetAlbum(musicaDTO.Album)
	musica.SetGenero(musicaDTO.Genero)
	musica.SetTituloCancion(musicaDTO.TituloCancion)
	musica.SetSelloDiscografico(musicaDTO.SelloDiscografico)
	musica.SetAnioLanzamiento(musicaDTO.AnioLanzamiento)
	this.repository.RegistrarMusica(musica)
}

// TODO registrarAudioLibro
func (this *MetadataAudioService) RegistrarAudioLibro(libroDTO dto.MetadataAudioLibrosDTO) {
	var libro entity.MetadataAudioLibros
	libro.SetTituloLibro(libroDTO.TituloLibro)
	libro.SetAutor(libroDTO.Autor)
	libro.SetNarrador(libroDTO.Narrador)
	libro.SetEditorial(libroDTO.Editorial)
	libro.SetIsbn(libroDTO.Isbn)
	libro.SetCapitulo(libroDTO.Capitulo)
	this.repository.RegistrarAudioLibro(libro)
}

// TODO registrarPodcast
func (this *MetadataAudioService) RegistrarPodcast(podcastDTO dto.MetadataPodcastDTO) {
	var podcast entity.MetadataPodcast
	podcast.SetNombrePodcast(podcastDTO.NombrePodcast)
	podcast.SetTituloEpisodio(podcastDTO.TituloEpisodio)
	podcast.SetAnfitrion(podcastDTO.Anfitrion)
	podcast.SetNumeroTemporada(podcastDTO.NumeroTemporada)
	podcast.SetNotasShow(podcastDTO.NotasShow)
	podcast.SetClasificacionContenido(podcastDTO.ClasificacionContenido)
	this.repository.RegistrarPodcast(podcast)
}

// TODO registrarRuidoBlanco
func (this *MetadataAudioService) RegistrarRuidoBlanco(ruidoDTO dto.MetadataRuidoBlancoDTO) {
	var ruido entity.MetadataRuidoBlanco
	ruido.SetTipoSonido(ruidoDTO.TipoSonido)
	ruido.SetFuenteAudio(ruidoDTO.FuenteAudio)
	ruido.SetUsoSugerido(ruidoDTO.UsoSugerido)
	ruido.SetProveedorContenido(ruidoDTO.ProveedorContenido)
	ruido.SetDuracionBucle(ruidoDTO.DuracionBucle)
	ruido.SetFrecuenciaDominante(ruidoDTO.FrecuenciaDominante)
	this.repository.RegistrarRuidoBlanco(ruido)
}

// ConsultarAudio recibe un título, busca el Entity en el repositorio y lo
// convierte a RespuestaMetadataAudioDTO con el código y mensaje según el
// resultado de la búsqueda.

func (this *MetadataAudioService) ConsultarAudio(titulo string) dto.RespuestaMetadataAudioDTO {
	var respuesta dto.RespuestaMetadataAudioDTO

	audio, encontrado := this.repository.BuscarAudio(titulo)

	if encontrado {

		var audioDTO dto.MetadataAudioDTO
		audioDTO.Titulo = audio.GetTitulo()
		audioDTO.Duracion = audio.GetDuracion()
		audioDTO.Tipo = audio.GetTipo()
		audioDTO.Disponible = audio.GetDisponible()

		respuesta.ObjAudio = audioDTO
		respuesta.Codigo = 200
		respuesta.Mensaje = "Metadata del audio encontrada"
	} else {
		respuesta.Codigo = 400
		respuesta.Mensaje = "La metadata del audio no se encontro"
	}

	return respuesta
}

// TODO ConsultarMusica
func (this *MetadataAudioService) ConsultarMusica(titulo string) dto.RespuestaMusicaDTO {
	var respuesta dto.RespuestaMusicaDTO
	musica, encontrado := this.repository.BuscarMusica(titulo)
	if encontrado {
		respuesta.VectorMusica = []dto.MetadataMusicaDTO{{ArtistaPrincipal: musica.GetArtistaPrincipal(), Album: musica.GetAlbum(), Genero: musica.GetGenero(), TituloCancion: musica.GetTituloCancion(), SelloDiscografico: musica.GetSelloDiscografico(), AnioLanzamiento: musica.GetAnioLanzamiento()}}
		respuesta.Codigo, respuesta.Mensaje = 200, "Metadata de la música encontrada"
	} else {
		respuesta.Codigo, respuesta.Mensaje = 400, "La metadata de la música no se encontro"
	}
	return respuesta
}

// TODO ConsultarAudioLibro
func (this *MetadataAudioService) ConsultarAudioLibro(titulo string) dto.RespuestaAudioLibrosDTO {
	var respuesta dto.RespuestaAudioLibrosDTO
	libro, encontrado := this.repository.BuscarAudioLibro(titulo)
	if encontrado {
		respuesta.VectorAudioLibros = []dto.MetadataAudioLibrosDTO{{TituloLibro: libro.TituloLibro(), Autor: libro.Autor(), Narrador: libro.Narrador(), Editorial: libro.Editorial(), Isbn: libro.Isbn(), Capitulo: libro.Capitulo()}}
		respuesta.Codigo, respuesta.Mensaje = 200, "Metadata del audiolibro encontrada"
	} else {
		respuesta.Codigo, respuesta.Mensaje = 400, "La metadata del audiolibro no se encontro"
	}
	return respuesta
}

// TODO ConsultarPodcast
func (this *MetadataAudioService) ConsultarPodcast(titulo string) dto.RespuestaPodcastDTO {
	var respuesta dto.RespuestaPodcastDTO
	podcast, encontrado := this.repository.BuscarPodcast(titulo)
	if encontrado {
		respuesta.VectorPodcast = []dto.MetadataPodcastDTO{{NombrePodcast: podcast.NombrePodcast(), TituloEpisodio: podcast.TituloEpisodio(), Anfitrion: podcast.Anfitrion(), NumeroTemporada: podcast.NumeroTemporada(), NotasShow: podcast.NotasShow(), ClasificacionContenido: podcast.ClasificacionContenido()}}
		respuesta.Codigo, respuesta.Mensaje = 200, "Metadata del podcast encontrada"
	} else {
		respuesta.Codigo, respuesta.Mensaje = 400, "La metadata del podcast no se encontro"
	}
	return respuesta
}

// TODO ConsultarRuidoBlanco
func (this *MetadataAudioService) ConsultarRuidoBlanco(tipoSonido string) dto.RespuestaRuidoBlancoDTO {
	var respuesta dto.RespuestaRuidoBlancoDTO
	ruido, encontrado := this.repository.BuscarRuidoBlanco(tipoSonido)
	if encontrado {
		respuesta.VectorRuidoBlanco = []dto.MetadataRuidoBlancoDTO{{TipoSonido: ruido.TipoSonido(), FuenteAudio: ruido.FuenteAudio(), UsoSugerido: ruido.UsoSugerido(), ProveedorContenido: ruido.ProveedorContenido(), DuracionBucle: ruido.DuracionBucle(), FrecuenciaDominante: ruido.FrecuenciaDominante()}}
		respuesta.Codigo, respuesta.Mensaje = 200, "Metadata del ruido blanco encontrada"
	} else {
		respuesta.Codigo, respuesta.Mensaje = 400, "La metadata del ruido blanco no se encontro"
	}
	return respuesta
}

//Obtener AudiosPorTipo retorna todo el vector de audios almacenados
func (this *MetadataAudioService) ObtenerAudiosPorTipo(tipo string) dto.RespuestaAudiosPorTipoDTO {
	var respuesta dto.RespuestaAudiosPorTipoDTO
	vecAudios := this.repository.ObtenerAudiosPorTipo(tipo)

	if len(vecAudios) > 0 {
		// Creamos un nuevo slice del tipo DTO con la misma capacidad
		vectorDTOs := make([]dto.MetadataAudioDTO, 0, len(vecAudios))

		// Mapeamos cada entidad a su DTO correspondiente
		for _, audio := range vecAudios {
			dtoAudio := dto.MetadataAudioDTO{
				Titulo:     audio.GetTitulo(),
				Duracion:   audio.GetDuracion(),
				Tipo:       audio.GetTipo(),
				Disponible: audio.GetDisponible(),
			}
			vectorDTOs = append(vectorDTOs, dtoAudio)
		}

		// Asignamos el nuevo slice de DTOs a la respuesta
		respuesta.VectorAudiosPorTipo = vectorDTOs
		respuesta.Codigo = 200
		respuesta.Mensaje = "Audios del tipo buscado encontrados"
	} else {
		respuesta.Codigo = 400
		respuesta.Mensaje = "No se encontraron audios del tipo buscado"
	}

	return respuesta
}

// TODO ObtenerMusica
func (this *MetadataAudioService) ObtenerMusica() dto.RespuestaMusicaDTO {
	var respuesta dto.RespuestaMusicaDTO
	for _, musica := range this.repository.ObtenerVectorMusica() {
		respuesta.VectorMusica = append(respuesta.VectorMusica, dto.MetadataMusicaDTO{ArtistaPrincipal: musica.GetArtistaPrincipal(), Album: musica.GetAlbum(), Genero: musica.GetGenero(), TituloCancion: musica.GetTituloCancion(), SelloDiscografico: musica.GetSelloDiscografico(), AnioLanzamiento: musica.GetAnioLanzamiento()})
	}
	if len(respuesta.VectorMusica) > 0 {
		respuesta.Codigo, respuesta.Mensaje = 200, "Música encontrada"
	} else {
		respuesta.Codigo, respuesta.Mensaje = 400, "No se encontro música"
	}
	return respuesta
}

// TODO ObtenerAudioLibros
func (this *MetadataAudioService) ObtenerAudioLibros() dto.RespuestaAudioLibrosDTO {
	var respuesta dto.RespuestaAudioLibrosDTO
	for _, libro := range this.repository.ObtenerVectorAudioLibros() {
		respuesta.VectorAudioLibros = append(respuesta.VectorAudioLibros, dto.MetadataAudioLibrosDTO{TituloLibro: libro.TituloLibro(), Autor: libro.Autor(), Narrador: libro.Narrador(), Editorial: libro.Editorial(), Isbn: libro.Isbn(), Capitulo: libro.Capitulo()})
	}
	if len(respuesta.VectorAudioLibros) > 0 {
		respuesta.Codigo, respuesta.Mensaje = 200, "Audiolibros encontrados"
	} else {
		respuesta.Codigo, respuesta.Mensaje = 400, "No se encontraron audiolibros"
	}
	return respuesta
}

// TODO ObtenerPodcast
func (this *MetadataAudioService) ObtenerPodcast() dto.RespuestaPodcastDTO {
	var respuesta dto.RespuestaPodcastDTO
	for _, podcast := range this.repository.ObtenerVectorPodcast() {
		respuesta.VectorPodcast = append(respuesta.VectorPodcast, dto.MetadataPodcastDTO{NombrePodcast: podcast.NombrePodcast(), TituloEpisodio: podcast.TituloEpisodio(), Anfitrion: podcast.Anfitrion(), NumeroTemporada: podcast.NumeroTemporada(), NotasShow: podcast.NotasShow(), ClasificacionContenido: podcast.ClasificacionContenido()})
	}
	if len(respuesta.VectorPodcast) > 0 {
		respuesta.Codigo, respuesta.Mensaje = 200, "Podcasts encontrados"
	} else {
		respuesta.Codigo, respuesta.Mensaje = 400, "No se encontraron podcasts"
	}
	return respuesta
}

// TODO ObtenerRuidoBlanco
func (this *MetadataAudioService) ObtenerRuidoBlanco() dto.RespuestaRuidoBlancoDTO {
	var respuesta dto.RespuestaRuidoBlancoDTO
	for _, ruido := range this.repository.ObtenerVectorRuidoBlanco() {
		respuesta.VectorRuidoBlanco = append(respuesta.VectorRuidoBlanco, dto.MetadataRuidoBlancoDTO{TipoSonido: ruido.TipoSonido(), FuenteAudio: ruido.FuenteAudio(), UsoSugerido: ruido.UsoSugerido(), ProveedorContenido: ruido.ProveedorContenido(), DuracionBucle: ruido.DuracionBucle(), FrecuenciaDominante: ruido.FrecuenciaDominante()})
	}
	if len(respuesta.VectorRuidoBlanco) > 0 {
		respuesta.Codigo, respuesta.Mensaje = 200, "Ruidos blancos encontrados"
	} else {
		respuesta.Codigo, respuesta.Mensaje = 400, "No se encontraron ruidos blancos"
	}
	return respuesta
}

// TODO ObtenerTiposDeAudio
func (this *MetadataAudioService) ObtenerTiposDeAudio() []dto.TipoAudioDTO {
	var tipos []dto.TipoAudioDTO
	for _, tipo := range this.repository.ObtenerVectorTiposDeAudio() {
		tipos = append(tipos, dto.TipoAudioDTO{Id: tipo.Id(), Tipo: tipo.Tipo()})
	}
	return tipos
}
