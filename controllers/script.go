package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/perseo_mid/models"
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
	var votacionRecivida map[string]interface{}
	var alertErr models.Alert
	alertas := append([]interface{}{"Response:"})

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &votacionRecivida); err == nil {

		votacionRespuesta, errVotacion := models.LlenaTablas(votacionRecivida)
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
