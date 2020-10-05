package cloudwatch

import (
	"context"
	"testing"

	"gitlab.com/digitalizm/grafana/pkg/models"
	"gitlab.com/digitalizm/grafana/pkg/tsdb"

	"gitlab.com/digitalizm/grafana/pkg/components/simplejson"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCloudWatch(t *testing.T) {
	Convey("CloudWatch", t, func() {

		Convey("executeQuery", func() {
			e := &CloudWatchExecutor{
				DataSource: &models.DataSource{
					JsonData: simplejson.New(),
				},
			}

			Convey("End time before start time should result in error", func() {
				_, err := e.executeQuery(context.Background(), &CloudWatchQuery{}, &tsdb.TsdbQuery{TimeRange: tsdb.NewTimeRange("now-1h", "now-2h")})
				So(err.Error(), ShouldEqual, "Invalid time range: Start time must be before end time")
			})

			Convey("End time equals start time should result in error", func() {
				_, err := e.executeQuery(context.Background(), &CloudWatchQuery{}, &tsdb.TsdbQuery{TimeRange: tsdb.NewTimeRange("now-1h", "now-1h")})
				So(err.Error(), ShouldEqual, "Invalid time range: Start time must be before end time")
			})
		})
	})
}
