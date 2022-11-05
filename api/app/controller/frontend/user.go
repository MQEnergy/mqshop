package frontend

import (
	"github.com/MQEnergy/mqshop/app/controller/base"
)

type UserController struct {
	*base.Controller
}

var User = &UserController{}
