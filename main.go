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
    counter := ratecounter.NewRateCounter(30 * time.Second)
    wallpaperApp := "feh"
    argCenter := "--bg-center"
    argOneScreen := "--no-xinerama"
    argColor := "--image-bg"
    pixelWP := "/home/guenther/Projects/lethani/wallpapers/pixel.jpg"

    for e := range events {
        switch e.Type {

        case keylogger.EvKey:

            if e.KeyPress() {
                counter.Incr(1)
                logrus.Println("30s-counterRate * 2 = ", counter.Rate() * 2)
                apm = float64(counter.Rate()) * 2

                // TBD: refactor into a very nice array
                // TBD: add mouse action tracking
                // TBD: show APM again
                // TBD: more intelligent APM attribution!

                if ( apm < 140 ) {
                    exec.Command(wallpaperApp, argCenter, argOneScreen, argColor, "#005F73", pixelWP).Output()
                } else if ( apm < 160 ) {
                    exec.Command(wallpaperApp, argCenter, argOneScreen, argColor, "#0A9396", pixelWP).Output()
                } else if ( apm < 180 ) {
                    exec.Command(wallpaperApp, argCenter, argOneScreen, argColor, "#94D2BD", pixelWP).Output()
                } else if ( apm < 200 ) {
                    exec.Command(wallpaperApp, argCenter, argOneScreen, argColor, "#EE9B00", pixelWP).Output()
                } else if ( apm < 220 ) {
                    exec.Command(wallpaperApp, argCenter, argOneScreen, argColor, "#CA6702", pixelWP).Output()
                } else if ( apm < 240 ) {
                    exec.Command(wallpaperApp, argCenter, argOneScreen, argColor, "#BB3E03", pixelWP).Output()
                } else if ( apm < 260 ) {
                    exec.Command(wallpaperApp, argCenter, argOneScreen, argColor, "#AE2012", pixelWP).Output()
                } else {
                    exec.Command(wallpaperApp, argCenter, argOneScreen, argColor, "#9B2226", pixelWP).Output()
                }
            }

            break
        }
    }
}
