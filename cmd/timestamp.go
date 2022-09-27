package cmd

import (
	"fmt"
	"time"

	"github.com/alfred_tools/alfred"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间转换工具",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			wf.SendFeedback()
		}()
		var query string
		if len(args) > 0 {
			query = args[0]
		}

		switch query {
		case "now": //返回当前时间
			now := time.Now().Unix()
			nowStr := alfred.FormateTime(now)
			wf.NewItem(fmt.Sprintf("时间：%s", nowStr)).Arg(nowStr).Valid(true)
			wf.NewItem(fmt.Sprintf("时间戳：%d", now)).Arg(fmt.Sprint(now)).Valid(true)
		default:
			val, err := cast.ToInt64E(query)
			if err == nil {
				nowStr := alfred.FormateTime(val)
				wf.NewItem(fmt.Sprintf("时间：%s", nowStr)).Arg(nowStr).Valid(true)
			} else {
				tm, err := time.ParseInLocation("2006-01-02 15:04:05", query, alfred.CstTimezone)
				if err != nil {
					wf.Fatal("输入参数(now/当前时间(2006-01-02 15:04:05)/时间戳)")
				} else {
					wf.NewItem(fmt.Sprintf("时间戳：%d", tm.Unix())).Arg(fmt.Sprint(tm.Unix())).Valid(true)
				}
			}
		}
	},
}
