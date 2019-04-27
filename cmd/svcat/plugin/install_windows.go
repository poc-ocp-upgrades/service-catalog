package plugin

import "os"

func getUserHomeDir() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return os.Getenv("USERPROFILE")
}
func getFileExt() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ".exe"
}
