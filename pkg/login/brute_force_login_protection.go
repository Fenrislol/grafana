package login

import (
	"time"

	"gitlab.com/digitalizm/grafana/pkg/bus"
	"gitlab.com/digitalizm/grafana/pkg/models"
	"gitlab.com/digitalizm/grafana/pkg/setting"
)

var (
	maxInvalidLoginAttempts int64 = 5
	loginAttemptsWindow           = time.Minute * 5
)

var validateLoginAttempts = func(username string) error {
	if setting.DisableBruteForceLoginProtection {
		return nil
	}

	loginAttemptCountQuery := models.GetUserLoginAttemptCountQuery{
		Username: username,
		Since:    time.Now().Add(-loginAttemptsWindow),
	}

	if err := bus.Dispatch(&loginAttemptCountQuery); err != nil {
		return err
	}

	if loginAttemptCountQuery.Result >= maxInvalidLoginAttempts {
		return ErrTooManyLoginAttempts
	}

	return nil
}

var saveInvalidLoginAttempt = func(query *models.LoginUserQuery) error {
	if setting.DisableBruteForceLoginProtection {
		return nil
	}

	loginAttemptCommand := models.CreateLoginAttemptCommand{
		Username:  query.Username,
		IpAddress: query.IpAddress,
	}

	return bus.Dispatch(&loginAttemptCommand)
}
