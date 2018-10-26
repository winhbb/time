package time

import "time"

type CancelableTimer struct {
	*time.Timer
	cancel chan struct{}
}

func NewCancelableTimer(d time.Duration) *CancelableTimer {
	return &CancelableTimer{
		time.NewTimer(d),
		make(chan struct{}),
	}
}

//结束等待有两种条件，一种是timer到时间自动结束，另一种是手动中取消
func (t *CancelableTimer) Wait() time.Time {
	select {
	case now := <-t.Timer.C:
		return now
	case <-t.cancel:
		return time.Now()
	}
}

func (t *CancelableTimer) Stop() bool {
	t.cancel <- struct{}{}
	return t.Timer.Stop()
}
