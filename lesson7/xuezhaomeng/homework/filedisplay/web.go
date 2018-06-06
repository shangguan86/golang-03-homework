package main

import (
	"fmt"
	"net/http"
)

func hander(w http.ResponseWriter, r *http.Request) {
	b, err := FileGet()
	CheckError(err)
	fmt.Fprintf(w, string(b))

}

func main() {

	//dir := os.Args[1]
	dir := "/tmp"
	ReadDir(&dir)

	http.HandleFunc("/", hander)
	http.ListenAndServe(":8080", nil)
}
