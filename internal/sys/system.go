package sys

import (
	"fmt"
	"os/exec"

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
