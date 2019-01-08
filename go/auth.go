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
func QqAuth(accessToken, openId string) (bool, error) {
	_, resp, err := xhttp.GetJSON(fmt.Sprintf(QqAuthApi, accessToken))
	if err != nil {
		return false, err
	}
	data := resp["openid"]
	if data == nil {
		return false, fmt.Errorf("QqAuth: qq_open_id is nil, %+v", resp)
	}
	qqOpenId := data.(string)
	if openId != qqOpenId {
		return false, fmt.Errorf("QqAuth: qq_open_id not match, %s != %s", openId, qqOpenId)
	}
	return true, nil
}

// WxAuth
func WxAuth(accessToken, openId string) (bool, error) {
	_, data, err := xhttp.GetJSON(fmt.Sprintf(WxAuthApi, accessToken, openId))
	if err != nil {
		return false, err
	}
	_, has_err := data["errcode"]
	if has_err {
		errcode := int(data["errcode"].(float64))
		if errcode != 0 {
			return false, fmt.Errorf("WxAuth: wx oauth error, errcode-> %d", errcode)
		}
	}
	return true, nil
}
