package testing

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/golibs-starter/golib"
	"gitlab.com/golibs-starter/golib-migrate"
	"gitlab.com/golibs-starter/golib-sample-public/bootstrap"
	"gitlab.com/golibs-starter/golib-sample-public/testing/properties"
	"gitlab.com/golibs-starter/golib-test"
	"gitlab.com/golibs-starter/golib/log"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"time"
)

type TestSuite struct {
	golibtest.FxTestSuite
	jwtSignKey *rsa.PrivateKey
	db         *gorm.DB
}

func (s *TestSuite) SetupSuite() {
	log.Info("Test App is initializing")
	s.Profile("testing")
	s.ProfilePath("../config/", "./config/")
	s.Options(bootstrap.All())
	s.Option(
		golibmigrate.MigrationOpt(),
		golib.ProvideProps(properties.NewJwtTestProperties),
		fx.Invoke(s.LoadJwtPrivateKey),
	)
	s.Populate(&s.db)
	s.SetupApp()
	log.Info("Test App is initialized")
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
