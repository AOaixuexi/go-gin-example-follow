package util

import "github.com/AOaixuexi/go-gin-example-follow/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}