package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("USAGE: %s.go URL\n", filepath.Base(os.Args[0]))
		os.Exit(-1)
	}

	URL, err := url.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}

	// NOTE: todavía no descubro porque tener solo la referencia
	client := &http.Client{Timeout: 10 * time.Second}

	// Aquí creamos solo la petición
	request, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		panic(err)
	}

	//  Enviamos la petición creada anteriormente a través de un cliente http
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	fmt.Println("STATUS CODE:", response.Status)
	fmt.Println(strings.Repeat("-", 30))

	// Obteniendo el header
	header, _ := httputil.DumpResponse(response, false)
	fmt.Println(string(header))
	fmt.Println(strings.Repeat("-", 30))

	// Obteniendo el response con body incluido
	bodyIncluded, _ := httputil.DumpResponse(response, true)
	fmt.Println(string(bodyIncluded))
	fmt.Println(strings.Repeat("-", 30))

	// Un mejor modo de explorar el header
	contentTypeHeaderData := strings.Split(response.Header.Get("Content-Type"), "charset=")
	if len(contentTypeHeaderData) != 2 {
		panic("No Content-Type header to show")
	}
	contentTypeHeader := contentTypeHeaderData[1]
	fmt.Println(contentTypeHeader)
	fmt.Println(strings.Repeat("-", 30))

	// Recorriendo todos los headers
	for k, v := range response.Header {
		fmt.Printf("%v => %v\n", k, v)
	}
	fmt.Println(strings.Repeat("-", 30))

	// Leyendo el tamaño del contenido en el body
	if response.ContentLength == -1 {
		panic("Cannot read response content or there is no content")
	} else {
		fmt.Println(response.ContentLength)
	}
	fmt.Println(strings.Repeat("-", 30))

	// Leyendo el contenido del body
	fmt.Println(response.Body)
	fmt.Println(strings.Repeat("-", 30))

	// NOTE: Esto es una técnica para encontrar el tamaño de un buffer

	// Crear contador y buffer para luego sacar la respuesta
	length := 0
	var buffer [1024]byte
	r := response.Body
	// Leer tanto como puede el buffer
	for {
		// Aquí se pide un slice del array que resulta ser el array completo
		// 	esto pasa el tipo de parámetros requeridos para r.Read
		n, err := r.Read(buffer[0:])
		// No tienes más que leer? Termina
		if err != nil {
			fmt.Println(err)
			break
		}
		// Suma el tamaño que tenías + lo nuevo
		length = length + n
	}
	// Voilá
	fmt.Println("Calculated response data length:", length)

}
