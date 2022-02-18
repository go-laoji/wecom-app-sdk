package wecom

//
// 企业微信部门管理接口
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/90204
import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/wecom-app-sdk/internal"
)

type Department struct {
	Id               int32    `json:"id"`
	Order            int      `json:"order,omitempty"`
	ParentId         int32    `json:"parentid" validate:"required"`
	Name             string   `json:"name" validate:"required,min=1,max=32"`
	NameEn           string   `json:"name_en,omitempty" validate:"omitempty,min=1,max=32"`
	DepartmentLeader []string `json:"department_leader"`
}

type DepartmentCreateResponse struct {
	internal.BizResponse
	Id int32 `json:"id"`
}

// DepartmentCreate 创建部门
func (app weCom) DepartmentCreate(department Department) (resp DepartmentCreateResponse) {
	if ok := validate.Struct(department); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getContactsAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/department/create?%s", queryParams.Encode()), department)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// DepartmentUpdate 更新部门
func (app weCom) DepartmentUpdate(department Department) (resp internal.BizResponse) {
	if ok := validate.Struct(department); ok != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getContactsAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/department/update?%s", queryParams.Encode()), department)
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

// DepartmentDelete 删除部门
func (app weCom) DepartmentDelete(id int32) (resp internal.BizResponse) {
	queryParams := app.buildBasicTokenQuery(app.getContactsAccessToken())
	queryParams.Add("id", fmt.Sprintf("%v", id))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/department/delete?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type DepartmentListResponse struct {
	internal.BizResponse
	Department []Department `json:"department"`
}

// DepartmentList 获取部门列表
func (app weCom) DepartmentList(id int32) (resp DepartmentListResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	queryParams.Add("id", fmt.Sprintf("%v", id))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/department/list?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return

}

type DepartmentSimpleListResponse struct {
	internal.BizResponse
	DepartmentId []struct {
		Id       int32 `json:"id"`
		ParentId int32 `json:"parentid"`
		Order    int   `json:"order"`
	}
}

// DepartmentSimpleList 获取子部门ID列表
func (app weCom) DepartmentSimpleList(id int32) (resp DepartmentSimpleListResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	if id > 0 {
		queryParams.Add("id", fmt.Sprintf("%v", id))
	}
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/department/simplelist?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}

type DepartmentGetResponse struct {
	internal.BizResponse
	Department Department `json:"department"`
}

// DepartmentGet 获取单个部门详情
func (app weCom) DepartmentGet(id int32) (resp DepartmentGetResponse) {
	queryParams := app.buildBasicTokenQuery(app.getAppAccessToken())
	if id > 0 {
		queryParams.Add("id", fmt.Sprintf("%v", id))
	} else {
		resp.ErrCode = 500
		resp.ErrorMsg = "部门ID异常"
		return
	}
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/department/get?%s", queryParams.Encode()))
	if err != nil {
		resp.ErrCode = 500
		resp.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &resp)
	}
	return
}
