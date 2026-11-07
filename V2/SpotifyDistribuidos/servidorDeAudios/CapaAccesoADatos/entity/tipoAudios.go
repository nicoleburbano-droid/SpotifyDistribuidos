package entity

type TipoAudio struct{
	id int
	tipo string
}

// getters

func (t *TipoAudio) Id() int {
	return t.id
}

func (t *TipoAudio) Tipo() string {
	return t.tipo
}

// setters

func (t *TipoAudio) SetId(id int) {
	t.id = id
}

func (t *TipoAudio) SetTipo(tipo string) {
	t.tipo = tipo
}
