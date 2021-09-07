package access

import (
	"qrcode/access/constant"
	"qrcode/access/rdbms"
	"qrcode/access/serviceline"
	"qrcode/environment"
)

type Access struct {
	ENV *environment.Properties
	RDBMS rdbms.FactoryInterface
	TEMPLATE constant.Templates
	SERVICELINE serviceline.FactoryServiceLine
	//GRPC grpc.FactoryInterface
}

func Initial(properties *environment.Properties) *Access {
	return &Access{
		ENV:   properties,
		RDBMS: rdbms.Create(properties),
		SERVICELINE: serviceline.ServiceLine(properties),
		//GRPC: grpc.Create(properties),
	}
}