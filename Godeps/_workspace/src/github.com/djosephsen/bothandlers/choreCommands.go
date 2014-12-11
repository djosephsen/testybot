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
		if len(res.Robot.Chores) == 0{
			reply=`No chores have been registered (sorry?)`
		}else{
			for _,c := range res.Robot.Chores{
				reply = fmt.Sprintf("%s\n%s:small_blue_diamond:%s:small_blue_diamond:%s:small_blue_diamond:%s",reply,c.Name, c.Schedule, c.Next, c.State)
			}
		}
		return res.Reply(reply)
	},
}
