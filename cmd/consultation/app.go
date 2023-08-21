package main

import (
	"HeidiTask/internal"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func init() {

}

type App struct {
}

// TODO: LocalAddr error when use a container as kafka host name.
func (app *App) run() {
	reader := internal.NewReader()
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("kafka ReadMessage failed:", err)
		}

		events := internal.ProfileUpdateEvents{}
		err = json.Unmarshal(msg.Value, &events)
		if err != nil {
			log.Fatal("parse event data failed:", err)
		}
		events.Sort()

		for _, eachEvent := range events.Events {
			fmt.Println(eachEvent)

			profile := internal.ReplicatePatientProfile{
				FirstName: eachEvent.FirstName,
				LastName:  eachEvent.LastName,
				Sex:       eachEvent.Sex,
				UpdatedAt: eachEvent.UpdateAt,
			}

			if eachEvent.PatientId == 0 {
				// create a new profile
				_ = internal.CreatePatientProfile(&profile)
			} else {
				// update profile
				_, _ = internal.UpdateReplicatePatientProfileById(eachEvent.PatientId, eachEvent.UpdateAt, &profile)
			}
		}

	}
}
