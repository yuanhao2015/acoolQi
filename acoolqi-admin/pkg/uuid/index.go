/*
@Time : 2021-10-15 16:47
@Author : acool
@File : index
*/
package uuid

import (
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
)

// 生成存入redis的key
func CreateRedisUuid(name string) (string, string) {
	id := uuid.NewV4()
	return fmt.Sprintf("users-%s-%s", name, strings.ReplaceAll(id.String(), "-", "")), strings.ReplaceAll(id.String(), "-", "")
}
