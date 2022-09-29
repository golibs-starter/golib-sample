package testing

import (
	"context"
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"gitlab.com/golibs-starter/golib-message-bus/testutil"
	"gitlab.com/golibs-starter/golib-sample-core/event"
	"gitlab.com/golibs-starter/golib-test"
	"gitlab.com/golibs-starter/golib/pubsub"
	"gitlab.com/golibs-starter/golib/web/constant"
	webContext "gitlab.com/golibs-starter/golib/web/context"
	webEvent "gitlab.com/golibs-starter/golib/web/event"
	"net/http"
	"testing"
	"time"
)

type SendOrderToDeliveryProviderHandlerTest struct {
	TestSuite
	messageCollector *golibmsgTestUtil.MessageCollector
}

func TestSendOrderToDeliveryProviderHandlerTest(t *testing.T) {
	s := SendOrderToDeliveryProviderHandlerTest{}
	s.Populate(&s.messageCollector)
	s.Decorate(func(httpClient *http.Client) *http.Client {
		httpmock.ActivateNonDefault(httpClient)
		return httpClient
	})
	suite.Run(t, &s)
}

func (s *SendOrderToDeliveryProviderHandlerTest) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

func (s *SendOrderToDeliveryProviderHandlerTest) TestWhenOrderCreated_ShouldSendToDeliveryService() {
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

	topic := "c1.order.order-created.test"
	golibtest.WaitUntilT(s.T(), func() bool { return s.messageCollector.Count(topic) >= 1 }, 20*time.Second)
	time.Sleep(1 * time.Second)
	s.Len(s.messageCollector.GetMessages(topic), 1)
	var actualEvent event.OrderCreatedEvent
	s.NoError(json.Unmarshal([]byte(s.messageCollector.GetMessages(topic)[0]), &actualEvent))
	s.Equal(e.Name(), actualEvent.Name())
	s.IsType(&event.OrderMessage{}, actualEvent.Payload())

	s.Equal(1, httpmock.GetCallCountInfo()["POST https://order.sample.api/v1/orders"])
}
