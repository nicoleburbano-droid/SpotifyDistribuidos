package vistas

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// lectorEntrada es único para todo el programa; crear varios bufio.Reader
// sobre os.Stdin puede hacer que se pierdan líneas escritas por el usuario.
var lectorEntrada = bufio.NewReader(os.Stdin)

// leerLinea lee una línea de la consola. Si la entrada se cierra (Ctrl+D)
// termina el programa para no quedar en un ciclo infinito.
func leerLinea() string {
	linea, err := lectorEntrada.ReadString('\n')
	if err == io.EOF && linea == "" {
		fmt.Println("\nHasta pronto.")
		os.Exit(0)
	}
	return strings.TrimSpace(linea)
}

// leerOpcion lee una opción entre 1 y maximo; retorna 0 si no es válida.
func leerOpcion(maximo int) int {
	fmt.Print("\nIngrese la opcion del menu: ")
	opcion, err := strconv.Atoi(leerLinea())
	if err != nil || opcion < 1 || opcion > maximo {
		fmt.Println("Opcion no valida, intentalo nuevamente")
		return 0
	}
	return opcion
}

// pedirTexto pide un texto obligatorio.
func pedirTexto(etiqueta string) string {
	for {
		fmt.Printf("%s: ", etiqueta)
		if texto := leerLinea(); texto != "" {
			return texto
		}
		fmt.Println("Este campo es obligatorio.")
	}
}

// pedirEntero pide un número entero mayor o igual a cero.
func pedirEntero(etiqueta string) int {
	for {
		fmt.Printf("%s: ", etiqueta)
		numero, err := strconv.Atoi(leerLinea())
		if err == nil && numero >= 0 {
			return numero
		}
		fmt.Println("Debe ingresar un número entero.")
	}
}
