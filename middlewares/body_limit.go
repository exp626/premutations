package middlewares

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

var limitSize int64 = 16 * 1024000

func BodyLimit(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, limitSize)
		err := next(c)
		return err
	}
}
