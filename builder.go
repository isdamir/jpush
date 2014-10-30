package jpush

//Base Builder
type Builder struct {
	Platform interface{} `json:"platform"`
	Audience interface{} `json:"audience"`
	Options  *Options    `json:"options"`
}

//MessageBuilder
type MessageBuilder struct {
	Builder
	Message interface{} `json:"message"`
}

//NoticeBuilder
type NoticeBuilder struct {
	Builder
	Notification map[string]interface{} `json:"notification"`
}

//MessageAndNotice
type MessageAndNoticeBuilder struct {
	Builder
	Notification map[string]interface{} `json:"notification"`
	Message      interface{}            `json:"message"`
}

//---------------------MessageBuilder --------------------
func NewMessageBuilder() *MessageBuilder {
	mb := &MessageBuilder{}
	mb.Options = NewOptions()
	return mb
}

func (this *MessageBuilder) SetMessage(m *Message) {
	this.Message = m
}

func (this *MessageBuilder) SetPlatform(pf *Platform) {
	this.Platform = pf.Object
}

func (this *MessageBuilder) SetAudience(ad *Audience) {
	this.Audience = ad.Object
}

func (this *MessageBuilder) SetOptions(o *Options) {
	this.Options = o
}

//---------------------NoticeBuilder----------------------
func NewNoticeBuilder() *NoticeBuilder {
	nb := &NoticeBuilder{}
	nb.Options = NewOptions()
	return nb
}

func (this *NoticeBuilder) SetPlatform(pf *Platform) {
	this.Platform = pf.Object
}

func (this *NoticeBuilder) SetAudience(ad *Audience) {
	this.Audience = ad.Object
}

func (this *NoticeBuilder) SetOptions(o *Options) {
	this.Options = o
}

/*
可以为每类Notice设置一个
*/
func (this *NoticeBuilder) SetNotice(o interface{}) {
	if this.Notification == nil {
		this.Notification = make(map[string]interface{})
	}
	switch o.(type) {
	case *NoticeAndroid:
		this.Notification[string(ANDROID)] = o
	case *NoticeWinphone:
		this.Notification[string(WINPHONE)] = o
	case *NoticeIos:
		this.Notification[string(IOS)] = o
	case *NoticeSimple:
		this.Notification["alert"] = o.(*NoticeSimple).Alert
	}
}

//------------------MessageAndNoticeBuilder------------------
func NewMessageAndNoticeBuilder() *MessageAndNoticeBuilder {
	mnb := &MessageAndNoticeBuilder{}
	mnb.Options = NewOptions()
	return mnb
}

func (this *MessageAndNoticeBuilder) SetPlatform(pf *Platform) {
	this.Platform = pf.Object
}

func (this *MessageAndNoticeBuilder) SetAudience(ad *Audience) {
	this.Audience = ad.Object
}

func (this *MessageAndNoticeBuilder) SetOptions(o *Options) {
	this.Options = o
}

func (this *MessageAndNoticeBuilder) SetMessage(m *Message) {
	this.Message = m
}

/*
可以为每类Notice设置一个
*/
func (this *MessageAndNoticeBuilder) SetNotice(o interface{}) {
	if this.Notification == nil {
		this.Notification = make(map[string]interface{})
	}
	switch o.(type) {
	case NoticeAndroid:
		this.Notification[string(ANDROID)] = o
	case NoticeWinphone:
		this.Notification[string(WINPHONE)] = o
	case NoticeIos:
		this.Notification[string(IOS)] = o
	case NoticeSimple:
		this.Notification["alert"] = o.(NoticeSimple).Alert
	}

}
