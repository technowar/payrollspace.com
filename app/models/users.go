package models

import (
	"regexp"
	"time"
)

var emailRegex *regexp.Regexp

func init() {
	emailRegex, _ = regexp.Compile(`^[^@]+@[^@]+$`)
}

// User ....
type User struct {
	ID                int64      `gorm:"AUTO_INCREMENT" json:"id"`
	UUID              string     `gorm:"type:varchar(32);unique_index" json:"uuid"`
	CreatedAt         time.Time  ` json:"created_at,omitempty"`
	UpdatedAt         time.Time  ` json:"updated_at,omitempty"`
	DeletedAt         *time.Time ` json:"deleted_at,omitempty"`
	Email             string     `gorm:"type:varchar(130);unique_index" json:"email,omitempty"`
	Password          string     ` json:"password,omitempty"`
	VerificationToken string     ` json:"verification_token,omitempty"`
	//ReferralCode      string     ` json:"referral_code,omitempty"`
	ReferralUserUUID string ` json:"referral_user_uuid,omitempty"`
}
