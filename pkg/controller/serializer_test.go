package controller

import (
	"encoding/json"
	"reflect"
	"testing"
	"bytes"
	"encoding/base64"
	"github.com/google/gofuzz"
)

const fuzzIters = 20

var fuzzer = fuzz.New()

func TestSerializeInt(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < fuzzIters; i++ {
		var intVal int
		fuzzer.Fuzz(&intVal)
		bytes, err := serialize(intVal)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		var intValPrime int
		err = json.Unmarshal(bytes, &intValPrime)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if intVal != intValPrime {
			t.Fatalf("Round trip failed; expected %v; got %v", intVal, intValPrime)
		}
	}
}
func TestSerializeFloat(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < fuzzIters; i++ {
		var floatVal float64
		fuzzer.Fuzz(&floatVal)
		bytes, err := serialize(floatVal)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		var floatValPrime float64
		err = json.Unmarshal(bytes, &floatValPrime)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if floatVal != floatValPrime {
			t.Fatalf("Round trip failed; expected %v; got %v", floatVal, floatValPrime)
		}
	}
}
func TestSerializeString(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < fuzzIters; i++ {
		var strVal string
		fuzzer.Fuzz(&strVal)
		bytes, err := serialize(strVal)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		strValPrime := string(bytes)
		if strVal != strValPrime {
			t.Fatalf("Round trip failed; expected %v; got %v", strVal, strValPrime)
		}
	}
}
func TestSerializeMap(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var mapVal map[string]string
	for i := 0; i < fuzzIters; i++ {
		fuzzer.Fuzz(&mapVal)
		bytes, err := serialize(mapVal)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		mapValPrime := make(map[string]string)
		err = json.Unmarshal(bytes, &mapValPrime)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if !reflect.DeepEqual(mapVal, mapValPrime) {
			t.Fatalf("Round trip failed; expected %v; got %v", mapVal, mapValPrime)
		}
	}
}
func TestSerializeSlice(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var sliceVal []string
	for i := 0; i < fuzzIters; i++ {
		fuzzer.Fuzz(&sliceVal)
		bytes, err := serialize(sliceVal)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		sliceValPrime := make([]string, 4)
		err = json.Unmarshal(bytes, &sliceValPrime)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if !reflect.DeepEqual(sliceVal, sliceValPrime) {
			t.Fatalf("Round trip failed; expected %v; got %v", sliceVal, sliceValPrime)
		}
	}
}
func TestSerializeByteSlice(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < fuzzIters; i++ {
		var byteSliceVal []byte
		fuzzer.Fuzz(&byteSliceVal)
		serializedBytes, err := serialize(byteSliceVal)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		base64EncodedBytes := []byte(`"` + base64.StdEncoding.EncodeToString(serializedBytes) + `"`)
		var byteSlicePrime []byte
		err = json.Unmarshal(base64EncodedBytes, &byteSlicePrime)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if !bytes.Equal(byteSliceVal, byteSlicePrime) {
			t.Fatalf("Round trip failed; expected %v; got %v", byteSliceVal, byteSlicePrime)
		}
	}
}
