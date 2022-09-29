package sys

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/nickmancari/vpn_remote/pkg/config"

)

func VpnSettingStatus() bool {

	status := config.Read().VpnSetting()

	if status == true {
		return true
	} else {
		return false
	}

}

func InternetProtocolAddress() string {
	//run system cmd to get external ip
	//return ip as string

	status := VpnSettingStatus()

	if status == true {
		cmd, err := exec.Command("curl", "ipv4.icanhazip.com").Output()
		if err != nil {
			fmt.Println(err)
		}

		return string(cmd)
	} else {
		return fmt.Sprint("IP Address Unaviable")
	}

}



func VpnDaemonStatus() string {

	status := VpnSettingStatus()
	var statusString string

	if status == true {
		systemctl, _ := exec.Command("systemctl", "status", "openvpn").Output()

		s := string(systemctl)

		statusRange := strings.Split(s, " ")

		i := 0
		for range statusRange {
			if strings.Contains(statusRange[i], "Active") {
				statusString = fmt.Sprintf(" %s ", strings.ToUpper(statusRange[i+1]))
				return statusString
			}
			i++
		}
	} else {
		return fmt.Sprint("VPN Status Unknown")
	}

	return ""
}

