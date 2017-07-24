package services

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/yl10/teambition/api"
	"github.com/yl10/teambition/model"
	"github.com/yl10/teambition/util"
)

var _ api.TBOAuth2 = &DefaultOAuth2Service{}

type AccessTokenProfile struct {
	CliendID     string
	ClientSecret string
	Code         string
	Token        string
	CreateAt     time.Time
	Expired      time.Time
}

// DefaultOAuth2Service 默认 TBOAuth2 接口实现
type DefaultOAuth2Service struct {
	cli      *util.HttpClient
	authURL  string
	apiURL   string
	cID      string
	cSec     string
	code     string
	callBack string
}

// NewDefaultOAuth2Service 创建实例
func NewDefaultOAuth2Service(cli *util.HttpClient, authURL, apiURL, cID, cSec, code, callBack string) (*DefaultOAuth2Service, error) {
	if cli == nil {
		return nil, util.NewError("没有 Http 客户端")
	}

	if authURL == "" || apiURL == "" {
		return nil, util.NewError("没有 验证域 或者 接口域")
	}

	if cID == "" || cSec == "" {
		return nil, util.NewError("没有 AppKey 或者 AppSecret")
	}

	defCli := &DefaultOAuth2Service{
		cli:      cli,
		authURL:  authURL,
		apiURL:   apiURL,
		cID:      cID,
		cSec:     cSec,
		code:     code,
		callBack: callBack,
	}

	return defCli, nil
}

// SetCode 设置 code
func (t *DefaultOAuth2Service) SetCode(code string) {
	t.code = code

	// 更新 context
	if tp, ok := getTokenProfile(t.cID); ok {
		tp.Code = code
		setTokenProfile(tp)
	}
}

// GetToken 获取 token
func (t *DefaultOAuth2Service) GetToken() (*model.AccessToken, error) {
	// 如果已经存在的
	if tp, ok := getTokenProfile(t.cID); ok {
		if tp != nil &&
			tp.Token != "" &&
			tp.Expired.After(time.Now()) {
			return t.profile2Token(tp), nil
		}
	}

	return t.RefreshToken()
}

// AttemptToSetTokenHead 尝试设置 token header
func (t *DefaultOAuth2Service) AttemptToSetTokenHead() {
	if tp, ok := getTokenProfile(t.cID); ok {
		if tp != nil &&
			tp.Token != "" &&
			tp.Expired.After(time.Now()) {
			t.cli.SetToken(tp.Token)
		}
	}
}

// RefreshToken 刷新
func (t *DefaultOAuth2Service) RefreshToken() (*model.AccessToken, error) {
	if t.code == "" {
		return nil, util.NewError("获取 Token 失败(code-404), 没有 code ")
	}

	grantType := "code"

	params := struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
		GrantType    string `json:"grant_type"`
	}{
		t.cID,
		t.cSec,
		t.code,
		grantType,
	}

	reqURL := t.authURL + "oauth2/access_token"

	dt, err := t.cli.Post(reqURL, params)
	if err != nil {
		return nil, util.TransErr2XErr(err)
	}

	var token model.AccessToken
	err = json.Unmarshal(dt, &token)
	if err != nil {
		return nil, util.TransErr2XErr(err)
	}
	// 保存 token 到 context
	tokenProfile := t.token2Profile(&token)
	setTokenProfile(tokenProfile)

	// token 添加到 header
	t.cli.SetToken(token.AccessToken)

	return &token, nil
}

// ValidToken 验证Token
func (t *DefaultOAuth2Service) ValidToken() error {
	tokenPro, ok := getTokenProfile(t.cID)
	if !ok {
		return util.NewError("获取 Context Token 失败")
	}

	reqURL := t.apiURL + "api/applications/" + t.cID + "/tokens/check"

	// OAuth2 后面有一个空格
	t.cli.SetToken(tokenPro.Token)
	res, err := t.cli.GetReq(reqURL, nil)
	if err != nil {
		return util.TransErr2XErr(err)
	}

	if res.StatusCode != http.StatusOK {
		return util.NewError("校验 Token 失败, 错误码 " + strconv.Itoa(res.StatusCode))
	}

	return nil
}

// RedirectURL 生成转向链接
// 注意 redirectURI 为非转码过的原始数据
func (t *DefaultOAuth2Service) RedirectURL(redirectURI string) string {
	if redirectURI == "" {
		return ""
	}

	query := url.Values{}
	query.Add("client_id", t.cID)
	query.Add("response_type", "code")
	query.Add("redirect_uri", t.callBack)

	// 注意这里是将跳转地址, 放置到 state 参数里面的
	query.Add("state", redirectURI)

	return t.authURL + "oauth2/authorize?" + query.Encode()
}

func (t *DefaultOAuth2Service) token2Profile(token *model.AccessToken) *AccessTokenProfile {
	return &AccessTokenProfile{
		CliendID:     t.cID,
		ClientSecret: t.cSec,
		Code:         t.code,
		Token:        token.AccessToken,
		CreateAt:     time.Now(),
		Expired:      time.Now().Add(365 * 24 * time.Hour),
	}
}

func (t *DefaultOAuth2Service) profile2Token(tp *AccessTokenProfile) *model.AccessToken {
	return &model.AccessToken{
		AccessToken: tp.Token,
	}
}

func setTokenProfile(tp *AccessTokenProfile) {
	ctx = context.WithValue(ctx, tp.CliendID, tp)
}

func getTokenProfile(clientID string) (*AccessTokenProfile, bool) {
	tp, ok := ctx.Value(clientID).(*AccessTokenProfile)
	return tp, ok
}

var ctx context.Context

func init() {
	ctx, _ = context.WithCancel(context.Background())
}
