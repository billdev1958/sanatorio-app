package pkg

import (
	"fmt"
	"reflect"
	"strings"
)

// TODO FILTERS ADD
type Filters map[string]interface{}

func Filter(params interface{}) (Filters, error) {
	filters := Filters{}

	value := reflect.ValueOf(params)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("params debe ser una estructura o un puntero a una estructura; se recibió %s", value.Kind())
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)

		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}

		// Dividir la etiqueta JSON para eliminar ",omitempty"
		tagParts := strings.Split(tag, ",")
		columnName := tagParts[0]

		if columnName == "" {
			continue
		}

		if fieldValue.IsValid() && !fieldValue.IsZero() {
			filters[columnName] = fieldValue.Interface()
		}
	}
	return filters, nil
}

func BuildWhereClause(filters map[string]interface{}) (string, []interface{}, error) {
	// Si no hay filtros, devolvemos una cláusula vacía y sin errores
	if len(filters) == 0 {
		return "", nil, nil
	}

	whereClause := ""
	args := []interface{}{}
	argIndex := 1

	// Iterar sobre los filtros
	for key, value := range filters {
		// Validar que las claves sean cadenas no vacías
		if key == "" {
			return "", nil, fmt.Errorf("clave de filtro vacía detectada")
		}

		// Validar que los valores no sean nulos
		if reflect.ValueOf(value).IsZero() {
			return "", nil, fmt.Errorf("valor nulo o vacío para la clave: %s", key)
		}

		// Construir la cláusula WHERE
		if whereClause == "" {
			whereClause = "WHERE"
		} else {
			whereClause += " AND"
		}

		// Agregar la condición con marcador de posición
		whereClause += fmt.Sprintf(" %s = $%d", key, argIndex)
		args = append(args, value)
		argIndex++
	}

	return whereClause, args, nil
}

func MapFiltersToColumns(filters map[string]interface{}, columnMapping map[string]string) (map[string]interface{}, error) {
	dbFilters := make(map[string]interface{})

	for key, value := range filters {
		if columnName, ok := columnMapping[key]; ok {
			dbFilters[columnName] = value
		} else {
			return nil, fmt.Errorf("filtro no reconocido: %s", key)
		}
	}

	return dbFilters, nil
}
