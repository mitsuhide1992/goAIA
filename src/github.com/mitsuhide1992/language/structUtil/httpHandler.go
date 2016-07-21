package structUtil

import (
	"fmt"
	// "log"
	"net/http"
)

type Hello struct {
}

func (Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, zly")
}

func sayHello() {
	fmt.Println("hello package success")
}

// func main() {
// 	var h Hell
// 	var err error = http.ListenAndServe("localhost:4000", h)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
