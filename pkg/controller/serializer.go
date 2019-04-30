package controller

import (
	"encoding/json"
)

func serialize(value interface{}) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if byteArrayVal, ok := value.([]byte); ok {
		return byteArrayVal, nil
	}
	if strVal, ok := value.(string); ok {
		return []byte(strVal), nil
	}
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return data, nil
}
