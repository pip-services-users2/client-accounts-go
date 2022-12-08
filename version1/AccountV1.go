package version1

import (
	"time"
)

type AccountV1 struct {
	/* Identification */
	Id    string `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`

	/* Activity tracking */
	CreateTime time.Time `json:"create_time"`
	Deleted    bool      `json:"deleted"`
	Active     bool      `json:"active"`

	/* User preferences */
	About    string `json:"about"`
	TimeZone string `json:"time_zone"`
	Language string `json:"language"`
	Theme    string `json:"theme"`

	/* Custom fields */
	CustomHdr interface{} `json:"custom_hdr"`
	CustomDat interface{} `json:"custom_dat"`
}

func EmptyAccountV1() *AccountV1 {
	return &AccountV1{}
}

func NewAccountV1(id string, login string, name string) *AccountV1 {
	return &AccountV1{
		Id:         id,
		Login:      login,
		Name:       name,
		CreateTime: time.Now(),
		Active:     true,
		Deleted:    false,
	}
}
