package access

import (
	"qrcode/access/file"
	"qrcode/access/rdbms"
	"qrcode/environment"
)

type Access struct {
	ENV *environment.Properties
	RDBMS rdbms.FactoryInterface
	FILE file.FactoryInterface
	//GRPC grpc.FactoryInterface
}

func Initial(properties *environment.Properties) *Access {
	return &Access{
		ENV:   properties,
		RDBMS: rdbms.Create(properties),
		FILE: file.Create(properties),
		//GRPC: grpc.Create(properties),
	}
}