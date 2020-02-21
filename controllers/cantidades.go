package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/perseo_mid/models"
)

// CantidadesController operations for Cantidades
type CantidadesController struct {
	beego.Controller
}

// URLMapping ...
func (c *CantidadesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Cantidades
// @Param	body		body 	models.Cantidades	true		"body for Cantidades content"
// @Success 201 {object} models.Cantidades
// @Failure 403 body is empty
// @router / [post]
func (c *CantidadesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Cantidades by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Cantidades
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CantidadesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	votacion, errVotacion := models.ObtenerCantidades(idStr)
	if votacion != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = votacion
	} else {
		alertErr.Type = "error"
		alertErr.Code = "404"
		alertas = append(alertas, errVotacion)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description get Cantidades
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Cantidades
// @Failure 403
// @router / [get]
func (c *CantidadesController) GetAll() {
	// var alertErr models.Alert
	// alertas := append([]interface{}{"Response:"})
	// votaciones, errVotaciones := models.ObtenerCantidades()
	// if votaciones != nil {
	// 	alertErr.Type = "OK"
	// 	alertErr.Code = "200"
	// 	alertErr.Body = votaciones
	// } else {
	// 	alertErr.Type = "error"
	// 	alertErr.Code = "404"
	// 	alertas = append(alertas, errVotaciones)
	// 	alertErr.Body = alertas
	// 	c.Ctx.Output.SetStatus(404)
	// }
	// c.Data["json"] = alertErr
	// c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Cantidades
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Cantidades	true		"body for Cantidades content"
// @Success 200 {object} models.Cantidades
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CantidadesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Cantidades
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CantidadesController) Delete() {

}
