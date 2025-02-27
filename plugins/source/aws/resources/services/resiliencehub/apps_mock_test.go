package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildApps(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockResiliencehubClient(ctrl)
	var l resiliencehub.ListAppsOutput
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	l.NextToken = nil
	mock.EXPECT().ListApps(
		gomock.Any(),
		&resiliencehub.ListAppsInput{},
		gomock.Any(),
	).Return(&l, nil)

	var d resiliencehub.DescribeAppOutput
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeApp(
		gomock.Any(),
		&resiliencehub.DescribeAppInput{AppArn: l.AppSummaries[0].AppArn},
		gomock.Any(),
	).Return(&d, nil)

	buildAppAssessments(t, mock)
	buildAppVersions(t, mock)
	return client.Services{Resiliencehub: mock}
}

func TestResiilencehubApps(t *testing.T) {
	client.AwsMockTestHelper(t, Apps(), buildApps, client.TestOptions{})
}
