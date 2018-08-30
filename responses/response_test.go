package responses

import (
	"encoding/json"
	"testing"
)

func TestAccountsResponse(t *testing.T) {
	data := `{"response":{"accounts":[{"seller":"2HStEl","from_buyer_mobile":"18280008445","money":"0.01","mobile":"18280008445","nickname":"－","created_at":"2018-08-10 10:08:39","order_num":6,"fans_id":0},{"seller":"3ELAST","from_buyer_mobile":"","money":"0.00","mobile":"15682138632","nickname":"乔治辣","created_at":"2018-06-08 15:10:01","order_num":0,"fans_id":5464052619},{"seller":"3kDRkn","from_buyer_mobile":"","money":"0.00","mobile":"15196382092","nickname":"曲原忻","created_at":"2018-03-05 14:54:00","order_num":0,"fans_id":4861323724},{"seller":"3kqYwh","from_buyer_mobile":"17602888986","money":"6.04","mobile":"18081092132","nickname":"喵味仙人派","created_at":"2018-08-27 13:59:50","order_num":19,"fans_id":4855716323},{"seller":"3dcYpJ","from_buyer_mobile":"17602888986","money":"57.00","mobile":"15281009123","nickname":"z🤡","created_at":"2018-08-27 16:35:53","order_num":18,"fans_id":4851134360},{"seller":"3bCthS","from_buyer_mobile":"","money":"1.04","mobile":"13208113135","nickname":"振宇","created_at":"2018-08-15 14:13:22","order_num":14,"fans_id":4863344245},{"seller":"3baVGV","from_buyer_mobile":"","money":"36.58","mobile":"17602888986","nickname":"有赞.展博","created_at":"2018-08-15 16:52:48","order_num":40,"fans_id":4673195512}],"total_results":7}}`
	response := &AccountsResponse{}
	err := json.Unmarshal([]byte(data), response)
	if err != nil {
		t.Errorf("json decode error: %v", err)
	}

	if response.Response.Accounts[0].Nickname != "－" {
		t.Errorf("expected < － >, got < %v >", response.Response.Accounts[0].Nickname)
	}

	if response.Response.Accounts[0].FromBuyerMobile != "18280008445" {
		t.Errorf("expected < 18280008445 >, got < %v >", response.Response.Accounts[0].FromBuyerMobile)
	}

	if response.Response.TotalResults != 7 {
		t.Errorf("expected < 7 >, got < %v >", response.Response.TotalResults)
	}
}
