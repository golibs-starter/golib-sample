package base

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"gitlab.id.vin/vincart/golib"
	"gitlab.id.vin/vincart/golib-gin"
	"gitlab.id.vin/vincart/golib-migrate"
	"gitlab.id.vin/vincart/golib-sample-public/bootstrap"
	"gitlab.id.vin/vincart/golib-sample-public/testing/properties"
	"gitlab.id.vin/vincart/golib-test"
	"go.uber.org/fx"
	"time"
)

type TestSuite struct {
	*golibtest.FxTestSuite
	jwtSignKey *rsa.PrivateKey
}

func NewTestSuite(tsOptions ...golibtest.TsOption) *TestSuite {
	ts := &TestSuite{}
	tsOptions = append(tsOptions,
		golibtest.WithFxOption(golib.ProvideProps(properties.NewJwtTestProperties)),
		golibtest.WithFxOption(fx.Invoke(ts.LoadJwtPrivateKey)),
		golibtest.WithFxOption(golibmigrate.MigrationOpt()),
		golibtest.WithInvokeStart(golibgin.StartTestOpt),
	)
	ts.FxTestSuite = golibtest.NewFxTestSuite(bootstrap.All(), tsOptions...)
	return ts
}

func (s TestSuite) CreateJwtToken(userId string) string {
	// create a signer for rsa 256
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	now := time.Now()
	token.Claims = &jwt.StandardClaims{
		Issuer:    "TESTER",
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(time.Minute * 1).Unix(),
		Subject:   userId,
	}

	// Creat token string
	jwtToken, err := token.SignedString(s.jwtSignKey)
	s.Require().NoError(err)
	return jwtToken
}

func (s *TestSuite) LoadJwtPrivateKey(props *properties.JwtTestProperties) {
	if len(props.PrivateKey) == 0 {
		s.Require().FailNow("Missing private key for test")
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(props.PrivateKey))
	s.Require().NoError(err)
	s.jwtSignKey = signKey
}
