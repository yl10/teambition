package api

import "github.com/yl10/teambition/model"

// TBOAuth2 认证接口
type TBOAuth2 interface {
	// SetCode 设置 code
	SetCode(code string)

	// GetToken 获取 token
	GetToken() (*model.AccessToken, error)

	// RefreshToken
	RefreshToken() (*model.AccessToken, error)

	// ValidToken 验证Token
	ValidToken() error

	// RedirectURL 生成转向链接
	RedirectURL(redirectURI string) string
}
