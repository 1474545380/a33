package help

import (
	"crypto/md5"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func GetUUID() string {
	return uuid.NewV4().String()
}

// md5生成密码
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
