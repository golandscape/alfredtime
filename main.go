package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Items struct {
	Rerun     float64 `json:"rerun"`
	Variables string  `json:"variables"`
	Items     []Item  `json:"items"`
}
type Item struct {
	Uid      string `json:"uid"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

func main() {
	var input = func() string {
		if len(os.Args) == 2 {
			return os.Args[1]
		}
		return ""
	}()

	if input != "" { //如果有输入内容
		atoi, err := strconv.Atoi(input)
		if err != nil { //不是数字
			parse, err := time.Parse("2006-01-02 15:04:05", input)
			if err != nil { //解析错误直接退出
				output("error", "", "格式:2006-01-02 15:04:05")
				return
			}
			ts := fmt.Sprintf("%d", parse.Unix())
			output(ts, ts, "复制")
		} else { //数字
			ts := time.Unix(int64(atoi), 0).Format("2006-01-02 15:04:05")
			output(ts, ts, "复制")
		}
		return
	}
	unix := time.Now().Unix()
	ts := fmt.Sprintf("%d", unix)
	output(ts, ts, "复制")
}
