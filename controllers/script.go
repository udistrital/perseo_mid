package controllers

import (
	"github.com/astaxie/beego"
)

// ScriptController operations for Script
type ScriptController struct {
	beego.Controller
}

// URLMapping ...
func (c *ScriptController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Script
// @Param	body		body 	{}	true		"body for Script content"
// @Success 201 {}
// @Failure 403 body is empty
// @router / [post]
func (c *ScriptController) Post() {

}
