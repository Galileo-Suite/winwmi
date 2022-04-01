package common

import (
	"embed"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
	"github.com/vgcrld/winwmi/common/maps"
	"github.com/yusufpapurcu/wmi"
)

//go:embed files/*
var files embed.FS

func GetAllMetrics(c *cli.Context) error {
	spew.Dump(c)
	p := getWmi()
	spew.Dump(p)
	return nil
}

func getWmi() (r []maps.Win32_PhysicalMemory) {
	q := wmi.CreateQuery(&r, "")
	err := wmi.Query(q, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

// GetFileString returns the embeded data in files by name
func GetFileString(n string) (desc string) {
	b, err := files.ReadFile("files/" + n)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
