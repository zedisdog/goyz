package goyz

import (
	"encoding/json"
	"errors"
	"github.com/zedisdog/goyz/responses"
	"github.com/zedisdog/goyz/util"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"
)

const baseUrl string = "https://open.youzan.com/api/oauthentry/"

func NewClient(token string) *client {
	return &client{token: token, httpClient: &http.Client{}}
}

type client struct {
	dontReport []int
	token      string
	httpClient *http.Client
}

func (c *client) DontReport(code int) {
	if !util.InSlice(c.dontReport, code) {
		c.dontReport = append(c.dontReport, code)
	}
}

func (c *client) post(method string, version string, params map[string]interface{}) (string, error) {
	response, err := c.httpClient.Do(c.postRequest(method, version, params))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if strings.Contains(string(content), "error_response") == true {
		errorResponse := &responses.ErrorResponse{}
		err := json.Unmarshal(content, errorResponse)
		if err != nil {
			panic(err)
		}

		err = errors.New(string(content))
		if util.InSlice(c.dontReport, errorResponse.ErrorResponse.Code) {
			return "", err
		} else {
			panic(err)
		}
	}

	return string(content), nil
}

func (c *client) postRequest(method string, version string, params map[string]interface{}) *http.Request {
	url := c.buildUrl(method, version)
	if params == nil {
		params = map[string]interface{}{}
	}
	params["access_token"] = c.token
	query := c.buildQuery(params)
	request, err := http.NewRequest("POST", url, strings.NewReader(query))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return request
}

func (client) buildUrl(method string, version string) string {
	methods := strings.Split(method, ".")
	return baseUrl + strings.Join(methods[:len(methods)-1], ".") + "/" + version + "/" + methods[len(methods)-1]
}

func (client) buildQuery(params map[string]interface{}) string {
	query := url2.Values{}
	for key, value := range params {

		switch v := value.(type) {
		case string:
			value = v
			break
		case int:
			value = strconv.Itoa(v)
			break
		}

		query.Set(key, value.(string))
	}
	return query.Encode()
}
