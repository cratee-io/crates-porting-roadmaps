package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var sepecials = map[string]string{
	"etcommon-rlp":        "rlp",
	"etcommon-block":      "block",
	"etcommon-hexutil":    "hexutil",
	"etcommon-trie":       "trie",
	"etcommon-block-core": "block-core",
	"etcommon-bloom":      "bloom",
	"etcommon-bigint":     "bigint",
	"aes":                 "aes/aes",
	"aesni":               "aes/aesni",
	"aes-soft":            "aes/aes-soft",
	"jh-x86_64":           "hashes/jh",
	"skein-hash":          "hashes/skein",
	"groestl-aesni":       "hashes/groestl",
	"threefish-cipher":    "block-ciphers/threefish",
	"c2-chacha":           "stream-ciphers/chacha",
	"blake-hash":          "hashes/blake",
	"ppv-lite86":          "utils-simd/ppv-lite86",
}

type DependencyManifest struct {
	Git    string
	Branch string
	Tag    string
}

func newTOML(reader io.Reader) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("toml")

	if err := v.ReadConfig(reader); err != nil {
		return nil, err
	}

	return v, nil
}

func readDependencies(url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	toml, err := newTOML(response.Body)
	if err != nil {
		return nil, err
	}

	if toml.IsSet("workspace") && !toml.IsSet("dependencies") {
		return nil, errors.New("workspace isn't crate")
	}

	uniques := make(map[string]struct{})
	dependencieTypes := []string{"dependencies", "build-dependencies", "dev-dependencies"}

	for _, d := range dependencieTypes {
		if !toml.IsSet(d) {
			continue
		}

		keys := toml.Sub(d).AllKeys()
		for _, v := range keys {
			x := strings.Split(v, ".")
			uniques[x[0]] = struct{}{}
		}
	}

	out := make([]string, 0, len(uniques))
	for k := range uniques {
		out = append(out, k)
	}

	return out, nil
}

func main() {
	f, err := os.Open("./dumb-all/Cargo.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	v := viper.New()
	// MUST set the config type
	v.SetConfigType("toml")

	if err := v.ReadConfig(f); err != nil {
		panic(err)
	}

	var dependencies map[string]DependencyManifest
	if err := v.UnmarshalKey("dependencies", &dependencies); err != nil {
		panic(err)
	}

	//fmt.Println(dependencies)

	//depsGraph := make(map[string]map[string]struct{})
	depsGraph := make(map[string]int)

	updateDepsGraph := func(crate, tomlURL string) error {
		deps, err := readDependencies(tomlURL)
		if err != nil {
			return err
		}

		fmt.Println(crate, tomlURL)
		//fmt.Println(deps)

		for _, d := range deps {
			if _, ok := dependencies[d]; !ok {
				continue
			}

			//if _, ok := depsGraph[d]; !ok {
			//	//depsGraph[d] = make(map[string]struct{})
			//}
			//depsGraph[d][k] = struct{}{}

			depsGraph[d]++
		}

		return nil
	}

	const github = "https://github.com"
	const rawcontent = "https://raw.githubusercontent.com"

	/*
		{
			const crate = "failure"
			vv := dependencies[crate]

			url := strings.ReplaceAll(vv.Git, github, rawcontent)
			url = strings.TrimSuffix(url, ".git")

			revision := vv.Tag
			if revision == "" {
				revision = vv.Branch
			}
			if revision == "" {
				revision = "master"
			}

			k := crate

			if kk, ok := sepecials[k]; ok {
				k = kk
			}

			if err := updateDepsGraph(k, url+"/"+revision+"/Cargo.toml"); err != nil {
				fmt.Println("err1", err)
			}

			if err := updateDepsGraph(k, url+"/"+revision+"/"+k+"/Cargo.toml"); err != nil {
				fmt.Println("err2", err)
			}

			crateName := strings.ReplaceAll(k, "-", "")
			if err := updateDepsGraph(k, url+"/"+revision+"/"+crateName+"/Cargo.toml"); err != nil {
				fmt.Println("err3", err)
			}
		}
		fmt.Println("--------- end test")
	*/

	fmt.Println("#(deps) =", len(dependencies))

	var failed []string
	for k, v := range dependencies {
		url := strings.ReplaceAll(v.Git, github, rawcontent)
		url = strings.TrimSuffix(url, ".git")

		revision := v.Tag
		if revision == "" {
			revision = v.Branch
		}
		if revision == "" {
			revision = "master"
		}

		depsGraph[k]++

		if err := updateDepsGraph(k, url+"/"+revision+"/Cargo.toml"); err == nil {
			continue
		}

		kk, ok := sepecials[k]
		if !ok {
			kk = k
		}
		if err := updateDepsGraph(k, url+"/"+revision+"/"+kk+"/Cargo.toml"); err == nil {
			continue
		}

		crateName := strings.ReplaceAll(k, "-", "")
		if err := updateDepsGraph(k, url+"/"+revision+"/"+crateName+"/Cargo.toml"); err == nil {
			continue
		}

		failed = append(failed, k)
	}

	for k, v := range depsGraph {
		fmt.Println(k, v)
	}

	fmt.Println("--------")
	fmt.Println("nOk:", len(depsGraph))
	fmt.Println("nErr:", len(failed))

	fmt.Println("--------")
	fmt.Println("failures")
	for _, v := range failed {
		fmt.Println(v)
	}

	for k := range depsGraph {
		if _, ok := dependencies[k]; !ok {
			fmt.Println("[-]", k)
		}
	}
}
