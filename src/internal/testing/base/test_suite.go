package base

import (
	"gitlab.id.vin/vincart/golib-gin"
	"gitlab.id.vin/vincart/golib-migrate"
	"gitlab.id.vin/vincart/golib-sample-internal/bootstrap"
	"gitlab.id.vin/vincart/golib-test"
)

type TestSuite struct {
	*golibtest.FxTestSuite
}

func NewTestSuite(tsOptions ...golibtest.TsOption) *TestSuite {
	ts := &TestSuite{}
	tsOptions = append(tsOptions,
		golibtest.WithFxOption(golibmigrate.MigrationOpt()),
		golibtest.WithInvokeStart(golibgin.StartTestOpt),
	)
	ts.FxTestSuite = golibtest.NewFxTestSuite(bootstrap.All(), tsOptions...)
	return ts
}