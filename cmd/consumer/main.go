package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	azbi "github.com/mkyc/go-structs-versioning-tests/azbi/v0"
	state "github.com/mkyc/go-structs-versioning-tests/state/v0"
)

const (
	configFile = "azbi-config.json"
	stateFile  = "state.json"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(filepath.Join(pwd, configFile))
	if err != nil {
		log.Fatal(err)
	}

	c := azbi.Config{}
	err = c.Load(b)
	if err != nil {
		log.Fatal(err)
	}

	p, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", string(p))

	b, err = ioutil.ReadFile(filepath.Join(pwd, stateFile))
	if err != nil {
		log.Fatal(err)
	}

	s := state.State{}
	err = s.Load(b)
	if err != nil {
		log.Fatal(err)
	}

	p, err = json.MarshalIndent(s, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", string(p))
}
