package robot

import (
	"ens-go/core"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
	"sync"
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
	println("robot is starting")
	var err error
	x.api, err = tgbotapi.NewBotAPI(x.token)
	if err != nil {
		return err
	}
	fmt.Printf("robot %s has started\n", x.api.Self.UserName)

	x.quit = make(chan struct{})
	x.running = true
	go x.loop()
	return nil
}

func (x *Robot) Stop() {
	fmt.Println("robot is stopping")
	x.mutex.Lock()
	defer x.mutex.Unlock()
	if x.running {
		close(x.quit)
		<-x.stopping
		x.running = false
	}
	fmt.Println("robot has stopped")
}

func (x *Robot) loop() {
	// telegram updates
	u := tgbotapi.NewUpdate(0)
	updates, err := x.api.GetUpdatesChan(u)
	if err != nil {
		fmt.Println("robot loop start failed", err)
		close(x.stopping)
		return
	}

	for {
		select {
		case <-x.quit:
			close(x.stopping)
			return
		case update := <-updates:
			if err := x.onUpdate(update); err != nil {
				fmt.Println("handle update failed", err)
			}
		}
	}
}

func (x *Robot) sendReply(msg *tgbotapi.Message, text string) {
	reply := tgbotapi.NewMessage(msg.Chat.ID, text)
	// reply.ReplyToMessageID = msg.MessageID
	if _, err := x.api.Send(reply); err != nil {
		fmt.Println("Send reply failed", "Err", err)
	}
}
