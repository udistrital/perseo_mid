package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/perseo_mid/models"
)

// VotacionController operations for Votacion
type VotacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *VotacionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title Create
// @Description create Votacion
// @Param	body		body 	models.Votacion	true		"body for Votacion content"
// @Success 201 {object} models.Votacion
// @Failure 403 body is empty
// @router / [post]
func (c *VotacionController) Post() {
	var votacionRecivida map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &votacionRecivida); err == nil {

		votacionRespuesta, errVotacion := models.PostVotaciones(votacionRecivida)
		// plantillaRespuerta, errPlantilla := models.IngresoPlantilla(votacionRecivida)

		if votacionRespuesta != nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = votacionRespuesta
			// alertErr.Body = models.CrearSuccess("la plantilla se ingreso con exito")
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, errVotacion)
			alertErr.Body = alertas
			c.Ctx.Output.SetStatus(400)
		}

	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err.Error())
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}

	c.Data["json"] = alertErr
	c.ServeJSON()

}

// GetOne ...
// @Title GetOne
// @Description get Votacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Votacion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *VotacionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	votacion, errVotacion := models.ObtenerTodasVotacionesID(idStr)
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
// @Description get Votacion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Votacion
// @Failure 403
// @router / [get]
func (c *VotacionController) GetAll() {
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})
	votaciones, errVotaciones := models.ObtenerTodasVotaciones()
	if votaciones != nil {
		alertErr.Type = "OK"
		alertErr.Code = "200"
		alertErr.Body = votaciones
	} else {
		alertErr.Type = "error"
		alertErr.Code = "404"
		alertas = append(alertas, errVotaciones)
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(404)
	}
	c.Data["json"] = alertErr
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Votacion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Votacion	true		"body for Votacion content"
// @Success 200 {object} models.Votacion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *VotacionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	var votacionRecivida map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &votacionRecivida); err == nil {

		votacionRespuesta, errVotacion := models.PutVotaciones(votacionRecivida, idStr)
		// plantillaRespuerta, errPlantilla := models.IngresoPlantilla(votacionRecivida)

		if votacionRespuesta != nil {
			alertErr.Type = "OK"
			alertErr.Code = "200"
			alertErr.Body = votacionRespuesta
			// alertErr.Body = models.CrearSuccess("la plantilla se ingreso con exito")
		} else {
			alertErr.Type = "error"
			alertErr.Code = "400"
			alertas = append(alertas, errVotacion)
			alertErr.Body = alertas
			c.Ctx.Output.SetStatus(400)
		}

	} else {
		alertErr.Type = "error"
		alertErr.Code = "400"
		alertas = append(alertas, err.Error())
		alertErr.Body = alertas
		c.Ctx.Output.SetStatus(400)
	}

	c.Data["json"] = alertErr
	c.ServeJSON()

}
