package cmd // Initialize terraform provider

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func InitTerraform() error {
	log.Println("Initializing Terraform Provider")

	// get current directory
	cmd := exec.Command("pwd")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	currentDir := string(output)
	workDir := strings.ReplaceAll(fmt.Sprintf("-chdir=%s/terraform", currentDir), "\n", "")

	// init terraform
	cmd = exec.Command("terraform", workDir, "init")
	output, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	log.Println(string(output))

	return nil

}
