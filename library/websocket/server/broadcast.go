package server

type broadcast struct {
	fromId    string
	fromIp    string
	toId      string
	msgType   int
	body      string
	receiveAt string
}
