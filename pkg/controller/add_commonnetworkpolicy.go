// MIT LICENSE

package controller

import (
	"github.com/bells17/common-network-policy-operator/pkg/controller/commonnetworkpolicy"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, commonnetworkpolicy.Add)
}
