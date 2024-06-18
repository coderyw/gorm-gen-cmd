package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm_gen_cmd/generate"
	"gorm_gen_cmd/model"
	"os"
)

var cfgFile string
var cfg *model.GenCfg

// 构建根 command 命令。前面我们介绍它还可以有子命令，这个command里没有构建子命令
var rootCmd = &cobra.Command{
	Use:   "gorm-gen",
	Short: "gorm从数据库生成model代码",
	Long:  `gorm从数据库生成model代码,支持设置数据库，多表名。同时生成dao基础代码`,
	Run: func(cmd *cobra.Command, args []string) {
		generate.GenFunc(cfg)
	},
}

// 执行 rootCmd 命令并检测错误
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// 加载运行初始化配置
	cobra.OnInitialize(initConfig)
	// rootCmd，命令行下读取配置文件，持久化的 flag，全局的配置文件
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件 (默认配置文件在 $HOME/.gorm_gen.yaml)")
	// local flag，本地化的配置
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().String("host", "", "数据库地址(不需要携带port)")
	rootCmd.PersistentFlags().String("port", "", "数据库端口")
	rootCmd.PersistentFlags().String("auth", "", "登录权限配置,格式为：用户名:密码")

	rootCmd.PersistentFlags().String("database", "", "需要生成model的数据库")
	rootCmd.PersistentFlags().String("tables", "", "需要生成model的表，支持多个，使用','隔开")
	rootCmd.PersistentFlags().String("outpath", "", "dao和model两个目录存放的path")

}

// 初始化配置的一些设置
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) // viper 设置配置文件
	} else { // 上面没有指定配置文件，下面就读取 home 下的 .gorm_gen.yaml文件
		// 配置文件参数设置
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gorm_gen")
	}

	// 解析flags到vipper
	f := rootCmd.Flags()
	if err := viper.BindPFlags(f); err != nil {
		cobra.CheckErr(err)
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil { // 读取配置文件
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
	cfg = new(model.GenCfg)
	if err := viper.Unmarshal(cfg); err != nil {

		cobra.CheckErr(err)
	}
	if len(cfg.Outpath) == 0 {
		path, err := os.Getwd()
		if err != nil {
			cobra.CheckErr(err)
		}
		cfg.Outpath = path
		rootCmd.Println("outpath:", cfg.Outpath)
	}
}
