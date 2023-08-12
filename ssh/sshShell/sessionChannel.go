package sshshell

import "fmt"

type sessionCHannel struct {
	uuid         string
	inputChannel chan string
	killChannel  chan bool
}

func NewSessionChannel(uuid string) sessionCHannel {
	return sessionCHannel{
		uuid:         uuid,
		inputChannel: make(chan string),
		killChannel:  make(chan bool),
	}
}

func (sc *sessionCHannel) WriteInput(msg string) {
	fmt.Printf("write input msg: %v\n", msg)
	fmt.Printf("sc: %v\n", sc.uuid)
	sc.inputChannel <- msg
}

func (sc *sessionCHannel) KillSession() {
	sc.killChannel <- true
}

func (sc *sessionCHannel) GetChannelInput() chan string {
	return sc.inputChannel
}
func (sc sessionCHannel) GetChannelKill() chan bool {
	return sc.killChannel
}
