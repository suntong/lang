////////////////////////////////////////////////////////////////////////////
// Porgram: Set-Go_Demo.go
// Purpose: Demo the golang-set usage
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://coderwall.com/p/latrzg/consider-not-looping-sets-in-go
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	set "github.com/deckarep/golang-set"
)

func main() {
	requiredClasses := set.NewSet()
	requiredClasses.Add("Cooking")
	requiredClasses.Add("English")
	requiredClasses.Add("Math")
	requiredClasses.Add("Biology")

	scienceClasses := set.NewSet()
	scienceClasses.Add("Biology")
	scienceClasses.Add("Chemistry")

	electiveClasses := set.NewSet()
	electiveClasses.Add("Welding")
	electiveClasses.Add("Music")
	electiveClasses.Add("Automotive")

	bonusClasses := set.NewSet()
	bonusClasses.Add("Go Programming")
	bonusClasses.Add("Python Programming")

	allClasses := requiredClasses.Union(scienceClasses).
		Union(electiveClasses).Union(bonusClasses)

	fmt.Println(scienceClasses.Contains("Cooking"))
	//Returns: false

	// what are all the classes I can take that are not science classes?
	fmt.Println(allClasses.Difference(scienceClasses))
	//Returns: Set{Music, Automotive, Go Programming, Python Programming, Cooking, English, Math, Welding}

	// Which of the science classes are also required classes?
	fmt.Println(scienceClasses.Intersect(requiredClasses))
	//Returns: Set{Biology}

	// How many bonus classes are offered for this winter schedule?
	fmt.Println(bonusClasses.Cardinality())
	//Returns: 2

	// Check exist
	fmt.Println(allClasses.Contains("Welding", "Automotive", "English"))
	//Returns: true
}
