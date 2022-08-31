package handlers

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

func Index(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func Status(w http.ResponseWriter, r *http.Request) {

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

func Stop(w http.ResponseWriter, r *http.Request) {
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

func Start(w http.ResponseWriter, r *http.Request) {
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


func Address(w http.ResponseWriter, r *http.Request) {
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

func Reboot(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sudo", "shutdown", "-r", "now")

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

        errors := tpl.ExecuteTemplate(w, "reboot.html", nil)
        if errors != nil {
                fmt.Println(errors)
        }

}

func Movies(w http.ResponseWriter, r *http.Request) {
	cmd, err := exec.Command("sudo", "ls", "/var/lib/transmission-daemon/downloads/").Output()
	if err != nil {
		fmt.Println(err)
	}

	files := string(cmd)

	data := struct{
		Movielist string
	}{
		Movielist: files,
	}

	errs := tpl.ExecuteTemplate(w, "movies.html", data)
	if errs != nil {
		fmt.Println(errs)
	}



}

func MoveMovies(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("/usr/bin/bash", "./scripts/move_movie.sh")


	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

        errors := tpl.ExecuteTemplate(w, "movemovies.html", nil)
        if errors != nil {
                fmt.Println(errors)
        }

}

func Purge(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("/usr/bin/bash", "./scripts/purge_files.sh")


	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

        errors := tpl.ExecuteTemplate(w, "prugefiles.html", nil)
        if errors != nil {
                fmt.Println(errors)
        }

}
