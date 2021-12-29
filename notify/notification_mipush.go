package notify

import (
	"fmt"
	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/logx"
	"github.com/xxgail/PushSDK"
)

func InitMiPushClient(cfg *config.ConfYaml, appSecret string, pkgName string) *PushSDK.Send {
	send := PushSDK.NewSend()
	send.SetChannel("mi")
	send.SetPlatForm(fmt.Sprintf("{\"app_secret\":\"%s\",\"restricted_package_name\":\"%s\"}", appSecret, pkgName))
	return send
}

// PushToMi provide send notification to Android server.
func PushToMiPush(req *PushNotification, cfg *config.ConfYaml) (resp *ResponsePush, err error) {
	logx.LogAccess.Debug("Start push notification for Mi")

	//req.HuaweiCollapseKey

	return resp, nil
}
