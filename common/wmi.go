package common

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
	"github.com/vgcrld/winwmi/maps"
	"github.com/yusufpapurcu/wmi"
)

func GetAllMetrics(c *cli.Context) error {
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
