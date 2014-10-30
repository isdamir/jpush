package jpush

type Options struct {
	/*
		纯粹用来作为 API 调用标识，API 返回时被原样返回，以方便 API 调用方匹配请求与返回。
	*/
	Sendno int `json:"sendno,omitempty"`
	/*
		推送当前用户不在线时，为该用户保留多长时间的离线消息，以便其上线时再次推送。默认 86400 （1 天），最长 10 天。设置为 0 表示不保留离线消息，只有推送当前在线的用户可以收到。
	*/
	Timelive int `json:"time_to_live,omitempty"`
	/*
		True 表示推送生产环境，False 表示要推送开发环境； 如果不指定则为推送生产环境。

		(消息) JPush 官方 API LIbrary (SDK) 默认设置为推送 “开发环境”。
	*/
	Apns_production bool `json:"apns_production,omitempty"`
	/*
		如果当前的推送要覆盖之前的一条推送，这里填写前一条推送的 msg_id 就会产生覆盖效果，即：1）该 msg_id 离线收到的消息是覆盖后的内容；2）即使该 msg_id Android 端用户已经收到，如果通知栏还未清除，则新的消息内容会覆盖之前这条通知；覆盖功能起作用的时限是：1 天。
		如果在覆盖指定时限内该 msg_id 不存在，则返回 1003 错误，提示不是一次有效的消息覆盖操作，当前的消息不会被推送。
	*/
	OverrideMsgID float32 `json:"override_msg_id,omitempty"`
	/*
		又名缓慢推送，把原本尽可能快的推送速度，降低下来，在给定的 n 分钟内，均匀地向这次推送的目标用户推送。最大值为 1440。未设置则不是定速推送。
	*/
	BigPushDuration int `json:"big_push_duration,omitempty"`
}

func NewOptions() *Options {
	op := &Options{}
	op.Timelive = 86400
	return op
}
