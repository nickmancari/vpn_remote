package sys

import (
	"testing"
)

func TestVpnSettingStatus(t *testing.T) {	
	got := VpnSettingStatus()
	want := false

	if want != got {
		t.Errorf("expected %v, but got %v", want, got)
	}
}
