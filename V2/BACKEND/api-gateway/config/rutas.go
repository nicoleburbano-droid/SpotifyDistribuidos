package config

// Ruta asocia un prefijo público del gateway con el nombre de un servicio en Consul.
type Ruta struct {
	Prefijo  string // ruta pública en el gateway, p. ej. /api/productos
	Servicio string // nombre con el que el microservicio se registró en Consul
	Destino  string // ruta interna en el microservicio, p. ej. /productos
}

var Rutas = []Ruta{
	{Prefijo: "/api/audios", Servicio: "servidorDeAudios", Destino: "/audios"},
	{Prefijo: "/api/musica", Servicio: "servidorDeAudios", Destino: "/musica"},
	{Prefijo: "/api/audiolibros", Servicio: "servidorDeAudios", Destino: "/audiolibros"},
	{Prefijo: "/api/podcasts", Servicio: "servidorDeAudios", Destino: "/podcasts"},
	{Prefijo: "/api/ruido-blanco", Servicio: "servidorDeAudios", Destino: "/ruido-blanco"},
	{Prefijo: "/api/tipos-audio", Servicio: "servidorDeAudios", Destino: "/tipos-audio"},
	{Prefijo: "/api/estadisticas", Servicio: "servidorEstadisticas", Destino: "/actuator"},
	{Prefijo: "/api/reacciones", Servicio: "servidorReacciones", Destino: "/"},
	// Envoy sigue escuchando en 8080 y conserva el destino gRPC 127.0.0.1:50051.
	{Prefijo: "/api/streaming", Servicio: "servidorStreaming", Destino: "/"},
	{Prefijo: "/api/pasarela", Servicio: "servidorSuscripciones", Destino: "/api"},
	{Prefijo: "/api/suscripciones", Servicio: "servidorSuscripciones", Destino: "/api/suscripciones"},
	{Prefijo: "/api/pagos", Servicio: "servidorSuscripciones", Destino: "/api/pagos"},
	{Prefijo: "/api/tokens", Servicio: "servidorSuscripciones", Destino: "/api/tokens"},
	{Prefijo: "/api/transacciones", Servicio: "servidorSuscripciones", Destino: "/api/transacciones"},
	{Prefijo: "/api/estado", Servicio: "servidorSuscripciones", Destino: "/api/estado"},
}
