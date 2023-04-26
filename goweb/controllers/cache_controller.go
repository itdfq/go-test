package controllers

import (
	"github.com/astaxie/beego"
	"goweb/cacheUtils"
	"log"
	"time"
)

type CacheController struct {
	beego.Controller
}

func (c *CacheController) Get() {
	key := c.GetString("key")
	log.Println("当前查询的参数为：", key)
	if key != "" {
		value := cacheUtils.GetStr(key)
		log.Println(key, "=", value)
		mystruct := &Result{200, "", value, 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
	}

}

func (c *CacheController) PutCache() {
	key := c.GetString("key")
	value := c.GetString("value")
	second, _ := c.GetInt("second")
	log.Println("当前的key:", key, "value:", value, "过期时间：", second, "秒")
	if second == 0 {
		second = 10
	}
	if key == "" || value == "" {
		mystruct := &Result{200, "", "key或value 不能为空", 0}
		c.Data["json"] = mystruct
		c.ServeJSON()
		return
	}
	timeSecond := time.Duration(second) * time.Second
	err := cacheUtils.SetStr(key, value, timeSecond)
	var result string
	if err != nil {
		result = ("保存失败")
	} else {
		result = ("保存成功")
	}
	mystruct := &Result{200, "", result, 0}
	c.Data["json"] = mystruct
	c.ServeJSON()
}
