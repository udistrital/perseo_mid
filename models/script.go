package models

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// LlenaTablas ..
func LlenaTablas(datosVotacion map[string]interface{}) (votaciones map[string]interface{}, outputError interface{}) {
	FechaCorte := ObtenerFechaCorte(datosVotacion["FechaCorte"])
	fmt.Println(FechaCorte)
	votacion := datosVotacion["datosVotacion"].(map[string]interface{})
	participantes := datosVotacion["cantidadParticipantes"].(map[string]interface{})
	if fmt.Sprintf("%v", votacion["TVI_ESTADO"]) == "true" {
		lista := AnalisisParticipantes(votacion, FechaCorte, participantes)
		respueta := map[string]interface{}{
			"lista": lista,
		}
		return respueta, nil
	}
	return nil, CrearError("Esta votacion esta inactiva")

}

// LlenarListaIndividual ...
func LlenarListaIndividual(fechaCorte string, key string, votacionID string) {
	var objeto map[string]interface{}
	url := key + "/" + fechaCorte + "/" + votacionID
	fmt.Println(beego.AppConfig.String("administrativa_amazon_jbpm_url") + beego.AppConfig.String("perseo_ns__proxy_service") + url)
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+beego.AppConfig.String("perseo_ns_service")+url, &objeto)
	if error != nil {
		// return nil, error
		logs.Error(error)
	}
	logs.Info(objeto)

}

// AnalisisParticipantes ...
func AnalisisParticipantes(votacion map[string]interface{}, fechaDECorte string, cantodadPersonas map[string]interface{}) string {
	listaParticipantes := ""
	for key, value := range votacion {
		keyPeticion := KeyPeticion(key)
		if value == true && key != "TVI_ESTADO" {
			// fmt.Println("entro a la validacion por value en true, para la key", key, "contrnaformacion ", keyPeticion)
			if fmt.Sprintf("%v", cantodadPersonas[keyPeticion]) == "0" {
				listaParticipantes = listaParticipantes + ", " + keyPeticion
				LlenarListaIndividual(fechaDECorte, keyPeticion, fmt.Sprintf("%v", votacion["TIV_CODIGO"]))
			} else {
				fmt.Println("LA KEY ", keyPeticion, "TENIA UN VALOR DIFERENTE A CERO PARTICIPANTES")
			}
		}
		// if value == false && key != "TVI_ESTADO" {
		// 	fmt.Println("entro a la validacion por value en false, para la key", key, "contrnaformacion ", keyPeticion)
		// }
	}
	return listaParticipantes

}

// ObtenerFechaCorte ...
func ObtenerFechaCorte(fecha interface{}) string {
	fechaString := fmt.Sprintf("%v", fecha)
	fechaString = strings.ToLower(fechaString)
	// fmt.Println(fecha, fechaString, strings.Split(fechaString, "t")[0])
	FechaCompleta := strings.Split(fechaString, "t")[0]
	arrayFecha := strings.Split(FechaCompleta, "-")
	fechaString = arrayFecha[0] + arrayFecha[1] + arrayFecha[2]
	return fechaString
}

// KeyPeticion ... trnaforma la key para que se pueda usar como parametro para la peticion y sea dinamico
func KeyPeticion(key string) string {
	var keyEnviada string
	keyString := strings.ToLower(key)
	arraykey := strings.Split(keyString, "tvi_")
	for i := 0; i < len(arraykey); i++ {
		if i == (len(arraykey) - 1) {
			keyEnviada = arraykey[i]
		} else {
			keyEnviada = arraykey[i] + "_"
		}
	}

	return keyEnviada
}
