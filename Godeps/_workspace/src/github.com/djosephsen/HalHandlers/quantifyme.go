package HalHandlers

import (
	"fmt"
	"github.com/djosephsen/hal"
	"math/rand"
	"time"
)

var Quantifyme = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `quantify (\S+)`,
	Usage: `*Quantifyme*: botname quantify <noun>: replies with a randomly generated quantification of funny but probably NSFW personal attributes like 'mads', 'fucks', horribleness and passive-aggresivity`,
	Run: func(res *hal.Response) error {
		user:=res.Match[1]
		if user==`me`{ user=`you` }
		now:=time.Now()
		rand.Seed(int64(now.Unix()))
		states:=[]string{`passive aggressive`, `mads`, `fucks`,`horrible`}
		state:=states[rand.Intn(len(states)-1)]
		var reply string
		switch state {
		case `horrible`,`passive aggressive`:
			if user==`you`{
				reply=fmt.Sprintf(`%s are currently %%%d.%04d %s`,user,rand.Intn(int(100)),rand.Intn(int(1000)),state)
			}else{
				reply=fmt.Sprintf(`%s is currently %%%d.%04d %s`,user,rand.Intn(int(100)),rand.Intn(int(1000)),state)
			}
		case `mads`:
			if user==`you`{
				reply=fmt.Sprintf(`%s are %d.%04d %s`,user,rand.Intn(int(4)),rand.Intn(int(1000)),state)
			}else{
				reply=fmt.Sprintf(`%s is %d.%04d %s`,user,rand.Intn(int(4)),rand.Intn(int(1000)),state)
			}
		case `fucks`:
			if user==`you`{
				reply=fmt.Sprintf(`%s give precisely %f %s`,user,rand.Float64(),state)
			}else{
				reply=fmt.Sprintf(`%s gives precisely %f %s`,user,rand.Float64(),state)
			}
		}

		return res.Reply(reply)
	},
}
