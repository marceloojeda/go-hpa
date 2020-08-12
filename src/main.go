package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	sqrt := Sqrt(0.0001, 10000000)
	greeting := greeting(sqrt)
	data := []byte(greeting)

	res.Write(data)
}

func greeting(term string) string {
	return fmt.Sprintf("<b>Code.education Rocks!</b></br><small>Square root: %s</small>", term)
}

func Sqrt(x float64, iterations int) string {

	var z float64 = 0.0

	for i := 0; i < iterations; i++ {
		z += math.Sqrt(x)
	}

	return strconv.FormatFloat(z, 'f', 6, 64)
}

func main() {
	// create a new handler
	handler := HttpHandler{}

	// listen and serve
	http.ListenAndServe(":8000", handler)
}
