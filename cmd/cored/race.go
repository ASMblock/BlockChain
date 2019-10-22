//+build race

package main

/*

Each build of cored also produces a race-enabled binary.
See builds/core/build for the details.

To use this facility, log in to a machine, stop the
normal binary via upstart, and start the race-enabled
binary directly by loading the config vars and running
command 'cored_race'. Race reports are printed to stderr.
See https://golang.org/doc/articles/race_detector.html
for more info.

*/

func init() {
	race = []interface{}{"race", true}
}
