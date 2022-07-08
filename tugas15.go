package main

import (
	"fmt"
	"net/http"
)

func angka(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= 100; i++ {
		fmt.Fprintln(w, i)
	}
}

func main() {
	// Setup Menu Utama
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Selamat datang di menu utama")
	})

	// Router
	http.HandleFunc("/angka", angka)

	// Setup server
	fmt.Println("Server 8000 activate")
	http.ListenAndServe(":8000", nil)
}
