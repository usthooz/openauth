package openauth

import (
	"fmt"

	"github.com/usthooz/gutil/http"
)

const (
	// QqAuthApi QQ授权api
	QqAuthApi = "https://graph.qq.com/oauth2.0/me?access_token=%s"
	// WxAuthApi 微信授权api
	WxAuthApi = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid==%s"
	// SinaAuthApi 新浪授权api
	SinaAuthApi = "https://api.weibo.com/2/account/get_uid.json?access_token=%s"
	// XiaomiAuthApi 小米授权api
	XiaomiAuthApi = "https://open.account.xiaomi.com/user/openidV2?token=%s&clientId=%s"
)

// QqAuth
func QqAuth(qqAccessToken, qqOpenId string) (bool, error) {
	_, resp, err := xhttp.GetJSON("https://graph.qq.com/oauth2.0/me?access_token=" + qqAccessToken)
	if err != nil {
		return false, err
	}
	data := resp["openid"]
	if data == nil {
		return false, fmt.Errorf("QqAuth: qq_open_id is nil, %+v", resp)
	}
	openId := data.(string)
	if qqOpenId != openId {
		return false, fmt.Errorf("QqAuth: qq_open_id not match, %s != %s", qqOpenId, openId)
	}
	return true, nil
}
