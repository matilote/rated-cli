package core

import (
	"time"
)

// Config represents the configuration of rated CLI.
type Config struct {
	ApiEndpoint           string        // endpoint to Rated network API
	WatcherRefreshRate    time.Duration // refresh rate of the watcher
	WatcherValidationKeys []string      // validation keys to watch
}
