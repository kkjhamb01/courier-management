package messaging

const (
	PubSubOffering = "Offering"
)

const (
	TopicRetryOfferRequest      = "TopicRetryOfferRequest"
	TopicOfferMaxRetries        = "TopicOfferMaxRetries"
	TopicNewOfferSentToCouriers = "TopicNewOfferSentToCouriers"
	TopicOfferCreationFailed    = "TopicOfferCreationFailed"
	TopicAcceptOfferFailed      = "TopicAcceptOfferFailed"
	TopicRejectOfferFailed      = "TopicRejectOfferFailed"

	TopicPushNotification = "TopicPushNotification"

	TopicCourierStatusUpdate = "TopicCourierStatusChanged"

	TopicDeliveryNewRequest                  = "TopicDeliveryNewRequest"
	TopicDeliveryRequestAccepted             = "TopicDeliveryRequestAccepted"
	TopicDeliveryRequestArrivedOrigin        = "TopicDeliveryRequestArrivedOrigin"
	TopicDeliveryRequestArrivedDestination   = "TopicDeliveryRequestArrivedDestination"
	TopicDeliveryRequestPickedUp             = "TopicDeliveryRequestPickedUp"
	TopicDeliveryRequestDelivered            = "TopicDeliveryRequestDelivered"
	TopicDeliveryRequestCompleted            = "TopicDeliveryRequestCompleted"
	TopicDeliveryRequestNavigatingToReceiver = "TopicDeliveryRequestNavigatingToReceiver"
	TopicDeliveryRequestNavigatingToSender   = "TopicDeliveryRequestNavigatingToSender"
	TopicDeliveryRequestSenderNotAvailable   = "TopicDeliveryRequestSenderNotAvailable"
	TopicDeliveryRequestReceiverNotAvailable = "TopicDeliveryRequestReceiverNotAvailable"
	TopicRequestCancelled                    = "TopicRequestCancelled"
	TopicRequestRejected                     = "TopicRequestRejected"
)
