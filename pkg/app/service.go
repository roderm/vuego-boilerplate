package app

import (
	"context"
)

//Service interface
type Service interface {
	Run(context.Context) error
}
