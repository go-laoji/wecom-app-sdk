package workchatapp

import "testing"

var testConfig = WorkChatConfig{
	CorpId: "ww49a9747f42745dc8",
	ContactSecret: "XTzeFr-ojsX7DshopWksorPLxAQr5qKLZpOUF0299LA"}
var testWorkChat = NewWorkChatApp(testConfig)

func TestWorkChat_GetCorpId(t *testing.T) {
	if testWorkChat.GetCorpId()!=testConfig.CorpId{
		t.Error("GetCorpId　Fail")
	}
}