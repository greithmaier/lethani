package main

import (
    "time"
    "github.com/MarinX/keylogger"
    "github.com/sirupsen/logrus"
    "github.com/paulbellamy/ratecounter"
    "os/exec"
)

func main() {

    keyboard := keylogger.FindKeyboardDevice()

    logrus.Println("Found a keyboard at", keyboard)
    k, err := keylogger.New(keyboard)
    if err != nil {
        logrus.Error(err)
        return
    }
    defer k.Close()

    var apm float64

    events := k.Read()
    counter := ratecounter.NewRateCounter(12 * time.Second)
    wallpaperApp := "feh"
    argScale := "--bg-scale"
    wallpaper1 := "~/Projects/lethani/wallpapers/1.jpg"
    wallpaper2 := "~/Projects/lethani/wallpapers/2.jpg"

    for e := range events {
        switch e.Type {

        case keylogger.EvKey:

            if e.KeyPress() {
                counter.Incr(1)
                apm = float64(counter.Rate()) * 5

                if ( apm < 120 ) {
                    logrus.Println("All good", apm)
                    exec.Command("bash", "-c", wallpaperApp, argScale, wallpaper1).Output()

                    //out, err := exec.Command("/bin/sh", "-c", "echo", "wp1").Output()
                    //output := string(out[:])
                    //logrus.Println(output)
                    //logrus.Println(out, err)

                } else {
                    logrus.Println("Chill my dude <3", apm)
                    exec.Command("bash", "-c", wallpaperApp, argScale, wallpaper2).Output()
                }
            }

            break
        }
    }
}
