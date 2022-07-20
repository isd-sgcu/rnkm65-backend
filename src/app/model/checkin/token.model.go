package checkin

type CheckinToken struct {
	Token       string
	UserId      string
	CheckinType string
}

type TokenInfo struct {
	Id          string
	CheckinType string
	EventType   int32
}
