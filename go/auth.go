package openauth

import (
	"fmt"
	"strconv"

	"github.com/usthooz/gutil/http"
)

const (
	// QqAuthApi QQ授权api
	QqAuthApi = "https://graph.qq.com/oauth2.0/me?access_token=%s"
	// WxAuthApi 微信授权api
	WxAuthApi = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid==%s"
	// SinaAuthApi 新浪授权api
	SinaAuthApi = "https://api.weibo.com/2/account/get_uid.json?access_token=%s&source=%s"
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
		return false, fmt.Errorf("QqAuth: qq_open_id not match-> openid: %s != resultOpenId: %s", openId, qqOpenId)
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

// SinaAuth
func SinaAuth(accessToken, sinaUid string) (bool, error) {
	_, data, err := xhttp.GetJSON(fmt.Sprintf(SinaAuthApi, accessToken, sinaUid))
	if err != nil {
		return false, err
	}
	_, has_err := data["errcode"]
	if has_err {
		errcode := int(data["errcode"].(float64))
		if errcode != 0 {
			return false, fmt.Errorf("SinaAuth: sina auth error, errcode-> %d", errcode)
		}
	}
	resultSinaUid := strconv.FormatFloat(data["uid"].(float64), 'f', 0, 64)
	if sinaUid != resultSinaUid {
		return false, fmt.Errorf("SinaAuth: sina_uid not match-> sina_uid: %s != resultSinaUid: %s", sinaUid, resultSinaUid)
	}
	return true, nil
}

// XiaomiAuth
func XiaomiAuth(accessToken, openId, appId string) (bool, error) {
	_, data, err := xhttp.GetJSON(fmt.Sprintf(XiaomiAuthApi, accessToken, appId))
	if err != nil {
		return false, err
	}
	result := data["result"].(string)
	if result == "error" {
		err_code := int(data["code"].(float64))
		err_msg := data["description"].(string)
		return false, fmt.Errorf("XiaomiAuth: xiaomi auth error, errorcode-> %d, err_description-> %v", err_code, err_msg)
	}
	openIDMap := data["data"].(map[string]interface{})
	openIDStr := openIDMap["openId"].(string)
	if openId != openIDStr {
		return false, fmt.Errorf("XiaomiAuth: xiaomi_open_id not match-> %s != %s", openId, openIDStr)
	}
	return true, nil
}
