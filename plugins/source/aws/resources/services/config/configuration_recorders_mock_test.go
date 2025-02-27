package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildConfigConfigurationRecorders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigserviceClient(ctrl)
	l := types.ConfigurationRecorder{}
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	sl := types.ConfigurationRecorderStatus{}
	if err := faker.FakeObject(&sl); err != nil {
		t.Fatal(err)
	}
	sl.Name = l.Name
	m.EXPECT().DescribeConfigurationRecorderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConfigurationRecorderStatusOutput{
			ConfigurationRecordersStatus: []types.ConfigurationRecorderStatus{sl},
		}, nil)
	m.EXPECT().DescribeConfigurationRecorders(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeConfigurationRecordersOutput{
			ConfigurationRecorders: []types.ConfigurationRecorder{l},
		}, nil)
	return client.Services{
		Configservice: m,
	}
}

func TestConfigConfigurationRecorders(t *testing.T) {
	client.AwsMockTestHelper(t, ConfigurationRecorders(), buildConfigConfigurationRecorders, client.TestOptions{})
}
