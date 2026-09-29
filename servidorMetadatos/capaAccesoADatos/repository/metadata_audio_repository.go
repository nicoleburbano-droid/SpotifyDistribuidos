package repository

import "microservicio/capaAccesoADatos/entity"

// MetadataAudioRepository es un repositorio que mantiene en memoria
// el slice de audios y expone las operaciones de búsqueda y registro.

type MetadataAudioRepository struct {
	vectorMetadataAudios []entity.MetadataAudio
}

// NewMetadataAudioRepository crea el repositorio y lo precarga con la
// metadata de audios de ejemplo.

func NewMetadataAudioRepository() *MetadataAudioRepository {
	this := &MetadataAudioRepository{}
	this.CargarMetadataAudios()
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

// RegistrarAudio agrega un nuevo audio al vector.

func (this *MetadataAudioRepository) RegistrarAudio(audio entity.MetadataAudio) {
	this.vectorMetadataAudios = append(this.vectorMetadataAudios, audio)
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
