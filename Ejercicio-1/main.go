package main

import (
	"fmt"
	"os"
)

type producto struct {
	id       string
	precio   float64
	cantidad int
}

func main() {
	var listaDeProductos []producto
	listaDeProductos = cargarListaDeProductos(listaDeProductos)
	texto := obtenerFormatoCSV(listaDeProductos)
	guardarArchivoCSV("./file.txt", texto)
}

func obtenerFormatoCSV(lista []producto) string {
	var texto string
	for _, prod := range lista {
		texto += fmt.Sprintf("%s,%.2f,%d\n", prod.id, prod.precio, prod.cantidad)
	}

	return texto
}

func cargarListaDeProductos(lista []producto) []producto {
	producto1 := producto{"001", 650, 5}
	producto2 := producto{"002", 1890, 3}
	producto3 := producto{"003", 14580, 10}

	lista = append(lista, producto1)
	lista = append(lista, producto2)
	lista = append(lista, producto3)

	return lista
}

func guardarArchivoCSV(filename, texto string) {
	bytesData := []byte(texto)
	os.WriteFile(filename, bytesData, os.FileMode(0644))
}
