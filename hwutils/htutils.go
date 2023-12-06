package htutils

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"strings"
)

/**
 * 成功的返回
 */
func Success(data interface{}) map[string]interface{} {
	return g.Map{
		"code": 200,
		"msg":  "ok",
		"data": data,
	}
}

/**
 * 失败的返回
 */

func Fail(args ...interface{}) map[string]interface{} {
	code := 500
	msg := "error"
	data := g.Map{}
	for _, v := range args {
		switch t := v.(type) {
		case string:
			msg = t
		case int:
			code = t
		case g.Map:
			data = t
		default:
		}
	}
	return g.Map{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

/**
** 判断字符串是否在数组中存在
 */
func InStringArray(needle string, args []string) bool {
	for i := 0; i < len(args); i++ {
		if needle == args[i] {
			return true
		}
	}
	return false
}

/**
**返回一个MD5字符串
 */
func GenerateMD5Str(prefix string) string {

	if prefix != "" {
		return strings.ToUpper(prefix + guid.S())
	} else {
		return strings.ToUpper(guid.S())
	}
}

func FormatTime(timestamp int64) string {
	time := gtime.NewFromTimeStamp(timestamp)
	return time.Format("2006-01-02 15:04:05")
}
