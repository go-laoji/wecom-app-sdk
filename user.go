package workchatapp

//
// 企业微信成员管理接口
//　TODO:
//		二次验证
//		邀请成员
//		获取加入企业二维码
//		获取企业活跃成员数
//
import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
)

// User 成员定义
//　可以参考连接：https://open.work.weixin.qq.com/api/doc/90000/90135/90195
type User struct {
	OpenUserId     string   `json:"open_userid,omitempty"` // 仅在查询时返回
	Userid         string   `json:"userid" validate:"required"`
	Name           string   `json:"name" validate:"required"`
	Alias          string   `json:"alias,omitempty"`
	Mobile         string   `json:"mobile"  validate:"required_without=Email,omitempty"`
	Department     []int32  `json:"department" validate:"required,max=100"`
	Order          []int32  `json:"order,omitempty"`
	Position       string   `json:"position,omitempty"`
	Gender         string   `json:"gender,omitempty" validate:"omitempty,oneof=1 2"`
	Email          string   `json:"email"  validate:"required_without=Mobile,omitempty,email"`
	IsLeaderInDept []int    `json:"is_leader_in_dept,omitempty"`
	DirectLeader   []string `json:"direct_leader"`
	Enable         int      `json:"enable"`
	AvatarMediaid  string   `json:"avatar_mediaid,omitempty"`
	Telephone      string   `json:"telephone,omitempty"`
	Address        string   `json:"address,omitempty"`
	MainDepartment int32    `json:"main_department,omitempty"`
	Extattr        struct {
		Attrs []Attrs `json:"attrs,omitempty"`
	} `json:"extattr,omitempty"`
	ToInvite         bool   `json:"to_invite,omitempty"`
	ExternalPosition string `json:"external_position,omitempty"`
	ExternalProfile  struct {
		ExternalCorpName string `json:"external_corp_name,omitempty"`
		WechatChannels   struct {
			Nickname string `json:"nickname,omitempty"`
		} `json:"wechat_channels,omitempty"`
		ExternalAttr []ExternalAttr `json:"external_attr,omitempty"`
	} `json:"external_profile,omitempty"`
}
type Attrs struct {
	Type int    `json:"type" validate:"required,oneof= 0 1 2"`
	Name string `json:"name" validate:"required"`
	Text struct {
		Value string `json:"value"`
	} `json:"text,omitempty"`
	Web struct {
		URL   string `json:"url" validate:"required"`
		Title string `json:"title" validate:"required"`
	} `json:"web,omitempty"`
}
type ExternalAttr struct {
	Type int    `json:"type"`
	Name string `json:"name"`
	Text struct {
		Value string `json:"value"`
	} `json:"text,omitempty"`
	Web struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"web,omitempty"`
	Miniprogram struct {
		Appid    string `json:"appid"`
		Pagepath string `json:"pagepath"`
		Title    string `json:"title"`
	} `json:"miniprogram,omitempty"`
}

// UserCreate 创建成员
func (app workChat) UserCreate(user User) (resp internal.BizResponse) {
	if ok := validate.Struct(user); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getContactsAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/create?%s", queryParams.Encode()), user)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UserGetResponse struct {
	internal.BizResponse
	User
}

// UserGet 读取成员
func (app workChat) UserGet(userId string) (resp UserGetResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("userid", userId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/user/get?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// UserUpdate 更新成员
func (app workChat) UserUpdate(user User) (resp internal.BizResponse) {
	if ok := validate.Struct(user); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getContactsAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/update?%s", queryParams.Encode()), user)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// UserDelete 删除成员
func (app workChat) UserDelete(userId string) (resp internal.BizResponse) {
	queryParams := app.buildBasicTokenQuery(app.getContactsAccessToken())
	queryParams.Add("userid", userId)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/user/delete?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// UserBatchDelete 批量删除成员
func (app workChat) UserBatchDelete(ids []string) (resp internal.BizResponse) {
	p := H{"useridlist": ids}
	queryParams := app.buildBasicTokenQuery(app.getContactsAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/batchdelete?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UserSimpleListResponse struct {
	internal.Error
	UserList []struct {
		UserId     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
	} `json:"userlist"`
}

// UserSimpleList 获取部门成员
func (app workChat) UserSimpleList(departId int32, fetchChild int) (resp UserSimpleListResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	if departId <= 0 {
		return UserSimpleListResponse{internal.Error{ErrorMsg: "部门ID必需大于0", ErrCode: 403}, nil}
	}
	queryParams.Add("department_id", fmt.Sprintf("%v", departId))
	queryParams.Add("fetch_child", fmt.Sprintf("%v", fetchChild))
	apiUrl := fmt.Sprintf("/cgi-bin/user/simplelist?%s", queryParams.Encode())
	body, err := internal.HttpGet(apiUrl)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(body, &resp); err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
		return
	}
	return
}

type UserListResponse struct {
	internal.BizResponse
	UserList []User `json:"userlist"`
}

// UserList 获取部门成员详情
func (app workChat) UserList(departId int32, fetchChild int) (resp UserListResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	if departId <= 0 {
		resp.ErrCode = 403
		resp.ErrorMsg = "部门ID必需大于0"
		return
	}
	queryParams.Add("department_id", fmt.Sprintf("%v", departId))
	queryParams.Add("fetch_child", fmt.Sprintf("%v", fetchChild))
	apiUrl := fmt.Sprintf("/cgi-bin/user/list?%s", queryParams.Encode())
	body, err := internal.HttpGet(apiUrl)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UserId2OpenIdResponse struct {
	internal.BizResponse
	OpenId string `json:"openid"`
}

// UserId2OpenId userid与openid互换
func (app workChat) UserId2OpenId(userId string) (resp UserId2OpenIdResponse) {
	p := H{"userid": userId}
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/user/convert_to_openid?%s",
		queryParams.Encode()), p)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type UserInfoResponse struct {
	internal.BizResponse
	UserId         string `json:"UserId"`
	DeviceId       string `json:"DeviceId"`
	OpenId         string `json:"OpenId"`
	ExternalUserId string `json:"external_userid"`
}

func (app workChat) GetUserInfo(code string) (resp UserInfoResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("code", code)
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/user/getuserinfo?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
