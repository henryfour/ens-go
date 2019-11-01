package robot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func (x *Robot) onUpdate(u tgbotapi.Update) error {
	if u.Message == nil {
		return nil
	}

	// private msg
	if u.Message.Chat.IsPrivate() {
		return x.handleMsg(u.Message)
	}

	// group msg
	if u.Message.Chat.IsGroup() || u.Message.Chat.IsSuperGroup() {
		return nil
	}
	return nil
}

func (x *Robot) handleMsg(msg *tgbotapi.Message) error {
	names := strings.Fields(msg.Text)
	domains := x.ens.GetDomainInfos(names)
	rsp := ""
	for _, d := range domains {
		ds := d.String()
		rsp += ds[:len(ds)-34] + "\n"
	}
	x.sendReply(msg, rsp)
	return nil
}
