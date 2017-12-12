package models

import "time"

// RoleAccess ....
type RoleAccess struct {
	UUID      int64      `gorm:"type:varchar(32)" json:"uuid"`
	URL       string     `json:"url,omitempty"`
	Status    string     `json:"status,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
