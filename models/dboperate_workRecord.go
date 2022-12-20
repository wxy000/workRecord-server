package models

import (
	"errors"
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

// 根据客户id获取客户信息
func GetCustomerInfoByCustomerid(customerid string) (int64, *Customers, error) {
	var cus Customers
	if customerid == "" {
		return 0, nil, errors.New("客户编号不得为空")
	}
	result := globals.DB.Where("customerid = ?", customerid).Limit(1).Find(&cus)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	if result.RowsAffected < 1 {
		return 0, nil, errors.New("客户不存在")
	}
	return result.RowsAffected, &cus, nil
}

// 根据产品id获取产品信息
func GetProductInfoByProductid(productid string) (int64, *Products, error) {
	var product Products
	if productid == "" {
		return 0, nil, errors.New("产品编号不得为空")
	}
	result := globals.DB.Where("productid = ?", productid).Limit(1).Find(&product)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	if result.RowsAffected < 1 {
		return 0, nil, errors.New("产品不存在")
	}
	return result.RowsAffected, &product, nil
}

// 根据id获取问题分类的信息
func GetIssuemainInfoByMainid(mainid string) (int64, *Issuemains, error) {
	var issuemain Issuemains
	if mainid == "" {
		return 0, nil, errors.New("问题分类不得为空")
	}
	result := globals.DB.Where("mainid = ?", mainid).Limit(1).Find(&issuemain)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	if result.RowsAffected < 1 {
		return 0, nil, errors.New("问题分类不存在")
	}
	return result.RowsAffected, &issuemain, nil
}

// 根据id获取问题类型的信息
func GetIssuedetailInfoByDetailid(detailid string) (int64, *Issuedetails, error) {
	var issuedetail Issuedetails
	if detailid == "" {
		return 0, nil, errors.New("问题类型不得为空")
	}
	result := globals.DB.Where("detailid = ?", detailid).Limit(1).Find(&issuedetail)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	if result.RowsAffected < 1 {
		return 0, nil, errors.New("问题类型不存在")
	}
	return result.RowsAffected, &issuedetail, nil
}
