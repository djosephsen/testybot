package bothandlers

import (
	"github.com/djosephsen/hal"
	"fmt"
	"reflect"
	)

var Help = &hal.Handler{
	Method:  hal.HEAR,
	Pattern: `help`,
	Usage: `*help*: prints this message when you type "botname help"`,
	Run: func(res *hal.Response) error {
		var reply string
		handlers:=res.Robot.Handlers()
		HandlerType:=reflect.ValueOf(new(hal.Handler)).Elem()
		//I can't reflect out the fullHandler type becaues hal doesn't export it
		//so this plugin can only print the usage of hal.Handler's
		//FullHandlerType:=reflect.ValueOf(new(hal.FullHandler)).Elem()
		for _,h:=range handlers{
			hval:=reflect.ValueOf(h).Elem()
			hal.Logger.Debug("this hval is a ",hval.Type())
			if hval.Type() == HandlerType.Type(){
					usage := hval.FieldByName(`Usage`)
					reply=fmt.Sprintf("%s\n%s",reply,usage)
			}
		}
		return res.Send(reply)
	},
}
