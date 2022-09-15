package services

func CloseAllConnections() {
	DeliveryConnection.Close()
}
