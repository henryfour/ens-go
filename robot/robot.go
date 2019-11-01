package robot

import (
	"ens-go/core"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/prometheus/common/log"
	"strings"
	"sync"
	"time"
)

// Robot is a telegram robot
type Robot struct {
	running  bool
	quit     chan struct{}
	stopping chan struct{}
	mutex    sync.Mutex

	ens   *core.Ens
	token string
	api   *tgbotapi.BotAPI
	users map[string]struct{}
}

func NewRobot(ens *core.Ens, token string, users []string) *Robot {
	x := &Robot{
		running:  false,
		quit:     make(chan struct{}),
		stopping: make(chan struct{}),
		mutex:    sync.Mutex{},
		ens:      ens,
		token:    token,
		api:      nil,
		users: make(map[string]struct{}),
	}
	for _, u := range users {
		x.users[strings.ToLower(u)] = struct{}{}
	}
	return x
}

func (x *Robot) Start() error {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	if x.running {
		return errors.New("robot is running already")
	}
	log.Info("robot is starting")
	var err error
	x.api, err = tgbotapi.NewBotAPI(x.token)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("robot %s has started", x.api.Self.UserName))

	x.quit = make(chan struct{})
	x.running = true
	go x.loop()
	return nil
}

func (x *Robot) Stop() {
	log.Info("robot is stopping")
	x.mutex.Lock()
	defer x.mutex.Unlock()
	if x.running {
		close(x.quit)
		<-x.stopping
		x.running = false
	}
	log.Info("robot has stopped")
}

func (x *Robot) loop() {
	// telegram updates
	u := tgbotapi.NewUpdate(0)
	updates, err := x.api.GetUpdatesChan(u)
	if err != nil {
		log.Error("robot loop start failed", err)
		close(x.stopping)
		return
	}

	tick := 10 * time.Millisecond
	timer := time.NewTimer(0)

	for {
		select {
		case <-x.quit:
			close(x.stopping)
			return
		case update := <-updates:
			if err := x.onUpdate(update); err != nil {
				log.Error("handle update failed", err)
			}
		case <-timer.C:
			timer.Reset(tick)
		}
	}
}

func (x *Robot) sendReply(msg *tgbotapi.Message, text string) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, text)
	// reply.ReplyToMessageID = msg.MessageID
	if _, err := x.api.Send(reply); err != nil {
		log.Error("Send reply failed", "Err", err)
	}
}
