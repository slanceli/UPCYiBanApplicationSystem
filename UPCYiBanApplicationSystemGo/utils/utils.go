package utils

import (
	"UPCYiBanApplicationSystemGo/dao"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Application struct {
	Name string	`json:"name"`
	PhomeNum string `json:"phome_num"`
	Gender string	`json:"gender"`
	MailAddress string	`json:"mail_address"`
	PoliticalFace string	`json:"political_face"`	//政治面貌
	Class string	`json:"class"`	//专业班级
	FirstVolunteer string	`json:"first_volunteer"`	//第一志愿
	SecondVolunteer string	`json:"second_volunteer"`	//第二志愿
	Transfers string	`json:"transfers"`	//服从调剂
	Profile string	`json:"profile"`	//个人简介
	Advantage string	`json:"advantage"`	//个人对该岗位的优势
	Cognition string	`json:"cognition"`	//个人对该岗位的认知和思路
	ReviewComments string	`json:"review_comments"`	//审核意见
}

func (application *Application) AddApplication (upcid string) bool {
	sqlStr := "DELETE FROM yibanapplication.application WHERE upcid = ?"
	_, err := dao.DB.Exec(sqlStr, upcid)
	if err != nil {
		fmt.Println("AddApplication failed, err: ", err)
		return false
	}
	sqlStr = "INSERT INTO yibanapplication.application VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = dao.DB.Exec(sqlStr, nil, application.Name, upcid, application.PhomeNum, application.Gender, application.MailAddress, application.PoliticalFace, application.Class, application.FirstVolunteer, application.SecondVolunteer, application.Transfers, application.Profile, application.Advantage, application.Cognition, "")
	if err != nil {
		fmt.Println("AddApplication failed, err: ", err)
		return false
	}
	return true
}

//数字石大登录黑盒
func UPClogin(username string, passwd string) bool {
	res := struct {
		E int	`json:"e"`
		M string	`json:"m"`
		D interface{}	`json:"d"`
	}{}
	client := &http.Client{}
	payload := make(url.Values)
	payload.Add("username", username)
	payload.Add("password", passwd)
	req, err := http.NewRequest("POST", "https://app.upc.edu.cn/uc/wap/login/check", strings.NewReader(payload.Encode()))
	if err != nil {
		fmt.Println("New Request failed, err: ", err)
		return false
		// handle error
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http do failed, err: ", err)
		return false
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ioutil.ReadAll failed, err: ", err)
		return false
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println("Unmarshal failed, err: ", err)
		return false
	}
	fmt.Println(res)
	if res.E == 0 {
		return true
	} else {
		return false
	}
}