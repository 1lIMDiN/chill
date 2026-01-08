package main

import (
	"fmt"
)

type Hero struct {
	Health int
	Attack int
	Def    int
}

type Inventory struct {
	Item map[string]int
}

type Mage struct {
	Hero
	Inventory
	Name string
}

func main() {
	Items := Inventory{make(map[string]int)}

	newHero := Mage{
		Hero: Hero{
			Health: 100,
			Attack: 25,
			Def:    20,
		},
		Inventory: Items,
		Name:      "Dragon Murder",
	}

	goblin := "Гоблин"
	ogre := "Огр"

	newHero.Fight(goblin)

	fmt.Println(newHero)

	newHero.Item["меч"] = 1

	if _, ok := newHero.Item["меч"]; ok {
		newHero.Attack += 25
	}

	newHero.Fight(ogre)

	fmt.Println(newHero)
}

func (h *Hero) Fight(unit string) {
	if unit == "Гоблин" {
		goblinHealth := 40
		goblinAttack := 5

		for goblinHealth > 0 && h.Health > 0{
			h.Health -= goblinAttack
			goblinHealth -= h.Attack
			fmt.Println(h.Health, goblinHealth)
		}
	} else {
		ogreHealth := 60
		ogreAttack := 10

		for ogreHealth > 0 && h.Health > 0 {
			ogreHealth -= h.Attack
			fmt.Printf("Герой атакует Огра на %d урона, здоровье Огра: %d\n", h.Attack, ogreHealth)
			h.Health -= ogreAttack
			
		}
	}
}
