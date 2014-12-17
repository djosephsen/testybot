package bothandlers

import "github.com/djosephsen/hal"

var Syn = &hal.Handler{
	Method:  hal.RESPOND,
	Usage: `*Syn*: bothaname syn: responds with an 'ack'`,
	Pattern: `syn`,
	Run: func(res *hal.Response) error {
		return res.Reply(`ack`)
	},
}
