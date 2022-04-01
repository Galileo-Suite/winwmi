package main

/*
	This code must run on Windows.
	It calls the WMI locally.

*/
import (
	"embed"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vgcrld/winwmi/common"
)

//go:embed files/*
var files embed.FS

func main() {

	app := &cli.App{
		Action:      common.GetAllMetrics,
		Usage:       getFileString("usage"),
		Description: getFileString("description"),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func getFileString(n string) (desc string) {
	b, err := files.ReadFile("files/" + n)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
