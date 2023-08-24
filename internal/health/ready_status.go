package health

import (
	"net/http"
	"sync/atomic"
)

type appStatus uint32

const (
	NotReady appStatus = iota
	Ready
)

type ReadyStatus struct {
	status uint32
}

func NewReadyStatus() *ReadyStatus {
	return &ReadyStatus{
		status: uint32(NotReady),
	}
}

func (s *ReadyStatus) IsReady() int {
	if atomic.LoadUint32(&s.status) == uint32(Ready) {
		return http.StatusOK
	}

	return http.StatusServiceUnavailable
}

func (s *ReadyStatus) SetReady() {
	atomic.StoreUint32(&s.status, uint32(Ready))
}
