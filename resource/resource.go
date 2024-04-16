package resource

import (
	"fmt"
	//"log"
)

type Resource struct {
	Name    string
	Region  string
	Filters string
	Account string
}

func (r *Resource) ToString() string {
	return fmt.Sprintf("Name: %s Region: %s Filters %s", r.Name, r.Region, r.Filters)
}

func (r *Resource) ToCliString() string {
	return fmt.Sprintf("--resources=%s --filter=%s --regions=%s", r.Name, r.Filters, r.Region)
}

func (r *Resource) ToCliList() ([]string, bool) {
	retString := []string{}

	isIam := false

	/*
		if r.Name == "iam" {
			isIam = true
		}
	*/

	retString = append(retString, "import")
	retString = append(retString, "aws")

	retString = append(retString, fmt.Sprintf("--resources=%s", r.Name))
	if r.Filters == "None" {
		if r.Account != "458305147808" {
			if r.Name != "iam" {
				retString = append(retString, fmt.Sprintf("--regions=%s", r.Region))
				retString = append(retString, fmt.Sprintf("--profile=%s", r.Account))
			} else {
				retString = append(retString, fmt.Sprintf("--profile=%s", r.Account))
				isIam = true
			}
		} else {
			retString = append(retString, fmt.Sprintf("--regions=%s", r.Region))
		}
	} else {
		if r.Account != "458305147808" {
			if !isIam {
				retString = append(retString, fmt.Sprintf("--filter=%s", r.Filters))
				retString = append(retString, fmt.Sprintf("--regions=%s", r.Region))
				retString = append(retString, fmt.Sprintf("--profile=%s", r.Account))
			} else {
				retString = append(retString, fmt.Sprintf("--filter=%s", r.Filters))
				retString = append(retString, fmt.Sprintf("--profile=%s", r.Account))
			}
		} else {
			retString = append(retString, fmt.Sprintf("--filter=%s", r.Filters))
			retString = append(retString, fmt.Sprintf("--regions=%s", r.Region))
		}
	}

	return retString, isIam
}
