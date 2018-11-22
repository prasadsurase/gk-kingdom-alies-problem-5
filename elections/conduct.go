package elections

import (
	"fmt"
	"math/rand"

	"github.com/prasadsurase/gk-kingdom-alies-problem-5/messages"
	"github.com/prasadsurase/gk-kingdom-alies-problem-5/search"
	"github.com/prasadsurase/gk-kingdom-alies-problem-5/structs"
)

var messageSet = [...]string{"Summer is coming", "a1d22n333a4444p", "oaaawaala", "zmzmzmzaztzozh", "Go risk it all",
	"Let's swing the sword together", "Die or play the tame of thrones", "Ahoy! Fight for me with men and money", "Drag on Martin!",
	"When you play the tame of thrones you win or you die.", "What could we say to the Lord of Death? Game on?",
	"Turn us away and we will burn you first", "Death is so terribly final while life is full of possibilities.", "You Win or You Die",
	"His watch is Ended", "Sphinx of black quartz judge my dozen vows", "Fear cuts deeper than swords My Lord.",
	"Different roads sometimes lead to the same castle.", "A DRAGON IS NOT A SLAVE.", "Do not waste paper", "Go ring all the bells",
	"Crazy Fredrick bought many very exquisite pearl emerald and diamond jewels.",
	"The quick brown fox jumps over a lazy dog multiple times.", "We promptly judged antique ivory buckles for the next prize.",
	"Walar Morghulis: All men must die."}

// Conduct elections
func Conduct(states []*structs.State) *structs.State {
	// reset the election results and start
	ResetAll(states)
	var ballotBox []*structs.Message

	for _, sender := range states {
		for _, receiver := range states {
			if sender != receiver {
				msg := getRandomMessage()
				ballotBox = append(ballotBox, &structs.Message{Sender: sender, Receiver: receiver, Message: msg})
			}
		}
	}

	// pick 6 random messages from the messages list
	for i := 0; i < 6; i++ {
		n := rand.Int() % len(ballotBox)
		msg := ballotBox[n]
		fmt.Println("Message:", msg)

		// add msg to sender's and receiver's respective queue.
		msg.Sender.OutgoingMessages = append(msg.Sender.OutgoingMessages, map[string]string{msg.Receiver.Name: msg.Message})
		msg.Receiver.IncomingMessages = append(msg.Receiver.IncomingMessages, map[string]string{msg.Sender.Name: msg.Message})

		// decide if the sender and receiver can be alies.
		result := messages.Compare(msg.Receiver.Emblem, msg.Message)

		// check if msg's receiver has already sent any message to msg's sender
		if len(msg.Receiver.OutgoingMessages) > 0 {
			for _, sentMessage := range msg.Receiver.OutgoingMessages {
				if sentMessage[msg.Sender.Name] == "" {
					fmt.Println("Aly?", result)
					if result == true {
						isAly := search.AlreadyAnAly(msg.Sender.Alies, msg.Receiver)
						if isAly == false {
							msg.Sender.Alies = append(msg.Sender.Alies, msg.Receiver)
						}
					}
				}
			}
		} else {
			fmt.Println("Aly?", result)
			if result == true {
				isAly := search.AlreadyAnAly(msg.Sender.Alies, msg.Receiver)
				if isAly == false {
					msg.Sender.Alies = append(msg.Sender.Alies, msg.Receiver)
				}
			}
		}
		for _, state := range states {
			fmt.Println("Allies for", state.Name, ":", len(state.Alies))
		}
	}

	statesWithAliances := []*structs.State{}
	for _, state := range states {
		if len(state.Alies) > 0 {
			fmt.Println("State:", state)
			statesWithAliances = append(statesWithAliances, state)
		}
	}

	// find state with max alies. if there are more than 1, rerun the func with the max states, else return the max state
	statesWithMaxAliances := []*structs.State{}
	for _, state := range statesWithAliances {
		if len(state.Alies) > 0 {
			// if slice of states of max alies is empty, add the state to slice if the state's alies > 0
			if len(statesWithMaxAliances) == 0 {
				statesWithMaxAliances = append(statesWithMaxAliances, state)
			} else {
				size := len(statesWithMaxAliances)
				// if state has more alies than the last state in the filtered slice, empty the filtred slice and add state to filtered slice.
				if len(state.Alies) > len(statesWithMaxAliances[size-1].Alies) {
					statesWithMaxAliances = []*structs.State{}
					statesWithMaxAliances = append(statesWithMaxAliances, state)
				} else if len(state.Alies) == len(statesWithMaxAliances[size-1].Alies) {
					statesWithMaxAliances = append(statesWithMaxAliances, state)
				}
			}
		}
	}

	for _, state := range statesWithMaxAliances {
		fmt.Println("State with max alies:", state)
	}

	if len(statesWithMaxAliances) == 0 {
		return Conduct(states)
	} else if len(statesWithMaxAliances) == 1 {
		return statesWithMaxAliances[0]
	} else {
		return Conduct(statesWithMaxAliances)
	}
}

// ResetAll reset all the data
func ResetAll(states []*structs.State) error {
	for _, state := range states {
		state.Alies = []*structs.State{}
		state.IncomingMessages = []map[string]string{}
		state.OutgoingMessages = []map[string]string{}
	}
	return nil
}
func getRandomMessage() string {
	n := rand.Int() % len(messageSet)
	return messageSet[n]
}
