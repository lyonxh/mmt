package gitlab

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	"mmt/utils/ci"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
)

type Gitlab struct {
	Client *resty.Client

	Address  string
	Username string
	Password string
}

func New(g Gitlab) *Gitlab {
	return &Gitlab{
		Client:   resty.New().SetTimeout(3 * time.Second).SetDisableWarn(true).SetRetryCount(3).SetRetryWaitTime(500 * time.Millisecond), // 忽略证书错误，设置超时时间为 3s
		Address:  g.Address,
		Username: g.Username,
		Password: g.Password,
	}
}

func (g *Gitlab) Healthy() error {
	resp, err := g.Client.R().
		Get(fmt.Sprintf("%s/-/healthy", g.Address))
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New("invide Username or Password")
	}
	return nil
}

// 获取token
func (g *Gitlab) Oauth2Token() (string, error) {
	config := &oauth2.Config{
		Endpoint: oauth2.Endpoint{
			AuthURL:  strings.TrimSuffix(fmt.Sprintf("%s/", g.Address), "/api/v4/") + "oauth/authorize",
			TokenURL: strings.TrimSuffix(fmt.Sprintf("%s/", g.Address), "/api/v4/") + "oauth/token",
		},
	}
	ctx := context.WithValue(context.TODO(), oauth2.HTTPClient, http.Client{Timeout: 3 * time.Second})
	t, err := config.PasswordCredentialsToken(ctx, g.Username, g.Password)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("request token error")
	}
	return t.AccessToken, nil
}

func (g *Gitlab) Auth() error {
	token, err := g.Oauth2Token()
	if err != nil {
		return err
	}
	resp, err := g.Client.R().SetAuthToken(token).
		Get(fmt.Sprintf("%s/api/v4/user", g.Address))
	if err != nil {
		return ci.GitlabRequestError
	}
	if resp.StatusCode() != 200 {
		return ci.GitlabUnauthorized
	}
	return nil
}

// CreateHook 创建webhook
func (g *Gitlab) CreateHook(webHookInjection ci.WebHookInjection) (int, error) {
	result := ci.GitlabCreateHook{}
	token, err := g.Oauth2Token()
	if err != nil {
		return -1, err
	}
	resp, err := g.Client.R().SetAuthToken(token).SetQueryParams(map[string]string{"url": webHookInjection.URL, "push_events": "true", "enable_ssl_verification": "false"}).
		SetResult(&result).
		Post(fmt.Sprintf("%s/api/v4/projects/%s%%2F%s/hooks", g.Address, webHookInjection.Owner, webHookInjection.Repo))
	if err != nil {
		return -1, ci.GitlabRequestError
	}
	if resp.StatusCode() != 201 {
		return -1, errors.New("注入失败")
	}
	return result.ID, nil
}

// DeleteHook 删除webhook
func (g *Gitlab) DeleteHook(webHookInjection ci.WebHookInjection) error {
	token, err := g.Oauth2Token()
	if err != nil {
		return err
	}
	resp, err := g.Client.R().SetAuthToken(token).
		Delete(fmt.Sprintf("%s/api/v4/projects/%s%%2F%s/hooks/%d", g.Address, webHookInjection.Owner, webHookInjection.Repo, webHookInjection.Id))
	if err != nil {
		return ci.GitlabRequestError
	}
	if resp.StatusCode() != 204 {
		return errors.New("删除失败")
	}
	return nil
}

func (g *Gitlab) CheckRepo(webhook ci.WebHookInjection) error {
	resp, err := g.Client.R().SetBasicAuth(g.Username, g.Password).
		Get(fmt.Sprintf("%s/api/v4/projects/%s%%2F%s", g.Address, webhook.Owner, webhook.Repo))
	if err != nil {
		return ci.GiteaRequestError
	}
	if resp.StatusCode() != 200 {
		return errors.New("校验仓库失败,确认仓库地址、用户、密码正确")
	}
	return nil
}

func (g *Gitlab) CheckHook(webhook ci.WebHookInjection) error {
	token, err := g.Oauth2Token()
	if err != nil {
		return err
	}
	resp, err := g.Client.R().SetAuthToken(token).
		Get(fmt.Sprintf("%s/api/v4/projects/%s%%2F%s/hooks/%d", g.Address, webhook.Owner, webhook.Repo, webhook.Id))
	if err != nil {
		return ci.GitlabRequestError
	}

	switch resp.StatusCode() {
	case 404:
		return ci.GitlabHookNotFound
	case 200:
		return nil
	}
	return nil
}
