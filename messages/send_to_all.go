package messages

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/prasadsurase/gk-kingdom-alies-problem-5/structs"
)

// SendToAll will send the messages to all states from the provided sender
func SendToAll(sender *structs.State, states []*structs.State) error {
	reader := bufio.NewReader(os.Stdin)
	for _, receiver := range states {
		if receiver != sender {
			fmt.Println("Enter message form", sender.Name, "to", receiver.Name)
			msg, _ := reader.ReadString([]byte("\n")[0])
			msg = strings.TrimSuffix(msg, "\n")
			sender.OutgoingMessages = append(sender.OutgoingMessages, map[string]string{receiver.Name: msg})
			receiver.IncomingMessages = append(receiver.IncomingMessages, map[string]string{sender.Name: msg})
			result := Compare(receiver.Emblem, msg)
			fmt.Println("Aly?", result)
			if result == true {
				sender.Alies = append(sender.Alies, receiver)
			}
		}
	}
	return nil
}

// Compare the messages to the provided emblem and return boolean result
func Compare(emblem string, msg string) bool {
	emblem = strings.ToLower(emblem)
	msg = strings.ToLower(msg)
	emblemSlice := strings.Split(emblem, "")
	result := true
	for _, char := range emblemSlice {
		if strings.Count(emblem, char) > strings.Count(msg, char) {
			result = false
			break
		}
	}
	return result
}
