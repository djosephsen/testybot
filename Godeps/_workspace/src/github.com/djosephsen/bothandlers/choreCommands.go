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
			reply=`Name  :small_blue_diamond:  Schedule  :small_blue_diamond:  Next Run  :small_blue_diamond: Current State`
			for _,c := range res.Robot.Chores{
				reply = fmt.Sprintf("%s\n%s:small_blue_diamond:%s:small_blue_diamond:%s:small_blue_diamond:%s",reply,c.Name, c.Schedule, c.Next, c.State)
			}
		}
		return res.Reply(reply)
	},
}
