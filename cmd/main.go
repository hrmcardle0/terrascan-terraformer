package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"gitlab.com/secops/development/aws/terrascan/resource"
	"gitlab.com/secops/development/aws/terrascan/terraformer"
)

const pathToTerraformer = "/usr/local/bin/terraformer"

func Setup(resource *resource.Resource) error {
	log.Printf("Setup function called with resource: %v\n", resource)

	// initialize terraform
	if err := InitTerraform(); err != nil {
		log.Println("Error Initializing Terraform Provider: %v", err)
	}

	// start terrawatch
	if err := terraformer.InitTerraformer(resource, pathToTerraformer); err != nil {
		log.Printf("Error Initializing Terraformer: %v: ", err)
	}

	return nil
}

// Initialize terraform provider
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
