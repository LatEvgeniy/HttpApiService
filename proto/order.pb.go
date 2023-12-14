// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: proto/order.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderDirection int32

const (
	OrderDirection_BUY  OrderDirection = 0
	OrderDirection_SELL OrderDirection = 1
)

// Enum value maps for OrderDirection.
var (
	OrderDirection_name = map[int32]string{
		0: "BUY",
		1: "SELL",
	}
	OrderDirection_value = map[string]int32{
		"BUY":  0,
		"SELL": 1,
	}
)

func (x OrderDirection) Enum() *OrderDirection {
	p := new(OrderDirection)
	*p = x
	return p
}

func (x OrderDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_order_proto_enumTypes[0].Descriptor()
}

func (OrderDirection) Type() protoreflect.EnumType {
	return &file_proto_order_proto_enumTypes[0]
}

func (x OrderDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderDirection.Descriptor instead.
func (OrderDirection) EnumDescriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

type OrderPair int32

const (
	OrderPair_USD_EUR OrderPair = 0
	OrderPair_USD_UAH OrderPair = 1
	OrderPair_UAH_EUR OrderPair = 2
)

// Enum value maps for OrderPair.
var (
	OrderPair_name = map[int32]string{
		0: "USD_EUR",
		1: "USD_UAH",
		2: "UAH_EUR",
	}
	OrderPair_value = map[string]int32{
		"USD_EUR": 0,
		"USD_UAH": 1,
		"UAH_EUR": 2,
	}
)

func (x OrderPair) Enum() *OrderPair {
	p := new(OrderPair)
	*p = x
	return p
}

func (x OrderPair) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderPair) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_order_proto_enumTypes[1].Descriptor()
}

func (OrderPair) Type() protoreflect.EnumType {
	return &file_proto_order_proto_enumTypes[1]
}

func (x OrderPair) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderPair.Descriptor instead.
func (OrderPair) EnumDescriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

type OrderType int32

const (
	OrderType_MARKET OrderType = 0
	OrderType_LIMIT  OrderType = 1
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0: "MARKET",
		1: "LIMIT",
	}
	OrderType_value = map[string]int32{
		"MARKET": 0,
		"LIMIT":  1,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_order_proto_enumTypes[2].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_proto_order_proto_enumTypes[2]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{2}
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string         `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Pair      OrderPair      `protobuf:"varint,2,opt,name=pair,proto3,enum=proto.OrderPair" json:"pair,omitempty"`
	Type      OrderType      `protobuf:"varint,3,opt,name=type,proto3,enum=proto.OrderType" json:"type,omitempty"`
	Direction OrderDirection `protobuf:"varint,4,opt,name=direction,proto3,enum=proto.OrderDirection" json:"direction,omitempty"`
	Price     float64        `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Volume    float64        `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateOrderRequest) GetPair() OrderPair {
	if x != nil {
		return x.Pair
	}
	return OrderPair_USD_EUR
}

func (x *CreateOrderRequest) GetType() OrderType {
	if x != nil {
		return x.Type
	}
	return OrderType_MARKET
}

func (x *CreateOrderRequest) GetDirection() OrderDirection {
	if x != nil {
		return x.Direction
	}
	return OrderDirection_BUY
}

func (x *CreateOrderRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrderRequest) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedOrder *Order    `protobuf:"bytes,1,opt,name=createdOrder,proto3" json:"createdOrder,omitempty"`
	Error        *ErrorDto `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetCreatedOrder() *Order {
	if x != nil {
		return x.CreatedOrder
	}
	return nil
}

func (x *CreateOrderResponse) GetError() *ErrorDto {
	if x != nil {
		return x.Error
	}
	return nil
}

type GetUserOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetUserOrdersRequest) Reset() {
	*x = GetUserOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserOrdersRequest) ProtoMessage() {}

func (x *GetUserOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetUserOrdersRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserOrdersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order  `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	Error  *ErrorDto `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetUserOrdersResponse) Reset() {
	*x = GetUserOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserOrdersResponse) ProtoMessage() {}

func (x *GetUserOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetUserOrdersResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *GetUserOrdersResponse) GetError() *ErrorDto {
	if x != nil {
		return x.Error
	}
	return nil
}

type RemoveOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	OrderId string `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
}

func (x *RemoveOrderRequest) Reset() {
	*x = RemoveOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveOrderRequest) ProtoMessage() {}

func (x *RemoveOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveOrderRequest.ProtoReflect.Descriptor instead.
func (*RemoveOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemoveOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type RemoveOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemovedOrder *Order    `protobuf:"bytes,1,opt,name=removedOrder,proto3" json:"removedOrder,omitempty"`
	Error        *ErrorDto `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RemoveOrderResponse) Reset() {
	*x = RemoveOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveOrderResponse) ProtoMessage() {}

func (x *RemoveOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveOrderResponse.ProtoReflect.Descriptor instead.
func (*RemoveOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveOrderResponse) GetRemovedOrder() *Order {
	if x != nil {
		return x.RemovedOrder
	}
	return nil
}

func (x *RemoveOrderResponse) GetError() *ErrorDto {
	if x != nil {
		return x.Error
	}
	return nil
}

type MatchOrdersEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedMatchedOrder *Order    `protobuf:"bytes,1,opt,name=createdMatchedOrder,proto3" json:"createdMatchedOrder,omitempty"`
	LimitMatchedOrder   *Order    `protobuf:"bytes,2,opt,name=limitMatchedOrder,proto3" json:"limitMatchedOrder,omitempty"`
	MatchedVolume       float64   `protobuf:"fixed64,3,opt,name=matchedVolume,proto3" json:"matchedVolume,omitempty"`
	Error               *ErrorDto `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MatchOrdersEvent) Reset() {
	*x = MatchOrdersEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchOrdersEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchOrdersEvent) ProtoMessage() {}

func (x *MatchOrdersEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchOrdersEvent.ProtoReflect.Descriptor instead.
func (*MatchOrdersEvent) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{6}
}

func (x *MatchOrdersEvent) GetCreatedMatchedOrder() *Order {
	if x != nil {
		return x.CreatedMatchedOrder
	}
	return nil
}

func (x *MatchOrdersEvent) GetLimitMatchedOrder() *Order {
	if x != nil {
		return x.LimitMatchedOrder
	}
	return nil
}

func (x *MatchOrdersEvent) GetMatchedVolume() float64 {
	if x != nil {
		return x.MatchedVolume
	}
	return 0
}

func (x *MatchOrdersEvent) GetError() *ErrorDto {
	if x != nil {
		return x.Error
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string         `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	OrderId        string         `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Pair           OrderPair      `protobuf:"varint,3,opt,name=pair,proto3,enum=proto.OrderPair" json:"pair,omitempty"`
	Direction      OrderDirection `protobuf:"varint,4,opt,name=direction,proto3,enum=proto.OrderDirection" json:"direction,omitempty"`
	Type           OrderType      `protobuf:"varint,5,opt,name=type,proto3,enum=proto.OrderType" json:"type,omitempty"`
	InitVolume     float64        `protobuf:"fixed64,6,opt,name=initVolume,proto3" json:"initVolume,omitempty"`
	InitPrice      float64        `protobuf:"fixed64,7,opt,name=initPrice,proto3" json:"initPrice,omitempty"`
	FilledPrice    float64        `protobuf:"fixed64,8,opt,name=filledPrice,proto3" json:"filledPrice,omitempty"`
	FilledVolume   float64        `protobuf:"fixed64,9,opt,name=filledVolume,proto3" json:"filledVolume,omitempty"`
	LockedVolume   float64        `protobuf:"fixed64,10,opt,name=lockedVolume,proto3" json:"lockedVolume,omitempty"`
	ExpirationDate int64          `protobuf:"varint,11,opt,name=expirationDate,proto3" json:"expirationDate,omitempty"`
	CreationDate   int64          `protobuf:"varint,12,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	UpdatedDate    int64          `protobuf:"varint,13,opt,name=updatedDate,proto3" json:"updatedDate,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{7}
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetPair() OrderPair {
	if x != nil {
		return x.Pair
	}
	return OrderPair_USD_EUR
}

func (x *Order) GetDirection() OrderDirection {
	if x != nil {
		return x.Direction
	}
	return OrderDirection_BUY
}

func (x *Order) GetType() OrderType {
	if x != nil {
		return x.Type
	}
	return OrderType_MARKET
}

func (x *Order) GetInitVolume() float64 {
	if x != nil {
		return x.InitVolume
	}
	return 0
}

func (x *Order) GetInitPrice() float64 {
	if x != nil {
		return x.InitPrice
	}
	return 0
}

func (x *Order) GetFilledPrice() float64 {
	if x != nil {
		return x.FilledPrice
	}
	return 0
}

func (x *Order) GetFilledVolume() float64 {
	if x != nil {
		return x.FilledVolume
	}
	return 0
}

func (x *Order) GetLockedVolume() float64 {
	if x != nil {
		return x.LockedVolume
	}
	return 0
}

func (x *Order) GetExpirationDate() int64 {
	if x != nil {
		return x.ExpirationDate
	}
	return 0
}

func (x *Order) GetCreationDate() int64 {
	if x != nil {
		return x.CreationDate
	}
	return 0
}

func (x *Order) GetUpdatedDate() int64 {
	if x != nil {
		return x.UpdatedDate
	}
	return 0
}

var File_proto_order_proto protoreflect.FileDescriptor

var file_proto_order_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x04,
	0x70, 0x61, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x70, 0x61,
	0x69, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x6e, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x44, 0x74, 0x6f, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2e, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x46, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6e, 0x0a, 0x13, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44,
	0x74, 0x6f, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xdb, 0x01, 0x0a, 0x10, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3e,
	0x0a, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x13, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x3a,
	0x0a, 0x11, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x11, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x74, 0x6f,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xd0, 0x03, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x12, 0x33, 0x0a, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x6c,
	0x65, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x2a, 0x23, 0x0a, 0x0e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03,
	0x42, 0x55, 0x59, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x2a,
	0x32, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x12, 0x0b, 0x0a, 0x07,
	0x55, 0x53, 0x44, 0x5f, 0x45, 0x55, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x53, 0x44,
	0x5f, 0x55, 0x41, 0x48, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x41, 0x48, 0x5f, 0x45, 0x55,
	0x52, 0x10, 0x02, 0x2a, 0x22, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_order_proto_rawDescOnce sync.Once
	file_proto_order_proto_rawDescData = file_proto_order_proto_rawDesc
)

func file_proto_order_proto_rawDescGZIP() []byte {
	file_proto_order_proto_rawDescOnce.Do(func() {
		file_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_order_proto_rawDescData)
	})
	return file_proto_order_proto_rawDescData
}

var file_proto_order_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_order_proto_goTypes = []interface{}{
	(OrderDirection)(0),           // 0: proto.OrderDirection
	(OrderPair)(0),                // 1: proto.OrderPair
	(OrderType)(0),                // 2: proto.OrderType
	(*CreateOrderRequest)(nil),    // 3: proto.CreateOrderRequest
	(*CreateOrderResponse)(nil),   // 4: proto.CreateOrderResponse
	(*GetUserOrdersRequest)(nil),  // 5: proto.GetUserOrdersRequest
	(*GetUserOrdersResponse)(nil), // 6: proto.GetUserOrdersResponse
	(*RemoveOrderRequest)(nil),    // 7: proto.RemoveOrderRequest
	(*RemoveOrderResponse)(nil),   // 8: proto.RemoveOrderResponse
	(*MatchOrdersEvent)(nil),      // 9: proto.MatchOrdersEvent
	(*Order)(nil),                 // 10: proto.Order
	(*ErrorDto)(nil),              // 11: proto.ErrorDto
}
var file_proto_order_proto_depIdxs = []int32{
	1,  // 0: proto.CreateOrderRequest.pair:type_name -> proto.OrderPair
	2,  // 1: proto.CreateOrderRequest.type:type_name -> proto.OrderType
	0,  // 2: proto.CreateOrderRequest.direction:type_name -> proto.OrderDirection
	10, // 3: proto.CreateOrderResponse.createdOrder:type_name -> proto.Order
	11, // 4: proto.CreateOrderResponse.error:type_name -> proto.ErrorDto
	10, // 5: proto.GetUserOrdersResponse.orders:type_name -> proto.Order
	11, // 6: proto.GetUserOrdersResponse.error:type_name -> proto.ErrorDto
	10, // 7: proto.RemoveOrderResponse.removedOrder:type_name -> proto.Order
	11, // 8: proto.RemoveOrderResponse.error:type_name -> proto.ErrorDto
	10, // 9: proto.MatchOrdersEvent.createdMatchedOrder:type_name -> proto.Order
	10, // 10: proto.MatchOrdersEvent.limitMatchedOrder:type_name -> proto.Order
	11, // 11: proto.MatchOrdersEvent.error:type_name -> proto.ErrorDto
	1,  // 12: proto.Order.pair:type_name -> proto.OrderPair
	0,  // 13: proto.Order.direction:type_name -> proto.OrderDirection
	2,  // 14: proto.Order.type:type_name -> proto.OrderType
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_proto_order_proto_init() }
func file_proto_order_proto_init() {
	if File_proto_order_proto != nil {
		return
	}
	file_proto_error_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserOrdersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserOrdersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveOrderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveOrderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchOrdersEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_order_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_order_proto_goTypes,
		DependencyIndexes: file_proto_order_proto_depIdxs,
		EnumInfos:         file_proto_order_proto_enumTypes,
		MessageInfos:      file_proto_order_proto_msgTypes,
	}.Build()
	File_proto_order_proto = out.File
	file_proto_order_proto_rawDesc = nil
	file_proto_order_proto_goTypes = nil
	file_proto_order_proto_depIdxs = nil
}
