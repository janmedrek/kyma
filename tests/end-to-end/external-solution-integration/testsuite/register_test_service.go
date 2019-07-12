package testsuite

import (
	"encoding/json"
	"github.com/kyma-project/kyma/tests/end-to-end/external-solution-integration/consts"
	"github.com/kyma-project/kyma/tests/end-to-end/external-solution-integration/step"
	"github.com/kyma-project/kyma/tests/end-to-end/external-solution-integration/testkit"
)

type RegisterTestService struct {
	testService    *testkit.TestService
	state          RegisterTestServiceState
}

type RegisterTestServiceState interface {
	SetServiceID(serviceID string)
	GetServiceID() string
	GetRegistryClient() *testkit.RegistryClient
}

var _ step.Step = &RegisterTestService{}

func NewRegisterTestService(testService *testkit.TestService, state RegisterTestServiceState) *RegisterTestService {
	return &RegisterTestService{
		testService:    testService,
		state:          state,
	}
}

func (s *RegisterTestService) Name() string {
	return "Register test service"
}

func (s *RegisterTestService) Run() error {
	url := s.testService.GetTestServiceURL()
	service := s.prepareService(url)

	id, err := s.state.GetRegistryClient().RegisterService(service)
	if err != nil {
		return err
	}

	s.state.SetServiceID(id)
	return nil
}

func (s *RegisterTestService) prepareService(targetURL string) *testkit.ServiceDetails {
	return &testkit.ServiceDetails{
		Provider:         consts.ServiceProvider,
		Name:             consts.ServiceName,
		Description:      consts.ServiceDescription,
		ShortDescription: consts.ServiceShortDescription,
		Identifier:       consts.ServiceIdentifier,
		Events: &testkit.Events{
			Spec: json.RawMessage(consts.ServiceEventsSpec),
		},
		Api: &testkit.API{
			TargetUrl: targetURL,
		},
	}
}

func (s *RegisterTestService) Cleanup() error {
	if serviceID := s.state.GetServiceID(); serviceID != "" {
		return s.state.GetRegistryClient().DeleteService(serviceID)
	}
	return nil
}
