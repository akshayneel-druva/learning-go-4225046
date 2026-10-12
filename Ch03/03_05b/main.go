package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps")
	states := make(map[string]string)
	fmt.Println("States:", states)

	states["NY"] = "New York"
	states["CA"] = "California"
	states["TX"] = "Texas"

	fmt.Println("States:", states)
	fmt.Println("Length of states:", len(states))

	fmt.Println(states["NY"])

	delete(states, "CA")
	fmt.Println("States:", states)

	states["FL"] = "Florida"
	fmt.Println("States:", states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	fmt.Println("Keys:", keys)
	sort.Strings(keys)
	fmt.Println("Sorted keys:", keys)

	for i := range keys {
		fmt.Println(keys[i], ":", states[keys[i]])
	}

	//sort the map by values
	values := make([]string, len(states))
	for i := 0; i < len(states); i++ {
		values[i] = states[keys[i]]
	}
	sort.Strings(values)
	fmt.Println("Sorted values:", values)

}
