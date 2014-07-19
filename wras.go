package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var schedule = map[string][]string{
	// day-hour: {duration, show name}
	"*-6":  {"120", "Rotation"},
	"1-20": {"120", "Georgia-Music-Show"},
	"3-20": {"120", "Distance+Lights"},
	"4-22": {"120", "Mighty-Aphrodite"},
	"6-12": {"120", "Deviltown"},
	"6-16": {"120", "QCS"},
	"6-20": {"120", "Soul-Kitchen"},
	"7-14": {"120", "Melodically-Challenged"},
}

func main() {
	now := time.Now()
	hour := now.Hour()
	day := now.Weekday()
	show_info := schedule[fmt.Sprintf("*-%d", hour)]
	if show_info == nil {
		show_info = schedule[fmt.Sprintf("%d-%d", day, hour)]
		if show_info == nil {
			return
		}
	}
	duration := fmt.Sprintf("%vmin", show_info[0])
	show_name := show_info[1]
	err := exec.Command("icecream", "--stop", duration, "http://www.publicbroadcasting.net/wras/ppr/wras2.m3u").Run()
	if err != nil {
	}
	filename := fmt.Sprintf("WRAS-%v-%v.mp3", show_name, now.Format("1-2-06"))
	err = os.Rename("WRAS2.mp3", filename)
	if err != nil {
	}
}
