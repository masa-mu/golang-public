package message

type Model interface {
	GetMessage() (err error)
}

type model struct {
	messageRepository Repository
}

func NewModel(messageRepository Repository) Model {
	return &model{messageRepository}
}

func (m model) GetMessage() (err error) {
	return nil
}
