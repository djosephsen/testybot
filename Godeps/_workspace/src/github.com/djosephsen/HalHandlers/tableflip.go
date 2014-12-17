package HalHandlers

import "github.com/djosephsen/hal"

var Tableflip = &hal.Handler{
	Method:  hal.HEAR,
	Usage: `*tableflip*: bot overhears the pattern: (table)*flip(table)* and responds by flipping a unicode table`,
	Pattern: `(table)*flip(table)*`,
	Run: func(res *hal.Response) error {
		return res.Send(`(╯°□°）╯︵ ┻━┻`)
	},
}
