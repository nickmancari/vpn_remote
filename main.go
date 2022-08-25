package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os/exec"

	"github.com/nickmancari/vpn_remote/handlers"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/status", handlers.Status)
	http.HandleFunc("/stop", handlers.Stop)
	http.HandleFunc("/start", handlers.Start)
	http.HandleFunc("/address", handlers.Address)
	http.HandleFunc("/reboot", handlers.Reboot)
	http.ListenAndServe(":8080", nil)

}
