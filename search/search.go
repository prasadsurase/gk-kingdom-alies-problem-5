package search

import (
	"fmt"
	"sort"

	"github.com/prasadsurase/gk-kingdom-alies-problem-5/structs"
)

// Ruler finds the ruler based on alies from the provided states
func Ruler(states []*structs.State) *structs.State {
	sort.Slice(states, func(i, j int) bool {
		return len(states[i].Alies) >= len(states[j].Alies)
	})
	for _, s := range states {
		fmt.Println("State: ", s)
	}
	ruler := states[0]
	if len(ruler.Alies) < 3 {
		return nil
	}
	return ruler
}

// AliesOfRuler finds the ruler from the provided states and returns its alies.
func AliesOfRuler(states []*structs.State) (*structs.State, []*structs.State) {
	ruler := Ruler(states)
	if ruler != nil {
		return ruler, ruler.Alies
	}
	return nil, nil
}

// AlreadyAnAly checks if the provided state is an aly
func AlreadyAnAly(alies []*structs.State, state *structs.State) bool {
	for _, aly := range alies {
		if aly == state {
			return true
		}
	}
	return false
}
