//go:generate generator-lifx
package main

import (
	"fmt"

	"github.com/AnthonyPoschen/lifx-home/pkg/protocol"
)

func main() {
	fmt.Println(protocol.BUTTON_GESTURE_PRESS)

}
