package wecom

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

var (
	validate *validator.Validate
	mutex    sync.RWMutex
)

type H map[string]interface{}

func init() {
	validate = validator.New()
}
