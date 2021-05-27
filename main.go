package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// https://mholt.github.io/json-to-go/
type Exception []struct {
	ID               string `json:"_id"`
	Isactive         bool   `json:"isActive"`
	Enviroment       string `json:"enviroment"`
	Organization     string `json:"organization"`
	PolsetName       string `json:"Polset-Name"`
	Workspace        string `json:"workspace"`
	RiskApproval     string `json:"risk_approval"`
	PocEmail         string `json:"poc_email"`
	Description      string `json:"description"`
	Created          string `json:"created"`
	Expires          string `json:"expires"`
	ExceptionDetails []struct {
		Policy           string `json:"policy"`
		EnforcementLevel string `json:"enforcement_level"`
	} `json:"exception_details"`
}

type HclPolicy struct {
	Policy []*struct {
		Source           string `hcl:"source"`
		EnforcementLevel string `hcl:"enforcement_level"`
	} `hcl:"policy,block"`
}

func LoadExceptionsFile(filename string) (Exception, error) {
	expFile, _ := ioutil.ReadFile(filename)
	var exceptions Exception
	err := json.Unmarshal(expFile, &exceptions)
	return exceptions, err
}

func readHCLFile(filePath string) {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	f, diags := hclwrite.ParseConfig(contents, "", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		fmt.Printf("errors: %s", diags)
		return
	}

	//Rename references of variable "a" to "z"
	for _, block := range f.Body().Blocks() {
		blockLabels := block.Labels()
		fmt.Println(blockLabels)
		blockAttr := block.Body().Attributes()
		fmt.Println(blockAttr)
		for _, attr := range f.Body().Attributes() {
			attr.Expr().RenameVariablePrefix(
				[]string{"a"},
				[]string{"z"},
			)
	}

	fmt.Printf("%s", f.Bytes())

	//fmt.Printf("File contents: %s", content)

	// file, diags := hclsyntax.ParseConfig(content, filePath, hcl.Pos{Line: 1, Column: 1})
	// if diags.HasErrors() {
	// 	fmt.Println("ParseConfig: %w", diags)
	// }

	// c := &HclPolicy{}
	// diags = gohcl.DecodeBody(file.Body, nil, c)
	// if diags.HasErrors() {
	// 	fmt.Println("DecodeBody: %w", diags)
	// }

	//fmt.Println(c)
}

func main() {
	// fmt.Println("Start of Script")
	// exceptions, _ := LoadExceptionsFile("exceptions.json")
	// for _, x := range exceptions {
	// 	for _, y := range x.ExceptionDetails {
	// 		fmt.Println(y.Policy)
	// 		fmt.Println(y)
	// 	}
	// }
	readHCLFile("sentinel.hcl")

}

