package consts

import "time"

// Topic
const (
	// Order
	RocketOrderTransactionTopic = "order_tx"

	RocketOrderDelayedTopic = "order_delay"
	RocketOrderDelayedTime  = time.Minute * 1

	RocketOrderNormalTopic = "order_normal"

	// Payment
	RocketPaymentTransactionTopic = "payment_tx"

	RocketPaymentDelayedTopic = "payment_delay"
	RocketPaymentNormalTopic  = "payment_normal"
)

// Tag
const (
	// Order
	RocketCreateOrderTag        = "create_order"
	RocketCreateOrderDelayedTag = "create_order_delayed"

	// Payment
	RocketCreatePaymentTag = "create_payment"
)

// ConsumerGroup
const (
	// Cart
	RocketClearCartConsumer = "clear_cart_consumer"
	// Order
	RocketDelayOrderCancelOrderConsumerGroup = "delayed_cancel_order_consumer"
	// Payment
	RocketDelayCancelPaymentConsumerGroup = "payment_consumer"
)
