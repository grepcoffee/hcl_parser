package exceptions

import (
	"io/ioutil"
)

type IExceptionsLoader interface {
	read(file:string) (string, error)
}
type ExceptionsLoader struct {}

func (e ExceptionsLoader)read(file:string) (string, error) {
	return ioutil.ReadFile(file)
}