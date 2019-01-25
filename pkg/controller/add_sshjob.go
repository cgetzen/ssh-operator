package controller

import (
	"github.com/cgetzen/ssh-operator/pkg/controller/sshjob"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sshjob.Add)
}
