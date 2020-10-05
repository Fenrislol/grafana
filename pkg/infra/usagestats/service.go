package usagestats

import (
	"context"
	"time"

	"gitlab.com/digitalizm/grafana/pkg/bus"
	"gitlab.com/digitalizm/grafana/pkg/login/social"
	"gitlab.com/digitalizm/grafana/pkg/models"
	"gitlab.com/digitalizm/grafana/pkg/services/alerting"
	"gitlab.com/digitalizm/grafana/pkg/services/sqlstore"

	"gitlab.com/digitalizm/grafana/pkg/infra/log"
	"gitlab.com/digitalizm/grafana/pkg/registry"
	"gitlab.com/digitalizm/grafana/pkg/setting"
)

var metricsLogger log.Logger = log.New("metrics")

func init() {
	registry.RegisterService(&UsageStatsService{})
}

type UsageStatsService struct {
	Cfg                *setting.Cfg               `inject:""`
	Bus                bus.Bus                    `inject:""`
	SQLStore           *sqlstore.SqlStore         `inject:""`
	AlertingUsageStats alerting.UsageStatsQuerier `inject:""`
	License            models.Licensing           `inject:""`

	log log.Logger

	oauthProviders map[string]bool
}

func (uss *UsageStatsService) Init() error {
	uss.log = log.New("infra.usagestats")
	uss.oauthProviders = social.GetOAuthProviders(uss.Cfg)
	return nil
}

func (uss *UsageStatsService) Run(ctx context.Context) error {
	uss.updateTotalStats()

	onceEveryDayTick := time.NewTicker(time.Hour * 24)
	everyMinuteTicker := time.NewTicker(time.Minute)
	defer onceEveryDayTick.Stop()
	defer everyMinuteTicker.Stop()

	for {
		select {
		case <-onceEveryDayTick.C:
			uss.sendUsageStats()
		case <-everyMinuteTicker.C:
			uss.updateTotalStats()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
