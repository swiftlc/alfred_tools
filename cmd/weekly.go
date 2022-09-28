package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/alfred_tools/alfred"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var weeklyCmd = &cobra.Command{
	Use:   "weekly",
	Short: "周报",
	Run: func(cmd *cobra.Command, args []string) {
		var query string
		if len(args) > 0 {
			query = args[0]
		}

		query = strings.TrimSpace(query)
		now := time.Now()
		weekday := now.Weekday()
		if weekday == time.Sunday {
			weekday = 7
		}
		thisWeekMonday := alfred.FormateTimeDate(now.AddDate(0, 0, int(time.Monday-weekday)).Unix())
		weeklyDir := viper.GetString("WEEKLY_DIR")
		if weeklyDir == "" {
			weeklyDir = "."
		}

		filepath := path.Join(weeklyDir, fmt.Sprintf("%s.weekly", thisWeekMonday))
		f, _ := os.Open(filepath)
		var data map[string][]string
		if f != nil {
			defer f.Close()
			json.NewDecoder(f).Decode(&data)
		}

		if data == nil {
			data = make(map[string][]string)
		}

		if query == "8F7291E6-6B8E-496F-B1D4-B4540456D4BD" {
			for k, v := range data {
				fmt.Println(k)
				for _, item := range v {
					fmt.Printf("\t%s\n", item)
				}
			}
			return
		}

		param := strings.SplitN(query, " ", 2)

		if len(param) != 2 {
			fmt.Printf("invalid param:%s", query)
			return
		}

		key := strings.TrimSpace(param[0])
		val := strings.TrimSpace(param[1])

		data[key] = append(data[key], val)

		d, _ := json.Marshal(data)

		ioutil.WriteFile(filepath, d, 0664)

		fmt.Print("周报记录成功")
	},
}
