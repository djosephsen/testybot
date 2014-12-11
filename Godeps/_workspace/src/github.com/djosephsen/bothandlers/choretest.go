package bothandlers

import (
	"github.com/djosephsen/hal"
)

var ChoreTest = &hal.Chore{
	Name:  `test-chore`,
	Sched: `0 * * * * * *`,
	Room: `C031P7M3E`,
	Run: func(res *hal.Response) error {
		return res.Send(`successful test!`)
	},
}
