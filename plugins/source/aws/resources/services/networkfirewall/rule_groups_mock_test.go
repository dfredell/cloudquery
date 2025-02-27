package networkfirewall

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockNetworkfirewallClient(ctrl)
	rgm := types.RuleGroupMetadata{}
	err := faker.FakeObject(&rgm)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&networkfirewall.ListRuleGroupsOutput{
			RuleGroups: []types.RuleGroupMetadata{rgm},
		}, nil)

	rg := networkfirewall.DescribeRuleGroupOutput{}
	if err := faker.FakeObject(&rg); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).Return(&rg, nil)

	return client.Services{
		Networkfirewall: m,
	}
}

func TestRuleGroups(t *testing.T) {
	client.AwsMockTestHelper(t, RuleGroups(), buildRuleGroupsMock, client.TestOptions{})
}
