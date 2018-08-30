package goyz

import (
	"github.com/zedisdog/goyz/responses"
	"testing"
)

const token string = "3b5a23c385163f1f8d8cb20cd1675fa1"

func TestSdk_SalesmanAccountsGet(t *testing.T) {
	sdk := NewSdk(token)
	response, err := sdk.SalesmanAccountsGet(map[string]interface{}{"page_size": 100, "page_no": 1})
	if err != nil {
		t.Errorf("get response error: %v", err)
	}

	if len(response.Response.Accounts) == 0 {
		t.Errorf("expected accounts not eq 0, make sure use the right token")
	}
}

func TestSdk_UserWeixinOpenidGet(t *testing.T) {
	mobile := "15281009123"
	sdk := NewSdk(token)
	response, err := sdk.UserWeixinOpenidGet(map[string]interface{}{"mobile": mobile})
	if err != nil {
		t.Errorf("get response error: %v", err)
	}

	if response != "oaePdw8j7mWPM-tEJ13T7xFPQ2Oc" {
		t.Errorf("expected < oaePdw8j7mWPM-tEJ13T7xFPQ2Oc >, got < %v >, make sure use the right token and mobile", response)
	}
}

func TestSdk_UsersWeixinFollowerGet(t *testing.T) {
	openId := "oaePdw8j7mWPM-tEJ13T7xFPQ2Oc"
	sdk := NewSdk(token)
	response, err := sdk.UsersWeixinFollowerGet(map[string]interface{}{"weixin_openid": openId})
	if err != nil {
		t.Errorf("get response error: %v", err)
	}
	if _, ok := interface{}(response).(*responses.WechatUserResponse); !ok {
		t.Errorf("got non-*responses.WechatUserResponse type")
	}

	if response.Response.User.Sex != "m" {
		t.Errorf("expected sex < m >, got sex < %v >", response.Response.User.Sex)
	}
}
