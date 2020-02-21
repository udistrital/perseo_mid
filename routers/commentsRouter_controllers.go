package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:CantidadesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:CantidadesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:CantidadesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:CantidadesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:CantidadesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:CantidadesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:CantidadesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:CantidadesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:CantidadesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:CantidadesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:FiltroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:FiltroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:FiltroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:FiltroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:FiltroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:FiltroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:FiltroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:FiltroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:FiltroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:FiltroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:ScriptController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:ScriptController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:VotacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:VotacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:VotacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:VotacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:VotacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:VotacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:VotacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/perseo_mid/controllers:VotacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
