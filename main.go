package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// State holds the name and the emblem of the state
type State struct {
	Name             string
	Emblem           string
	Alies            []*State
	IncomingMessages []map[string]string
	OutgoingMessages []map[string]string
}

// Message struct for election
type Message struct {
	Sender   *State
	Receiver *State
	Message  string
}

// MessageSet is the collection of messages used during election.
var MessageSet = [...]string{"Summer is coming", "a1d22n333a4444p", "oaaawaala", "zmzmzmzaztzozh", "Go risk it all",
	"Let's swing the sword together", "Die or play the tame of thrones", "Ahoy! Fight for me with men and money", "Drag on Martin!",
	"When you play the tame of thrones you win or you die.", "What could we say to the Lord of Death? Game on?",
	"Turn us away and we will burn you first", "Death is so terribly final while life is full of possibilities.", "You Win or You Die",
	"His watch is Ended", "Sphinx of black quartz judge my dozen vows", "Fear cuts deeper than swords My Lord.",
	"Different roads sometimes lead to the same castle.", "A DRAGON IS NOT A SLAVE.", "Do not waste paper", "Go ring all the bells",
	"Crazy Fredrick bought many very exquisite pearl emerald and diamond jewels.",
	"The quick brown fox jumps over a lazy dog multiple times.", "We promptly judged antique ivory buckles for the next prize.",
	"Walar Morghulis: All men must die."}

func main() {
	land := State{Name: "Land", Emblem: "Panda"}
	water := State{Name: "Water", Emblem: "Octopus"}
	ice := State{Name: "Ice", Emblem: "Mammoth"}
	air := State{Name: "Air", Emblem: "Owl"}
	fire := State{Name: "Fire", Emblem: "Dragon"}

	states := []*State{&land, &water, &ice, &air, &fire}
	reader := bufio.NewReader(os.Stdin)

	var input int
	for {
		fmt.Println("\n============================================================================")
		fmt.Println("Please choose an option")
		fmt.Println("1) Who is the ruler of Southeros?")
		fmt.Println("2) Send messages to all kingdoms.")
		fmt.Println("3) Hold election till ruler is selected")
		fmt.Println("4) Allies of Ruler?")
		fmt.Println("5) Exit")
		fmt.Scan(&input)
		switch input {
		case 1:
			ruler, err := findRuler(states)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if ruler == nil {
				fmt.Println("Ruler: None")
			} else {
				fmt.Println("Ruler:", ruler.Name)
			}
		case 2:
			resetAll(states)
			err := sendMessagesToAll(&land, states)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Sent messages to all states.")
		case 3:
			resetAll(states)
			// hold election.
			fmt.Println("Enter the kingdoms competing to be the ruler:")
			msg, _ := reader.ReadString([]byte("\n")[0])
			msg = strings.TrimSuffix(msg, "\n")

			holdElection(states)
			fmt.Println("Elections complete.")
		case 4:
			// alies of the ruler
			ruler, alies, err := findAliesOfRuler(states)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			if ruler != nil {
				fmt.Print("Alies of Kingdom ", ruler.Name, " are: ")
				for _, aly := range alies {
					fmt.Print(" ", aly.Name)
				}
			} else {
				fmt.Println("No Ruler, No Alies.")
			}
		case 5:
			return
		default:
			fmt.Println("Wrong choice.")
		}
	}
}

func findRuler(states []*State) (*State, error) {
	var ruler *State
	for _, state := range states {
		if ruler == nil {
			ruler = state
		}
		if (len(state.Alies) > 0) && (len(state.Alies) > len(ruler.Alies)) {
			ruler = state
		}
	}
	if len(ruler.Alies) == 0 {
		return nil, nil
	}
	return ruler, nil
}

func sendMessagesToAll(sender *State, states []*State) error {
	reader := bufio.NewReader(os.Stdin)
	for _, receiver := range states {
		if receiver != sender {
			fmt.Println("Enter message form", sender.Name, "to", receiver.Name)
			msg, _ := reader.ReadString([]byte("\n")[0])
			msg = strings.TrimSuffix(msg, "\n")
			sender.OutgoingMessages = append(sender.OutgoingMessages, map[string]string{receiver.Name: msg})
			receiver.IncomingMessages = append(receiver.IncomingMessages, map[string]string{sender.Name: msg})
			result := compare(receiver.Emblem, msg)
			fmt.Println("Aly?", result)
			if result == true {
				sender.Alies = append(sender.Alies, receiver)
			}
		}
	}
	return nil
}

func compare(emblem string, msg string) bool {
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

func findAliesOfRuler(states []*State) (*State, []*State, error) {
	var ruler *State
	for _, state := range states {
		if ruler == nil {
			ruler = state
		}
		if (len(state.Alies) > 0) && (len(state.Alies) > len(ruler.Alies)) {
			ruler = state
		}
	}
	if len(ruler.Alies) == 0 {
		return nil, nil, nil
	}
	return ruler, ruler.Alies, nil
}

func resetAll(states []*State) error {
	for _, state := range states {
		state.Alies = []*State{}
		state.IncomingMessages = []map[string]string{}
		state.OutgoingMessages = []map[string]string{}
	}
	return nil
}

func getRandomMessage() string {
	n := rand.Int() % len(MessageSet)
	return MessageSet[n]
}

func holdElection(states []*State) (*State, error) {
	resetAll(states)
	var ballotBox []*Message

	for _, sender := range states {
		for _, receiver := range states {
			if sender != receiver {
				msg := getRandomMessage()
				ballotBox = append(ballotBox, &Message{Sender: sender, Receiver: receiver, Message: msg})
			}
		}
	}

	for i := 0; i < 6; i++ {
		n := rand.Int() % len(ballotBox)
		msg := ballotBox[n]
		fmt.Println("Message:", msg)

		// add msg to sender's and receiver's respective queue.
		msg.Sender.OutgoingMessages = append(msg.Sender.OutgoingMessages, map[string]string{msg.Receiver.Name: msg.Message})
		msg.Receiver.IncomingMessages = append(msg.Receiver.IncomingMessages, map[string]string{msg.Sender.Name: msg.Message})

		// decide if the sender and receiver can be alies.
		result := compare(msg.Receiver.Emblem, msg.Message)

		// check if msg's receiver has already sent any message to msg's sender
		if len(msg.Receiver.OutgoingMessages) > 0 {
			for _, sentMessage := range msg.Receiver.OutgoingMessages {
				if sentMessage[msg.Sender.Name] == "" {
					fmt.Println("Aly?", result)
					if result == true {
						msg.Sender.Alies = append(msg.Sender.Alies, msg.Receiver)
					}
				}
			}
		} else {
			fmt.Println("Aly?", result)
			if result == true {
				msg.Sender.Alies = append(msg.Sender.Alies, msg.Receiver)
			}
		}
		for _, state := range states {
			fmt.Println("Allies for", state.Name, ":", len(state.Alies))
		}
	}

	statesWithAliances := []*State{}
	for _, state := range states {
		if len(state.Alies) > 0 {
			fmt.Println("State:", state)
			statesWithAliances = append(statesWithAliances, state)
		}
	}

	// find state with max alies. if there are more than 1, rerun the func with the max states, else return the max state
	statesWithMaxAliances := []*State{}
	for _, state := range statesWithAliances {
		if len(state.Alies) > 0 {
			// if slice of states of max alies is empty, add the state to slice if the state's alies > 0
			if len(statesWithMaxAliances) == 0 {
				statesWithMaxAliances = append(statesWithMaxAliances, state)
			} else {
				size := len(statesWithMaxAliances)
				// if state has more alies than the last state in the filtered slice, empty the filtred slice and add state to filtered slice.
				if len(state.Alies) > len(statesWithMaxAliances[size-1].Alies) {
					statesWithMaxAliances = []*State{}
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
		return holdElection(states)
	} else if len(statesWithMaxAliances) == 1 {
		return statesWithMaxAliances[0], nil
	} else {
		return holdElection(statesWithMaxAliances)
	}
}
