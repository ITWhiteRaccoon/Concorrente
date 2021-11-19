//Eduardo C. Andrade e Julia A. Maia
package main

import (
	"fmt"
	"log"
)

//region Definições

const N = 8

//Channels each node will use to read messages
var receiverChannels [N]chan Message

var allowedConnections = [][]int{
	/* 1 */ {4},
	/* 2 */ {3, 4},
	/* 3 */ {2, 5, 6},
	/* 4 */ {1, 2, 5},
	/* 5 */ {3, 4, 6, 7, 8},
	/* 6 */ {3, 5},
	/* 7 */ {5, 8},
	/* 8 */ {5, 7},
}

type MessageType int

const (
	Data         MessageType = 'd'
	Confirmation             = 'c'
)

type Message struct {
	id          string
	messageType MessageType
	sender      int
	receiver    int
	content     string
	route       []int
}

type Mail struct {
	receiver int
	message  Message
}

//endregion Definições

func createMessage(messageType MessageType, messageNumber, sender, receiver int, route []int) Message {
	return Message{
		id:          fmt.Sprintf("m%df%dt%d", messageNumber, sender, receiver), //data_message(i)from(id)to(i)
		messageType: messageType,
		sender:      sender,
		receiver:    receiver,
		content:     fmt.Sprintf("Hi node %d, I'm node %d!", receiver, sender),
		route:       route,
	}
}

//Creates a message to each node from 0 to N-1 and sends it to all nodes connected to the sender
func createAndSendToAll(sender int) {
	for i := 0; i < N; i++ {
		if i != sender {
			msg := createMessage(Data, i, sender, i, []int{sender})
			go sendToConnected(sender, msg)
		}
	}
}

//Sends a message to all nodes connected to the sender
func sendToConnected(sender int, msg Message) {
	for _, connection := range allowedConnections[sender] {
		receiverChannels[connection-1] <- msg
	}
}

func sendTo(sender, receiver int, msg Message) {
	sent := false
	for _, connection := range allowedConnections[sender] {
		if connection-1 == receiver { //Sends the message only if nodes are connected
			receiverChannels[connection-1] <- msg
			sent = true
		}
	}
	if !sent {
		log.Fatalf("receiver %d is not connected to sender %d\n", receiver, sender)
	}
}

func node(id int, receiverChan chan Message) {
	readMessages := make(map[string]bool)

	go createAndSendToAll(id)

	for {
		select {
		case receivedMessage := <-receiverChan:
			switch receivedMessage.messageType {
			case Data:
				if !readMessages[receivedMessage.id] { //If received message is not in read messages list
					readMessages[receivedMessage.id] = true //Add to read list
					if receivedMessage.receiver == id {     //If the message is for this node
						fmt.Printf("Received message id:%s\tcontent%s\n", receivedMessage.id, receivedMessage.content)

						lastNode := receivedMessage.route[len(receivedMessage.route)-1]
						confirmationMsg := Message{
							id:          receivedMessage.id,
							messageType: Confirmation,
							sender:      id,
							receiver:    receivedMessage.sender,
							content:     "",
							route:       receivedMessage.route[:len(receivedMessage.route)-1],
						}
						go sendTo(id, lastNode, confirmationMsg)
					} else { //If the message is for another node
						receivedMessage.route = append(receivedMessage.route, id)
						go sendToConnected(id, receivedMessage)
					}
				}
			case Confirmation:
				if receivedMessage.receiver == id { //If confirmation is for this node
					fmt.Printf("Confirmed message id:%s\n", receivedMessage.id)
				} else { //If confirmation is for another node
					lastNode := receivedMessage.route[len(receivedMessage.route)-1]              //Reads last node from route
					receivedMessage.route = receivedMessage.route[:len(receivedMessage.route)-1] //Removes it from the message
					go sendTo(id, lastNode, receivedMessage)                                     //Sends the updated message to last node
				}
			}
		default:

		}
	}
}
func inundacao() {
	for i := 0; i < N; i++ {
		receiverChannels[i] = make(chan Message, N*2)
	}

	for i := 0; i < N; i++ {
		go node(i, receiverChannels[i])
	}
	fin := make(chan struct{})
	<-fin
}
