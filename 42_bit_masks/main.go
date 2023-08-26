package main

import "fmt"

// KeySet is a set of keys in the game.
type KeySet uint64

const (
	Copper  KeySet = 1 << iota // 1
	Jade                       // 2
	Crystal                    // 4
	// max 64 (uint64)
)

// Player is a player in the game
type Player struct {
	Name string
	Keys KeySet
}

// AddKey adds a key to the player keys
func (p *Player) AddKey(key KeySet) {
	p.Keys |= key
}

// RemoveKey removes key from player
func (p *Player) RemoveKey(key KeySet) {
	p.Keys &= ^key
}

// HasKey returns true if player has a key
func (p *Player) HasKey(key KeySet) bool {
	return p.Keys&key != 0
}

func main() {
	p := Player{Name: "Foo"}
	p.AddKey(Copper)
	p.AddKey(Crystal)
	p.AddKey(Jade)
	p.RemoveKey(Jade)
	fmt.Println(p.HasKey(Copper))
	fmt.Println(p.HasKey(Crystal))
	fmt.Println(p.HasKey(Jade))
}
