package sys

import (
	"testing"
)

func TestVpnSettingStatus(t *testing.T) {	
	got := VpnSettingStatus()
	want := false

	if want != got {
		t.Errorf("eExpected '%v', but got '%v'", want, got)
	}
}

func TestInternetProtocolAddress(t *testing.T) {
	got := InternetProtocolAddress()
	want := "IP Address Unaviable"

	if got != want {
		t.Errorf("Expected '%v', but got '%v'", want, got)
	}
}

func TestVpnDaemonStatus(t *testing.T) {
	got := VpnDaemonStatus()
	want := "VPN Status Unknown"

	if got != want {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
