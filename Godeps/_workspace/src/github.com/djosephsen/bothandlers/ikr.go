package bothandlers

import (
	"time"
	"math/rand"
	"strings"
	"github.com/djosephsen/hal"
)

type ikr struct{
}

func (h *ikr) Method() string {
	return hal.HEAR
}

func (h *ikr) Usage() string {
	return `*ikr* - listens for enthusiasm; responds automatically with validation`
}

func (h *ikr) Pattern() string {
	triggers:=[]string{
		"best.*ev(er|ar)",
		"so good",
		"they have the best",
		"awesome",
		"I love",
		"fantastic|wonderful|outstanding|magnificent|brilliant|genius|amazing|epic|nice!",
		"ZOMG|OMG|OMFG",
		"(so|pretty) great",
		"off the hook",
	}
	pat := "(?i)"+strings.Join(triggers,"|")
	return pat
}

func (h *ikr) Run(res *hal.Response) error {
   now:=time.Now()
	rand.Seed(int64(now.Unix()))
	replies := []string{
		"*I know right?!*",
		"*OMG* couldn't agree more",
		":+1:",
		"+1",
		":arrow_up: THAT",
		":arrow_up: you complete me :arrow_up:",
		"so true",
		"agreed.",
		"that's the fact jack",
		"YUUUUUUP",
		"that's what I'm talkin bout",
		"*IKR?!*",
		"singit",
		"^droppin the truth bombs :boom: :boom: :boom:",
		"#legit",
		"/me nodds emphatically in agreement",
		"for REALZ though",
		"FOR REALSIES",
		"it's like you *literally* just read my mind right now",
	}
	reply := replies[rand.Intn(len(replies)-1)]
	hal.Logger.Debug(" *** ikr:Sending response: %s", reply)
	res.Send(reply)
	return nil
}

// Ping exports
var IKR = &ikr{}
