package main

import (
	"log"
	"os"

	"github.com/karronqiu/GoInAction/chapter02/search"

	_ "github.com/karronqiu/GoInAction/chapter02/matchers"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
