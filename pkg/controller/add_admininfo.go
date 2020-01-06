package controller

import (
	"github.com/wanghh2000/operator_sdk_pro1/pkg/controller/admininfo"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, admininfo.Add)
}
