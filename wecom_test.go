package wecom

import (
	"testing"
)

var testConfig = WorkChatConfig{
	CorpId:        "ww190690c489d2eb53",
	ContactSecret: "08tnu5LGrsbKwvEDfTGlBMFMw3CsUCwRMavxvkLZSH8",
	AppId:         "1000002",
	AppSecret:     "pedn4nqraARPFOG_A-aVFz1F9pp1sdR-3K1fsCpTwg0",
}
var testWorkChat = NewWeComApp(testConfig)

func TestWorkChat_GetCorpId(t *testing.T) {
	if testWorkChat.GetCorpId() != testConfig.CorpId {
		t.Error("GetCorpId　Fail")
	}
}
