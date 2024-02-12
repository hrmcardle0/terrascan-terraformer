package cmd

import (
	//"bytes"
	//"encoding/json"
	//"fmt"
	"log"
	//"net/http"
	//"os/exec"
	"strings"

	"gitlab.com/secops/development/aws/terrascan/helpers"
	"gitlab.com/secops/development/aws/terrascan/resource"
	"gitlab.com/secops/development/aws/terrascan/terraformer"
)

var (
	region = "eu-west-1"
)

const pathToTerraformer = "/usr/local/bin/terraformer"

func Init(e string) error {

	event, err := helpers.ToJson(e)
	if err != nil {
		log.Println("Error parsing json for event: ", e)
		return err
	}

	// setup target resource block
	targetResource := resource.Resource{
		Name:    strings.Split(event.EventSource, ".amazonaws.com")[0],
		Region:  region,
		Filters: "Name=id;Value=" + event.BucketName,
	}

	/*
		* Cli Configuration
		// retreive input variables
		// resource-name: string of resources in terraformer format, ex...s3,ec2
		// filters: string representations of filters, ex..vpc=vpc_id1:vpc_id2:vpc_id3
		flag.StringVar(&targetResource.Name, "resource-name", "None", "Resource Name Flag Set")
		flag.StringVar(&targetResource.Filters, "filters", "None", "Resource Name Flag Set")
		flag.Parse()
	*/

	if err := Setup(&targetResource); err != nil {
		return err
	}

	return nil

}

func Setup(resource *resource.Resource) error {
	log.Printf("Setup function called with resource: %v\n", resource)

	// initialize terraform
	if err := InitTerraform(); err != nil {
		log.Println("Error Initializing Terraform Provider: %v", err)
		return err
	}

	// start terraformer
	if err := terraformer.InitTerraformer(resource, pathToTerraformer); err != nil {
		log.Printf("Error Initializing Terraformer: %v: ", err)
		return err
	}

	// create string with generated terraform
	terraformString, err := terraformer.GenerateString()
	if err != nil {
		log.Printf("Error generating master terraform string: %v: ", err)
		return err
	}

	// call tfsec endpoint
	if err = InitHttp(terraformString); err != nil {
		log.Printf("Error calling HTTP endpoint: %v\n", err)
		return err
	}
	return nil
}
