package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ConfigDefaults(cmdList ...*cobra.Command) {
	ConfigEnv()
	ConfigSet()
	ConfigFlags(cmdList...)
}

func ConfigSet() {
	for _, item := range ConfigOptions {
		SetDefault(item.(map[string]interface{}))
	}
}

func SetDefault(flagMap map[string]interface{}) {
	for index, item := range flagMap {
		opt := item.(map[string]interface{})
		_, ok := opt["value"]
		if !ok {
			SetDefault(opt)
			continue
		}
		viper.SetDefault(index, opt["value"])
	}
}

func ConfigFlags(cmdList ...*cobra.Command) {
	for _, cmd := range cmdList {
		for _, item := range ConfigOptions[cmd.Use].(map[string]interface{}) {
			BindFlags(cmd, item.(map[string]interface{}))
		}
	}
}

func BindFlags(cmd *cobra.Command, flagMap map[string]interface{}) {
	for index, item := range flagMap {
		opt := item.(map[string]interface{})
		_, ok := opt["help"]
		if !ok {
			BindFlags(cmd, opt)
			continue
		}
		help := opt["help"].(string)
		switch value := opt["value"].(type) {
		case int:
			cmd.Flags().Int(index, value, help)
		case bool:
			cmd.Flags().Bool(index, value, help)
		case string:
			cmd.Flags().String(index, value, help)
		case float32:
			cmd.Flags().Float32(index, value, help)
		case float64:
			cmd.Flags().Float64(index, value, help)
		}
	}
}

func ConfigBindFlags(cmd *cobra.Command) {
	for index, _ := range ConfigOptions[cmd.Use].(map[string]interface{}) {
		viper.BindPFlag(index, cmd.Flags().Lookup(index))
	}
}

func ConfigEnv() {
	viper.SetEnvPrefix("authit")
	viper.AutomaticEnv()
}
