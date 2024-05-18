package data

import "github.com/teranixbq/gfunc"

func InitJson() (*gfunc.Query, error ){
	JsonConnect, err := gfunc.NewJsonFile("app/data/data.json")
	if err != nil {
		 return nil,err
	}

	return JsonConnect,nil
}