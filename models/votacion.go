package models

import "github.com/astaxie/beego"

// ObtenerTodasVotaciones ...
func ObtenerTodasVotaciones() (votaciones map[string]interface{}, outputError interface{}) {
	var votacionesCenso map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"contratoSuscritoProxyService/dependencias_sic/", &votacionesCenso)
	if error != nil {
		return nil, error
	} else {
		return votacionesCenso, nil

	}
	// return nil, nil
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
