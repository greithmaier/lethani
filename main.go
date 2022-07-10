package main

import (
    "time"
    "github.com/MarinX/keylogger"
    "github.com/sirupsen/logrus"
    "github.com/paulbellamy/ratecounter"
    "github.com/reujab/wallpaper"
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
    counter := ratecounter.NewRateCounter(15 * time.Second)

    background, err := wallpaper.Get()
    logrus.Println("WP", background)

    for e := range events {
        switch e.Type {

        case keylogger.EvKey:

            if e.KeyPress() {
                counter.Incr(1)
                apm = float64(counter.Rate()) * 4
                logrus.Println("APM", apm)

                if ( apm < 70 ) {
                    logrus.Println("All good", apm)
                    err = wallpaper.SetFromFile("/usr/share/backgrounds/gnome/adwaita-day.jpg")

                } else {
                    logrus.Println("Chill my dude <3", apm)

                }
            }

            break
        }
    }
}
