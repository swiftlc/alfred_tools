package cmd

import (
	"fmt"
	"strconv"

	"github.com/alfred_tools/alfred"
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

		v, err := strconv.ParseInt(query, 10, 64)
		if err != nil || v < 0 {
			wf.Fatal("输入格式有误")
			return
		}

		t := alfred.FormateTime(v)

		wf.NewItem(fmt.Sprintf("时间：%s", t)).Arg(t).Valid(true)
		wf.NewItem(fmt.Sprintf("时间戳：%d", v)).Arg(fmt.Sprint(v)).Valid(true)
	},
}
