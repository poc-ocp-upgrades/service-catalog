package output

import (
	"fmt"
	"io"
)

func WriteClientVersion(w io.Writer, client string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(w, "Client Version: %s\n", client)
}
func WriteServerVersion(w io.Writer, server string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(w, "Server Version: %s\n", server)
}
