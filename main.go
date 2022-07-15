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

    var currentApm float64

    events := k.Read()
    counter := ratecounter.NewRateCounter(60 * time.Second)
    wallpaperApp := "feh"
    argCenter := "--bg-center"
    argOneScreen := "--no-xinerama"
    argColor := "--image-bg"
    pixelWP := "/home/guenther/Projects/lethani/wallpapers/pixel.jpg"
    // cannot use map, as map iteration is in random order
    apms := []float64{100, 115, 130, 150, 180, 220, 250, 300}
    colors := []string{"#005F73", "#0A9396", "#94D2BD", "#EE9B00", "#CA6702", "#BB3E03", "#AE2012", "#9B2226"}

    for e := range events {
        switch e.Type {

        case keylogger.EvKey:

            if e.KeyPress() {
                counter.Incr(1)
                logrus.Println("60s-counterRate = ", counter.Rate())
                currentApm = float64(counter.Rate())

                for i, apm := range apms {
                    if ( currentApm < apm ) {
                        exec.Command(wallpaperApp, argCenter, argOneScreen, argColor, colors[i], pixelWP).Output()
                        break
                    }
                }
            }

            break
        }
    }
}
