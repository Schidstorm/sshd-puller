package cli

import (
	"github.com/schidstorm/sshd-puller/config"
	"github.com/schidstorm/sshd-puller/puller"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Run() error {
	rootCmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			configFilePath, err := cmd.PersistentFlags().GetString("config")
			if err != nil {
				logrus.Errorln(err)
				return err
			}
			configFileData, err := ioutil.ReadFile(configFilePath)
			if err != nil {
				logrus.Errorln(err)
				return err
			}

			cfg := &config.Config{}
			err = yaml.Unmarshal(configFileData, cfg)
			if err != nil {
				logrus.Errorln(err)
				return err
			}

			return puller.Run(cfg)
		},
	}

	rootCmd.PersistentFlags().String("config", "/etc/sshd/puller.yml", "Config file")
	err := rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}