package main

import (
	"golang_practice/Gophercizes/Task/cmd"
	"golang_practice/Gophercizes/Task/db"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	dbPath := filepath.Join(home, "tasks.db")
	if err = db.Init(dbPath); err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}
