package jpush

type NoticeSimple struct {
	/*
		通知的内容在各个平台上，都可能只有这一个最基本的属性 "alert"
	*/
	Alert string `json:"alert"`
}
type Notice struct {
	NoticeSimple
	/*
		这里自定义 JSON 格式的 Key/Value 信息，以供业务使用
	*/
	Extras map[string]interface{} `json:"extras,omitempty"`
}

type NoticeAndroid struct {
	Notice
	/*
		如果指定了，则通知里原来展示 App名称的地方，将展示成这个字段。
	*/
	Title string `json:"title"`
	/*
		Android SDK 可设置通知栏样式，这里根据样式 ID 来指定该使用哪套样式。
	*/
	BuilderId int `json:"builder_id,omitempty"`
}
type NoticeIos struct {
	Notice
	/*
		如果无此字段，则此消息无声音提示；有此字段，如果找到了指定的声音就播放该声音，否则播放默认声音,如果此字段为空字符串，iOS 7 为默认声音，iOS 8 为无声音。
		(消息) 说明：JPush 官方 API Library (SDK) 会默认填充声音字段。提供另外的方法关闭声音。
	*/
	Sound string `json:"sound,omitempty"`
	/*
		如果不填，表示不改变角标数字；否则把角标数字改为指定的数字；为 0 表示清除。
		新增支持 "+1" 功能，详情参考：http://blog.jpush.cn/ios_apns_badge_plus/

		(消息) 说明：JPush 官方 API Library (SDK) 会默认填充 badge 值为 "+1"。提供另外的方法不变更 badge 值。
	*/
	Badge string `json:"badge,omitempty"`
	/*
		如果为 true 表示要静默推送。
	*/
	ContentAvailable bool `json:"content-available"`
	/*
		设置APNs payload中的"category"字段值。

		(消息) 说明：ios8才支持该字段。
	*/
	Category string `json:"category,omitempty"`
}
type NoticeWinphone struct {
	Notice
	/*
		会填充到 toast 类型 text1 字段上。
	*/
	Title string `json:"title"`
	/*
		点击打开的页面。会填充到推送信息的 param 字段上，表示由哪个 App 页面打开该通知。可不填，则由默认的首页打开。
	*/
	OpenPage string `json:"_open_page,omitempty"`
}

func NewNoticeSimple() *NoticeSimple {
	return &NoticeSimple{}
}
func NewNoticeAndroid() *NoticeAndroid {
	return &NoticeAndroid{}
}
func NewNoticeIos() *NoticeIos {
	return &NoticeIos{}
}
func NewNoticeWinphone() *NoticeWinphone {
	return &NoticeWinphone{}
}
func (this *Notice) AddExtras(key string, value interface{}) {
	if this.Extras == nil {
		this.Extras = make(map[string]interface{})
	}
	this.Extras[key] = value
}
