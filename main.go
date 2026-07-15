package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/gunnu", gunnu_diva)
	fmt.Println("on the port 8000")
	http.ListenAndServe(":8000", nil)
}
func gunnu_diva(w http.ResponseWriter, r *http.Request) {
	fmt.Println("gunnu is the diva")
}
