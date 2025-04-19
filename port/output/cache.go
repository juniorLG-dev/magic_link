package output

import (
	"magic_link/adapter/output/model"
)

type PortCache interface {
	Set(model.UserCode) error
	Get(string) (string, error)
}