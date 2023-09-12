package mascot_test

import (
	"testing"

	"github.com/michelle083/4143_PLC/Assignments/P01/mascot"
)

// function to test the mascot package
func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
	  t.Fatal("Wrong Mascot")
	}
}
