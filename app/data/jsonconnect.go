package data

import "github.com/teranixbq/gfunc"

func InitJson() (*gfunc.Query, error ){
	JsonConnect, err := gfunc.NewJsonUrl("https://storage.apicode.my.id/data.json")
	if err != nil {
		 return nil,err
	}

	return JsonConnect,nil
}