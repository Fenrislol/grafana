package sqlstore

import (
	"gitlab.com/digitalizm/grafana/pkg/bus"
	"gitlab.com/digitalizm/grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *models.GetDBHealthQuery) error {
	return x.Ping()
}
