package consts

// Topic
const (
	// Order
	RocketOrderTransactionTopic = "order_tx"
	RocketOrderNormalTopic      = "order_normal"

	// Payment
	RocketPaymentTransactionTopic = "payment_tx"
	RocketPaymentNormalTopic      = "payment_normal"
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
	RocketDelayCancelOrderConsumerGroup = "delayed_cancel_order_consumer"
	// Payment
	RocketDelayCancelPaymentConsumerGroup = "payment_consumer"
)
