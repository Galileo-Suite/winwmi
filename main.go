package main

/*
	This code must run on Windows.
	It calls the WMI locally.

*/
import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vgcrld/winwmi/common"
)

func main() {

	app := &cli.App{
		Action:      common.GetAllMetrics,
		Usage:       common.GetFileString("usage"),
		Description: common.GetFileString("description"),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
