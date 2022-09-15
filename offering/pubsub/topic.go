package pubsub

const (
	TopicNewOffer    = "new_offer"
	TopicCancelOffer = "cancel_offer"
	TopicAcceptOffer = "accept_offer"

	TopicRetryOffer          = "retry_offer"
	TopicMaxOfferRetries     = "max_offer_retries"
	TopicOfferSentToCouriers = "offer_sent_to_couriers"
	TopicOfferCreationFailed = "offer_creation_failed"
	TopicAcceptOfferFailed   = "accept_offer_failed"
	TopicRejectOfferFailed   = "reject_offer_failed"
)
