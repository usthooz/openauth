# openauth
主流第三方登录Token授权校验.

### 功能
用于客户端第三方授权登录后，服务器进行进一步的授权校验。

### 平台
- QQ
- 微信
- 新浪
- 小米
- Google
- Facebook

### 使用

```go
go get github.com/usthooz/openauth
```

#### Example

```go
package openauth

import (
	"github.com/usthooz/openauth/facebook"
	"github.com/usthooz/openauth/google"
	"testing"
)

func TestQqAuth(t *testing.T) {
	success, err := QqAuth("", "")
	t.Logf("Success: %v, Err: %v", success, err)
}

func TestSinaAuth(t *testing.T) {
	success, err := SinaAuth("", "")
	t.Logf("Success: %v, Err: %v", success, err)
}

func TestWxAuth(t *testing.T) {
	success, err := WxAuth("", "")
	t.Logf("Success: %v, Err: %v", success, err)
}

func TestXiaomiAuth(t *testing.T) {
	success, err := XiaomiAuth("", "", "")
	t.Logf("Success: %v, Err: %v", success, err)
}

func TestGoogleVerify(t *testing.T) {
	result, err := google.GoogleVerify("", "")
	t.Logf("Success: %v, Err: %v", result, err)
}

func TestFacebookVerify(t *testing.T) {
	result, err := facebook.FacebookVerify("", "", "")
	t.Logf("Success: %v, Err: %v", result, err)
}
```