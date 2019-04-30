package command

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"
)

type HasWaitFlags interface{ ApplyWaitFlags() error }
type Waitable struct {
	Wait		bool
	rawTimeout	string
	Timeout		*time.Duration
	rawInterval	string
	Interval	time.Duration
}

func NewWaitable() *Waitable {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Waitable{}
}
func (c *Waitable) AddWaitFlags(cmd *cobra.Command) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd.Flags().BoolVar(&c.Wait, "wait", false, "Wait until the operation completes.")
	cmd.Flags().StringVar(&c.rawTimeout, "timeout", "5m", "Timeout for --wait, specified in human readable format: 30s, 1m, 1h. Specify -1 to wait indefinitely.")
	cmd.Flags().StringVar(&c.rawInterval, "interval", "1s", "Poll interval for --wait, specified in human readable format: 30s, 1m, 1h")
}
func (c *Waitable) ApplyWaitFlags() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !c.Wait {
		return nil
	}
	if c.rawTimeout != "-1" {
		timeout, err := time.ParseDuration(c.rawTimeout)
		if err != nil {
			return fmt.Errorf("invalid --timeout value (%s)", err)
		}
		c.Timeout = &timeout
	}
	interval, err := time.ParseDuration(c.rawInterval)
	if err != nil {
		return fmt.Errorf("invalid --interval value (%s)", err)
	}
	c.Interval = interval
	return nil
}
