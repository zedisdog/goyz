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

type OpenIdResponse struct {
	Response struct {
		OpenId string `json:"open_id"`
	} `json:"response"`
}

type WechatUserResponse struct {
	Response struct {
		User struct {
			IsFollow  bool   `json:"is_follow"`
			City      string `json:"city"`
			Sex       string `json:"sex"`
			Avatar    string `json:"avatar"`
			TradedNum int    `json:"traded_num"`
			Points    int    `json:"points"`
			Tags      []struct {
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"tags"`
			Nick       string `json:"nick"`
			FollowTime int    `json:"follow_time"`
			Province   string `json:"province"`
			UserId     int    `json:"user_id"`
			UnionId    string `json:"union_id"`
			LevelInfo  struct {
				LevelId   int    `json:"level_id"`
				LevelName string `json:"level_name"`
			} `json:"level_info"`
			TradedMoney  string `json:"traded_money"`
			WeixinOpenid string `json:"weixin_openid"`
		} `json:"user"`
	} `json:"response"`
}

type ErrorResponse struct {
	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error_response"`
}
