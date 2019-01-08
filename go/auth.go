package openauth

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
