package service

import (
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestNotify(t *testing.T) {
	t.Run("Send is called with correct message", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()

		msg := "hello world!"
		MockNotifier := NewMockNotifier(controller)

		MockNotifier.EXPECT().Send(msg)
		us := &UserService{Notifier: MockNotifier}
		us.Notify(msg)
	})
}
