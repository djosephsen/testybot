package main

import (
	"github.com/djosephsen/hal"
	_ "github.com/djosephsen/hal/adapter/slack"
	"github.com/djosephsen/bothandlers"
	_ "github.com/djosephsen/hal/store/memory"
	"os"
)

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	robot.Handle(
		bothandlers.Syn,
		bothandlers.Tableflip,
		bothandlers.IKR,
//		bothandlers.ListChores,
//		bothandlers.ListRooms,
	)

//	robot.Schedule(
//		bothandlers.Secondtest,
//		)

	if err := robot.Run(); err != nil {
		hal.Logger.Error(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
