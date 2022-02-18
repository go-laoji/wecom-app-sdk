package wecom

import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-app-sdk/internal"
)

type AgentGetResponse struct {
	internal.BizResponse
	Agentid        int    `json:"agentid"`
	Name           string `json:"name"`
	SquareLogoURL  string `json:"square_logo_url"`
	Description    string `json:"description"`
	AllowUserinfos struct {
		User []struct {
			Userid string `json:"userid"`
		} `json:"user"`
	} `json:"allow_userinfos"`
	AllowPartys struct {
		Partyid []int `json:"partyid"`
	} `json:"allow_partys"`
	AllowTags struct {
		Tagid []int `json:"tagid"`
	} `json:"allow_tags"`
	Close              int    `json:"close"`
	RedirectDomain     string `json:"redirect_domain"`
	ReportLocationFlag int    `json:"report_location_flag"`
	Isreportenter      int    `json:"isreportenter"`
	HomeURL            string `json:"home_url"`
}

func (app weCom) AgentGet() (resp AgentGetResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("agentid", app.appId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/agent/get?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type AgentListResponse struct {
	internal.BizResponse
	AgentList []struct {
		AgentId       int    `json:"agentid"`
		Name          string `json:"name"`
		SquareLogoUrl string `json:"square_logo_url"`
	}
}

func (app weCom) AgentList() (resp AgentListResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/agent/list?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
