package access

import (
	"qrcode/access/constant"
	"qrcode/access/rdbms"
	"qrcode/environment"
)

type Access struct {
	ENV *environment.Properties
	RDBMS rdbms.FactoryInterface
	TEMPLATE constant.Templates
	//GRPC grpc.FactoryInterface
}

func Initial(properties *environment.Properties) *Access {
	return &Access{
		ENV:   properties,
		RDBMS: rdbms.Create(properties),
		//GRPC: grpc.Create(properties),
	}
}