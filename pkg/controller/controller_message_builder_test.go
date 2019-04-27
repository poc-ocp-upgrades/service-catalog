package controller

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
)

func normalEventBuilder(reason string) *MessageBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return new(MessageBuilder).normal().reason(reason)
}
func warningEventBuilder(reason string) *MessageBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return new(MessageBuilder).warning().reason(reason)
}

type MessageBuilder struct {
	eventMessage	string
	reasonMessage	string
	message		string
}

func (mb *MessageBuilder) normal() *MessageBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mb.eventMessage = corev1.EventTypeNormal
	return mb
}
func (mb *MessageBuilder) warning() *MessageBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mb.eventMessage = corev1.EventTypeWarning
	return mb
}
func (mb *MessageBuilder) reason(reason string) *MessageBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mb.reasonMessage = reason
	return mb
}
func (mb *MessageBuilder) msg(msg string) *MessageBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	space := ""
	if mb.message > "" {
		space = " "
	}
	mb.message = fmt.Sprintf(`%s%s%s`, mb.message, space, msg)
	return mb
}
func (mb *MessageBuilder) msgf(format string, a ...interface{}) *MessageBuilder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	msg := fmt.Sprintf(format, a...)
	return mb.msg(msg)
}
func (mb *MessageBuilder) stringArr() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{mb.String()}
}
func (mb *MessageBuilder) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := ""
	space := ""
	if mb.eventMessage > "" {
		s += fmt.Sprintf("%s%s", space, mb.eventMessage)
		space = " "
	}
	if mb.reasonMessage > "" {
		s += fmt.Sprintf("%s%s", space, mb.reasonMessage)
		space = " "
	}
	if mb.message > "" {
		s += fmt.Sprintf("%s%s", space, mb.message)
		space = " "
	}
	return s
}
