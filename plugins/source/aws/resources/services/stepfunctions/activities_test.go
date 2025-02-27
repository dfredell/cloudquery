package stepfunctions

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildActivities(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockSfnClient(ctrl)
	ali := types.ActivityListItem{}
	err := faker.FakeObject(&ali)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListActivities(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sfn.ListActivitiesOutput{
			Activities: []types.ActivityListItem{ali},
		}, nil)

	return client.Services{
		Sfn: m,
	}
}

func TestActivities(t *testing.T) {
	client.AwsMockTestHelper(t, Activities(), buildActivities, client.TestOptions{})
}
