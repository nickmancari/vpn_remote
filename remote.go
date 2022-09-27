package main

import (
	"net/http"

	"github.com/nickmancari/vpn_remote/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/stop", handlers.Stop)
	http.HandleFunc("/start", handlers.Start)
	http.HandleFunc("/reboot", handlers.Reboot)
	http.HandleFunc("/media", handlers.MediaController)
	http.HandleFunc("/move", handlers.Move)
	http.HandleFunc("/delete", handlers.Delete)
	http.HandleFunc("/remove", handlers.Remove)
	http.ListenAndServe(":8080", nil)

}
