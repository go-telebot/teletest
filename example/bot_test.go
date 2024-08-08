package main

import (
	"testing"

	"github.com/go-telebot/teletest"
	tele "gopkg.in/telebot.v3"

	. "github.com/smartystreets/goconvey/convey"
)

var pref = tele.Settings{
	Offline:     true,
	Synchronous: true,
}

func TestBot(t *testing.T) {
	tt := teletest.New(t)

	Convey("Given a new bot", tt, func() {
		b, err := Bot(pref)
		So(err, ShouldBeNil)

		Convey("When user sends /start", func() {
			c := tt.FromCommand("/start")

			Convey("Bot should reply with a greeting", func() {
				expect := tt.Expect(b)
				expect.Send(expect.Any, "Hello, teletest!")

				So(expect.Trigger(c), ShouldBeNil)
				So(db, ShouldContainKey, c.Sender().ID)
			})
		})
	})
}
