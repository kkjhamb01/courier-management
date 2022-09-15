package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/delivery/db"
	"gitlab.artin.ai/backend/courier-management/delivery/model"
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	uaa "gitlab.artin.ai/backend/courier-management/uaa/proto"
)

func GetCourierRequestsHistory(ctx context.Context, userId string, pageSize int, pageNumber int) ([]*deliveryPb.RequestHistoryItem, error) {

	logger.Infof("GetCourierRequestsHistory userId = %v, pageSize = %v, pageNumber = %v", userId, pageSize, pageNumber)

	offset := pageSize * (pageNumber - 1)

	var courierRequests []model.Request
	err := db.MariaDbClient().
		Preload("Locations").
		Where("requests.courier_id = ?", userId).
		Order("updated_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&courierRequests).
		Error
	if err != nil {
		logger.Error("failed to fetch courier requests", err)
		return nil, uaa.Internal.Error(err)
	}

	requestHistoryItems := make([]*deliveryPb.RequestHistoryItem, len(courierRequests))

	for i, courierRequest := range courierRequests {
		var requestStatuses []model.RideStatus
		err = db.MariaDbClient().
			Where("request_id = ?", courierRequest.ID).
			Order("created_at asc").
			Find(&requestStatuses).
			Error
		if err != nil {
			logger.Error("failed to fetch request statuses", err)
			return nil, uaa.Internal.Error(err)
		}

		requestStatusesProto := make([]*deliveryPb.RideStatus, len(requestStatuses))
		for i, requestStatus := range requestStatuses {
			requestStatusesProto[i] = requestStatus.ToProtoP()
		}
		requestHistoryItems[i] = &deliveryPb.RequestHistoryItem{
			Request: courierRequest.ToProtoP(),
			Events:  requestStatusesProto,
		}
	}

	return requestHistoryItems, nil
}

func GetCustomerRequestsHistory(ctx context.Context, userId string, pageSize int, pageNumber int) ([]*deliveryPb.RequestHistoryItem, error) {

	logger.Infof("GetCustomerRequestsHistory userId = %v, pageSize = %v, pageNumber = %v", userId, pageSize, pageNumber)

	offset := pageSize * (pageNumber - 1)

	var customerRequests []model.Request
	err := db.MariaDbClient().
		Preload("Locations").
		Where("requests.customer_id = ?", userId).
		Order("updated_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&customerRequests).
		Error
	if err != nil {
		logger.Error("failed to fetch customer requests", err)
		return nil, uaa.Internal.Error(err)
	}

	requestHistoryItems := make([]*deliveryPb.RequestHistoryItem, len(customerRequests))

	for i, customerRequest := range customerRequests {
		var requestStatuses []model.RideStatus
		err = db.MariaDbClient().
			Where("request_id = ?", customerRequest.ID).
			Order("created_at asc").
			Find(&requestStatuses).
			Error
		if err != nil {
			logger.Error("failed to fetch request statuses", err)
			return nil, uaa.Internal.Error(err)
		}

		requestStatusesProto := make([]*deliveryPb.RideStatus, len(requestStatuses))
		for i, requestStatus := range requestStatuses {
			requestStatusesProto[i] = requestStatus.ToProtoP()
		}
		requestHistoryItems[i] = &deliveryPb.RequestHistoryItem{
			Request: customerRequest.ToProtoP(),
			Events:  requestStatusesProto,
		}
	}

	return requestHistoryItems, nil
}

func GetRequestHistory(ctx context.Context, requestId string) (*deliveryPb.GetRequestHistoryResponse, error) {

	//logger.Infof("GetRequestHistory requestId = %v", requestId)

	var request model.Request
	err := db.MariaDbClient().
		Preload("Locations").
		Where("ID = ?", requestId).
		First(&request).
		Error
	if err != nil {
		logger.Error("failed to fetch request", err)
		return nil, uaa.Internal.Error(err)
	}

	var requestHistoryItems *deliveryPb.RequestHistoryItem

	var requestStatuses []model.RideStatus
	err = db.MariaDbClient().
		Where("request_id = ?", requestId).
		Order("created_at asc").
		Find(&requestStatuses).
		Error
	if err != nil {
		logger.Error("failed to fetch request statuses", err)
		return nil, uaa.Internal.Error(err)
	}

	requestStatusesProto := make([]*deliveryPb.RideStatus, len(requestStatuses))
	for i, requestStatus := range requestStatuses {
		requestStatusesProto[i] = requestStatus.ToProtoP()
	}
	requestHistoryItems = &deliveryPb.RequestHistoryItem{
		Request: request.ToProtoP(),
		Events:  requestStatusesProto,
	}

	return &deliveryPb.GetRequestHistoryResponse{
		Items: requestHistoryItems,
	}, nil
}
