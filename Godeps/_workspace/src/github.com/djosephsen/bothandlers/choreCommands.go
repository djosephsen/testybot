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
				reply = fmt.Sprintf("%s\n%s:small_blue_diamond:%s:small_blue_diamond:%s:small_blue_diamond:%s",reply,c.Name, c.Sched, c.Next, c.State)
			}
		}
		return res.Reply(reply)
	},
}

var ListRooms = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `(what room)|(list *room)|(room *list)`,
	Run: func(res *hal.Response) error {
		room := res.Message.Room
		reply := fmt.Sprintf("Current room is: %s",room)
		return res.Send(reply)
	},
}
