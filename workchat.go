package workchatapp

import (
	"encoding/json"
	"fmt"
	badger "github.com/dgraph-io/badger/v2"
	"github.com/go-laoji/workchatapp/internal"
	"net/url"
	"os"
	"time"
)

type IWorkChat interface {
	getContactsAccessToken() string
	getAppAccessToken() string
	GetCorpId() string

	//通讯录管理－成员管理 ↓

	UserCreate(User) internal.BizResponse
	UserGet(string) UserGetResponse
	UserUpdate(User) internal.BizResponse
	UserDelete(string) internal.BizResponse
	UserBatchDelete([]string) internal.BizResponse
	UserSimpleList(int32, int) UserSimpleListResponse
	UserList(int32, int) UserListResponse
	UserId2OpenId(string) UserId2OpenIdResponse

	//通讯录管理－部门管理 ↓

	DepartmentCreate(Department) DepartmentCreateResponse
	DepartmentUpdate(Department) internal.BizResponse
	DepartmentDelete(int32) internal.BizResponse
	DepartmentList(int32) DepartmentListResponse

	//通讯录管理－标签管理 ↓

	TagCreate(Tag) TagCreateResponse
	TagUpdate(Tag) internal.BizResponse
	TagDelete(int) internal.BizResponse
	TagList() TagListResponse
	TagUserList(int) TagUserListResponse
	TagAddUsers(int, []string, []int32) TagAddOrDelUsersResponse
	TagDelUsers(int, []string, []int32) TagAddOrDelUsersResponse

	//客户联系－联系我 ↓

	ExternalAddContactWay(ContactMe) ContactMeAddResponse
	ExternalUpdateContactWay(ContactMe) internal.BizResponse
	ExternalGetContactWay(string) ContactMeGetResponse
	ExternalListContactWay(int64, int64, string, int) ContactMeListResponse
	ExternalDeleteContactWay(string) internal.BizResponse
	ExternalCloseTempChat(string, string) internal.BizResponse

	ExternalContactGetFollowUserList() ExternalContactGetFollowUserListResponse

	//客户联系－规则组管理 ↓

	ExternalContactCustomerStrategyList(string, int) ExternalContactCustomerStrategyListResponse
	ExternalContactCustomerStrategyGet(int) ExternalContactCustomerStrategyGetResponse
	ExternalContactCustomerStrategyGetRange(int, string, int) ExternalContactCustomerStrategyGetRangeResponse
	ExternalContactCustomerStrategyCreate(ExternalContactCustomerStrategyCreateRequest) ExternalContactCustomerStrategyCreateResponse
	ExternalContactCustomerStrategyEdit(ExternalContactCustomerStrategyEditRequest) internal.BizResponse
	ExternalContactCustomerStrategyDelete(int) internal.BizResponse

	//应用管理 ↓

	AgentGet() AgentGetResponse
	AgentList() (resp AgentListResponse)

	//企业标签管理(客户) ↓

	CorpTagList([]string, []string) CorpTagListResponse
	CorpTagAdd(CorpTagGroup) CorpTagAddResponse
	CorpTagUpdate(CorpTag) internal.BizResponse
	CorpTagDelete([]string, []string) internal.BizResponse

	//在职继承 ↓

	TransferCustomer(TransferCustomerRequest) TransferCustomerResponse
	TransferResult(TransferResultRequest) TransferResultResponse

	//离职继承 ↓

	GetUnassignedList(request UnAssignedRequest) (resp UnAssignedResponse)
	TransferCustomerResigned(request TransferCustomerRequest) (resp TransferCustomerResponse)
	TransferResultResigned(request TransferResultRequest) (resp TransferResultResponse)
	TransferGroupChat(request GroupChatTransferRequest) (resp GroupChatTransferResponse)

	//客户群管理 ↓

	GroupChatList(GroupChatListFilter) GroupChatListResponse
	GroupChat(GroupChatRequest) GroupChatResponse
	GroupOpenId2ChatId(string) GroupOpenId2ChatIdResponse

	//客户朋友圈 ↓
	// TODO: 测试未完成，接口不完整

	AddMomentTask(task MomentTask) (resp AddMomentTaskResponse)
	GetMomentTaskResult(jobId string) (resp GetMomentTaskResultResponse)
	GetMomentList(filter MomentListFilter) (resp GetMomentListResponse)
	GetMomentTask(filter MomentTaskFilter) (resp GetMomentTaskResponse)

	//素材管理 ↓

	MediaUploadAttachment(Media) MediaUploadResponse
	MediaUpload(fileType MediaType, filePath string, fileName string) (resp MediaUploadResponse)
	MediaUploadImg(filePath string, fileName string) (resp MediaUploadImgResponse)

	//消息推送-客户群发 ↓

	AddMsgTemplate(msg ExternalMsg) (resp AddMsgTemplateResponse)
	GetGroupMsgListV2(filter GroupMsgListFilter) (resp GetGroupMsgListV2Response)
	GetGroupMsgTask(filter GroupMsgTaskFilter) (resp GetGroupMsgTaskResponse)
	GetGroupMsgSendResult(filter GroupMsgSendResultFilter) (resp GetGroupMsgSendResultResponse)

	//消息推送(内部成员) ↓

	MessageSend(msg interface{}) (resp MessageSendResponse)
	MessageReCall(msgId string) (resp internal.BizResponse)
	MessageUpdateTemplateCard(msg TemplateCardUpdateMessage) (resp MessageUpdateTemplateCardResponse)
	AppChatCreate(request AppChatCreateRequest) (resp AppChatCreateResponse)
	AppChatUpdate(request AppChatUpdateRequest) (resp internal.BizResponse)
	AppChatGet(chatId string) (resp AppChatGetResponse)
	AppChatSend(msg interface{}, chatId string) (resp internal.BizResponse)
}

type WorkChatConfig struct {
	CorpId        string
	ContactSecret string
	AppId         string
	AppSecret     string
}

type workChat struct {
	IWorkChat
	corpId        string
	contactSecret string
	appId         string
	appSecret     string
	cache         *badger.DB
}

func NewWorkChatApp(c WorkChatConfig) IWorkChat {
	app := new(workChat)
	app.corpId = c.CorpId
	app.contactSecret = c.ContactSecret
	app.appId = c.AppId
	app.appSecret = c.AppSecret
	app.cache, _ = badger.Open(badger.DefaultOptions("").WithInMemory(true))
	return app
}

func (app workChat) GetCorpId() string {
	return app.corpId
}

type accessTokenResponse struct {
	internal.BizResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (app workChat) requestAccessToken(secret string) (resp accessTokenResponse) {
	apiUrl := fmt.Sprintf("/cgi-bin/gettoken?corpid=%s&corpsecret=%s", app.corpId, secret)
	var data []byte
	var err error
	if data, err = internal.HttpGet(apiUrl); err != nil {
		resp.ErrCode = 400
		resp.ErrorMsg = err.Error()
		return
	}
	if err = json.Unmarshal(data, &resp); err != nil {
		resp.ErrCode = 400
		resp.ErrorMsg = err.Error()
		return
	}
	return resp
}

func (app *workChat) getContactsAccessToken() (token string) {
	mutex.Lock()
	defer mutex.Unlock()
	var err error
	var item *badger.Item
	err = app.cache.View(func(txn *badger.Txn) error {
		item, err = txn.Get([]byte("contactToken"))
		if err == badger.ErrKeyNotFound {
			return err
		}
		item.Value(func(val []byte) error {
			token = string(val)
			return nil
		})
		return err
	})
	if err == badger.ErrKeyNotFound {
		if resp := app.requestAccessToken(app.contactSecret); resp.ErrCode != 0 {
			panic(resp)
			//logger.Error(err.Error())
		} else {
			token = resp.AccessToken
			app.cache.Update(func(txn *badger.Txn) error {
				entry := badger.NewEntry([]byte("contactToken"), []byte(token)).WithTTL(time.Second * 7200)
				err = txn.SetEntry(entry)
				return err
			})
		}
	}
	return token
}

func (app *workChat) getAppAccessToken() (token string) {
	mutex.Lock()
	defer mutex.Unlock()
	var err error
	var item *badger.Item
	err = app.cache.View(func(txn *badger.Txn) error {
		item, err = txn.Get([]byte("appToken"))
		if err == badger.ErrKeyNotFound {
			return err
		}
		item.Value(func(val []byte) error {
			token = string(val)
			return nil
		})
		return err
	})
	if err == badger.ErrKeyNotFound {
		if resp := app.requestAccessToken(app.appSecret); resp.ErrCode != 0 {
			panic(resp)
			//logger.Error(err.Error())
		} else {
			token = resp.AccessToken
			app.cache.Update(func(txn *badger.Txn) error {
				entry := badger.NewEntry([]byte("appToken"), []byte(token)).WithTTL(time.Second * 7200)
				err = txn.SetEntry(entry)
				return err
			})
		}
	}
	return token
}

func (app workChat) buildBasicTokenQuery(token string) url.Values {
	queryParams := url.Values{}
	queryParams.Add("access_token", token)
	if os.Getenv("debug") != "" {
		queryParams.Add("debug", "1")
	}
	return queryParams
}
