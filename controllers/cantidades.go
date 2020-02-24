package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/perseo_mid/models"
)

// CantidadesController cantidad de personas segun cada rol participante
type CantidadesController struct {
	beego.Controller
}

// URLMapping ...
func (c *CantidadesController) URLMapping() {

	c.Mapping("GetOne", c.GetOne)
}

// GetOne ...
// @Title GetOne
// @Description get Cantidades by id
// @Param	id		path 	string	true		"id de la votacion"
// @Success 200 {}
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
