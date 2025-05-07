package main

import (
	"fmt"
	"strings"
)

// Función genérica para filtrar
// se puede optimizar para usar un algoritmo de busqueda optimo en caso que se requiera buscar eficiencia en cambio del for normal
// Función genérica para filtrar utilizando concurrencia
// Función genérica para filtrar utilizando un simple for loop
func Where[T any](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list)) // Preallocate memory to avoid multiple allocations
	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

type User struct {
	Cedula   string
	Nombre   string
	Apellido string
}

func main() {
	// Lista de usuarios
	usuarios := []User{
		{Cedula: "8014334567", Nombre: "Juan", Apellido: "Pérez"},
		{Cedula: "1234567890", Nombre: "Ana", Apellido: "Gómez"},
		{Cedula: "8014987654", Nombre: "Luis", Apellido: "Martínez"},
		{Cedula: "5678123456", Nombre: "María", Apellido: "López"},
	}

	// Filtrar las cédulas que comienzan con "80143 enviando a la funcion where la callback func que se encarga de validar si un usuario cumple la condicion"
	result := Where(usuarios, func(u User) bool {
		return strings.Contains(u.Cedula, "80143")
	})

	// Imprimir los resultados
	for _, u := range result {
		fmt.Printf("Usuario: %s %s, Cédula: %s\n", u.Nombre, u.Apellido, u.Cedula)
	}
}
