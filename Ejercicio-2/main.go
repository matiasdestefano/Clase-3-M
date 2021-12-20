package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type producto struct {
	id       string
	precio   float64
	cantidad int
}

func main() {
	file, _ := os.Open("./file.txt")
	texto := escribirCabeceraArchivo()
	scanner := bufio.NewScanner(file)

	var total float64
	for scanner.Scan() {
		var prod producto = obtenerProducto(scanner.Text())
		texto += fmt.Sprintf("\n%s%19.2f%20d", prod.id, prod.precio, prod.cantidad)
		total += float64(prod.cantidad) * prod.precio
	}

	texto += fmt.Sprintf("\n%22.2f", total)
	guardarArchivo("./resultado.txt", texto)
}

func escribirCabeceraArchivo() string {
	return fmt.Sprintf("ID%20s%20s", "Precio", "Cantidad")
}

func obtenerProducto(linea string) producto {
	datos := strings.Split(linea, ",")
	var prod producto
	prod.id = datos[0]
	fmt.Sscanf(datos[1], "%f", &prod.precio)
	fmt.Sscanf(datos[2], "%d", &prod.cantidad)
	return prod
}

func guardarArchivo(filename, texto string) {
	bytesData := []byte(texto)
	os.WriteFile(filename, bytesData, os.FileMode(0644))
}
