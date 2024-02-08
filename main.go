package main

import (
	"flag"
	"gitlab.com/secops/development/aws/terrascan/cmd"
	"gitlab.com/secops/development/aws/terrascan/resource"
	"log"
	//"github.com/spf13/viper"
)

var (
	region = "eu-west-1"
)

func main() {

	log.Println("Terrascan Started")

	// target resource block
	targetResource := resource.Resource{
		Region: region,
	}

	// retreive input variables
	// resource-name: string of resources in terraformer format, ex...s3,ec2
	// filters: string representations of filters, ex..vpc=vpc_id1:vpc_id2:vpc_id3
	flag.StringVar(&targetResource.Name, "resource-name", "None", "Resource Name Flag Set")
	flag.StringVar(&targetResource.Filters, "filters", "None", "Resource Name Flag Set")
	flag.Parse()

	if err := cmd.Setup(&targetResource); err != nil {
		log.Println(err)
	}

}
