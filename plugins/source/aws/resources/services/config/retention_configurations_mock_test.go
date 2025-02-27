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

func buildRetentionConfigurations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockConfigserviceClient(ctrl)
	l := types.RetentionConfiguration{}
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeRetentionConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&configservice.DescribeRetentionConfigurationsOutput{
			RetentionConfigurations: []types.RetentionConfiguration{l},
		}, nil)
	return client.Services{
		Configservice: m,
	}
}

func TestRetentionConfigurations(t *testing.T) {
	client.AwsMockTestHelper(t, RetentionConfigurations(), buildRetentionConfigurations, client.TestOptions{})
}
