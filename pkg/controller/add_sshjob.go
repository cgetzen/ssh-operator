package controller

import (
	"ssh-operator/pkg/controller/sshjob"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sshjob.Add)
}
