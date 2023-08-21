package internal

import "sort"

type ProfileUpdateData struct {
	PatientId uint   `from:"patientId" json:"patientId"`
	FirstName string `from:"firstName" json:"firstName"`
	LastName  string `from:"lastName" json:"lastName"`
	Sex       string `from:"sex" json:"sex"`
	UpdateAt  int64  `from:"updatedAt" json:"updatedAt"`
}

type ProfileUpdateEvents struct {
	Events []ProfileUpdateData
}

// Sort : sort the multiple events in the order below
// 1. PatientId == 0
// 2. less UpdateAt
func (events *ProfileUpdateEvents) Sort() {
	sort.Slice(events.Events, func(i, j int) bool {
		return events.Events[i].PatientId == 0 || events.Events[i].UpdateAt < events.Events[j].UpdateAt
	})
}
