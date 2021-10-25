package workchatapp

import (
	"fmt"
	"testing"
	"time"
)

var testTextMsgId = ""

func TestWorkChat_MessageSend(t *testing.T) {
	var testTextMessage = TextMessage{
		Message: Message{
			ToUser: "jifengwei",
		},
		Text: Text{
			Content: "message",
		},
	}
	resp := testWorkChat.MessageSend(testTextMessage)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	testTextMsgId = resp.MsgId
}

// 其它类型消息撤回未做测试，有时候会触发　41063　错误码，因为异步发送
func TestWorkChat_MessageReCall(t *testing.T) {
	resp := testWorkChat.MessageReCall(testTextMsgId)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
	t.Log(resp.ErrorMsg)
}

func TestWorkChat_MessageSend_Template_Card_TextNotice(t *testing.T) {
	var testTemplateCard = TemplateCardMessage{
		Message: Message{
			ToUser: "jifengwei",
		},
		TemplateCard: TemplateCard{
			CardType: CardTypeTextNotice,
			EmphasisContent: &EmphasisContent{
				Title: "template card",
				Desc:  "template card text_notice desc",
			},
			Source: Source{
				Desc:      "source desc",
				IconURL:   "http://wework.qpic.cn/bizmail/c0t2sVyzn1DcaYct1MwlVvB8sqCYcQ9Pzhic8qd3xGBNOWWHXK9Bnsw/0",
				DescColor: 1,
			},
			ActionMenu: &ActionMenu{
				Desc: "交互辅助",
				ActionList: []ActionList{
					{Text: "接受发送", Key: "A"},
					{Text: "不接受发送", Key: "B"},
				},
			},
			MainTitle: MainTitle{
				Title: "欢迎使用企业微信",
			},
			TaskID: fmt.Sprintf("t_%v", time.Now().Unix()),
			JumpList: []JumpList{
				{Type: 1, Title: "企业微信官网", URL: "https://work.weixin.qq.com"},
			},
			CardAction: CardAction{Type: 1, URL: "https://work.weixin.qq.com"},
		},
	}
	resp := testWorkChat.MessageSend(testTemplateCard)
	if resp.ErrCode != 0 {
		t.Error(resp.ErrorMsg)
		return
	}
}
