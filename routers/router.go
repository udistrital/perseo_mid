// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/perseo_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/votacion",
			beego.NSInclude(
				&controllers.VotacionController{},
			),
		),
		beego.NSNamespace("/script",
			beego.NSInclude(
				&controllers.ScriptController{},
			),
		),
		beego.NSNamespace("/filtro",
			beego.NSInclude(
				&controllers.FiltroController{},
			),
		),
		beego.NSNamespace("/cantidades",
			beego.NSInclude(
				&controllers.CantidadesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
