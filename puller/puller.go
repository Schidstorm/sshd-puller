package puller

import (
	"context"
	"fmt"
	api2 "github.com/schidstorm/sshd-puller/api"
	"github.com/schidstorm/sshd-puller/config"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func RunLoop(ctx context.Context, cfg *config.Config) error {
	var counter uint64 = 0
	counter = attemptRun(ctx, cfg, counter)
	isDone := false

	for !isDone {
		select {
		case <-ctx.Done():
			isDone = true
		case <-time.After(cfg.LoopTime):
			counter = attemptRun(ctx, cfg, counter)
		}
	}

	return nil
}

func attemptRun(ctx context.Context, cfg *config.Config, counter uint64) uint64 {
	failedAttempts := 0
	for {
		select {
		case <-ctx.Done():
			return counter
		case <-time.After(cfg.RetryTimeout):
			err := Run(ctx, cfg, counter)
			counter++
			if err != nil {
				failedAttempts++
				logrus.Errorln(err, map[string]string{
					"failedAttempts": fmt.Sprintf("%d", failedAttempts),
				})
				if failedAttempts >= cfg.Tries {
					logrus.Errorln(err, "Giving up")
					return counter
				}
			} else {
				failedAttempts = 0
				return counter
			}
		}
	}
}

func Run(ctx context.Context, cfg *config.Config, counter uint64) error {
	api := &api2.Api{Endpoint: cfg.Endpoints[counter%uint64(len(cfg.Endpoints))]}
	keys, err := api.GetKeys(ctx, cfg.ServerKey)
	if err != nil {
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
