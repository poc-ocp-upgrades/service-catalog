package controller

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
)

func TestBuildOriginatingIdentity(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	userInfo := v1beta1.UserInfo{Username: "person@place.com", UID: "abcd-1234", Groups: []string{"stuff-dev", "main-eng"}, Extra: map[string]v1beta1.ExtraValue{"foo": {"bar", "baz"}}}
	e := osb.OriginatingIdentity{Platform: "kubernetes", Value: `{extra: {"foo":["bar","baz"]},"groups":["stuff-dev","main-eng"],"uid":"abcd-1234","username":"person@place.com"}`}
	g, err := buildOriginatingIdentity(&userInfo)
	if err != nil {
		t.Fatalf("Unexpected Error, %+v", err)
	}
	if e.Platform != g.Platform {
		t.Fatalf("Unexpected Platform, %s", expectedGot(e.Platform, g.Platform))
	}
	var retUserInfo v1beta1.UserInfo
	err = json.Unmarshal([]byte(g.Value), &retUserInfo)
	if err != nil {
		t.Fatalf("Unexpected Error, %+v", err)
	}
	if userInfo.Username != retUserInfo.Username {
		t.Fatalf("Unexpected Value Username, %s", expectedGot(userInfo.Username, retUserInfo.Username))
	}
	if userInfo.UID != retUserInfo.UID {
		t.Fatalf("Unexpected Value UID, %s", expectedGot(userInfo.UID, retUserInfo.UID))
	}
	if !reflect.DeepEqual(userInfo.Groups, retUserInfo.Groups) {
		t.Fatalf("Unexpected Value Groups, %s", expectedGot(fmt.Sprintf("%#v", userInfo.Groups), fmt.Sprintf("%#v", retUserInfo.Groups)))
	}
	if extras, ok := retUserInfo.Extra["foo"]; !ok {
		t.Fatalf("Unexpected Value extras, %s", expectedGot(fmt.Sprintf("%#v", userInfo.Extra), fmt.Sprintf("%#v", retUserInfo.Extra)))
	} else {
		if !reflect.DeepEqual(extras, userInfo.Extra["foo"]) {
			t.Fatalf("Unexpected Value extras, %s", expectedGot(fmt.Sprintf("%#v", userInfo.Extra), fmt.Sprintf("%#v", retUserInfo.Extra)))
		}
	}
}
