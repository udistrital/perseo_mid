package models

import "github.com/astaxie/beego"

// ObtenerTodasVotaciones ...
func ObtenerTodasVotaciones() (votaciones []map[string]interface{}, outputError interface{}) {
	var votacionesCenso []map[string]interface{}
	// error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"contratoSuscritoProxyService/dependencias_sic/", &votacionesCenso)
	// if error != nil {
	// 	return nil, error
	// } else {
	// 	return votacionesCenso, nil

	// }
	votacionesCenso = append(votacionesCenso, map[string]interface{}{
		"Id":             1,
		"Nombre":         "aaaaaaaaaaa",
		"Observacion":    "sssssssssssssss",
		"Año":            "2020-02-20T05:00:00.000Z",
		"Fechaejecucion": "2020-02-14T05:00:00.000Z",
		"Estado":         true,
		"DocentesPlanta": true,
		"DocentesVe":     true,
		"Funcionarios":   true,
		"Estudiantes":    true,
		"Egresados":      true,
		"Contratistas":   false,
		"Exrectores":     true,
	})
	votacionesCenso = append(votacionesCenso, map[string]interface{}{
		"Id":             2,
		"Nombre":         "prueba",
		"Observacion":    "reolucion x",
		"Año":            "2020-02-20T05:00:00.000Z",
		"Fechaejecucion": "2020-02-14T05:00:00.000Z",
		"Estado":         true,
		"DocentesPlanta": true,
		"DocentesVe":     true,
		"Funcionarios":   false,
		"Estudiantes":    true,
		"Egresados":      true,
		"Contratistas":   false,
		"Exrectores":     false,
	})
	return votacionesCenso, nil
}

// ObtenerTodasVotacionesID ...
func ObtenerTodasVotacionesID(idVotacion string) (votaciones map[string]interface{}, outputError interface{}) {
	var votacionesCenso map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"contratoSuscritoProxyService/dependencias_sic/"+idVotacion, &votacionesCenso)
	if error != nil {
		return nil, error
	} else {
		return votacionesCenso, nil

	}
	// return nil, nil
}
