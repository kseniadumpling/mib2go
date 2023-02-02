// Code generated by mib2go. DO NOT EDIT.
package mibs

import (
	"github.com/sleepinggenius2/gosmi/models"
	"github.com/sleepinggenius2/gosmi/types"
)

type snmpv2TcModule struct {
	RowStatusType   models.Type
	StorageTypeType models.Type
}

var Snmpv2Tc = snmpv2TcModule{
	RowStatusType:   snmpv2TcRowStatusType,
	StorageTypeType: snmpv2TcStorageTypeType,
}

var snmpv2TcRowStatusType = models.Type{
	BaseType: types.BaseTypeEnum,
	Enum: &models.Enum{
		BaseType: types.BaseTypeInteger32,
		Values: []models.NamedNumber{
			models.NamedNumber{Name: "active", Value: 1},
			models.NamedNumber{Name: "notInService", Value: 2},
			models.NamedNumber{Name: "notReady", Value: 3},
			models.NamedNumber{Name: "createAndGo", Value: 4},
			models.NamedNumber{Name: "createAndWait", Value: 5},
			models.NamedNumber{Name: "destroy", Value: 6},
		},
	},
	Name: "RowStatus",
}

var snmpv2TcStorageTypeType = models.Type{
	BaseType: types.BaseTypeEnum,
	Enum: &models.Enum{
		BaseType: types.BaseTypeInteger32,
		Values: []models.NamedNumber{
			models.NamedNumber{Name: "other", Value: 1},
			models.NamedNumber{Name: "volatile", Value: 2},
			models.NamedNumber{Name: "nonVolatile", Value: 3},
			models.NamedNumber{Name: "permanent", Value: 4},
			models.NamedNumber{Name: "readOnly", Value: 5},
		},
	},
	Name: "StorageType",
}

