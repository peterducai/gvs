package services

import (
	"os"
	"time"
)

//Item in watchlist
type Item struct {
	path          string
	lastseen      string //TODO: figure type
	sleepseconds  int
	sleeptreshold int
}

//Watchlist list of paths to watch
type Watchlist struct {
	Items []Item
}

//WatchManager change sleep var depending on how often file changes
func WatchManager() {

}

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

//WatchDir watch directory
func WatchDir(path string) {

}
