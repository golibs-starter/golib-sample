package testing

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"gitlab.com/golibs-starter/golib-message-bus"
	golibmsgTestUtil "gitlab.com/golibs-starter/golib-message-bus/testutil"
	"gitlab.com/golibs-starter/golib-sample-adapter/repository/mysql/model"
	"gitlab.com/golibs-starter/golib-sample-core/event"
	"gitlab.com/golibs-starter/golib-test"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type CreateOrderControllerTest struct {
	TestSuite
	db               *gorm.DB
	messageCollector *golibmsgTestUtil.MessageCollector
}

func TestCreateOrderControllerTest(t *testing.T) {
	s := new(CreateOrderControllerTest)
	s.Option(
		golibmsgTestUtil.MessageCollectorOpt(),
		golibmsg.KafkaConsumerOpt(),
		golibmsg.KafkaConsumerReadyWaitOpt(),
	)
	s.Populate(&s.db, &s.messageCollector)
	suite.Run(t, s)
}

func (s *CreateOrderControllerTest) TestCreateOrder_GiveValidBody_WhenRepoResponseSuccess_ShouldReturnSuccess() {
	// Execute request
	ra := golibtest.NewRestAssuredSuite(s).
		When().
		Post("/v1/orders").
		BearerToken(s.jwtTestUtil.CreateJwtToken("10")).
		Body(`{"total_amount": 85000}`).
		Then()

	// Assert inserted data
	var actualOrder model.Order
	s.NoError(s.db.First(&actualOrder).Error)
	s.NotNil(actualOrder)
	s.EqualValues("10", actualOrder.UserId)
	s.EqualValues(85000, actualOrder.TotalAmount)

	// Assert request response
	ra.Status(http.StatusOK).
		Body("data.id", actualOrder.Id).
		Body("data.user_id", actualOrder.UserId).
		Body("data.total_amount", actualOrder.TotalAmount).
		BodyCb("data.created_at", func(value interface{}) {
			s.InDelta(actualOrder.CreatedAt.Unix(), value, 1)
		})

	topic := "c1.order.order-created.test"
	golibtest.WaitUntil(func() bool { return s.messageCollector.Count(topic) >= 1 }, 20*time.Second)
	time.Sleep(1 * time.Second)
	s.Len(s.messageCollector.GetMessages(topic), 1)
	var actualEvent event.OrderCreatedEvent
	s.NoError(json.Unmarshal([]byte(s.messageCollector.GetMessages(topic)[0]), &actualEvent))
	s.Equal("OrderCreatedEvent", actualEvent.Name())
	s.IsType(&event.OrderMessage{}, actualEvent.Payload())
	actualEventPayload := actualEvent.Payload().(*event.OrderMessage)
	s.EqualValues(actualOrder.Id, actualEventPayload.Id)
	s.EqualValues(actualOrder.UserId, actualEventPayload.UserId)
	s.EqualValues(actualOrder.TotalAmount, actualEventPayload.TotalAmount)
	s.InDelta(actualOrder.CreatedAt.Unix(), actualEventPayload.CreatedAt, 1)
}
