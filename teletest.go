package teletest

import (
	"sync"
	"testing"

	"github.com/go-telebot/teletest/mocks"
	"github.com/stretchr/testify/mock"

	tele "gopkg.in/telebot.v3"
)

//go:generate mockery --srcpkg gopkg.in/telebot.v3 --name API --filename api.go

type T struct {
	*testing.T
	a    *mocks.API
	err  *error
	once *sync.Once
}

func New(t *testing.T) *T {
	return &T{
		T:    t,
		a:    mocks.NewAPI(t),
		err:  new(error),
		once: new(sync.Once),
	}
}

var User = &tele.User{
	ID:           1,
	FirstName:    "Tele",
	LastName:     "Test",
	Username:     "teletest",
	LanguageCode: "en",
}

func (t *T) FromCommand(text string) tele.Context {
	return tele.NewContext(t.a, tele.Update{
		Message: &tele.Message{
			Sender: User,
			Text:   text,
		},
	})
}

func (t *T) Expect(b *tele.Bot) *Expect {
	t.once.Do(func() {
		b.Use(func(next tele.HandlerFunc) tele.HandlerFunc {
			return func(c tele.Context) error {
				err := next(c)
				*t.err = err
				return err
			}
		})
	})
	return &Expect{
		Any: mock.Anything,
		t:   t,
		b:   b,
	}
}

type Expect struct {
	Any any
	t   *T
	b   *tele.Bot
}

func (e *Expect) Trigger(c tele.Context) error {
	e.b.ProcessContext(c)
	return *e.t.err
}

func (e *Expect) Send(args ...any) *mock.Call {
	return e.t.a.On("Send", args...).Return(nil, nil)
}

func (e *Expect) Assert() bool {
	return e.t.a.AssertExpectations(e.t)
}
