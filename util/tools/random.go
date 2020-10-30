package tools

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

//生成UUID
//is是否去除-符号 默认去除
func UUID(is ...bool) string {
	if len(is) > 0 && is[0] {
		return fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), nil))
	}
	return strings.Replace(fmt.Sprintf("%s", uuid.Must(uuid.NewV4(), nil)), "-", "", -1)
}

//生成指定长度的随机字符串
func RangeString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//生成指定长度的随机数字
func RangeCode(length int) string {
	var container string
	for i := 0; i < length; i++ {
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		container += fmt.Sprintf("%01v", rnd.Int31n(10))
	}
	return container
}

//生成指定范围内随机值
//[min,max)
func RangeNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return randNum
}
