package strto

import uuid "github.com/satori/go.uuid"

func GetUUID1() string {
	return uuid.NewV1().String()
}

// DomainPerson = iota
//	DomainGroup
//	DomainOrg
func GetUUID2(domain byte) string {
	return uuid.NewV2(domain).String()
}

func GetUUID3() string {
	return uuid.NewV3(uuid.UUID{}, "").String()
}

func GetUUID4() string {
	return uuid.NewV4().String()
}

func GetUUID5() string {
	return uuid.NewV5(uuid.UUID{}, "").String()
}
