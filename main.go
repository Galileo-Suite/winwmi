package main

/*
	This code must run on Windows.
	It calls the WMI locally.

*/
import (
	"embed"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/yusufpapurcu/wmi" // On Linux/Mac, build constraints - it must be built on Win

	"github.com/urfave/cli/v2"

	"github.com/vgcrld/winwmi/maps"

	_ "github.com/vgcrld/winwmi/common" // Inclue for fun?
)

//go:embed files/*
var files embed.FS

func main() {

	description, err := files.ReadFile("files/description")
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Action:      getAllMetrics,
		Usage:       "Easy to use metrics getter",
		Description: string(description),
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func getAllMetrics(c *cli.Context) error {
	p := getWmi(maps.Win32_Process{})
	spew.Dump(len(p), "processe(s)")
	return nil
}

func getWmi(m maps.Win32_Process) (r []maps.Win32_Process) {
	q := wmi.CreateQuery(&r, "")
	err := wmi.Query(q, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
