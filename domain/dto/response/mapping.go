package response

import (
	"encoding/json"

	"hokapi/domain/model"
)

func ModelToResponse(data model.Hero) ResponseHero {
	return ResponseHero{
		Id:        data.Id,
		Name:      data.Name,
		Picture:   data.Picture,
		Role:      data.Role,
		Type_Hero: data.Type_Hero,
		Skill:     data.Skill,
	}
}

func ListModelToResponse(data []model.Hero) []ResponseHero {
	list := []ResponseHero{}
	for _, v := range data {
		result := ModelToResponse(v)
		list = append(list, result)
	}
	return list
}

// Mapping interface
func InterfaceToResponse(data interface{}) ResponseHero {
	dataModel := model.Hero{}
	resultMap, _ := data.(map[string]interface{})
	_ = mapToStruct(resultMap, &dataModel)
	return ResponseHero{
		Id:        dataModel.Id,
		Name:      dataModel.Name,
		Picture:   dataModel.Picture,
		Role:      dataModel.Role,
		Type_Hero: dataModel.Type_Hero,
		Skill:     dataModel.Skill,
	}
}

func mapToStruct(m map[string]interface{}, s interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}

func ListInterfaceToResponse(data []interface{}) []ResponseHero {
	list := []ResponseHero{}
	for _, v := range data {
		if m, ok := v.(map[string]interface{}); ok {
			result := InterfaceToResponse(m)
			list = append(list, result)
		}
	}
	return list
}
