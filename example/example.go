package main

import (
	"fmt"
	"github.com/isdamir/jpush"
)

const (
	appKey = "修改为您自己的"
	secret = "修改为您自己的"
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
	notice.Alert = "测试"

	//NoticeBuilder
	nb := jpush.NewNoticeBuilder()
	nb.SetPlatform(jpush.AllPlatform())
	nb.SetAudience(jpush.AllAudience())
	nb.SetNotice(notice)
	//push
	ret, err := c.Send(nb)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		if ret.Error.Code == 0 {
			fmt.Println("ok:", ret.Sendno)
		} else {
			fmt.Println("err:", ret.Error.Message)
		}
	}
	noticeAndroid := jpush.NewNoticeAndroid()
	noticeAndroid.Alert = "测试android"
	noticeAndroid.Title = "标题"
	nb = jpush.NewNoticeBuilder()
	nb.SetPlatform(jpush.AllPlatform())
	nb.SetAudience(jpush.AllAudience())
	nb.SetNotice(noticeAndroid)
	//push
	ret, err = c.Send(nb)
	if err != nil {
		fmt.Println("err:", err.Error())
	} else {
		if ret.Error.Code == 0 {
			fmt.Println("ok:", ret.Sendno)
		} else {
			fmt.Println("err:", ret.Error.Message)
		}
	}
	var msg jpush.Message
	msg.Title = "Hello"
	msg.Content = "祝大家工作顺利"

	mb := jpush.NewMessageBuilder()
	mb.SetPlatform(jpush.AllPlatform())
	mb.SetAudience(jpush.AllAudience())
	mb.SetMessage(&msg)

	//push
	ret, err = c.Send(mb)
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
