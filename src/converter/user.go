package converter

import (
	"github.com/sky0621/cv-admin/src/ent"
	"github.com/sky0621/cv-admin/src/rest"
)

// goverter:converter
type UserConverter interface {
	ConvertUserAttributeFromRESTToEnt(src rest.UserAttribute) *ent.UserCreate
}
