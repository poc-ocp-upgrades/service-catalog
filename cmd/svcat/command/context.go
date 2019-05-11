package command

import (
	"io"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat"
	"github.com/spf13/viper"
)

type Context struct {
	Output	io.Writer
	App		*svcat.App
	Viper	*viper.Viper
}
