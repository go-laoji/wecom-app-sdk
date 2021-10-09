package workchatapp

//
// 企业微信部门管理接口
// 参考连接　https://open.work.weixin.qq.com/api/doc/90000/90135/90204
import (
	"encoding/json"
	"fmt"
	"github.com/go-laoji/workchatapp/internal"
)

type Department struct {
	Id       int32  `json:"id"`
	Order    int    `json:"order,omitempty"`
	ParentId int32  `json:"parentid" validate:"required"`
	Name     string `json:"name" validate:"required,min=1,max=32"`
	NameEn   string `json:"name_en,omitempty" validate:"omitempty,min=1,max=32"`
}

type DepartmentCreateResponse struct {
	internal.BizResponse
	Id int32 `json:"id"`
}

// DepartmentCreate 创建部门
func (app workChat) DepartmentCreate(department Department) (resp DepartmentCreateResponse) {
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
func (app workChat) DepartmentUpdate(department Department) (e internal.Error) {
	if ok := validate.Struct(department); ok != nil {
		e.ErrCode = 500
		e.ErrorMsg = ok.Error()
		return
	}
	queryParams := app.buildBasicTokenQuery(app.getContactsAccessToken())
	body, err := internal.HttpPost(fmt.Sprintf("/cgi-bin/department/update?%s", queryParams.Encode()), department)
	if err != nil {
		e.ErrCode = 500
		e.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &e)
	}
	return
}

// DepartmentDelete 删除部门
func (app workChat) DepartmentDelete(id int32) (e internal.Error) {
	queryParams := app.buildBasicTokenQuery(app.getContactsAccessToken())
	queryParams.Add("id", fmt.Sprintf("%v", id))
	body, err := internal.HttpGet(fmt.Sprintf("/cgi-bin/department/delete?%s", queryParams.Encode()))
	if err != nil {
		e.ErrCode = 500
		e.ErrorMsg = err.Error()
	} else {
		json.Unmarshal(body, &e)
	}
	return
}

type DepartmentListResponse struct {
	internal.BizResponse
	Department []Department `json:"department"`
}

// DepartmentList 获取部门列表
func (app workChat) DepartmentList(id int32) (resp DepartmentListResponse) {
	queryParams := app.buildBasicTokenQuery(app.getContactsAccessToken())
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
