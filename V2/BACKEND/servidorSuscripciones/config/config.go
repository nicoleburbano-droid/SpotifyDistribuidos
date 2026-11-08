package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config agrupa la configuración del servidor leída desde variables de entorno.
type Config struct {
	Puerto          string
	DirectorioDatos string
}

// Cargar lee las variables de entorno y aplica valores por defecto.
//
//	PUERTO            puerto HTTP (por defecto 8080)
//	MODO_FALLO        NINGUNO | PETICION | RESPUESTA (por defecto NINGUNO)
//	CANTIDAD_FALLOS   número de peticiones de pago, por token, que fallarán (por defecto 0)
//	RETARDO_FALLO_MS  tiempo que el servidor "cuelga" la conexión antes de cerrarla (por defecto 5000)
//	DIRECTORIO_DATOS  carpeta donde se persisten tokens y transacciones (por defecto ./datos)
func Cargar() Config {
	cfg := Config{
		Puerto:          valorOPorDefecto("PUERTO", "8080"),
		DirectorioDatos: valorOPorDefecto("DIRECTORIO_DATOS", "./datos"),
	}

	return cfg
}

func valorOPorDefecto(clave, porDefecto string) string {
	if v := strings.TrimSpace(os.Getenv(clave)); v != "" {
		return v
	}
	return porDefecto
}
