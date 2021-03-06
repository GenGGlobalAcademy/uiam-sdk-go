package uiamsdk

import (
	"time"
)

// CreateUserReq CreateUserReq
type CreateUserReq struct {
	Username    string    `json:"username"`
	Type        string    `json:"type"`
	PhoneCode   string    `json:"phone_code"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	AvatarURL   string    `json:"avatar_url"`
	Description string    `json:"description"`
	Attributes  Attribute `json:"attributes"`
}

// UpdateBasicUpdateRequest UpdateRequest
type UpdateBasicUpdateRequest struct {
	UserID      string    `json:"user_id"`
	Username    string    `json:"username"`
	Description string    `json:"description"`
	AvatarURL   string    `json:"avatar_url"`
	Attributes  Attribute `json:"attributes"`
}

// UserFullUpdateRequest UserModifyReq
type UserFullUpdateRequest struct {
	UserID      uint64    `json:"user_id"`
	Username    string    `json:"username"`
	Type        string    `json:"type"`
	PhoneCode   string    `json:"phone_code"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	AvatarURL   string    `json:"avatar_url"`
	Description string    `json:"description"`
	Attributes  Attribute `json:"attributes"`
}

// UserResetPasswordRequest UpdateRequest
type UserResetPasswordRequest struct {
	UserID   string `json:"user_id"`
	RealmID  string `json:"realm_id"`
	Password string `json:"password"`
}

// LoginRequest LoginRequest
type LoginRequest struct {
	ID          uint64 `json:"id"`
	PhoneCode   string `json:"phone_code"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

// KycSyncRequest KycSyncRequest
type KycSyncRequest struct {
	UserID      uint64     `json:"user_id"`
	KycLevel    int        `json:"kyc_level"`
	KycStatus   string     `json:"kyc_status"`
	KycError    string     `json:"kyc_error"`
	CertifiedAt *time.Time `json:"certified_at"`
	IDNo        string     `json:"id_no"`
	Name        string     `json:"name"`
	Country     string     `json:"country"`
	IDType      string     `json:"id_type"`
}

// AuthListRequest 用户表求
type AuthListRequest struct {
	Offset   uint64               `json:"offset" `
	Limit    uint64               `json:"limit"`
	Provider AuthProviderTypeEnum `json:"provider"`
}

// TokenCreateRequest TokenCreateRequest
type TokenCreateRequest struct {
	Type   AuthTypeEnum   `json:"type"`
	Scheme AuthSchemeEnum `json:"scheme"`

	Audience  string `json:"aud,omitempty"` // 这个表示 UIAM_ID
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`

	// Custom
	Provider  string `json:"oap"` // Wechat / Dingding
	UID       string `json:"uid"` // 这个 UID 表示的是第三方系统的 ID，像 ZOTC 里的那个自己的
	SessionID string `json:"sid"`
	Sign      string `json:"sig"`
	SignAlg   string `json:"sal"`

	// ExpriedAt
	Duration time.Duration `json:"duration"` //  有效时长
}

// PhoneCodeVerifyRequest PhoneCodeVerifyRequest
type PhoneCodeVerifyRequest struct {
	PhoneCode   string `json:"phone_code"`
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}

// AuthBindingRequest  auth 绑定
type AuthBindingRequest struct {
	UserID      uint64               `json:"user_id"`
	Provider    AuthProviderTypeEnum `json:"provider"`
	OauthID     string               `json:"oauth_id"`
	UnionID     string               `json:"union_id"`
	Credential  string               `json:"credential"`
	AppUserName string               `json:"app_user_name"`
	AppUserID   string               `json:"app_user_id"`
}
