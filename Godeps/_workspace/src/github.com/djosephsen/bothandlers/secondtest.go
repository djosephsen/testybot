package bothandlers

import (
	"github.com/djosephsen/hal"
)

var Secondtest = &hal.Chore{
	Name:  `Second Test`,
	Sched: `0 * * * * * *`,
	Room: `C031P7M3E`,
	Run: func(res *hal.Response) error {
		return res.Send(`successful test!`)
	},
}
