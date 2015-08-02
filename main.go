package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/darshandzend/yt2audio/yt"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//get the video metdata
	video, err := yt.Get("tyy2fJiOZDE")
	check(err)

	fmt.Println(video)

	//download the video and write to file
	video.Download(0, "video.mp4")

	//convert to mp3
	avconvPath, pathErr := exec.LookPath("avconv")
	check(pathErr)

	args := []string{"avconv", "-i", "video.mp4", "audio.mp3"}

	env := os.Environ()
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
	}()
	fmt.Println(avconvPath, args)
	avconvErr := syscall.Exec(avconvPath, args, env)
	check(avconvErr)

}
