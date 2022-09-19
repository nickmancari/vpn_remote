package handlers

import (
	"fmt"
	"net/http"
	"html/template"
	"os/exec"
	"strings"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {

	statusOutput, _ := exec.Command("systemctl", "status", "openvpn").Output()
/*	Not handle error here for now
	if err != nil {
		fmt.Printf("Status Function Error: %s\n", err)
	}
*/
	s := string(statusOutput)

	statusRange := strings.Split(s, " ")

	i := 0
	var statusString string
	for range statusRange {
		if strings.Contains(statusRange[i], "Active") {
			statusString = fmt.Sprintf(" %s %s", statusRange[i], statusRange[i+1])
		}
		i++
	}

	vpnStatus := struct{
		Stat	string
	}{
		Stat:	statusString,
	}

	errs := tpl.ExecuteTemplate(w, "index.html", vpnStatus)
	if errs != nil {
		fmt.Println(errs)
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
			statusString = fmt.Sprintf(" %s %s", statusRange[i], statusRange[i+1])
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


func MediaController(w http.ResponseWriter, r *http.Request) {
	cmd, err := exec.Command("sudo", "ls", "/media/tux/MOTHERSHIP/TV/").Output()
	if err != nil {
		fmt.Println(err)
	}

	dwnld, errs := exec.Command("sudo", "ls", "/var/lib/transmission-daemon/downloads/").Output()
	if err != nil {
		fmt.Println(errs)
	}

	d := string(dwnld)
	downloadFiles := strings.Split(d, "\n")


	s := string(cmd)
	files := strings.Split(s, "\n")

	data := struct{
		Showlist []string
		Downloadlist []string
	}{
		Showlist: files,
		Downloadlist: downloadFiles,
	}

	errors := tpl.ExecuteTemplate(w, "media.html", data)
	if errs != nil {
		fmt.Println(errors)
	}



}

func Move(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	directory := r.FormValue("directory")
	media := r.FormValue("current")


	if directory == "Movies" {

		files, errs := os.ReadDir("/var/lib/transmission-daemon/downloads/"+media)
		if errs != nil {
			fmt.Println(errs)
		}

		for _, file := range files {

			if strings.Contains(file.Name(), ".mp4") {
				cmd := exec.Command("mv", "/var/lib/transmission-daemon/downloads/"+media+"/"+file.Name(), "/media/tux/MOTHERSHIP/Movies/")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
			}

		}
	} else {
		cmd := exec.Command("mv", "/var/lib/transmission-daemon/downloads/"+media, "/media/tux/MOTHERSHIP/TV/"+directory)

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	err := tpl.ExecuteTemplate(w, "move.html", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {

	dwnld, errs := exec.Command("sudo", "ls", "/var/lib/transmission-daemon/downloads/").Output()
	if errs != nil {
		fmt.Println(errs)
	}

	d := string(dwnld)
	downloadFiles := strings.Split(d, "\n")


	data := struct{
		Downloadlist []string
	}{
		Downloadlist: downloadFiles,
	}

	errors := tpl.ExecuteTemplate(w, "delete.html", data)
	if errs != nil {
		fmt.Println(errors)
	}

}

func Remove(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	media := r.FormValue("current")

	cmd := exec.Command("sudo", "rm", "-rf",  "/var/lib/transmission-daemon/downloads/"+media)

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	errors := tpl.ExecuteTemplate(w, "removed.html", nil)
	if err != nil {
		fmt.Println(errors)
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
