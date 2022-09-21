package testing

import (
	assert "github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitlab.com/golibs-starter/golib-message-bus"
	"gitlab.com/golibs-starter/golib-sample-adapter/repository/mysql/model"
	"gitlab.com/golibs-starter/golib-sample-core/event"
	"gitlab.com/golibs-starter/golib-sample-public/testing/handler"
	"gitlab.com/golibs-starter/golib-test"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type CreateOrderControllerTest struct {
	TestSuite
	db             *gorm.DB
	dummyCollector *handler.OrderEventDummyCollector
}

func TestCreateOrderControllerTest(t *testing.T) {
	s := new(CreateOrderControllerTest)
	s.Option(
		golibmsg.KafkaConsumerOpt(),
		golibmsg.KafkaConsumerReadyWaitOpt(),
		golibmsg.ProvideConsumer(handler.NewOrderCreatedEventDummyHandler),
		fx.Provide(handler.NewOrderEventDummyCollector),
	)
	s.Populate(&s.db, &s.dummyCollector)
	s.Invoke(TruncateTablesInvoker("orders"))
	suite.Run(t, s)
}

func (s CreateOrderControllerTest) TestCreateOrder_GiveValidBody_WhenRepoResponseSuccess_ShouldReturnSuccess() {
	// Execute request
	ra := golibtest.NewRestAssured(s.T()).
		When().
		Post("/v1/orders").
		BearerToken(s.CreateJwtToken("10")).
		Body(`{"total_amount": 85000}`).
		Then()

	// Assert inserted data
	var actualOrder model.Order
	assert.NoError(s.T(), s.db.First(&actualOrder).Error)
	assert.NotNil(s.T(), actualOrder)
	assert.EqualValues(s.T(), "10", actualOrder.UserId)
	assert.EqualValues(s.T(), 85000, actualOrder.TotalAmount)

	// Assert request response
	ra.Status(http.StatusOK).
		Body("data.id", actualOrder.Id).
		Body("data.user_id", actualOrder.UserId).
		Body("data.total_amount", actualOrder.TotalAmount).
		BodyCb("data.created_at", func(value interface{}) {
			assert.InDelta(s.T(), actualOrder.CreatedAt.Unix(), value, 1)
		})

	// Collect & assert published event
	golibtest.WaitUntil(func() bool { return len(s.dummyCollector.CreatedEvents()) >= 1 }, 20*time.Second)
	assert.Len(s.T(), s.dummyCollector.CreatedEvents(), 1)
	actualEvent := s.dummyCollector.CreatedEvents()[0]
	assert.Equal(s.T(), "OrderCreatedEvent", actualEvent.Name())
	assert.IsType(s.T(), &event.OrderMessage{}, actualEvent.Payload())
	actualEventPayload := actualEvent.Payload().(*event.OrderMessage)
	assert.EqualValues(s.T(), actualOrder.Id, actualEventPayload.Id)
	assert.EqualValues(s.T(), actualOrder.UserId, actualEventPayload.UserId)
	assert.EqualValues(s.T(), actualOrder.TotalAmount, actualEventPayload.TotalAmount)
	assert.InDelta(s.T(), actualOrder.CreatedAt.Unix(), actualEventPayload.CreatedAt, 1)
}
