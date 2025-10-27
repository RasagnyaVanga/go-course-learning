package service

// mockgen creates fake implementation for this interface
// i.e, MockNotifier (a struct for this interface that implements Send method)
// and a constructor - NewMockNotifier(controller)
// also has Send's base method and expected call.
type Notifier interface {
	Send(msg string) error
}

type UserService struct {
	Notifier Notifier
}

func (u *UserService) Notify(msg string) error {
	return u.Notifier.Send(msg)
}
