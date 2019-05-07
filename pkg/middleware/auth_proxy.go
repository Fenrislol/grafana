package middleware

import (
	"github.com/maksimmernikov/grafana/pkg/infra/remotecache"
	authproxy "github.com/maksimmernikov/grafana/pkg/middleware/auth_proxy"
	m "github.com/maksimmernikov/grafana/pkg/models"
)

const (

	// cachePrefix is a prefix for the cache key
	cachePrefix = authproxy.CachePrefix
)

func initContextWithAuthProxy(store *remotecache.RemoteCache, ctx *m.ReqContext, orgID int64) bool {
	auth := authproxy.New(&authproxy.Options{
		Store: store,
		Ctx:   ctx,
		OrgID: orgID,
	})

	// Bail if auth proxy is not enabled
	if auth.IsEnabled() == false {
		return false
	}

	// If the there is no header - we can't move forward
	if auth.HasHeader() == false {
		return false
	}

	// Check if allowed to continue with this IP
	if result, err := auth.IsAllowedIP(); result == false {
		ctx.Handle(407, err.Error(), err.DetailsError)
		return true
	}

	// Try to get user id from various sources
	id, err := auth.GetUserID()
	if err != nil {
		ctx.Handle(500, err.Error(), err.DetailsError)
		return true
	}

	// Get full user info
	user, err := auth.GetSignedUser(id)
	if err != nil {
		ctx.Handle(500, err.Error(), err.DetailsError)
		return true
	}

	// Add user info to context
	ctx.SignedInUser = user
	ctx.IsSignedIn = true

	// Remember user data it in cache
	if err := auth.Remember(); err != nil {
		ctx.Handle(500, err.Error(), err.DetailsError)
		return true
	}

	return true
}
