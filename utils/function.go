package utils

import "time"

// GetFormatTime 获取格式化时间，如：2006-01-02 15:04:05
func GetFormatTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
