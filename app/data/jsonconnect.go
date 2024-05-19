package data

import "github.com/teranixbq/gfunc"

func InitJson() (*gfunc.Query, error ){
	JsonConnect, err := gfunc.NewJsonUrl("https://raw.githubusercontent.com/teranixbq/HOK-API/main/app/data/data.json")
	if err != nil {
		 return nil,err
	}

	return JsonConnect,nil
}