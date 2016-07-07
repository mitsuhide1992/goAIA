package main

import (
	"fmt"
	"log"
	"net/http"
)

type zString string

type zStruct struct {
	who   string
	where string
	when  string
}

func (zs zString) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<button>zlu</button>")
}

func (zs *zStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, zs.when, zs.where, zs.who)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", zString("I'm a frayed knot."))
	http.Handle("/struct", &zStruct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
