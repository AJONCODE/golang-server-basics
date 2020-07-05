package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AJONCODE/golangServerMisc/02_hittingExternalApi/route"
)

func main() {
	http.HandleFunc("/people", route.GetPeople())

	fmt.Println("Server on :8080")

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
