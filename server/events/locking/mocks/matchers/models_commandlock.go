// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
)

func AnyModelsCommandLock() models.CommandLock {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(models.CommandLock))(nil)).Elem()))
	var nullValue models.CommandLock
	return nullValue
}

func EqModelsCommandLock(value models.CommandLock) models.CommandLock {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue models.CommandLock
	return nullValue
}
