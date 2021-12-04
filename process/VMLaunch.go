package process

import (
	"fmt"
	"log"
	"os/exec"
)

func LaunchVM() {
	out, err := exec.Command("bash", "./process/echo.sh").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
