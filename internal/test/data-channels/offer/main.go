package main

import (
	"bufio"
	"fmt"
	"github.com/pion/webrtc/v3"
	"log"
	"meeting/internal/helper"
	"os"
	"strconv"
	"time"
)

func main() {
	// 1. create peer connection
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		return
	}
	defer func() {
		if err := peerConnection.Close(); err != nil {
			log.Println(err.Error())
		}
	}()
	// 2. create data channel
	dataChannel, err := peerConnection.CreateDataChannel("foo", nil)
	dataChannel.OnOpen(func() {
		log.Println("data channel has opened")
		i := -1000
		for range time.NewTicker(time.Second * 5).C {
			if err := dataChannel.SendText("offer : hello world " + strconv.Itoa(i)); err != nil {
				log.Println(err.Error())
			}
		}
	})
	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		fmt.Println(string(msg.Data))
	})
	// 3. create offer
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		return
	}
	// 4. set local description
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		return
	}
	// 5. print offer
	println("OFFER:")
	println(helper.Encode(offer))
	// 6. input answer
	println("请输入ANSWER:")
	var answer webrtc.SessionDescription
	answerStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	helper.Decode(answerStr, &answer)
	// 7. set remote description
	if err := peerConnection.SetRemoteDescription(answer); err != nil {
		panic(err)
	}
	select {}
}
