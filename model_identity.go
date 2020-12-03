package uiamsdk

import "time"

// User User
type User struct {
	ID          uint64      `mapstructure:"id" json:"id,omitempty"`
	UUID        string      `mapstructure:"uuid" json:"uuid"`
	Name        string      `mapstructure:"name" json:"name,omitempty"`
	PhoneCode   string      `mapstructure:"phone_code" json:"phone_code"`
	PhoneNumber string      `mapstructure:"phone_number" json:"phone_number"`
	Email       string      `mapstructure:"email" json:"email"`
	Description string      `mapstructure:"description" json:"description"`
	Attributes  string      `mapstructure:"attributes" json:"attributes"`
	Auths       *AuthObject `mapstructure:"authorizations" json:"authorizations,omitempty"`
	Profile     *Profile    `mapstructure:"profile" json:"profile,omitempty"`
	Status      string      `mapstructure:"status" json:"status"`
	CreatedAt   time.Time   `mapstructure:"created_at" json:"created_at"`
	AvatarUrl   string      `mapstructure:"avatar_url" json:"avatar_url"`
}

// Profile profile
type Profile struct {
	UserID      string    `json:"user_id"`
	Name        string    `json:"name,omitempty"`
	KycLevel    int       `json:"kyc_level,omitempty"`
	KycStatus   string    `json:"kyc_status,omitempty"`
	KycError    string    `json:"kyc_error,omitempty"`
	IDDigest    string    `json:"id_digest,omitempty"`
	CertifiedAt time.Time `json:"certified_at"`
}
