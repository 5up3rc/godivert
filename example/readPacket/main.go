package main

import (
	"github.com/williamfhe/godivert"
	"fmt"
)

func main() {
	winDivert, err := godivert.NewWinDivertHandle("true")
	if err != nil {
		panic(err)
	}

	packet, err := winDivert.Recv()
	if err != nil {
		panic(err)
	}

	fmt.Println(packet)

    packet.Send(winDivert)

	winDivert.Close()
}