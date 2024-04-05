package mascot_test

import (
	"mods/mascot"
	"testing"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Gopher" {
		t.Error("Expected Gopher")
	}
}
