// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TinyDependency struct {
	Dependents  []string
	DependencyN int
}

func main() {
	data, err := ioutil.ReadFile("dependencies.json")
	if err != nil {
		panic(err)
	}

	var dependencies map[string]TinyDependency
	if err := json.Unmarshal(data, &dependencies); err != nil {
		panic(err)
	}

	//fmt.Println(dependencies)

	var queue []string
	for crate, v := range dependencies {
		if v.DependencyN == 0 {
			queue = append(queue, crate)
		}
	}

	oks := make(map[string]struct{})
	for i := 0; i < len(dependencies); i++ {
		if len(queue) == 0 {
			fmt.Println("cyclic at", i)
			break
		}

		ok := queue[0]
		queue = queue[1:]

		for _, d := range dependencies[ok].Dependents {
			x := dependencies[d]
			x.DependencyN--
			dependencies[d] = x

			if x.DependencyN == 0 {
				queue = append(queue, d)
			}
		}

		oks[ok] = struct{}{}
		fmt.Printf("%d: %s\n", i, ok)
	}

	fmt.Println("-------")
	for k, v := range dependencies {
		if _, done := oks[k]; done {
			continue
		}

		fmt.Println(k, v)
	}
}
