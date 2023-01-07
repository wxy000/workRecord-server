package models

import (
	"server/globals"
)

type AnalysisRecordList1 struct {
	Feedbackdate       string  `json:"feedbackdate"`
	Handleestimatetime float32 `json:"handleestimatetime"`
	Handleactualtime   float32 `json:"handleactualtime"`
}

type AnalysisRecordList2 struct {
	Cname              string  `json:"cname"`
	Handleestimatetime float32 `json:"handleestimatetime"`
	Handleactualtime   float32 `json:"handleactualtime"`
}

type AnalysisRecordList3 struct {
	Mainname           string  `json:"mainname"`
	Detailname         string  `json:"detailname"`
	Handleestimatetime float32 `json:"handleestimatetime"`
	Handleactualtime   float32 `json:"handleactualtime"`
}

type AnalysisRecordList3_1 struct {
	Classname          string  `json:"classname"`
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

func GetAnalysisRecordList2(handlerid string, feedbackdatestart string, feedbackdateend string) (bool, *[]AnalysisRecordList2, int64) {
	var d1 []AnalysisRecordList2
	res := globals.DB.Table("records a").
		Select("b.cname cname,sum(a.handleestimatetime) handleestimatetime,sum(a.handleactualtime) handleactualtime").
		Joins("LEFT JOIN customers b ON b.customerid = a.customerid").
		Where("handlerid = ? AND feedbackdate BETWEEN ? AND ?", handlerid, feedbackdatestart, feedbackdateend).
		Group("cname").
		Order("handleestimatetime desc")
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

func GetAnalysisRecordList3(handlerid string, feedbackdatestart string, feedbackdateend string) (bool, *[]AnalysisRecordList3, int64) {
	var d1 []AnalysisRecordList3
	res := globals.DB.Table("records a").
		Select("CONCAT_WS('.',c.mainid,c.mainname) mainname,CONCAT_WS('.',b.detailid,b.detailname) detailname,sum(a.handleestimatetime) handleestimatetime,sum(a.handleactualtime) handleactualtime").
		Joins("LEFT JOIN issuedetails b ON b.detailid = a.issuedetailid").
		Joins("LEFT JOIN issuemains c ON c.mainid = b.mainid").
		Where("handlerid = ? AND feedbackdate BETWEEN ? AND ?", handlerid, feedbackdatestart, feedbackdateend).
		Group("a.issuedetailid").
		Order("a.issuedetailid")
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

func GetAnalysisRecordList3_1(handlerid string, feedbackdatestart string, feedbackdateend string) (bool, *[]AnalysisRecordList3_1, int64) {
	var d1 []AnalysisRecordList3_1
	res := globals.DB.Table("records a").
		Select("c.classname classname,sum(a.handleestimatetime) handleestimatetime,sum(a.handleactualtime) handleactualtime").
		Joins("LEFT JOIN issuedetails b ON b.detailid = a.issuedetailid").
		Joins("LEFT JOIN classes c ON c.classid = b.classid").
		Where("handlerid = ? AND feedbackdate BETWEEN ? AND ?", handlerid, feedbackdatestart, feedbackdateend).
		Group("c.classid")
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}
