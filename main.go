package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/budhirajamadhav/url-shortener/router"
)

func main() {
	
	r := router.Router()
	fmt.Println("Server is getting started ....")
	log.Fatal(http.ListenAndServe(":4000", r))

}
