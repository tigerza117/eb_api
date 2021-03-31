package user

import (
	"github.com/tigerza117/eb_api/models"
)

type controller struct {}

func New() models.UserController {
	return controller{}
}
