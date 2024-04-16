package terraformer

import (
	"fmt"
	"gitlab.com/secops/development/aws/terrascan/resource"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	cliList, isIam := resource.ToCliList()
	log.Printf("Running: terraformer %v\n", cliList)

	// if we are running IAM Checks and the source is not the cyber account, run the custom script
	if isIam {
		log.Println("IAM is set")
		cmd = exec.Command("/home/ec2-user/environment/terrascan/iam.sh", resource.Account)
	} else {
		cmd = exec.Command(fmt.Sprintf("%s", path), cliList...)
	}
	cmd.Dir = "./terraform"
	output, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	log.Println(string(output))
	return nil
}

func GenerateString() (string, error) {

	// Directory containing directories with .tf files
	rootDir := "./terraform/generated"
	var masterString string

	// Walk through root directory
	err := filepath.Walk(rootDir, func(dirPath string, dirInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip processing the root directory itself
		if dirPath == rootDir {
			return nil
		}

		// Check if the current entry is a directory
		if dirInfo.IsDir() {
			// Iterate over files in the directory
			files, err := ioutil.ReadDir(dirPath)
			if err != nil {
				return err
			}

			// Iterate over each file in the directory
			for _, file := range files {
				// Check if file is a .tf file
				if filepath.Ext(file.Name()) == ".tf" {
					// Read file contents
					contents, err := ioutil.ReadFile(filepath.Join(dirPath, file.Name()))
					if err != nil {
						return err
					}

					// Append file contents to master string
					masterString += string(contents)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// Print master string
	return masterString, nil
}
