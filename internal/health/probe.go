package health

import (
	"net/http"
	"sync/atomic"
)

type probeStatus uint32

const (
	ProbeStatusOK probeStatus = iota
	ProbeStatusFailed
)

type Probe struct {
	status uint32
}

func NewProbe() *Probe {
	return &Probe{
		status: uint32(ProbeStatusOK),
	}
}

func (p *Probe) GetStatus() int {
	if atomic.LoadUint32(&p.status) == uint32(ProbeStatusOK) {
		return http.StatusOK
	}

	return http.StatusServiceUnavailable
}

func (p *Probe) SetStatus(status probeStatus) {
	atomic.StoreUint32(&p.status, uint32(status))
}
