package bothandlers

import (
	"fmt"
	"github.com/djosephsen/hal"
	"time"
)

var ListRooms = &hal.Handler{
	Method:  hal.RESPOND,
	Usage:	`*Chores/ListRooms*: botname (what room)|(list *room)|(room *list): prints the Name of the current chatroom`,
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
	Usage:	`*Chores/ListChores*: botname (list chores)|(chore list): lists all registered chores`,
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

var ManageChores = &hal.Handler{
	Method:  hal.RESPOND,
	Usage:	`*Chores/ManageChores*: botname (start|stop) chore [chorename]: stops or starts the named chore`,
	Pattern: `(start|stop) chore (.*)`,
	Run: func(res *hal.Response) error {
		var reply string
		act:=res.Match[1]
		cname:=res.Match[2]
		c:=hal.GetChoreByName(cname,res.Robot)
		if c == nil{ 
			reply = fmt.Sprintf("Chore not found: %s",(cname[3]))
		}else{
			if act==`stop `{
				hal.KillChore(c)
			}else{
				hal.StartChore(c)
			}
			reply = fmt.Sprintf("%s\n%s:small_blue_diamond:%s:small_blue_diamond:%v:small_blue_diamond:%s",reply,c.Name, c.Sched, c.Next.Sub(time.Now()), c.State)
		}
		return res.Reply(reply)
	},
}
