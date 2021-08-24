package control

import (
	"qrcode/access"
	"sync"
)

var (
	apiControl APIControl
	apiOnce    sync.Once
)

type APIControl struct {
	access *access.Access
}


func APICreate(access *access.Access) *APIControl {
	apiOnce.Do(func() {
		apiControl = APIControl{access: access}
	})
	return &apiControl
}

//func (ctrl *APIControl) LogMode() protobuf.GRPCLogMode {
//	return ctrl.access.ENV.LogMode
//}
