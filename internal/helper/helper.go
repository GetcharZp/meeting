package helper

import uuid "github.com/satori/go.uuid"

// GetUUID
// 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
}
