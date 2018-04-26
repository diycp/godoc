package controllers

import (
	"github.com/astaxie/beego"
	"showdoc/consts"
	"showdoc/models"
)

// Operations about Users
type ItemController struct {
	beego.Controller
}



// @Title MyList
// @Description mylist
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /myList [get]
func (u *ItemController) MyList() {
	json := consts.Json{}
	//var data [0]int

	uid := u.GetSession(consts.SESSION_UID)
	if uid == nil {
		json.Set(10000, "用户未登录")
		u.Data["json"] = json.VendorError()
		u.ServeJSON()
	} else {

		userId :=uid.(int64)
		myItem := models.GetMyItem(userId)
		data := [len(myItem)]map[string]interface{}

		//var data = [](map[string]interface{})
		for i, value := range myItem {
			println(i)
			if value != nil {
				data[i] = value.Format()
			}
		}

		json.SetData(data)
		u.Data["json"] = json.VendorOk()
		u.ServeJSON()
	}


}


// @Title add item
// @Description add item
// @Param   item_type     formData    int   true        "项目类型 1常规项目  2单页项目"
// @Param   item_name     formData    string  true        "项目名称"
// @Param   password     formData    string  false        "查看密码"
// @Param   item_description     formData    string  false        "项目描述"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /add [post]
func (u *ItemController) Add() {

	var err error
	var json consts.Json
	var userId int64
	uid := u.GetSession(consts.SESSION_UID)
	if uid == nil {
		json.Set(10000, "用户未登录")
		u.Data["json"] = json.VendorError()
		u.ServeJSON()
		return
	} else {
		userId = uid.(int64)
	}

	item_type,_ := u.GetInt("item_type")
	item_name := u.GetString("item_name")
	password := u.GetString("password")
	item_description := u.GetString("item_description")

	var item models.Item
	item.Title = item_name
	item.Type = item_type
	item.Password = password
	item.Description = item_description
	item.UserId = userId
	item.Id,err = item.Create()

	if err != nil {
		json.Set(500, err.Error())
		u.Data["json"] = json.VendorError()
		u.ServeJSON()
		return
	}
	json.Set(0,"成功")
	u.Data["json"] = json.VendorOk()
	u.ServeJSON()
}
