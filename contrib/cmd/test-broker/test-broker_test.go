package main

import "testing"

func TestInit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
