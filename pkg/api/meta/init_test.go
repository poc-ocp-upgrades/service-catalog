package meta

import (
	"testing"
)

func TestGetAccessor(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if GetAccessor() != accessor {
		t.Fatalf("GetAccessor didn't return the pre-initialized accessor")
	}
}
