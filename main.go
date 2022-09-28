package main

import (
	"os"

	"github.com/teksoftgroup/database"
)

func main() {
	runTest()
}

func runTest() {

	// initialize db
	dal, _ := database.NewLayer("db.db", os.Getpagesize())

	// create a new page
	p := dal.AllocEmptyPage()
	p.Number = dal.Manager.GetNextPage()
	copy(p.Data[:], "data")

	_ = dal.WritePage(p)
}
