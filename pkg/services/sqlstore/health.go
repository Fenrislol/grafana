package sqlstore

import (
	"github.com/Fenrislol/grafana/pkg/bus"
	"github.com/Fenrislol/grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *models.GetDBHealthQuery) error {
	return x.Ping()
}
