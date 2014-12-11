package bothandlers

import (
	"github.com/djosephsen/hal"
)

var Secondtest = &hal.Chore{
	Name:  `Second Test`,
	Schedule: `* * * * * *`,
	Room: `C02JM295Z`,
	Run: func(res *hal.Response) error {
		return res.Send(`successful test!`)
	},
}
