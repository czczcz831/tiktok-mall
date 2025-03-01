package consts

import "time"

// Topic
const (
	// Order
	RocketOrderTransactionTopic = "order_tx"

	RocketOrderDelayedTopic = "order_delay"
	RocketOrderDelayedTime  = time.Minute * 10

	RocketOrderNormalTopic = "order_normal"

	// Payment
	RocketPaymentTransactionTopic = "payment_tx"

	RocketPaymentDelayedTopic = "payment_delay"
	RocketPaymentDelayedTime  = time.Minute * 1

	RocketPaymentNormalTopic = "payment_normal"
)

// Tag
const (
	// Order
	RocketCreateOrderTag        = "create_order"
	RocketCreateOrderDelayedTag = "create_order_delayed"

	// Payment
	RocketCreatePaymentTag = "create_payment"
	RocketPaidSuccessTag   = "paid_success"
)

// ConsumerGroup
const (
	// Cart
	RocketClearCartConsumer = "clear_cart_consumer"
	// Order
	RocketDelayOrderCancelOrderConsumerGroup = "delayed_cancel_order_consumer"
	RocketMarkOrderPaidConsumerGroup         = "mark_order_paid_consumer"
	// Payment
	RocketDelayCancelPaymentConsumerGroup = "payment_consumer"
	RocketDBDecreaseStockConsumerGroup    = "db_decrease_stock_consumer"
)
