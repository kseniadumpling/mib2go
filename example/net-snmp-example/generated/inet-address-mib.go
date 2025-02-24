// Code generated by mib2go. DO NOT EDIT.
package mibs

import (
	"github.com/sleepinggenius2/gosmi/models"
	"github.com/sleepinggenius2/gosmi/types"
)

type inetAddressMibModule struct {
	InetAddressType     models.Type
	InetAddressTypeType models.Type
}

var InetAddressMib = inetAddressMibModule{
	InetAddressType:     inetAddressMibInetAddressType,
	InetAddressTypeType: inetAddressMibInetAddressTypeType,
}

var inetAddressMibInetAddressType = models.Type{
	BaseType: types.BaseTypeOctetString,
	Name:     "InetAddress",
	Ranges: []models.Range{
		models.Range{BaseType: types.BaseTypeUnsigned32, MinValue: 0, MaxValue: 255},
	},
}

var inetAddressMibInetAddressTypeType = models.Type{
	BaseType: types.BaseTypeEnum,
	Enum: &models.Enum{
		BaseType: types.BaseTypeInteger32,
		Values: []models.NamedNumber{
			models.NamedNumber{Name: "unknown", Value: 0},
			models.NamedNumber{Name: "ipv4", Value: 1},
			models.NamedNumber{Name: "ipv6", Value: 2},
			models.NamedNumber{Name: "ipv4z", Value: 3},
			models.NamedNumber{Name: "ipv6z", Value: 4},
			models.NamedNumber{Name: "dns", Value: 16},
		},
	},
	Name: "InetAddressType",
}

