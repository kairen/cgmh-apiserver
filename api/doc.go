package swagger

import (
	"github.com/swaggo/swag"
)

type s struct{}

func (s *s) ReadDoc() string {
	d, _ := Asset("api/swagger-spec/v1.yml")
	return string(d)
}

func init() {
	swag.Register(swag.Name, &s{})
}
