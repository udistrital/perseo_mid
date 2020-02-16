package models

import (
	"fmt"

	"github.com/astaxie/beego"
)

// ObtenerTodasVotaciones ...
func ObtenerTodasVotaciones() (votaciones []map[string]interface{}, outputError interface{}) {
	var votacionesCenso map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+beego.AppConfig.String("perseo_ns_service")+"votaciones", &votacionesCenso)
	if error != nil {
		return nil, error
	}
	votacionesArray := votacionesCenso["votaciones"].(map[string]interface{})
	ArrayDeVotaciones, errArray := GetElementoMaptoStringToMapArray(votacionesArray["votacion"])
	if ArrayDeVotaciones != nil {
		votacionesConvertido := conversionVotacionCliente(ArrayDeVotaciones)
		return votacionesConvertido, nil
	}
	return nil, errArray
}

// ObtenerTodasVotacionesID ...
func ObtenerTodasVotacionesID(idVotacion string) (votaciones []map[string]interface{}, outputError interface{}) {
	// var votacionesCenso []map[string]interface{}
	var votacionesCenso map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+beego.AppConfig.String("perseo_ns_service")+"votaciones_id/"+idVotacion, &votacionesCenso)
	if error != nil {
		return nil, error
	}
	votacionesArray := votacionesCenso["votacion_id"].(map[string]interface{})
	ArrayDeVotaciones, errArray := GetElementoMaptoStringToMapArray(votacionesArray["votacion"])
	if ArrayDeVotaciones != nil {
		votacionesConvertido := conversionVotacionCliente(ArrayDeVotaciones)
		return votacionesConvertido, nil
	}
	return nil, errArray

	// votacionesCenso = append(votacionesCenso, map[string]interface{}{
	// 	"Id":             2,
	// 	"Nombre":         "prueba",
	// 	"Observacion":    "reolucion x",
	// 	"Año":            "2020-02-20T05:00:00.000Z",
	// 	"Fechaejecucion": "2020-02-14T05:00:00.000Z",
	// 	"Estado":         true,
	// 	"DocentesPlanta": true,
	// 	"DocentesVe":     true,
	// 	"Funcionarios":   false,
	// 	"Estudiantes":    true,
	// 	"Egresados":      true,
	// 	"Contratistas":   false,
	// 	"Exrectores":     false,
	// })
	// return votacionesCenso, nil
}

// PostVotaciones ...
func PostVotaciones(votacion map[string]interface{}) (votacionEnviada map[string]interface{}, outputError interface{}) {
	// var votacionesCenso []map[string]interface{}
	// error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+"contratoSuscritoProxyService/dependencias_sic/"+idVotacion, &votacionesCenso)
	// if error != nil {
	// 	return nil, error
	// } else {
	// 	return votacionesCenso, nil

	// }
	// votacionesCenso = append(votacionesCenso, map[string]interface{}{
	// 	"Id":             2,
	// 	"Nombre":         "prueba",
	// 	"Observacion":    "reolucion x",
	// 	"Año":            "2020-02-20T05:00:00.000Z",
	// 	"Fechaejecucion": "2020-02-14T05:00:00.000Z",
	// 	"Estado":         true,
	// 	"DocentesPlanta": true,
	// 	"DocentesVe":     true,
	// 	"Funcionarios":   false,
	// 	"Estudiantes":    true,
	// 	"Egresados":      true,
	// 	"Contratistas":   false,
	// 	"Exrectores":     false,
	// })
	// return votacionesCenso, nil
	votacionConvertida := conversionVotacionJBPM(votacion)

	return votacionConvertida, nil
}

func conversionVotacionJBPM(votacion map[string]interface{}) (votacionconvertida map[string]interface{}) {
	// votacion["Contratistas"] = "N"
	for key, value := range votacion {
		fmt.Println("Key:", key, "Value:", value)
		if fmt.Sprintf("%v", votacion[key]) == "true" {
			votacion[key] = "S"
		}
		if fmt.Sprintf("%v", votacion[key]) == "false" {
			votacion[key] = "N"
		}
	}

	return votacion
}

func conversionVotacionCliente(votacion []map[string]interface{}) (votacionconvertida []map[string]interface{}) {
	for i := 0; i < len(votacion); i++ {
		for key := range votacion[i] {
			// fmt.Println("Key:", key, "Value:", value)
			if fmt.Sprintf("%v", votacion[i][key]) == "S" {
				votacion[i][key] = true
			}
			if fmt.Sprintf("%v", votacion[i][key]) == "A" {
				votacion[i][key] = true
			}
			if fmt.Sprintf("%v", votacion[i][key]) == "N" {
				votacion[i][key] = false
			}
			if fmt.Sprintf("%v", votacion[i][key]) == "I" {
				votacion[i][key] = false
			}
		}
	}

	return votacion
}

// PutVotaciones ...
func PutVotaciones(votacion map[string]interface{}, votacionID string) (votacionEnviada map[string]interface{}, outputError interface{}) {

	votacionConvertida := conversionVotacionJBPM(votacion)

	return votacionConvertida, nil
}
