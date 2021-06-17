package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

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

func main() {
	strings := []string{"sentinel.hcl", "sentinel1.hcl"}
	for _, filepath := range strings {
		diagWr := hcl.NewDiagnosticTextWriter(os.Stderr, nil, 78, false)

		inFile := filepath //"sentinel.hcl"
		src, err := ioutil.ReadFile(inFile)
		if err != nil {
			log.Fatalf("failed to read %s: %s", inFile, err)
		}
		f, diags := hclwrite.ParseConfig(src, inFile, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			diagWr.WriteDiagnostics(diags)
			os.Exit(1)
		}

		for _, block := range f.Body().Blocks() {
			if block.Type() != "policy" {
				continue
			}
			labels := block.Labels()
			if len(labels) != 1 {
				diags = diags.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "Invalid policy block",
					Detail:   "A policy block must have only one label, giving the policy name.",
				})
				continue
			}
			policyName := labels[0]

			elAttr := block.Body().GetAttribute("enforcement_level")
			if elAttr == nil {
				continue
			}

			toks := elAttr.Expr().BuildTokens(nil)
			// We're looking for specifically a string literal "advisory",
			// which will appear as three tokens: OQuote, QuotedLit, CQuote.
			if len(toks) != 3 || toks[0].Type != hclsyntax.TokenOQuote || toks[1].Type != hclsyntax.TokenQuotedLit || toks[2].Type != hclsyntax.TokenCQuote {
				diags = diags.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "Unrecognized enforcement_level expression",
					Detail:   fmt.Sprintf("Can't process enforcement_level for policy %q: this tool only recognizes expressions that are literal strings.", policyName),
				})
				continue
			}

			el := string(toks[1].Bytes)
			//fmt.Printf("testing %q", el)
			fmt.Println(el)
			exceptions, _ := LoadExceptionsFile("exceptions.json")
			for _, x := range exceptions {
				active := x.Isactive
				if active == true {
					for _, y := range x.ExceptionDetails {
						exp_policy_name := y.Policy
						if exp_policy_name == policyName {
							switch el {
							case "hard-mandatory":
								fmt.Println("Case hard mandatory")
								if exp_policy_name == policyName {
									newEL := y.EnforcementLevel //"soft-mandatory"
									log.Printf("rewriting policy %q enforcement level to %q", policyName, newEL)
									block.Body().SetAttributeValue("enforcement_level", cty.StringVal(newEL))
								}
							case "soft-mandatory":
								fmt.Println("Case soft mandatory")
								if exp_policy_name == policyName {
									newEL := y.EnforcementLevel //"soft-mandatory"
									log.Printf("rewriting policy %q enforcement level to %q", policyName, newEL)
									block.Body().SetAttributeValue("enforcement_level", cty.StringVal(newEL))
								}
							case "advisory":
								fmt.Println("Case advisory")
								if exp_policy_name == policyName {
									newEL := y.EnforcementLevel //"soft-mandatory"
									log.Printf("rewriting policy %q enforcement level to %q", policyName, newEL)
									block.Body().SetAttributeValue("enforcement_level", cty.StringVal(newEL))
								}
							}
						}
					}
				}
			}
			diagWr.WriteDiagnostics(diags)
			if diags.HasErrors() {
				os.Exit(1)
			}
		}

		if err := ioutil.WriteFile(filepath, f.Bytes(), 0644); err != nil {
			log.Fatal(err)
		}
	}
}
