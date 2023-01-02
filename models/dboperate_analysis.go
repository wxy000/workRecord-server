package models

import (
	"server/globals"
)

type AnalysisRecordList1 struct {
	Feedbackdate       string  `json:"feedbackdate"`
	Handleestimatetime float32 `json:"handleestimatetime"`
	Handleactualtime   float32 `json:"handleactualtime"`
}

func GetAnalysisRecordList1(handlerid string, feedbackdatestart string, feedbackdateend string) (bool, *[]AnalysisRecordList1, int64) {
	var d1 []AnalysisRecordList1
	res := globals.DB.Table("records").
		Select("date_format(feedbackdate,'%Y-%m-%d') feedbackdate, sum(handleestimatetime) handleestimatetime, sum(handleactualtime) handleactualtime").
		Where("handlerid = ? AND feedbackdate BETWEEN ? AND ?", handlerid, feedbackdatestart, feedbackdateend).
		Group("date_format(feedbackdate,'%Y-%m-%d')").
		Order("date_format(feedbackdate,'%Y-%m-%d')")
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}
