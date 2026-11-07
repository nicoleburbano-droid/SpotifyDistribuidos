package capaaccesodatos

import (
	"fmt"
	"os"
)

// abre el archivo del audio dado su titulo y lo devuelve como *os.File
func AbrirArchivo(titulo string) (*os.File, error) {
	file, err := os.Open(titulo)
	if err != nil {
		fmt.Println("Error Audio abierto:", titulo)
		return nil, fmt.Errorf("No se pudo abrir el archivo: %w", err)
	}
	fmt.Println("Audio abierto: ", titulo)
	return file, nil
}
