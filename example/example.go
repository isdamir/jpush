package main

import (
	"fmt"
	"github.com/isdamir/jpush"
)

const (
	appKey = "fc4e7512bfd31c3dba9bad3e"
	secret = "570d3442dc88d1382fe93b89"
)

const jsons = `{
   "platform" : "all",
   "audience" : "all",
   "notification" : {
      "alert" : "Hi, go!" 
   }
}`

func main() {
	c := jpush.NewPushClient(secret, appKey)

	//Platform
	//var pf jpush.Platform
	//pf.Add(ANDROID)

	//Audience
	//var ad jpush.Audience
	//s := []string{"1", "2", "3"}
	//ad.SetID(s)

	//Notice
	notice := jpush.NewNoticeSimple()
	notice.Alert = "alert_test"

	//NoticeBuilder
	nb := jpush.NewNoticeBuilder()
	nb.SetPlatform(jpush.AllPlatform())
	nb.SetAudience(jpush.AllAudience())
	nb.SetNotice(notice)

	var msg jpush.Message
	msg.Title = "Hello"
	msg.Content = "祝大家工作顺利"

	mb := jpush.NewMessageBuilder()
	mb.SetPlatform(jpush.AllPlatform())
	mb.SetAudience(jpush.AllAudience())
	mb.SetMessage(&msg)

	//push
	ret, err := c.Send(mb)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		if ret.Error.Code == 0 {
			fmt.Println("ok:", ret.Sendno)
		} else {
			fmt.Println("err:", ret.Error.Message)
		}
	}

}
