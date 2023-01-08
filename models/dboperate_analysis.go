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

type AnalysisRecordListSum struct {
	Handleestimatetime float32 `json:"handleestimatetime"`
	Handleactualtime   float32 `json:"handleactualtime"`
}

type AnalysisRecordList4 struct {
	Customername       string  `json:"customername"`
	Number             string  `json:"number"`
	Title              string  `json:"title"`
	Content            string  `json:"content"`
	Feedbackname       string  `json:"feedbackname"`
	Feedbackdate       string  `json:"feedbackdate"`
	Productname        string  `json:"productname"`
	Urgent             string  `json:"urgent"`
	Issuemainname      string  `json:"issuemainname"`
	Issuedetailname    string  `json:"issuedetailname"`
	Handlername        string  `json:"handlername"`
	Handleestimatetime float32 `json:"handleestimatetime"`
	Handleactualtime   float32 `json:"handleactualtime"`
	Handlereply        string  `json:"handlereply"`
	Casestatus         uint    `json:"casestatus"`
	Onsite             uint    `json:"onsite"`
	Closetime          string  `json:"closetime"`
	Mark               string  `json:"mark"`
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
		Group("c.classid").
		Order("handleestimatetime desc")
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

func GetAnalysisRecordListSum(handlerid string, feedbackdatestart string, feedbackdateend string) (bool, *[]AnalysisRecordListSum, int64) {
	var d1 []AnalysisRecordListSum
	res := globals.DB.Table("records a").
		Select("sum(handleestimatetime) handleestimatetime,sum(handleactualtime) handleactualtime").
		Where("handlerid = ? AND feedbackdate BETWEEN ? AND ?", handlerid, feedbackdatestart, feedbackdateend)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

func GetAnalysisRecordList4(handlerid string, feedbackdatestart string, feedbackdateend string) (bool, *[]AnalysisRecordList4, int64) {
	var d1 []AnalysisRecordList4
	res := globals.DB.Table("records a").
		Select("b.customername, a.`number`, a.title, a.content, c.realName feedbackname, a.feedbackdate, CONCAT(a.productid,\".\",d.productname) productname, a.urgent, CONCAT(a.issuemainid,\".\",e.mainName) issuemainname, CONCAT(a.issuedetailid,\".\",f.detailName) issuedetailname, g.realName handlername, a.handleestimatetime, a.handleactualtime, a.handlereply, a.casestatus, a.onsite, a.closetime, a.mark").
		Joins("left join customers b on b.customerid = a.customerid").
		Joins("left join users c on c.username = a.feedbackid").
		Joins("LEFT JOIN products d on d.productid = a.productid").
		Joins("LEFT JOIN issuemains e on e.mainid = a.issuemainid").
		Joins("left join issuedetails f on f.detailid = a.issuedetailid").
		Joins("left join users g on g.username = a.handlerid").
		Where("handlerid = ? AND feedbackdate BETWEEN ? AND ?", handlerid, feedbackdatestart, feedbackdateend).
		Order("a.feedbackdate desc")
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}
