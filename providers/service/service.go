package service

import (
	"github.com/cerana/cerana/acomm"
	"github.com/cerana/cerana/provider"
)

// Provider is a provider of service management functionality.
type Provider struct {
	config  *provider.Config
	tracker *acomm.Tracker
}

// New creates a new instance of Provider.
func New(config *provider.Config, tracker *acomm.Tracker) *Provider {
	return &Provider{
		config:  config,
		tracker: tracker,
	}
}

// RegisterTasks registers all of the provider task handlers with the server.
func (p *Provider) RegisterTasks(server *provider.Server) {
	server.RegisterTask("service-create", p.Create)
	server.RegisterTask("service-get", p.Get)
	server.RegisterTask("service-list", p.List)
	server.RegisterTask("service-restart", p.Restart)
	server.RegisterTask("service-remove", p.Remove)
}

func (p *Provider) executeRequests(requests []*acomm.Request) error {
	for _, req := range requests {
		doneChan := make(chan *acomm.Response, 1)
		defer close(doneChan)
		rh := func(req *acomm.Request, resp *acomm.Response) {
			doneChan <- resp
		}
		req.SuccessHandler = rh
		req.ErrorHandler = rh

		if err := p.tracker.TrackRequest(req, p.config.RequestTimeout()); err != nil {
			return err
		}
		if err := acomm.Send(p.config.CoordinatorURL(), req); err != nil {
			return err
		}

		resp := <-doneChan
		if resp.Error != nil {
			return resp.Error
		}
	}
	return nil
}
