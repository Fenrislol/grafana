package remotecache

import (
	"testing"

	"gitlab.com/digitalizm/grafana/pkg/services/sqlstore"
	"gitlab.com/digitalizm/grafana/pkg/setting"
)

// NewFakeStore creates store for testing
func NewFakeStore(t *testing.T) *RemoteCache {
	t.Helper()

	opts := &setting.RemoteCacheOptions{
		Name:    "database",
		ConnStr: "",
	}

	SQLStore := sqlstore.InitTestDB(t)

	dc := &RemoteCache{
		SQLStore: SQLStore,
		Cfg: &setting.Cfg{
			RemoteCacheOptions: opts,
		},
	}

	err := dc.Init()
	if err != nil {
		t.Fatalf("failed to init remote cache for test. error: %v", err)
	}

	return dc
}
