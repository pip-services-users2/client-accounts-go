package version1

import (
	"encoding/json"

	"github.com/pip-services-users2/client-accounts-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]interface{}) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]interface{} {
	var r map[string]interface{}

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value interface{}) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) interface{} {
	if value == "" {
		return nil
	}

	var m interface{}
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromAccount(account *AccountV1) *protos.Account {
	if account == nil {
		return nil
	}

	obj := &protos.Account{
		Id:         account.Id,
		Login:      account.Login,
		Name:       account.Name,
		About:      account.About,
		CreateTime: convert.StringConverter.ToString(account.CreateTime),
		Deleted:    account.Deleted,
		Active:     account.Active,
		TimeZone:   account.TimeZone,
		Language:   account.Language,
		Theme:      account.Theme,
		CustomHdr:  toJson(account.CustomHdr),
		CustomDat:  toJson(account.CustomDat),
	}

	return obj
}

func toAccount(obj *protos.Account) *AccountV1 {
	if obj == nil {
		return nil
	}

	account := &AccountV1{
		Id:         obj.Id,
		Login:      obj.Login,
		Name:       obj.Name,
		About:      obj.About,
		CreateTime: convert.DateTimeConverter.ToDateTime(obj.CreateTime),
		Deleted:    obj.Deleted,
		Active:     obj.Active,
		TimeZone:   obj.TimeZone,
		Language:   obj.Language,
		Theme:      obj.Theme,
		CustomHdr:  fromJson(obj.CustomHdr),
		CustomDat:  fromJson(obj.CustomDat),
	}

	return account
}

func fromAccountPage(page *data.DataPage[*AccountV1]) *protos.AccountPage {
	if page == nil {
		return nil
	}

	obj := &protos.AccountPage{
		Total: int64(page.Total),
		Data:  make([]*protos.Account, len(page.Data)),
	}

	for i, account := range page.Data {
		obj.Data[i] = fromAccount(account)
	}

	return obj
}

func toAccountPage(obj *protos.AccountPage) data.DataPage[*AccountV1] {
	if obj == nil {
		return *data.NewEmptyDataPage[*AccountV1]()
	}

	accounts := make([]*AccountV1, len(obj.Data))

	for i, v := range obj.Data {
		accounts[i] = toAccount(v)
	}

	page := data.NewDataPage(accounts, int(obj.Total))

	return *page
}
