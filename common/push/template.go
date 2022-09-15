package push

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
)

type Templater interface {
	GetCategory() string
	GetSound() string
	GetPhoneNumbers() []string
	String() string
	GetTitle() string
	GetData() *structpb.Struct
	GetMessage() string
}