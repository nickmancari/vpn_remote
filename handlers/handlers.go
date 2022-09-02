package handlers

import (
	"fmt"
	"net/http"
	"html/template"
	"os/exec"
	"strings"
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
		fmt.Printf("Status Function Error: %s\n", err)
	}

	s := string(statusOutput)

	statusRange := strings.Split(s, " ")

	i := 0
	var statusString string
	for range statusRange {
		if strings.Contains(statusRange[i], "Active") {
			statusString = fmt.Sprintf("Systemctl Status:: %s %s", statusRange[i], statusRange[i+1])
		}
		i++
	}

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

func ListedShows(w http.ResponseWriter, r *http.Request) {
	cmd, err := exec.Command("sudo", "ls", "/media/tux/MOTHERSHIP/TV/").Output()
	if err != nil {
		fmt.Println(err)
	}

	s := string(cmd)
	files := strings.Split(s, "\n")

	data := struct{
		Showlist []string
	}{
		Showlist: files,
	}

	errs := tpl.ExecuteTemplate(w, "tvshows.html", data)
	if errs != nil {
		fmt.Println(errs)
	}



}

func Purge(w http.ResponseWriter, r *http.Request) {

	cmd := exec.Command("sudo", "/usr/bin/bash", "./scripts/purge_files.sh")


	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

        errors := tpl.ExecuteTemplate(w, "purge.html", nil)
        if errors != nil {
                fmt.Println(errors)
        }

}
