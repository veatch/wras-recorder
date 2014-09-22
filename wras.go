package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var schedule = map[string][]string{
	// day-hour: {duration, show name}
	"*-6":  {"120", "Rotation"},
	"0-14": {"120", "Melodically-Challenged"},
	"1-20": {"120", "Georgia-Music-Show"},
	"3-20": {"120", "Distance+Lights"},
	"4-22": {"120", "Mighty-Aphrodite"},
	"6-12": {"120", "Deviltown"},
	"6-16": {"120", "QCS"},
	"6-20": {"120", "Soul-Kitchen"},
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
	filename := fmt.Sprintf("WRAS-%v-%v.mp3", now.Format("01-02-06"), show_name)
	err = os.Rename("WRAS2.mp3", filename)
	if err != nil {
	}
	new_path := fmt.Sprintf("/www/veatch/wras/%v", filename)
	err = os.Link(filename, new_path)
	if err != nil {
	}
	err = os.Chdir("/www/veatch/wras/")
	if err != nil {
	}
	files, _ := filepath.Glob("*.mp3")
	body, _ := ioutil.ReadFile("head")
	for i := range files {
		link := fmt.Sprintf("<a href=\"/wras/%v\">%v</a><br/>\n", files[i], files[i])
		body = append(body, []byte(link)...)
	}
	ioutil.WriteFile("/www/veatch/wras/index.html", body, 0644)
}
