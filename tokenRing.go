// TOKEN RING
package main

import "fmt"

const N = 4

type Packet struct {
	sender   int
	receiver int
	message  string
}
type Token struct {
	hasToken bool
	packet   Packet
}

func sender(id int, send chan Packet) {
	for i := 0; i < N; i++ {
		send <- Packet{id, i, "msg"}
	}
}
func receiver(id int, rec chan Packet) {
	for {
		p := <-rec
		println("Pacote recebido : ", id, p.sender, p.receiver, p.message)
	}
}
func node(id int, hasToken bool, send chan Packet, receive chan Packet, ringMy chan Token, ringNext chan Token) {
	fmt.Println("node ", id)
	go sender(id, send)
	go receiver(id, receive)
	sent := false
	for {
		if hasToken { //se tem token
			if !sent {
				select {
				case s := <-send: //se tem mensagem do usuário: manda
					sent = true
					ringNext <- Token{false, s}
					//println(id, " sent ", s.message, " from ", s.sender, " to ", s.receiver)
				default: //senão: repassa token
					continue
				}
			} else {
				select {
				case msg := <-ringMy:
					if msg.packet.receiver == id {
						receive <- msg.packet
					}
					hasToken = false
					sent = false
					ringNext <- Token{true, Packet{0, 0, ""}}
				default:
					continue
				}
			}
		} else { //se não tem token
			t := <-ringMy
			if !t.hasToken { //se recebe mensagem
				if t.packet.receiver == id { //para si: repassa ao usuário E no anel
					receive <- t.packet
					ringNext <- Token{false, t.packet}
				} else { //para outro: repassa no anel
					ringNext <- t
				}
			} else { //se recebe token: vá para o caso anterior
				hasToken = true
			}
		}
	}
}
func tokenRing() {
	var chanRing [N]chan Token
	var chanSend [N]chan Packet
	var chanRec [N]chan Packet
	for i := 0; i < N; i++ {
		chanRing[i] = make(chan Token)
		chanSend[i] = make(chan Packet)
		chanRec[i] = make(chan Packet)
	}
	for i := 0; i < (N - 1); i++ {
		go node(i, false, chanSend[i], chanRec[i], chanRing[i], chanRing[i+1])
	}
	go node(N-1, true, chanSend[N-1], chanRec[N-1], chanRing[N-1], chanRing[0])
	fin := make(chan struct{})
	<-fin
}
