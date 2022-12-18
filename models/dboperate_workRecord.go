package models

import (
	"server/common"
	"server/globals"

	"github.com/gin-gonic/gin"
)

type RecordList struct {
	Id                 uint    `json:"id"`
	Customerid         string  `json:"customerid"`
	Customername       string  `json:"customername"`
	Number             string  `json:"number"`
	Title              string  `json:"title"`
	Content            string  `json:"content"`
	Feedbackid         string  `json:"feedbackid"`
	Feedbackname       string  `json:"feedbackname"`
	Feedbackdate       string  `json:"feedbackdate"`
	Productid          string  `json:"productid"`
	Productname        string  `json:"productname"`
	Urgent             string  `json:"urgent"`
	Issuemainid        string  `json:"issuemainid"`
	Issuemainname      string  `json:"issuemainname"`
	Issuedetailid      string  `json:"issuedetailid"`
	Issuedetailname    string  `json:"issuedetailname"`
	Handlerid          string  `json:"handlerid"`
	Handlername        string  `json:"handlername"`
	Handleestimatetime float32 `json:"handleestimatetime"`
	Handleactualtime   float32 `json:"handleactualtime"`
	Handlereply        string  `json:"handlereply"`
	Casestatus         uint    `json:"casestatus"`
	Onsite             uint    `json:"onsite"`
	Closetime          string  `json:"closetime"`
	Bugpeopleid        string  `json:"bugpeopleid"`
	Bugpeoplename      string  `json:"bugpeoplename"`
	Mark               string  `json:"mark"`
}

// 根据处理人获取工作记录
func GetRecordByHandlerid(handlerid string, c *gin.Context) (bool, *[]RecordList, int64) {
	var d1 []RecordList
	res := globals.DB.Table("records a").
		Select("a.id, a.customerid, b.customername, a.`number`, a.title, a.content, a.feedbackid, c.realName feedbackname, a.feedbackdate, a.productid, CONCAT(a.productid,\".\",d.productname) productname, a.urgent, a.issuemainid, CONCAT(a.issuemainid,\".\",e.mainName) issuemainname, a.issuedetailid, CONCAT(a.issuedetailid,\".\",f.detailName) issuedetailname, a.handlerid, g.realName handlername, a.handleestimatetime, a.handleactualtime, a.handlereply, a.casestatus, a.onsite, a.closetime, a.bugpeopleid, h.realName bugpeoplename, a.mark").
		Joins("left join customers b on b.customerid = a.customerid").
		Joins("left join users c on c.userid = a.feedbackid").
		Joins("LEFT JOIN products d on d.productid = a.productid").
		Joins("LEFT JOIN issuemains e on e.mainid = a.issuemainid").
		Joins("left join issuedetails f on f.detailid = a.issuedetailid").
		Joins("left join users g on g.userid = a.handlerid").
		Joins("left join users h on h.userid = a.bugpeopleid").
		Where("a.handlerid = ?", handlerid).
		Order("a.feedbackdate desc")

	// 这个是取条数的
	r1 := res.Scan(&d1)
	// 这个是按分页返回数据
	res.Scopes(common.Paginate(c)).Scan(&d1)

	if r1.Error != nil {
		return false, nil, 0
	}
	return true, &d1, r1.RowsAffected
}
