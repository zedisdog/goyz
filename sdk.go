package goyz

import (
	"encoding/json"
	"github.com/zedisdog/goyz/responses"
)

func NewSdk(token string) *sdk {
	client := NewClient(token)
	return &sdk{client: client}
}

type sdk struct {
	client *client
}

func (s *sdk) DontReport(code int) {
	s.client.DontReport(code)
}

func (s *sdk) SalesmanAccountsGet(params map[string]interface{}) (*responses.AccountsResponse, error) {
	content, err := s.client.post("youzan.salesman.accounts.get", "3.0.0", params)
	if err != nil {
		return nil, err
	}
	response := &responses.AccountsResponse{}
	err = json.Unmarshal(content, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *sdk) UserWeixinOpenidGet(params map[string]interface{}) (string, error) {
	content, err := s.client.post("youzan.user.weixin.openid.get", "3.0.0", params)
	if err != nil {
		return "", err
	}
	response := &responses.OpenIdResponse{}
	err = json.Unmarshal(content, response)
	if err != nil {
		return "", err
	}

	return response.Response.OpenId, nil
}

func (s *sdk) UsersWeixinFollowerGet(params map[string]interface{}) (*responses.WechatUserResponse, error) {
	content, err := s.client.post("youzan.users.weixin.follower.get", "3.0.0", params)
	if err != nil {
		panic(err)
	}
	response := &responses.WechatUserResponse{}
	err = json.Unmarshal(content, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
