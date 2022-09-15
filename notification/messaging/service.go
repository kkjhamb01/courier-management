package messaging

import (
	"context"
	"errors"
	"fmt"
	"github.com/appleboy/gorush/rpc/proto"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/push"
	"gitlab.artin.ai/backend/courier-management/notification/db"
	"gitlab.artin.ai/backend/courier-management/notification/domain"
	npb "gitlab.artin.ai/backend/courier-management/notification/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type PushClient struct {
	config                config.NotificationData
	db					  *gorm.DB
}

type DeviceInfo struct {
	DeviceToken  string
	DeviceOs int32
}

func (d DeviceInfo) String() string {
	return fmt.Sprintf("OS : %v , Token : %v", d.DeviceOs, d.DeviceToken)
}

func (d DeviceInfo) platform() int32 {
	if d.DeviceOs == int32(npb.DeviceOS_DEVICE_OS_IOS){
		return 1
	}
	if d.DeviceOs == int32(npb.DeviceOS_DEVICE_OS_ANDROID){
		return 2
	}
	return 0
}

func (c *PushClient) OnNewPushEvent(ctx context.Context, event push.Templater) {
	logger.Debugf("OnNewPushEvent start processing new push event %v", event)

	phoneNumbers := event.GetPhoneNumbers()

	if phoneNumbers != nil{
		for _,phoneNumber := range phoneNumbers {
			logger.Debugf("OnNewPushEvent try get registered devices of phoneNumber %v", phoneNumber)

			var devices []DeviceInfo
			err := c.db.Model(&domain.Device{}).Select("DeviceToken", "DeviceOs").Where("phone_number=?", phoneNumber).Scan(&devices).Error

			if err == nil && len(devices) == 0{
				err = errors.New(fmt.Sprintf("no device found for %v", phoneNumber))
			}

			if err != nil{
				logger.Errorf("OnNewPushEvent error in querying devices %v", err)
				return
			}

			for _,deviceInfo := range devices{
				if err = c.sendNotification(deviceInfo, event); err != nil{
					logger.Errorf("OnNewPushEvent error in sending notification to " + deviceInfo.String(), err)
					return
				}
			}
		}
	}
}

func (c *PushClient) sendNotification(device DeviceInfo, event push.Templater) error {
	logger.Debugf("sendNotification start sending notification %s to device %s", event.String(), device.String())

	messageText := event.GetMessage()

	conn, err := grpc.Dial(c.config.Gorush, grpc.WithInsecure())
	if err != nil{
		logger.Errorf("sendNotification cannot connect to gorush", err)
		return err
	}
	defer conn.Close()

	client := proto.NewGorushClient(conn)

	notificationRequest := proto.NotificationRequest{
		Title: event.GetTitle(),
		Platform: device.platform(),
		Tokens:   []string{device.DeviceToken},
		Message:  messageText,
		Badge:    1,
		Category: event.GetCategory(),
		Sound:    event.GetSound(),
	}

	data := event.GetData()

	if data != nil {
		notificationRequest.Data = data
	}

	logger.Debugf("sendNotification send push to gorush %v", notificationRequest)

	r, err := client.Send(context.Background(), &notificationRequest)

	if err != nil {
		logger.Errorf("sendNotification error in sending push: %v", err)
		return err
	}
	if r == nil{
		logger.Debugf("sendNotification error in sending push - cannot connect to gorush")
		return errors.New("cannot connect to gorush")
	}

	if r.Success{
		logger.Debugf("sendNotification successful push, counts = %v", r.Counts)
	} else {
		logger.Debugf("sendNotification unsuccessful push, counts = %v", r.Counts)
	}

	return nil
}

func NewPushClient(config config.NotificationData) *PushClient {
	db,err := db.NewOrm(config.Database)
	if err != nil{
		logger.Fatalf("NewPushClient cannot connect to database", err)
	}

	return &PushClient{
		config: config,
		db: db,
	}
}