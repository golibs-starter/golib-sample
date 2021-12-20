package base

import (
	"gitlab.com/golibs-starter/golib-gin"
	"gitlab.com/golibs-starter/golib-sample-worker/bootstrap"
	"gitlab.com/golibs-starter/golib-test"
)

type TestSuite struct {
	*golibtest.FxTestSuite
}

func NewTestSuite(tsOptions ...golibtest.TsOption) *TestSuite {
	ts := &TestSuite{}
	tsOptions = append(tsOptions,
		golibtest.WithInvokeStart(golibgin.StartTestOpt),
	)
	ts.FxTestSuite = golibtest.NewFxTestSuite(bootstrap.All(), tsOptions...)
	return ts
}
