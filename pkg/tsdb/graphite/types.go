package graphite

import "gitlab.com/digitalizm/grafana/pkg/tsdb"

type TargetResponseDTO struct {
	Target     string                `json:"target"`
	DataPoints tsdb.TimeSeriesPoints `json:"datapoints"`
}
