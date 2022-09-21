package testing

import (
	"context"
	"github.com/jarcoal/httpmock"
	assert "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.com/golibs-starter/golib-message-bus"
	"gitlab.com/golibs-starter/golib-sample-core/event"
	"gitlab.com/golibs-starter/golib-sample-worker/testing/dummy"
	"gitlab.com/golibs-starter/golib-test"
	"gitlab.com/golibs-starter/golib/pubsub"
	"gitlab.com/golibs-starter/golib/web/constant"
	webContext "gitlab.com/golibs-starter/golib/web/context"
	webEvent "gitlab.com/golibs-starter/golib/web/event"
	"go.uber.org/fx"
	"net/http"
	"testing"
	"time"
)

type SendOrderToDeliveryProviderHandlerTest struct {
	TestSuite
	httpClient *http.Client
	collector  *dummy.OrderEventDummyCollector
}

func TestSendOrderToDeliveryProviderHandlerTest(t *testing.T) {
	s := SendOrderToDeliveryProviderHandlerTest{}
	s.Option(
		golibmsg.KafkaAdminOpt(),
		golibmsg.KafkaProducerOpt(),
		golibmsg.ProvideConsumer(dummy.NewOrderCreatedEventDummyHandler),
		fx.Provide(dummy.NewOrderEventDummyCollector),
		fx.Populate(&s.httpClient, &s.collector),
	)
	suite.Run(t, &s)
}

func (s SendOrderToDeliveryProviderHandlerTest) TestWhenOrderCreated_ShouldSendToDeliveryService() {
	httpmock.ActivateNonDefault(s.httpClient)
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://order.sample.api/v1/orders",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(http.StatusCreated, `{
 "meta": {
   "code": 201,
   "message": "Successful"
 },
 "data": {
   "id": "20",
   "created_at": 1637415974
 }
}`), nil
		},
	)

	requestAttrs := webContext.RequestAttributes{
		DeviceId:        "TEST-DEVICE-ID",
		DeviceSessionId: "TEST-DEVICE-SESSION-ID",
		CorrelationId:   "TEST-REQUEST-ID",
		SecurityAttributes: webContext.SecurityAttributes{
			UserId:            "test-user-id",
			TechnicalUsername: "test-technical-username",
		},
	}
	ctx := context.WithValue(context.Background(), constant.ContextReqAttribute, &requestAttrs)
	e := &event.OrderCreatedEvent{
		AbstractEvent: webEvent.NewAbstractEvent(ctx, "OrderCreatedEvent"),
		PayloadData: &event.OrderMessage{
			Id:          2,
			UserId:      "10",
			TotalAmount: 85000,
			CreatedAt:   time.Now().Unix(),
		},
	}

	pubsub.Publish(e)

	golibtest.WaitUntil(func() bool { return len(s.collector.CreatedEvents()) >= 1 }, 20*time.Second)
	time.Sleep(1 * time.Second)
	assert.Len(s.T(), s.collector.CreatedEvents(), 1)
	expectedEvent := s.collector.CreatedEvents()[0]
	assert.Equal(s.T(), "OrderCreatedEvent", expectedEvent.Name())
	assert.IsType(s.T(), &event.OrderMessage{}, expectedEvent.Payload())

	assert.Equal(s.T(), 1, httpmock.GetCallCountInfo()["POST https://order.sample.api/v1/orders"])
}
