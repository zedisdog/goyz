package responses

type AccountsResponse struct {
	Response struct {
		Accounts []struct {
			Seller          string `json:"seller"`
			FromBuyerMobile string `json:"from_buyer_mobile"`
			Money           string `json:"money"`
			Mobile          string `json:"mobile"`
			Nickname        string `json:"nickname"`
			CreateAt        string `json:"create_at"`
			OrderNum        int    `json:"order_num"`
			FansId          int    `json:"fans_id"`
		} `json:"accounts"`
		TotalResults int `json:"total_results"`
	} `json:"response"`
}

type ErrorResponse struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error_response"`
}
