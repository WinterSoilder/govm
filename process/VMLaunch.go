package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func LaunchVM(vm_data map[string]interface{}) {
	godotenv.Load("../project.env")
	TEMPLATE := os.Getenv("TEMPLATE")
	content, err := ioutil.ReadFile("/home/shashank/packer-qemu-templates/ubuntu/ubuntu.json")
	if err != nil {
		log.Fatal(err)
	}

	// Declared an empty interface
	var result interface{}
	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(content), &result)
	m := result.(map[string]interface{})
	variables := m["variables"].(map[string]interface{})

	variables["vm_name"] = vm_data["VM_name"]
	variables["cpus"] = vm_data["CPUs"]
	variables["disk_size"] = vm_data["Disk"]
	variables["memory"] = vm_data["Memory"]
	variables["ssh_fullname"] = vm_data["SSH_username"]
	variables["ssh_password"] = vm_data["SSH_password"]
	variables["ssh_username"] = vm_data["SSH_username"]
	byteValue, _ := json.Marshal(m)

	// // Write back to file
	ioutil.WriteFile(TEMPLATE+"ubuntu.json", byteValue, 0644)

	cmd := exec.Command("packer", "build", "ubuntu.json")
	cmd.Dir = TEMPLATE
	out, cmdErr := cmd.Output()
	if cmdErr != nil {
		fmt.Print("err", cmdErr)
	}
	fmt.Print("out", out)
}
