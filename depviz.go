package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var sepecials = map[string]string{
	"etcommon-rlp":              "rlp",
	"etcommon-block":            "block",
	"etcommon-hexutil":          "hexutil",
	"etcommon-trie":             "trie",
	"etcommon-block-core":       "block-core",
	"etcommon-bloom":            "bloom",
	"etcommon-bigint":           "bigint",
	"aes":                       "aes/aes",
	"aesni":                     "aes/aesni",
	"aes-soft":                  "aes/aes-soft",
	"jh-x86_64":                 "hashes/jh",
	"skein-hash":                "hashes/skein",
	"groestl-aesni":             "hashes/groestl",
	"threefish-cipher":          "block-ciphers/threefish",
	"c2-chacha":                 "stream-ciphers/chacha",
	"blake-hash":                "hashes/blake",
	"ppv-lite86":                "utils-simd/ppv-lite86",
	"parity-scale-codec-derive": "derive",
}

type DependencyManifest struct {
	Git    string
	Branch string
	Tag    string
}

type TinyDependency struct {
	URL         string
	Dependents  []string
	DependencyN int
}

func newTOML(reader io.Reader) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("toml")

	if err := v.ReadConfig(reader); err != nil {
		return nil, err
	}

	return v, nil
}

func readDependencies(crate, url string) ([]string, error) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	toml, err := newTOML(response.Body)
	if err != nil {
		return nil, err
	}

	pkgName := toml.GetString("package.name")
	if pkgName != crate {
		return nil, fmt.Errorf("not the right pkg('%s') for crate(%s)", pkgName, crate)
	}

	uniques := make(map[string]struct{})
	//dependencieTypes := []string{"dependencies", "build-dependencies", "dev-dependencies"}
	dependencieTypes := []string{"dependencies"}

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
	depsGraph := make(map[string]TinyDependency)

	const github = "https://github.com"
	const rawcontent = "https://raw.githubusercontent.com"

	updateDepsGraph := func(crate, revision, tomlURL string) error {
		deps, err := readDependencies(crate, tomlURL)
		if err != nil {
			return err
		}

		crateDependency := depsGraph[crate]

		fmt.Println(crate, tomlURL)
		//fmt.Println(deps)
		//crateDependency.URL = tomlURL
		crateDependency.URL = strings.ReplaceAll(
			strings.ReplaceAll(tomlURL, "/"+revision+"/", "/blob/"+revision+"/"),
			rawcontent, github)

		for _, d := range deps {
			if _, ok := dependencies[d]; !ok {
				continue
			}

			/*
				if _, ok := depsGraph[crate]; !ok {
					depsGraph[crate] = make(map[string]struct{})
				}
				depsGraph[crate][d] = struct{}{}
			*/

			crateDependency.DependencyN++

			y := depsGraph[d]
			y.Dependents = append(y.Dependents, crate)
			depsGraph[d] = y
		}

		depsGraph[crate] = crateDependency

		return nil
	}

	fmt.Println("#(deps) =", len(dependencies))

	cared := "parity-scale-codec-derive"

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

		if _, ok := depsGraph[k]; !ok {
			depsGraph[k] = TinyDependency{}
		}

		if _, ok := sepecials[k]; !ok {
			if err := updateDepsGraph(k, revision, url+"/"+revision+"/Cargo.toml"); err == nil {
				continue
			} else if k == cared {
				fmt.Println("err0:", err, url+"/"+revision+"/"+"/Cargo.toml")
			}
		}

		kk, ok := sepecials[k]
		if !ok {
			kk = k
		}
		if err := updateDepsGraph(k, revision, url+"/"+revision+"/"+kk+"/Cargo.toml"); err == nil {
			continue
		} else if k == cared {
			fmt.Println("err1:", err, url+"/"+revision+"/"+kk+"/Cargo.toml")
		}

		crateName := strings.ReplaceAll(k, "-", "")
		if err := updateDepsGraph(k, revision, url+"/"+revision+"/"+crateName+"/Cargo.toml"); err == nil {
			continue
		} else if k == cared {
			fmt.Println("err2:", err, url+"/"+revision+"/"+crateName+"/Cargo.toml")
		}

		crateName = strings.ReplaceAll(k, "-", "_")
		if err := updateDepsGraph(k, revision, url+"/"+revision+"/"+crateName+"/Cargo.toml"); err == nil {
			continue
		} else if k == cared {
			fmt.Println("err3:", err, url+"/"+revision+"/"+crateName+"/Cargo.toml")
		}

		failed = append(failed, k)
	}

	//for k, v := range depsGraph {
	//	fmt.Println(k, v)
	//}

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

	outJSON, err := json.MarshalIndent(depsGraph, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("dependencies.json", outJSON, 0644); err != nil {
		panic(err)
	}
}
