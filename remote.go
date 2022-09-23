package main

import (
	"net/http"

	"github.com/nickmancari/vpn_remote/handlers"
)

/*
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}
*/
func main() {

	/*
	static := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", static))
*/
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/stop", handlers.Stop)
	http.HandleFunc("/start", handlers.Start)
	http.HandleFunc("/address", handlers.Address)
	http.HandleFunc("/reboot", handlers.Reboot)
	http.HandleFunc("/media", handlers.MediaController)
	http.HandleFunc("/move", handlers.Move)
	http.HandleFunc("/delete", handlers.Delete)
	http.HandleFunc("/remove", handlers.Remove)
	http.ListenAndServe(":8080", nil)

}
