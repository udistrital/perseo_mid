package models

import "github.com/astaxie/beego"

// ObtenerCantidades ...
func ObtenerCantidades(codigo string) (votaciones []map[string]interface{}, outputError interface{}) {
	ArrayCantidades := []map[string]interface{}{}

	DocPlanta, errPLanta := ObtenerCantDocPlanta(codigo)
	if DocPlanta != nil {
		DocVE, errVE := ObtenerCantDocVE(codigo)
		if DocVE != nil {
			Egresados, errEgre := ObtenerEgresados(codigo)
			if Egresados != nil {
				Administrativos, errAdm := ObtenerAdministrativos(codigo)
				if Administrativos != nil {
					Estudiantes, errEst := ObtenerEstudiantes(codigo)
					if Estudiantes != nil {
						ArrayCantidades = append(ArrayCantidades, map[string]interface{}{
							"docentes_planta": DocPlanta["cantidades"],
							"docentes_ve":     DocVE["cantidades"],
							"egresados":       Egresados["cantidades"],
							"administrativos": Administrativos["cantidades"],
							"estudiantes":     Estudiantes["cantidades"],
						})
						return ArrayCantidades, nil
					}
					return nil, errEst

				}
				return nil, errAdm

			}
			return nil, errEgre

		}
		return nil, errVE

	}
	return nil, errPLanta
}

// ObtenerCantDocPlanta ...
func ObtenerCantDocPlanta(codigo string) (cantidad map[string]interface{}, outputError interface{}) {
	var CantidadPersonas map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+beego.AppConfig.String("perseo_ns_service")+"cantidadDocentesPlanta/"+codigo, &CantidadPersonas)
	if error != nil {
		return nil, error
	}
	cantDocPlanta, errCantDocPlanta := ObtenerValorCantidad(CantidadPersonas["docentes_planta"])
	if cantDocPlanta != nil {
		return cantDocPlanta, nil
	}
	return nil, errCantDocPlanta
}

// ObtenerCantDocVE ...
func ObtenerCantDocVE(codigo string) (cantidad map[string]interface{}, outputError interface{}) {
	var CantidadPersonas map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+beego.AppConfig.String("perseo_ns_service")+"cantidadDocentesVE/"+codigo, &CantidadPersonas)
	if error != nil {
		return nil, error
	}
	cantDocPlanta, errCantDocPlanta := ObtenerValorCantidad(CantidadPersonas["docentes_ve"])
	if cantDocPlanta != nil {
		return cantDocPlanta, nil
	}
	return nil, errCantDocPlanta
}

// ObtenerEgresados ...
func ObtenerEgresados(codigo string) (cantidad map[string]interface{}, outputError interface{}) {
	var CantidadPersonas map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+beego.AppConfig.String("perseo_ns_service")+"cantidadEgresados/"+codigo, &CantidadPersonas)
	if error != nil {
		return nil, error
	}
	cantDocPlanta, errCantDocPlanta := ObtenerValorCantidad(CantidadPersonas["egresados"])
	if cantDocPlanta != nil {
		return cantDocPlanta, nil
	}
	return nil, errCantDocPlanta
}

// ObtenerAdministrativos ...
func ObtenerAdministrativos(codigo string) (cantidad map[string]interface{}, outputError interface{}) {
	var CantidadPersonas map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+beego.AppConfig.String("perseo_ns_service")+"cantidadAdministrativos/"+codigo, &CantidadPersonas)
	if error != nil {
		return nil, error
	}
	cantDocPlanta, errCantDocPlanta := ObtenerValorCantidad(CantidadPersonas["administrativos"])
	if cantDocPlanta != nil {
		return cantDocPlanta, nil
	}
	return nil, errCantDocPlanta
}

// ObtenerEstudiantes ...
func ObtenerEstudiantes(codigo string) (cantidad map[string]interface{}, outputError interface{}) {
	var CantidadPersonas map[string]interface{}
	error := GetJSONJBPM(beego.AppConfig.String("administrativa_amazon_jbpm_url")+beego.AppConfig.String("perseo_ns_service")+"cantidadEstudiantes/"+codigo, &CantidadPersonas)
	if error != nil {
		return nil, error
	}
	cantDocPlanta, errCantDocPlanta := ObtenerValorCantidad(CantidadPersonas["estudiantes_cant"])
	if cantDocPlanta != nil {
		return cantDocPlanta, nil
	}
	return nil, errCantDocPlanta
}

// ObtenerValorCantidad ...
func ObtenerValorCantidad(objeto interface{}) (cantidad map[string]interface{}, outputError interface{}) {
	arraycantidad, errCant := GetElementoMaptoStringToMapArray(objeto.(map[string]interface{})["cantidad"])
	if arraycantidad != nil {

		return arraycantidad[0], nil
	}
	return nil, errCant
}
