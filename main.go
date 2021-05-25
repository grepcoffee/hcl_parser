package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// type ExceptionDetail struct {
// 	Policy           string `json:"policy"`
// 	EnforcementLevel string `json:"enforcement_level"`
// }

// type Exception struct {
// 	Id               string `json:"_id"`
// 	IsActive         bool   `json:"is_Active"`
// 	Environment      string `json:"enviroment"`
// 	Organization     string `json:"organization"`
// 	PolsetName       string `json:"polset_name"`
// 	Workspace        string `json:"workspace"`
// 	RiskApproval     string `json:"risk_approval"`
// 	PocEmail         string `json:"poc_email"`
// 	Description      string `json:"description"`
// 	Created          string `json:"created"`
// 	Expires          string `json:"expires"`
// 	ExceptionDetails struct {
// 		Policy           string `json:"policy"`
// 		EnforcementLevel string `json:"enforcement_level"`
// 	} `json:"exception_details"`
// }

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

func LoadExceptionsFile(filename string) (Exception, error) {
	expFile, _ := ioutil.ReadFile(filename)
	var exceptions Exception
	err := json.Unmarshal(expFile, &exceptions)
	return exceptions, err
}

func main() {
	fmt.Println("Start of Script")
	exceptions, _ := LoadExceptionsFile("exceptions.json")
	for _, x := range exceptions {
		for _, y := range x.ExceptionDetails {
			fmt.Println(y.Policy)
		}
	}

	//fmt.Println(exceptions.)

}

// file, log_err := os.OpenFile("exceptions.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
// if log_err != nil {
// 	log.Fatal(log_err)
// }
// log.SetOutput(file)
// log.Print("Start of Script")

// if err != nil {
// 	log.Print(err.Error())
// }
// log.Print("Opened Excp json file")

// err2 := json.Unmarshal(expFile, &exceptions)
// return exceptions, err2}

// // Start of Script
// file, log_err := os.OpenFile("exceptions.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
// if log_err != nil {
// 	log.Fatal(log_err)
// }
// defer file.Close()
// log.SetOutput(file)
// log.Print("Start of Script")

// expFile, err := ioutil.ReadFile("exceptions.json")
// if err != nil {
// 	log.Print(err.Error())
// }
// log.Print("Opened Excp json file")

// var exceptions []Exception

// err2 := json.Unmarshal(expFile, &exceptions)
// if err2 != nil {
// 	log.Print("Error unmarshalling file")
// 	log.Print(err2.Error())
// }
// for _, x := range exceptions {
// 	fmt.Println(x.Environment)
// 	fmt.Println(x[exceptiondetails].Policy)
// }

// // logger config
// file, log_err := os.OpenFile("exceptions.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
// if log_err != nil {
// 	log.Fatal(log_err)
// }
// defer file.Close()
// log.SetOutput(file)
// log.Print("Starting Exceptions Parser Job")

// loader := exceptions.ExceptionsLoader()
// expFile, err := loader.readFile("exceptions.json")
// if err != nil {
// 	fmt.Println(err.Error())
// 	//panic()
// }
// log.Print("Successfully Opened exceptions.json")

// var exceptions []Exception
// err2 := json.Unmarshal(expFile, &exceptions)
// if err2 != nil {
// 	log.Print("Error unmarshalling exceptions json")
// 	log.Print(err2.Error())
// }
// for _, x := range exceptions {
// 	fmt.Println(x.Environment)

// 	for _, v := range x.ExceptionDetails {
// 		fmt.Println(v.Policy)
// 	}
// }

// Resoures used
//https://www.youtube.com/watch?v=y_eIBmt3JdY&t=171s

// expFile, err := os.Open(filename)
// defer expFile.Close()
// fmt.Println("Loading")
// if err != nil {
// 	fmt.Println("Could not open file")
// 	return exceptions, err
// }
// expParser := json.NewDecoder(expFile)
// err = expParser.Decode(&exceptions)
// fmt.Println("JsonDecode")
// return exceptions, err
