// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TinyDependency struct {
	URL         string
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

	fmt.Println("no. | crate | versions")
	fmt.Println("----|-------|---------")

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
		fmt.Printf("%03d | [%s](%s) | \n", i, ok, dependencies[ok].URL)
	}

	fmt.Println("-------")
	for k, v := range dependencies {
		if _, done := oks[k]; done {
			continue
		}

		fmt.Println(k, v)
	}
}
