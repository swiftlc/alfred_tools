package cmd

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "alfred_tools",
		Short: "alfred_tools",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

var wf *aw.Workflow

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err:%+v\n", err)
		}
	}()

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.alfred_tools.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.SetDefault("name", "zhangsan")

	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(vscodeOpenCmd)
	rootCmd.AddCommand(weeklyCmd)

	wf = aw.New()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".alfred_tools")
	}

	//环境变量需要全大写
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}
