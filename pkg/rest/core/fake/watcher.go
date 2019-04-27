package fake

import (
	"fmt"
	"time"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

type Watcher struct{ ch chan watch.Event }

func NewWatcher() *Watcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Watcher{ch: make(chan watch.Event)}
}
func (w *Watcher) SendObject(evtType watch.EventType, obj runtime.Object, timeout time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	evt := watch.Event{Type: evtType, Object: obj}
	select {
	case w.ch <- evt:
	case <-time.After(timeout):
		return fmt.Errorf("couldn't send after %s", timeout)
	}
	return nil
}
func (w *Watcher) ReceiveChan() <-chan watch.Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return w.ch
}
func (w *Watcher) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(w.ch)
}
