package main

import (
	"os"
	"time"
)

//WatchFile watches file
func WatchFile(filePath string) error {
	initialStat, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	for {
		stat, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			break
		}

		time.Sleep(5 * time.Second)
	}

	return nil
}
