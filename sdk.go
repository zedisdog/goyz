package goyz

import (
	"encoding/json"
	"goyz/responses"
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
	response := &responses.AccountsResponse{};
	err = json.Unmarshal([]byte(content), response)
	if err != nil {
		return nil, err
	}

	return response, nil
}