package main

import (
	"fmt"

	"github.com/prasadsurase/gk-kingdom-alies-problem-5/elections"
	"github.com/prasadsurase/gk-kingdom-alies-problem-5/messages"
	"github.com/prasadsurase/gk-kingdom-alies-problem-5/search"
	"github.com/prasadsurase/gk-kingdom-alies-problem-5/structs"
)

func main() {
	space := structs.State{Name: "Space", Emblem: "Gorilla"}
	land := structs.State{Name: "Land", Emblem: "Panda"}
	water := structs.State{Name: "Water", Emblem: "Octopus"}
	ice := structs.State{Name: "Ice", Emblem: "Mammoth"}
	air := structs.State{Name: "Air", Emblem: "Owl"}
	fire := structs.State{Name: "Fire", Emblem: "Dragon"}

	states := []*structs.State{&space, &land, &water, &ice, &air, &fire}

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
			// find the ruler among the states.
			ruler := search.Ruler(states)
			if ruler == nil {
				fmt.Println("Ruler: None")
			} else {
				fmt.Println("Ruler:", ruler.Name)
			}
		case 2:
			// reset the results and start making alies again
			elections.ResetAll(states)
			err := messages.SendToAll(&space, states)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Sent messages to all states.")
		case 3:
			//reset the results and conduct elections till a ruler is found.
			elections.ResetAll(states)

			ruler := elections.Conduct(states)
			fmt.Println("Elections complete.")
			fmt.Println("Ruler:", ruler.Name)
		case 4:
			// alies of the ruler
			ruler, alies := search.AliesOfRuler(states)
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
