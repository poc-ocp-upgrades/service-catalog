package controller

import (
	"encoding/json"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
)

const (
	originatingIdentityPlatform = "kubernetes"
)

func buildOriginatingIdentity(userInfo *v1beta1.UserInfo) (*osb.OriginatingIdentity, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if userInfo == nil {
		return nil, nil
	}
	oiValue, err := json.Marshal(userInfo)
	if err != nil {
		return nil, err
	}
	oi := &osb.OriginatingIdentity{Platform: originatingIdentityPlatform, Value: string(oiValue)}
	return oi, nil
}
