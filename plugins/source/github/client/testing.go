package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v49/github"
	"github.com/rs/zerolog"
)

type TestOptions struct{}

func GithubMockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) GithubServices, _ TestOptions) {
	version := "vDev"
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	var cs github.Repository
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	someId := int64(5555555)
	cs.Parent = &github.Repository{ID: &someId}
	cs.TemplateRepository = &github.Repository{ID: &someId}
	cs.Source = &github.Repository{ID: &someId}

	b := builder(t, ctrl)
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		return &Client{
			logger: l,
			Github: b,
			orgServices: map[string]GithubServices{
				"":        b,
				"testorg": b,
			},
			orgs:            []string{"testorg"},
			orgRepositories: map[string][]*github.Repository{"testorg": {&cs}},
		}, nil
	}
	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	p.SetLogger(l)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
