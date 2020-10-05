package dtos

import "gitlab.com/digitalizm/grafana/pkg/models"

type UpdateDashboardAclCommand struct {
	Items []DashboardAclUpdateItem `json:"items"`
}

type DashboardAclUpdateItem struct {
	UserId     int64                 `json:"userId"`
	TeamId     int64                 `json:"teamId"`
	Role       *models.RoleType      `json:"role,omitempty"`
	Permission models.PermissionType `json:"permission"`
}
