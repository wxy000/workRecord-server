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
	Cname              string  `json:"cname"`
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

type AnalysisRecordList5 struct {
	Typea              string  `json:"typea"`
	Ny                 string  `json:"ny"`
	Handleestimatetime float32 `json:"handleestimatetime"`
	Handleactualtime   float32 `json:"handleactualtime"`
}

type AnalysisRecordList6 struct {
	Cname              string  `json:"cname"`
	Ny                 string  `json:"ny"`
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
		Select("b.cname, a.`number`, a.title, a.content, c.realName feedbackname, a.feedbackdate, CONCAT(a.productid,\".\",d.productname) productname, a.urgent, CONCAT(a.issuemainid,\".\",e.mainName) issuemainname, CONCAT(a.issuedetailid,\".\",f.detailName) issuedetailname, g.realName handlername, a.handleestimatetime, a.handleactualtime, a.handlereply, a.casestatus, a.onsite, a.closetime, a.mark").
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

func GetAnalysisRecordList5(handlerid string, feedbackdatestart string, feedbackdateend string) (bool, *[]AnalysisRecordList5, int64) {
	var d1 []AnalysisRecordList5
	// 签单
	sql1 := ` SELECT '1' typea,CONCAT(nian, '-', LPAD(yue, 2, 0)) ny,SUM(IFNULL(handleestimatetime, 0)) handleestimatetime,SUM(IFNULL(handleactualtime, 0)) handleactualtime
				FROM rangemonths
				LEFT JOIN (SELECT *
							 FROM records a
							 LEFT JOIN issuedetails b ON b.detailid = a.issuedetailid
							WHERE b.classid IN ('C01','C02','C03','C11')
							  AND a.handlerid = ?
							  AND a.feedbackdate BETWEEN ? AND ?
						  ) records ON feedbackdate BETWEEN datestart AND dateend
			   GROUP BY nian,yue `
	// 返工
	sql2 := ` SELECT '2' typea,CONCAT(nian, '-', LPAD(yue, 2, 0)) ny,SUM(IFNULL(handleestimatetime, 0)) handleestimatetime,SUM(IFNULL(handleactualtime, 0)) handleactualtime
				FROM rangemonths
				LEFT JOIN (SELECT *
							 FROM records a
							 LEFT JOIN issuedetails b ON b.detailid = a.issuedetailid
							WHERE b.classid IN ('C04')
							  AND a.handlerid = ?
							  AND a.feedbackdate BETWEEN ? AND ?
						  ) records ON feedbackdate BETWEEN datestart AND dateend
			   GROUP BY nian,yue `
	// 问题处理
	sql3 := ` SELECT '3' typea,CONCAT(nian, '-', LPAD(yue, 2, 0)) ny,SUM(IFNULL(handleestimatetime, 0)) handleestimatetime,SUM(IFNULL(handleactualtime, 0)) handleactualtime
				FROM rangemonths
				LEFT JOIN (SELECT *
							 FROM records a
							 LEFT JOIN issuedetails b ON b.detailid = a.issuedetailid
							WHERE b.classid NOT IN ('C01','C02','C03','C04','C11')
							  AND a.handlerid = ?
							  AND a.feedbackdate BETWEEN ? AND ?
						  ) records ON feedbackdate BETWEEN datestart AND dateend
			   GROUP BY nian,yue `
	sql := " SELECT * FROM ( " + sql1 + " UNION ALL " + sql2 + " UNION ALL " + sql3 + " ) xx WHERE handleestimatetime <> 0 ORDER BY typea,ny "
	res := globals.DB.Raw(sql, handlerid, feedbackdatestart, feedbackdateend, handlerid, feedbackdatestart, feedbackdateend, handlerid, feedbackdatestart, feedbackdateend)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}

func GetAnalysisRecordList6(handlerid string, feedbackdatestart string, feedbackdateend string) (bool, *[]AnalysisRecordList6, int64) {
	var d1 []AnalysisRecordList6
	sql := ` SELECT cname, CONCAT(nian,'-',LPAD( yue, 2, 0 )) ny,SUM(IFNULL( handleestimatetime, 0 )) handleestimatetime,SUM(IFNULL( handleactualtime, 0 )) handleactualtime 
			   FROM rangemonths
			   LEFT JOIN (SELECT * FROM records a
				 		   WHERE a.handlerid = ?
							 AND a.feedbackdate BETWEEN ? AND ?) records ON feedbackdate BETWEEN datestart AND dateend
			   LEFT JOIN customers ON customers.customerid = records.customerid
			  GROUP BY nian,yue,cname 
			  ORDER BY cname,ny,handleestimatetime DESC `
	res := globals.DB.Raw(sql, handlerid, feedbackdatestart, feedbackdateend)
	r1 := res.Scan(&d1)
	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}
