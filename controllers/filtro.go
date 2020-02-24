package controllers

import (
	"github.com/astaxie/beego"
)

// FiltroController operations for Filtro
type FiltroController struct {
	beego.Controller
}

// URLMapping ...
func (c *FiltroController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Filtro
// @Param	body		body 	models.Filtro	true		"body for Filtro content"
// @Success 201 {object} models.Filtro
// @Failure 403 body is empty
// @router / [post]
func (c *FiltroController) Post() {

}

// Delete ...
// @Title Delete
// @Description delete the Filtro
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FiltroController) Delete() {

}
