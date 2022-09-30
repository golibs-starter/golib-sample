package testing

import (
	"context"
	"encoding/json"
	"github.com/jarcoal/httpmock"
	assert "github.com/stretchr/testify/require"
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

func TestSendOrderToDeliveryProvider_WhenOrderCreated_ShouldSendToDeliveryService(t *testing.T) {
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
	golibtest.WaitUntilT(t, func() bool { return messageCollector.Count(topic) >= 1 }, 20*time.Second)
	time.Sleep(1 * time.Second)
	assert.Len(t, messageCollector.GetMessages(topic), 1)
	var actualEvent event.OrderCreatedEvent
	assert.NoError(t, json.Unmarshal([]byte(messageCollector.GetMessages(topic)[0]), &actualEvent))
	assert.Equal(t, e.Name(), actualEvent.Name())
	assert.IsType(t, &event.OrderMessage{}, actualEvent.Payload())
	assert.Equal(t, 1, httpmock.GetCallCountInfo()["POST https://order.sample.api/v1/orders"])
}
