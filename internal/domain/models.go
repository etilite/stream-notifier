package domain

// todo possibly these models should be simple DTO
type Follower struct {
	id     string
	chatID string
}

func NewFollower(id string, chatID string) *Follower {
	return &Follower{
		id:     id,
		chatID: chatID,
	}
}

func (f *Follower) ID() string {
	return f.id
}

func (f *Follower) ChatID() string {
	return f.chatID
}

type Stream struct {
	id        string
	followers []Follower
}

func NewStream(id string, followers []Follower) *Stream {
	return &Stream{
		id:        id,
		followers: followers,
	}
}

func (s *Stream) ID() string {
	return s.id
}

func (s *Stream) Followers() []Follower {
	return s.followers
}

type Notification struct {
	id        string
	messageID string
}

func NewNotification(id string, messageID string) *Notification {
	return &Notification{
		id:        id,
		messageID: messageID,
	}
}

func (n *Notification) ID() string {
	return n.id
}

func (n *Notification) MessageID() string {
	return n.messageID
}
