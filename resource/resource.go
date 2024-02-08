package resource

import (
	"fmt"
)

type Resource struct {
	Name    string
	Region  string
	Filters string
}

func (r *Resource) ToString() string {
	return fmt.Sprintf("Name: %s Region: %s Filters %s", r.Name, r.Region, r.Filters)
}

func (r *Resource) ToCliString() string {
	return fmt.Sprintf("--resources=%s --filter=%s --regions=%s", r.Name, r.Filters, r.Region)
}

func (r *Resource) ToCliList() []string {
	retString := []string{}

	retString = append(retString, "import")
	retString = append(retString, "aws")

	retString = append(retString, fmt.Sprintf("--resources=%s", r.Name))
	if r.Filters == "None" {
		retString = append(retString, fmt.Sprintf("--regions=%s", r.Region))
	} else {
		retString = append(retString, fmt.Sprintf("--filter=%s", r.Filters))
		retString = append(retString, fmt.Sprintf("--regions=%s", r.Region))
	}

	return retString
}
