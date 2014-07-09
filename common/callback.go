package common

import (
	"net/http"
)

type Callback interface {
	Parse(*http.Request) error
}
