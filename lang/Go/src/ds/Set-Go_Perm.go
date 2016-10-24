////////////////////////////////////////////////////////////////////////////
// Porgram: Set-Go_Perm.go
// Purpose: Demo saving the golang-set
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"

	set "github.com/deckarep/golang-set"
)

const persistName = "Set-Go_Perm.gob"

func main() {
	requiredClasses := set.NewSet()
	requiredClasses.Add("Cooking")
	requiredClasses.Add("English")
	requiredClasses.Add("Math")
	requiredClasses.Add("Biology")

	// SaveState failed:gob: type sync.RWMutex has no exported fields
	// if err := SaveState(persistName, requiredClasses); err != nil {
	// 	log.Fatal("SaveState failed:", err)
	// }

	{
		// Create JSON from the set.
		// ... Ignore errors.
		b, _ := json.Marshal(requiredClasses)
		// Convert bytes to string.
		s := string(b)
		fmt.Println(s)

		actual := set.NewSet()
		json.Unmarshal(b, actual)
		//fmt.Printf("%#v\n", actual)
		fmt.Println(actual)
	}

	X := []interface{}{10, 12, 12, 12, 13}
	Sx := set.NewSetFromSlice(X)
	{
		//fmt.Println(Sx)
		b, _ := json.Marshal(Sx)
		s := string(b)
		fmt.Println(s)

		actual := set.NewSet()
		json.Unmarshal(b, actual)
		//fmt.Println(actual)
		fmt.Printf("%v\n", actual)
	}

}

func SaveState(persistName string, state interface{}) error {
	// create persistence file
	f, err := os.Create(persistName)
	if err != nil {
		return err
	}
	defer f.Close()

	// write persistemce file
	e := gob.NewEncoder(f)
	if err = e.Encode(state); err != nil {
		return err
	}
	return nil
}

func RestoreState(persistName string, state interface{}) error {
	// open persistence file
	f, err := os.Open(persistName)
	if err != nil {
		return err
	}
	defer f.Close()

	// read persistemce file
	e := gob.NewDecoder(f)
	if err = e.Decode(state); err != nil {
		return err
	}
	return nil
}
