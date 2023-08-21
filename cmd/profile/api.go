package main

import (
	"HeidiTask/internal"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func initRouter() *gin.Engine {
	router := gin.New()

	profileApi := router.Group("/profile")
	{
		profileApi.POST("/add", addOneProfile)
		profileApi.PUT("/update/one", updateOneProfile)
		profileApi.PUT("/update/multiple", updateMultipleProfile)
		profileApi.GET("/query", queryProfileById)
	}

	return router
}

/*
create a new profile data
*/
func addOneProfile(c *gin.Context) {
	var profile internal.ProfileUpdateData
	if err := c.ShouldBindJSON(&profile); err != nil {
		internal.Fail(c, internal.ApiParameterError, "request body error.")
		return
	}

	curNano := time.Now().UnixNano()
	newProfile := internal.PatientProfile{
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Sex:       profile.Sex,
		UpdatedAt: curNano,
	}

	err := internal.CreatePatientProfile(&newProfile)
	if err != nil {
		internal.Fail(c, internal.ApiServerError, err.Error())
		return
	}

	profile.UpdateAt = curNano

	// send kafka event message
	sendUpdateKafkaMessages([]internal.ProfileUpdateData{profile})

	// response
	internal.Success(c, newProfile)
}

// update one profile
func updateOneProfile(c *gin.Context) {
	var profile internal.ProfileUpdateData
	if err := c.ShouldBindJSON(&profile); err != nil {
		internal.Fail(c, internal.ApiParameterError, "request body error.")
		return
	}

	obj, err := internal.QueryPatientProfileById(profile.PatientId)
	if err != nil {
		internal.Fail(c, internal.ApiObjectNotFound, "can't find the profile")
		return
	}

	curNano := time.Now().UnixNano()
	updateProfile := internal.PatientProfile{
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Sex:       profile.Sex,
		UpdatedAt: curNano,
	}

	rows, err := internal.PartialUpdatePatientProfile(profile.PatientId, curNano, &updateProfile)
	if rows == 0 || err != nil {
		internal.Fail(c, internal.ApiServerError, "update profile failed")
		return
	}

	profile.UpdateAt = curNano

	// send kafka event message
	sendUpdateKafkaMessages([]internal.ProfileUpdateData{profile})

	// response
	internal.Success(c, obj)
}

// update multiple profiles
func updateMultipleProfile(c *gin.Context) {
	var profileArray []internal.ProfileUpdateData
	if err := c.ShouldBindJSON(&profileArray); err != nil {
		internal.Fail(c, internal.ApiParameterError, "request body error.")
		return
	}

	var updatedProfiles []internal.PatientProfile
	var updateKafkaEvents []internal.ProfileUpdateData

	for _, profile := range profileArray {
		obj, err := internal.QueryPatientProfileById(profile.PatientId)
		if err != nil {
			continue
		}

		curNano := time.Now().UnixNano()
		updateProfile := internal.PatientProfile{
			FirstName: profile.FirstName,
			LastName:  profile.LastName,
			Sex:       profile.Sex,
			UpdatedAt: curNano,
		}

		rows, err := internal.PartialUpdatePatientProfile(profile.PatientId, curNano, &updateProfile)
		if rows == 0 || err != nil {
			continue
		}

		updatedProfiles = append(updatedProfiles, obj)

		profile.UpdateAt = curNano
		updateKafkaEvents = append(updateKafkaEvents, profile)
	}

	// send kafka event message
	sendUpdateKafkaMessages(updateKafkaEvents)

	// response
	internal.Success(c, true)
}

// query profile by patient id
func queryProfileById(c *gin.Context) {
	id := c.Query("id")
	if len(id) == 0 {
		internal.Fail(c, internal.ApiParameterError, "request body error.")
		return
	}
	patientID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		internal.Fail(c, internal.ApiParameterError, "request body error.")
		return
	}
	profile, err := internal.QueryPatientProfileById(uint(patientID))
	if err != nil {
		internal.Fail(c, internal.ApiServerError, "can't find the profile")
		return
	}

	// response
	internal.Success(c, profile)
}
