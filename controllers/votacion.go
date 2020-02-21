package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/perseo_mid/models"
)

// VotacionController crud para votaciones de perseo
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
// @Description Creacion de votacion
// @Param	body		body 	{}	true		"se requiere el cuerpo de la votacion, en jbpm se puede obtener o en el cliente"
// @Success 201 {} votacion creada
// @Failure 403 body is empty
// @Failure 400 Bad Request
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
// @Description Obtiene una votacion segun un id especifico de la misma
// @Param	id		path 	string	true		"se requiere el ID de la votacion"
// @Success 200 {}
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
// @Description Obtiene todas la votaciones
// @Success 200 {}
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
// @Description Actualizacion de votaciones
// @Param	id		path 	string	true		"se requiere el id de la votacion a axctualizar"
// @Param	body		body 	{}	true		"se requiere el cuerpo de la votacion, en jbpm se puede obtener o en el cliente"
// @Success 200 {}
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
