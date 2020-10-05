package middleware

import (
	"net/http"

	"gitlab.com/digitalizm/grafana/pkg/models"
	"gopkg.in/macaron.v1"
)

func MeasureRequestTime() macaron.Handler {
	return func(res http.ResponseWriter, req *http.Request, c *models.ReqContext) {
	}
}
