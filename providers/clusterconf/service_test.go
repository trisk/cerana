package clusterconf_test

import (
	"path"

	"github.com/cerana/cerana/acomm"
	"github.com/cerana/cerana/providers/clusterconf"
	"github.com/pborman/uuid"
)

func (s *clusterConf) TestGetService() {
	service, err := s.addService()
	s.Require().NoError(err)

	tests := []struct {
		id  string
		err string
	}{
		{"", "missing arg: id"},
		{"does-not-exist", "service config not found"},
		{service.ID, ""},
	}

	for _, test := range tests {
		req, err := acomm.NewRequest(acomm.RequestOptions{
			Task: "get-service",
			Args: &clusterconf.IDArgs{ID: test.id},
		})
		s.Require().NoError(err, test.id)
		result, streamURL, err := s.clusterConf.GetService(req)
		s.Nil(streamURL, test.id)
		if test.err != "" {
			s.Contains(err.Error(), test.err, test.id)
			s.Nil(result, test.id)
		} else {
			s.NoError(err, test.id)
			if !s.NotNil(result, test.id) {
				continue
			}
			servicePayload, ok := result.(*clusterconf.ServicePayload)
			s.True(ok, test.id)
			s.Equal(test.id, servicePayload.Service.ID, test.id)
		}
	}
}

func (s *clusterConf) TestUpdateService() {
	service, err := s.addService()
	s.Require().NoError(err)
	service2, err := s.addService()
	s.Require().NoError(err)

	tests := []struct {
		desc     string
		id       string
		modIndex uint64
		err      string
	}{
		{"no id", "", 0, ""},
		{"new id", uuid.New(), 0, ""},
		{"create existing id", service.ID, 0, "CAS failed"},
		{"update existing id", service2.ID, service2.ModIndex, ""},
	}

	for _, test := range tests {
		req, err := acomm.NewRequest(acomm.RequestOptions{
			Task: "update-service",
			Args: &clusterconf.ServicePayload{
				Service: &clusterconf.Service{
					ServiceConf: clusterconf.ServiceConf{ID: test.id},
					ModIndex:    test.modIndex,
				},
			},
		})
		s.Require().NoError(err, test.desc)
		result, streamURL, err := s.clusterConf.UpdateService(req)
		s.Nil(streamURL, test.desc)
		if test.err != "" {
			s.Contains(err.Error(), test.err, test.desc)
			s.Nil(result, test.desc)
		} else {
			s.NoError(err, test.desc)
			if !s.NotNil(result, test.desc) {
				continue
			}
			servicePayload, ok := result.(*clusterconf.ServicePayload)
			s.True(ok, test.id)
			if test.id == "" {
				s.NotEmpty(servicePayload.Service.ID, test.desc)
			} else {
				s.Equal(test.id, servicePayload.Service.ID, test.desc)
			}
			s.NotEqual(test.modIndex, servicePayload.Service.ModIndex, test.desc)
		}
	}
}

func (s *clusterConf) TestDeleteService() {
	service, err := s.addService()
	s.Require().NoError(err)

	tests := []struct {
		id  string
		err string
	}{
		{"", "missing arg: id"},
		{"does-not-exist", ""},
		{service.ID, ""},
	}

	for _, test := range tests {
		req, err := acomm.NewRequest(acomm.RequestOptions{
			Task: "delete-service",
			Args: &clusterconf.IDArgs{ID: test.id},
		})
		s.Require().NoError(err, test.id)
		result, streamURL, err := s.clusterConf.DeleteService(req)
		s.Nil(streamURL, test.id)
		s.Nil(result, test.id)
		if test.err != "" {
			s.Contains(err.Error(), test.err, test.id)
		} else {
			s.NoError(err, test.id)
		}
	}
}

func (s *clusterConf) addService() (*clusterconf.Service, error) {
	// Populate a service
	service := &clusterconf.Service{ServiceConf: clusterconf.ServiceConf{
		ID:      uuid.New(),
		Dataset: "testds",
	}}
	key := path.Join("services", service.ID, "config")
	data := map[string]interface{}{
		key: service,
	}
	indexes, err := s.loadData(data)
	if err != nil {
		return nil, err
	}

	service.ModIndex = indexes[key]
	return service, nil
}
