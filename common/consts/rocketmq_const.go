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
	RocketCartConsumerGroup = "cart_consumer"
	// Order
	RocketOrderConsumerGroup = "order_consumer"
	// Payment
	RocketPaymentConsumerGroup = "payment_consumer"
)
