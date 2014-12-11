package bothandlers

import (
	"fmt"
	"github.com/djosephsen/hal"
	"time"
	"strings"
)

var ListRooms = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `(what room)|(list *room)|(room *list)`,
	Run: func(res *hal.Response) error {
		room := res.Message.Room
		reply := fmt.Sprintf("Current room is: %s",room)
		return res.Send(reply)
	},
}


var ListChores = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `(list chores)|(chore list)`,
	Run: func(res *hal.Response) error {
		var reply string
		if len(res.Robot.Chores) == 0{
			reply=`No chores have been registered (sorry?)`
		}else{
			reply=`Name  :small_blue_diamond:  Schedule  :small_blue_diamond:  Firing in  :small_blue_diamond: Current State`
			for _,c := range res.Robot.Chores{
				reply = fmt.Sprintf("%s\n%s:small_blue_diamond:%s:small_blue_diamond:%v:small_blue_diamond:%s",reply,c.Name, c.Sched, c.Next.Sub(time.Now()), c.State)
			}
		}
		return res.Reply(reply)
	},
}

var StopChore = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `(stop chore)|(chore stop)`,
	Run: func(res *hal.Response) error {
		var reply string
		cname:=strings.SplitAfterN(res.Match[0],` `,3)
		c:=hal.GetChoreByName(cname[2],res.Robot)
		hal.KillChore(c)
		reply = fmt.Sprintf("%s\n%s:small_blue_diamond:%s:small_blue_diamond:%v:small_blue_diamond:%s",reply,c.Name, c.Sched, c.Next.Sub(time.Now()), c.State)
		return res.Reply(reply)
	},
}
