package helpers

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") //for windows only
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
