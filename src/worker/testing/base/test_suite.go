package base

import (
    "gitlab.com/golibs-starter/golib-sample-worker/bootstrap"
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
