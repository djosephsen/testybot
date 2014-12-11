package bothandlers

import "github.com/djosephsen/hal"

var Tableflip = &hal.Handler{
	Method:  hal.HEAR,
	Pattern: `(table)*flip(table)*`,
	Run: func(res *hal.Response) error {
		return res.Send(`(╯°□°）╯︵ ┻━┻`)
	},
}
