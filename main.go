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
		fmt.Println(blockLabels) // Policy Names
		// blockAttr := block.Body().Attributes()
		// fmt.Println(blockAttr)
		// Rename references of variable "a" to "z"
	}
}

func replaceHCLFile(filePath string) {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	f, diags := hclwrite.ParseConfig(contents, "", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		fmt.Printf("errors: %s", diags)
		return
	}

	for _, attr := range f.Body().Attributes() {
		attr.Expr().RenameVariablePrefix(
			[]string{"advisory"},
			[]string{"12345678"},
		)
	}
	fmt.Printf("%s", f.Bytes())
}

func main() {
	// fmt.Println("Start of Script")
	// exceptions, _ := LoadExceptionsFile ("exceptions.json")
	// for _, x := range exceptions {
	// 	for _, y := range x.ExceptionDetails {
	// 		fmt.Println(y.Policy)
	// 		fmt.Println(y)
	// 	}
	// }
	readHCLFile("sentinel.hcl")
	replaceHCLFile("sentinel.hcl")
}

// https://pkg.go.dev/github.com/hashicorp/hcl/v2/hclwrite

// f := hclwrite.NewEmptyFile()
// rootBody := f.Body()
// rootBody.SetAttributeValue("string", cty.StringVal("bar")) // this is overwritten later
// rootBody.AppendNewline()
// rootBody.SetAttributeValue("object", cty.ObjectVal(map[string]cty.Value{
// 	"foo": cty.StringVal("foo"),
// 	"bar": cty.NumberIntVal(5),
// 	"baz": cty.True,
// }))
// rootBody.SetAttributeValue("string", cty.StringVal("foo"))
// rootBody.SetAttributeValue("bool", cty.False)
// rootBody.SetAttributeTraversal("path", hcl.Traversal{
// 	hcl.TraverseRoot{
// 		Name: "env",
// 	},
// 	hcl.TraverseAttr{
// 		Name: "PATH",
// 	},
// })
// rootBody.AppendNewline()
// fooBlock := rootBody.AppendNewBlock("foo", nil)
// fooBody := fooBlock.Body()
// rootBody.AppendNewBlock("empty", nil)
// rootBody.AppendNewline()
// barBlock := rootBody.AppendNewBlock("bar", []string{"a", "b"})
// barBody := barBlock.Body()

// fooBody.SetAttributeValue("hello", cty.StringVal("world"))

// bazBlock := barBody.AppendNewBlock("baz", nil)
// bazBody := bazBlock.Body()
// bazBody.SetAttributeValue("foo", cty.NumberIntVal(10))
// bazBody.SetAttributeValue("beep", cty.StringVal("boop"))
// bazBody.SetAttributeValue("baz", cty.ListValEmpty(cty.String))

// fmt.Printf("%s", f.Bytes())
