package util

import (
    "github.com/gin-gonic/gin"
    "github.com/unknwon/com"

    "github.com/AOaixuexi/go-gin-example-follow/pkg/setting"
)

// GetPage gets page from the request.
func GetPage(c *gin.Context) int {
    result := 0
    page, _ := com.StrTo(c.Query("page")).Int()
    if page > 0 {
        result = (page - 1) * setting.PageSize
    }

    return result
}