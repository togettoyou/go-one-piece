package tools

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

// UUID 生成UUID
// is是否去除-符号 默认去除
func UUID(is ...bool) string {
	if len(is) > 0 && is[0] {
		return fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), nil))
	}
	return strings.Replace(fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), nil)), "-", "", -1)
}
