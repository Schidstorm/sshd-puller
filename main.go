package main

import (
	"github.com/ghodss/yaml"
	api2 "github.com/schidstorm/sshd-puller/api"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	ServerKey string `yaml:"serverKey"`
	Endpoint string `yaml:"endpoint"`
	AuthorizedKeysFile string `yaml:"authorizedKeysFile"`
}

func main() {
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

			config := &Config{}
			err = yaml.Unmarshal(configFileData, config)
			if err != nil {
				logrus.Errorln(err)
				return err
			}

			api := &api2.Api{Endpoint: config.Endpoint}
			keys, err := api.GetKeys(config.ServerKey)
			if err != nil {
				logrus.Errorln(err)
				return err
			}

			var fileMode os.FileMode
			stats, err := os.Stat(config.AuthorizedKeysFile)
			if err != nil {
				fileMode = 0600
			} else {
				fileMode = stats.Mode()
			}

			err = ioutil.WriteFile(config.AuthorizedKeysFile, []byte(strings.Join(keys, "\n")), fileMode)
			if err != nil {
				logrus.Errorln(err)
				return err
			}

			return nil
		},
	}

	rootCmd.PersistentFlags().String("config", "/etc/sshd/puller.yml", "Config file")
	err := rootCmd.Execute()
	if err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
}