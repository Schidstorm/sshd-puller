package puller

import (
	api2 "github.com/schidstorm/sshd-puller/api"
	"github.com/schidstorm/sshd-puller/config"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

func Run(cfg *config.Config) error {
	api := &api2.Api{Endpoint: cfg.Endpoint}
	keys, err := api.GetKeys(cfg.ServerKey)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	var fileMode os.FileMode
	stats, err := os.Stat(cfg.AuthorizedKeysFile)
	if err != nil {
		fileMode = 0600
	} else {
		fileMode = stats.Mode()
	}

	err = ioutil.WriteFile(cfg.AuthorizedKeysFile, []byte(strings.Join(keys, "\n")), fileMode)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}
