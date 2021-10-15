/*
@Time : 2021-10-14 11:47
@Author : acool
@File : index
*/
package times

import (
	"time"
)

var (
	//全局时间格式带秒
	FormatTimeWithSecond = "2006-01-02 15:04:05"

	//全局时间格式带分
	FormatTimeWithMinu = "2006-01-02 15:04"
)

//获取相差时间
func GetHourDiffer(start_time, end_time string) int64 {
	t1, err := time.ParseInLocation(FormatTimeWithSecond, start_time, time.Local)
	t2, err := time.ParseInLocation(FormatTimeWithSecond, end_time, time.Local)
	if err == nil && t1.Before(t2) {
		return t2.Unix() - t1.Unix()
	} else {
		return t1.Unix() - t2.Unix()
	}
}
