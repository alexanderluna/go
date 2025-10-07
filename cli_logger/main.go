package main

import "cli_logger/pocketlog"

func main() {
	lgr := pocketlog.New(pocketlog.LevelDebug)

	lgr.Debugf("Here is some helpful information")
	lgr.Infof("Here is something you might not have known")
	lgr.Warnf("This is causing some odd behavior")
	lgr.Errorf("This is fatal to the current operation")
	lgr.Fatalf("This caused the app or service to shut down")

}
