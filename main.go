package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	log "github.com/sirupsen/logrus"
	"github.com/zclconf/go-cty/cty"
)

type Exception []struct { //Exceptions JSON Struct
	ID               string `json:"_id"`           //Randomly Generated ID number for tracking purposes.
	IsActive         bool   `json:"isActive"`      //Is Exception Active or Inactive
	Enviroment       string `json:"enviroment"`    //What Enviroment prod or dev.
	Organization     string `json:"organization"`  //What Organization does this exception apply to.
	PolsetName       string `json:"Polset-Name"`   //Policyset Name
	Workspace        string `json:"workspace"`     //Sentinel Workspace Name
	RiskApproval     string `json:"risk_approval"` //Risk approval link
	PocEmail         string `json:"poc_email"`     //Point of contact email
	Description      string `json:"description"`   //Description of Risk approval and policies being exempted.
	Created          string `json:"created"`       //Exception Created date. (For Documentation Purposes)
	Expires          string `json:"expires"`       //Exception Expiry date. (For Documentation Purposes)
	ExceptionDetails []struct {
		Policy           string `json:"policy"`            //Exception Policy Name
		EnforcementLevel string `json:"enforcement_level"` //Exception new enforcement level
	} `json:"exception_details"`
}

type HclPolicy struct {
	Policy []*struct {
		Source           string `hcl:"source"`            //Sentinel Policy location
		EnforcementLevel string `hcl:"enforcement_level"` //Sentinel Policy Enforcement Level.
	} `hcl:"policy,block"`
}

func LoadExceptionsFile(filename string) (Exception, error) {
	expFile, _ := ioutil.ReadFile(filename)
	var exceptions Exception
	err := json.Unmarshal(expFile, &exceptions)
	return exceptions, err
}

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Info("Starting Sentinel Exceptions Parser")

	strings := []string{"sentinel.hcl", "sentinel1.hcl"}
	for _, filePath := range strings {
		diagWr := hcl.NewDiagnosticTextWriter(os.Stderr, nil, 78, false)
		src, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal("failed to ", err)
		}
		f, diags := hclwrite.ParseConfig(src, filePath, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			diagWr.WriteDiagnostics(diags)
			log.Fatal("Error in Hcl Parse Config")
			os.Exit(1)
		}

		for _, block := range f.Body().Blocks() {
			log.Info("Examining Policy Structure")
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
			//We're looking for specifically a string literal "advisory",
			//which will appear as three tokens: OQuote, QuotedLit, CQuote.
			if len(toks) != 3 || toks[0].Type != hclsyntax.TokenOQuote || toks[1].Type != hclsyntax.TokenQuotedLit || toks[2].Type != hclsyntax.TokenCQuote {
				diags = diags.Append(&hcl.Diagnostic{
					Severity: hcl.DiagWarning,
					Summary:  "Unrecognized enforcement_level expression",
					Detail:   fmt.Sprintf("Can't process enforcement_level for policy %q: this tool only recognizes expressions that are literal strings.", policyName),
				})
				continue
			}

			el := string(toks[1].Bytes)
			exceptions, _ := LoadExceptionsFile("exceptions.json")
			for _, exceptionList := range exceptions {
				if exceptionList.IsActive == true {
					for _, exceptionDetailsList := range exceptionList.ExceptionDetails {
						exp_policy_name := exceptionDetailsList.Policy
						if exp_policy_name == policyName {
							switch el {
							case "hard-mandatory":
								fmt.Println("Case hard mandatory")
								if exp_policy_name == policyName {
									newEL := exceptionDetailsList.EnforcementLevel
									log.Info("rewriting policy ", policyName, " enforcement level to ", newEL)
									block.Body().SetAttributeValue("enforcement_level", cty.StringVal(newEL))
								}
							case "soft-mandatory":
								fmt.Println("Case soft mandatory")
								if exp_policy_name == policyName {
									newEL := exceptionDetailsList.EnforcementLevel
									log.Info("rewriting policy ", policyName, " enforcement level to ", newEL)
									block.Body().SetAttributeValue("enforcement_level", cty.StringVal(newEL))
								}
							case "advisory":
								fmt.Println("Case advisory")
								if exp_policy_name == policyName {
									newEL := exceptionDetailsList.EnforcementLevel
									log.Info("rewriting policy ", policyName, " enforcement level to ", newEL)
									block.Body().SetAttributeValue("enforcement_level", cty.StringVal(newEL))
								}
							}
						}
					}
				}
			}
			diagWr.WriteDiagnostics(diags)
			if diags.HasErrors() {
				log.Fatal("Failed to Write Policies")
			}
		}

		if err := ioutil.WriteFile(filePath, f.Bytes(), 0644); err != nil {
			log.Error("Error Writing to File")
			log.Fatal(err)
		}
		log.Info("End of Exceptions Parsing")
	}
}
