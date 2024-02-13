package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Download(songId string, query string) {
	args := []string{"-f", "bestaudio", "-x", "--audio-format", "mp3", "--audio-quality", "320k", "--add-metadata", "--output", "" + songId + ".%(ext)s", "" + YtResource + "/search?q=" + query + "#songs", "--playlist-items", "1"}
	cmd := exec.Command("./yt-dlp.exe", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println("Result: " + out.String())
}
