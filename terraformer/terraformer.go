package terraformer

import (
	"fmt"
	"gitlab.com/secops/development/aws/terrascan/resource"
	"log"
	"os/exec"
	//"strings"
)

func InitTerraformer(resource *resource.Resource, path string) error {

	// remove old generated
	cliList := []string{"-rf", "generated/"}
	log.Println("Cleaning up old directory...")

	cmd := exec.Command("rm", cliList...)
	cmd.Dir = "./terraform"
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	log.Println(string(output))

	// run terraformer
	cliList = resource.ToCliList()
	log.Printf("Running: terraformer %v\n", cliList)

	cmd = exec.Command(fmt.Sprintf("%s", path), cliList...)
	cmd.Dir = "./terraform"
	output, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	log.Println(string(output))
	return nil
}
