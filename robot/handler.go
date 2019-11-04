package robot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func (x *Robot) onUpdate(u tgbotapi.Update) error {
	if u.Message != nil && u.Message.Chat.IsPrivate() {
		return x.handleMsg(u.Message)
	}
	return nil
}

func (x *Robot) handleMsg(msg *tgbotapi.Message) error {
	u := strings.ToLower(msg.From.UserName)
	if _, ok := x.users[u]; !ok {
		return nil
	}
	names := strings.Fields(msg.Text)
	println("query ens: " + strings.Join(names, ","), "user", u)
	domains := x.ens.GetDomainInfos(names)
	rsp := ""
	for _, d := range domains {
		ds := d.String()
		rsp += ds[:len(ds)-34] + "\n"
	}
	x.sendReply(msg, rsp)
	return nil
}
