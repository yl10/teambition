package teambition

import (
	"net/http"

	"github.com/yl10/teambition/model"
	"github.com/yl10/teambition/services"
	"github.com/yl10/teambition/util"
)

// TBClient 快捷客户端
type TBClient interface {
	GetToken(code string) (*model.AccessToken, error)
	ValidToken() error
	RedirectURL() string
}

// DefTBClient 快捷客户端默认实现
type DefTBClient struct {
	req  *http.Request
	cli  *util.HttpClient
	cID  string
	cSec string
	tbCb string
	*services.DefaultOAuth2Service
	*services.DefaultOrgService
	*services.DefaultTasklistService
	*services.DefaultTaskService
	*services.DefaultProjectService
	*services.DefaultTagService
	*services.DefaultStageService
}

// NewTBClient 返回一个默认的快捷客户端
func NewTBClient(req *http.Request, clientID, clientSecret, callBack string) (*DefTBClient, error) {
	if req == nil {
		return nil, util.NewError("没有 http 请求")
	}

	if clientID == "" || clientSecret == "" {
		return nil, util.NewError("没有 AppKey 或者 AppSecret")
	}

	// 新建请求客户端
	defTBCli := DefTBClient{
		req:  req,
		cli:  util.NewHttpClient(),
		cID:  clientID,
		cSec: clientSecret,
		tbCb: callBack,
	}

	if err := initProxys(&defTBCli); err != nil {
		return nil, err
	}

	// 设置 header
	defTBCli.AttemptToSetTokenHead()

	return &defTBCli, nil
}

func initProxys(t *DefTBClient) error {
	t.req.ParseForm()
	code := t.req.Form.Get("code")

	var err error
	t.DefaultOAuth2Service, err = services.NewDefaultOAuth2Service(t.cli, AuthDomain, APIDomain, t.cID, t.cSec, code, t.tbCb)
	if err != nil {
		return err
	}

	t.DefaultOrgService, err = services.NewDefaultOrgService(t.cli, APIDomain)
	if err != nil {
		return err
	}

	t.DefaultTasklistService, err = services.NewDefaultTasklistService(t.cli, APIDomain)
	if err != nil {
		return err
	}

	t.DefaultTaskService, err = services.NewDefaultTaskService(t.cli, APIDomain)
	if err != nil {
		return err
	}

	t.DefaultProjectService, err = services.NewDefaultProjectService(t.cli, APIDomain)
	if err != nil {
		return err
	}

	t.DefaultTagService, err = services.NewDefaultTagService(t.cli, APIDomain)
	if err != nil {
		return err
	}

	t.DefaultStageService, err = services.NewDefaultStageService(t.cli, APIDomain)
	if err != nil {
		return err
	}

	return nil
}
