// https://code.tutsplus.com/tutorials/lets-go-object-oriented-programming-in-golang--cms-26540

package main

import "fmt"

type Creature struct {
	Name string
	Real bool
}

func Dump(c *Creature) {
	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

func (c Creature) Dump() {
	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

type FlyingCreature struct {
	Creature
	WingSpan int
}

func (fc FlyingCreature) Dump() {
	fmt.Printf("Name: '%s', Real: %t, WingSpan: %d\n",
		fc.Name,
		fc.Real,
		fc.WingSpan)
}

func (fc *FlyingCreature) Grow() {
	fc.WingSpan++
}

type Unicorn struct {
	Creature
}

type Dragon struct {
	FlyingCreature
}

type Pterodactyl struct {
	FlyingCreature
}

func NewPterodactyl(wingSpan int) *Pterodactyl {
	pet := &Pterodactyl{
		FlyingCreature{
			Creature{
				"Pterodactyl",
				true,
			},
			wingSpan,
		},
	}
	return pet
}

type Dumper interface {
	Dump()
}

type Door struct {
	Thickness int
	Color     string
}

func (d Door) Dump() {
	fmt.Printf("Door => Thickness: %d, Color: %s", d.Thickness, d.Color)
}

func main() {
	creature := &Creature{
		"some creature",
		false,
	}

	uni := Unicorn{
		Creature{
			"Unicorn",
			false,
		},
	}

	pet1 := &Pterodactyl{
		FlyingCreature{
			Creature{
				"Pterodactyl",
				true,
			},
			5,
		},
	}

	pet2 := NewPterodactyl(8)

	door := &Door{3, "red"}

	Dump(creature)
	creature.Dump()
	uni.Dump()
	pet1.Dump()
	pet2.Dump()

	creatures := []Creature{
		*creature,
		uni.Creature,
		pet1.Creature,
		pet2.Creature}
	fmt.Println("\nDump() through Creature embedded type")
	for _, creature := range creatures {
		creature.Dump()
	}

	dumpers := []Dumper{creature, uni, pet1, pet2, door}
	fmt.Println("\nDump() through Dumper interface")
	for _, dumper := range dumpers {
		dumper.Dump()
	}

	fmt.Println("\n")
	pet2.Grow()
	pet2.Dump()

}

/*

Name: 'some creature', Real: false
Name: 'some creature', Real: false
Name: 'Unicorn', Real: false
Name: 'Pterodactyl', Real: true, WingSpan: 5
Name: 'Pterodactyl', Real: true, WingSpan: 8

Dump() through Creature embedded type
Name: 'some creature', Real: false
Name: 'Unicorn', Real: false
Name: 'Pterodactyl', Real: true
Name: 'Pterodactyl', Real: true

Dump() through Dumper interface
Name: 'some creature', Real: false
Name: 'Unicorn', Real: false
Name: 'Pterodactyl', Real: true, WingSpan: 5
Name: 'Pterodactyl', Real: true, WingSpan: 8
Door => Thickness: 3, Color: red

Name: 'Pterodactyl', Real: true, WingSpan: 9

*/
