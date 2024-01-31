package main

import (
	"log"
	"net/http"

	"pijar_camp/config"
	"pijar_camp/controllers/provinsicontroller"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/provinsi", provinsicontroller.Index)
	http.HandleFunc("/provinsi/add", provinsicontroller.Add)
	http.HandleFunc("/provinsi/edit", provinsicontroller.Edit)
	http.HandleFunc("/provinsi/delete", provinsicontroller.Delete)

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
