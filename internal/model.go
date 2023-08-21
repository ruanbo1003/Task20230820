package internal

import (
	"HeidiTask/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	mysqlDB *gorm.DB
)

func init() {
	mysqlDB, _ = gorm.Open(mysql.Open(config.MySqlDsn), &gorm.Config{})
}

func MysqlMigrate() {
	db, err := gorm.Open(mysql.Open(config.MySqlDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("gorm open mysql failed:", err.Error())
	}

	err = db.AutoMigrate(&PatientProfile{}, &ReplicatePatientProfile{})
	if err != nil {
		log.Fatal("mysql db migrate failed:", err.Error())
	}
}

type PatientProfile struct {
	PatientID uint   `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Sex       string `gorm:"not null"`
	UpdatedAt int64  `gorm:"autoUpdateTime:nano"`
}

type ReplicatePatientProfile struct {
	PatientID uint   `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Sex       string `gorm:"not null"`
	UpdatedAt int64  `gorm:"autoUpdateTime:nano"`
}

func CreatePatientProfile(profile interface{}) error {
	ret := mysqlDB.Create(profile)
	return ret.Error
}

func QueryPatientProfileById(id uint) (PatientProfile, error) {
	var profile PatientProfile
	ret := mysqlDB.First(&profile, id)

	return profile, ret.Error
}

func PartialUpdatePatientProfile(id uint, nano int64, updateData *PatientProfile) (int64, error) {
	ret := mysqlDB.Model(PatientProfile{}).
		Where("patient_id = ?", id).
		Where("updated_at < ?", nano).
		Updates(*updateData)
	return ret.RowsAffected, ret.Error
}

func UpdateReplicatePatientProfileById(id uint, nano int64, updateData *ReplicatePatientProfile) (int64, error) {
	ret := mysqlDB.Model(ReplicatePatientProfile{}).
		Where("patient_id = ?", id).
		Where("updated_at < ?", nano).
		Updates(*updateData)
	return ret.RowsAffected, ret.Error
}
