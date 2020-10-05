package utils

import (
	"errors"
	"fmt"

	"gitlab.com/digitalizm/grafana/pkg/bus"
	"gitlab.com/digitalizm/grafana/pkg/models"
)

func CheckOrgExists(orgID int64) error {
	query := models.GetOrgByIdQuery{Id: orgID}
	if err := bus.Dispatch(&query); err != nil {
		if errors.Is(err, models.ErrOrgNotFound) {
			return err
		}
		return fmt.Errorf("failed to check whether org. with the given ID exists: %w", err)
	}
	return nil
}