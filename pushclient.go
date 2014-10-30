package jpush

import (
	"encoding/base64"
	"encoding/json"
	//"fmt"
	"github.com/isdamir/jpush/common"
)

const (
	SUCCESS_FLAG  = "msg_id"
	HOST_NAME_SSL = "https://api.jpush.cn/v3/push"
	BASE64_TABLE  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var base64Coder = base64.NewEncoding(BASE64_TABLE)

type PushClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewPushClient(secret, appKey string) *PushClient {
	//base64
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &PushClient{secret, appKey, auth, HOST_NAME_SSL}
	return pusher
}

func (this *PushClient) Send(builder interface{}) (*common.RetData, error) {

	content, err := json.Marshal(builder)
	//fmt.Println(string(content))
	if err != nil {
		return nil, err
	}

	return this.SendPushBytes(content)
}

func (this *PushClient) SendPushString(content string) (ret *common.RetData, err error) {
	ret, err = common.SendPostString(this.BaseUrl, content, this.AuthCode)
	return
}

func (this *PushClient) SendPushBytes(content []byte) (ret *common.RetData, err error) {
	ret, err = common.SendPostBytes(this.BaseUrl, content, this.AuthCode)
	return
}
