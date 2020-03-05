package mt

import (
	"context"
	"github.com/sirupsen/logrus"
	"net"
	"sync"
)

type targetsServer struct {
	mu          *sync.Mutex
	listenerMap map[string]net.Listener
}

func (t targetsServer) List(context.Context, *ListTargetParams) (*ListTargetResponse, error) {
	logrus.Info("List targets")
	var resp ListTargetResponse
	for addr := range t.listenerMap {
		resp.Targets = append(resp.Targets, &Target{Addr: addr})
	}
	return &resp, nil
}

func (t *targetsServer) New(_ context.Context, p *NewTargetParams) (*NewTargetResponse, error) {
	logrus.WithField("target", p.Addr).Info("New target")
	var resp NewTargetResponse

	t.mu.Lock()
	defer t.mu.Unlock()

	_, ok := t.listenerMap[p.Addr]
	if ok {
		logrus.WithField("target", p.Addr).Warn("target already exist")
		return &resp, nil
	}

	ln, err := net.Listen("tcp", p.Addr)
	if err != nil {
		return nil, err
	}
	logrus.WithField("target", p.Addr).Info("run target")
	go t.runTarget(p.Addr, ln)
	t.listenerMap[p.Addr] = ln
	return &resp, nil
}

func (t *targetsServer) runTarget(addr string, ln net.Listener) {
	logger := logrus.WithField("target", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Error(err)
			continue
		}
		logger.
			WithField("remote_addr", conn.RemoteAddr()).
			Info("connection accepted")
		conn.Write([]byte("OK"))
		conn.Close()
	}
}

func NewServer() TargetsServer {
	return &targetsServer{
		mu:          new(sync.Mutex),
		listenerMap: map[string]net.Listener{},
	}
}
