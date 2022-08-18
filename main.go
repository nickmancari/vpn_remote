package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os/exec"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/status", status)
	http.HandleFunc("/stop", stop)
	http.HandleFunc("/start", start)
	http.HandleFunc("/address", address)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func status(w http.ResponseWriter, r *http.Request) {

	statusOutput, err := exec.Command("systemctl", "status", "openvpn").Output()
	if err != nil {
		fmt.Println(err)
	}

	statusString := string(statusOutput)

	vpnStatus := struct{
		Stat	string
	}{
		Stat:	statusString,
	}

	errs := tpl.ExecuteTemplate(w, "status.html", vpnStatus)
	if errs != nil {
		fmt.Println(err)
	}

}

func stop(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sudo", "systemctl", "stop", "openvpn")

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	errors := tpl.ExecuteTemplate(w, "stop.html", nil)
	if errors != nil {
		fmt.Println(errors)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sudo", "systemctl", "start", "openvpn")

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	errors := tpl.ExecuteTemplate(w, "start.html", nil)
	if errors != nil {
		fmt.Println(errors)
	}
}


func address(w http.ResponseWriter, r *http.Request) {
	cmd, err := exec.Command("curl", "ipv4.icanhazip.com").Output()
	if err != nil {
		fmt.Println(err)
	}

	location := string(cmd)

	data := struct{
		Address string
	}{
		Address: location,
	}

	errs := tpl.ExecuteTemplate(w, "address.html", data)
	if errs != nil {
		fmt.Println(errs)
	}
}




