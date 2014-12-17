package bothandlers

import (
	"github.com/djosephsen/hal"
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
)

type gifyout struct{
	Meta	interface{}
	Data	struct{
		Tags []string
		Caption string
		Username string
	 	Image_width string
    	Image_frames string
    	Image_mp4_url string
    	Image_url string
    	Image_original_url string
    	Url string
    	Id string
    	Type string
    	Image_height string
    	Fixed_height_downsampled_url string
    	Fixed_height_downsampled_width string
    	Fixed_height_downsampled_height string
    	Fixed_width_downsampled_url string
    	Fixed_width_downsampled_width string
    	Fixed_width_downsampled_height string
    	Rating string
	}
}


var Gifme = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `gif me (.*)`,
	Usage: `*Gifme*: botname gif me freddie mercury: returns a random rated:PG gif of freddy mercury via the giphy API`,
	Run: func(res *hal.Response) error {
	
		search:=res.Match[1]
		q:=url.QueryEscape(search)
		myurl:=fmt.Sprintf("http://api.giphy.com/v1/gifs/random?rating=pg&api_key=dc6zaTOxFJmzC&tag=%s",q)
		hal.Logger.Debug(`myurl is`,myurl)
		g:=new(gifyout)
		resp,_:=http.Get(myurl)
		dec:= json.NewDecoder(resp.Body)
		dec.Decode(g)
		return res.Send(g.Data.Image_url)
	},
}
