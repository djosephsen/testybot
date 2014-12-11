package bothandlers

import (
	"fmt"
	"github.com/djosephsen/hal"
)

var ListChores = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `(list chores)|(chore list)`,
	Run: func(res *hal.Response) error {
		var reply string
		for _,c := range res.Robot.Chores{
			reply = fmt.Sprintf("%s\n%s:%s:%s:%s",reply,c.Name, c.Schedule, c.Next, c.State)
		}
		return res.Reply(reply)
	},
}
