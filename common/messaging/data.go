package messaging

import (
	"github.com/golang/protobuf/proto"
	"github.com/kkjhamb01/courier-management/common/logger"
	"google.golang.org/protobuf/runtime/protoiface"
)

func decodeProto(data []byte, msg protoiface.MessageV1) error {
	err := proto.Unmarshal(data, msg)
	if err != nil {
		logger.Error("failed to unmarshal message", err)
		return err
	}

	return nil
}

func encodeProto(msg protoiface.MessageV1) ([]byte, error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		logger.Error("failed to unmarshal message", err)
		return nil, err
	}

	return data, nil
}
