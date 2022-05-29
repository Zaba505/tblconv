// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: plugin.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTs     uint64        `protobuf:"varint,1,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	Query       string        `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Args        []*NamedValue `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	ReturnsRows bool          `protobuf:"varint,4,opt,name=returns_rows,json=returnsRows,proto3" json:"returns_rows,omitempty"`
	Txn         *TxnContext   `protobuf:"bytes,5,opt,name=txn,proto3" json:"txn,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetStartTs() uint64 {
	if x != nil {
		return x.StartTs
	}
	return 0
}

func (x *Request) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Request) GetArgs() []*NamedValue {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Request) GetReturnsRows() bool {
	if x != nil {
		return x.ReturnsRows
	}
	return false
}

func (x *Request) GetTxn() *TxnContext {
	if x != nil {
		return x.Txn
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastInsertId int64       `protobuf:"varint,1,opt,name=last_insert_id,json=lastInsertId,proto3" json:"last_insert_id,omitempty"`
	RowsAffected int64       `protobuf:"varint,2,opt,name=rows_affected,json=rowsAffected,proto3" json:"rows_affected,omitempty"`
	Columns      []string    `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	Rows         []*Row      `protobuf:"bytes,4,rep,name=rows,proto3" json:"rows,omitempty"`
	Txn          *TxnContext `protobuf:"bytes,5,opt,name=txn,proto3" json:"txn,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetLastInsertId() int64 {
	if x != nil {
		return x.LastInsertId
	}
	return 0
}

func (x *Response) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

func (x *Response) GetColumns() []string {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *Response) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *Response) GetTxn() *TxnContext {
	if x != nil {
		return x.Txn
	}
	return nil
}

type TxnContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTs   int64 `protobuf:"varint,1,opt,name=start_ts,json=startTs,proto3" json:"start_ts,omitempty"`
	CommitTs  int64 `protobuf:"varint,2,opt,name=commit_ts,json=commitTs,proto3" json:"commit_ts,omitempty"`
	Committed bool  `protobuf:"varint,3,opt,name=committed,proto3" json:"committed,omitempty"`
	Aborted   bool  `protobuf:"varint,4,opt,name=aborted,proto3" json:"aborted,omitempty"`
	Isolation int64 `protobuf:"varint,5,opt,name=isolation,proto3" json:"isolation,omitempty"`
	ReadOnly  bool  `protobuf:"varint,6,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
}

func (x *TxnContext) Reset() {
	*x = TxnContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxnContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxnContext) ProtoMessage() {}

func (x *TxnContext) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxnContext.ProtoReflect.Descriptor instead.
func (*TxnContext) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *TxnContext) GetStartTs() int64 {
	if x != nil {
		return x.StartTs
	}
	return 0
}

func (x *TxnContext) GetCommitTs() int64 {
	if x != nil {
		return x.CommitTs
	}
	return 0
}

func (x *TxnContext) GetCommitted() bool {
	if x != nil {
		return x.Committed
	}
	return false
}

func (x *TxnContext) GetAborted() bool {
	if x != nil {
		return x.Aborted
	}
	return false
}

func (x *TxnContext) GetIsolation() int64 {
	if x != nil {
		return x.Isolation
	}
	return 0
}

func (x *TxnContext) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

type NamedValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If the Name is not empty it should be used for the parameter identifier and
	// not the ordinal position.
	//
	// Name will not have a symbol prefix.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Ordinal position of the parameter starting from one and is always set.
	Ordinal int64 `protobuf:"varint,2,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	// Value is the parameter value.
	Value *Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NamedValue) Reset() {
	*x = NamedValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamedValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedValue) ProtoMessage() {}

func (x *NamedValue) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedValue.ProtoReflect.Descriptor instead.
func (*NamedValue) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *NamedValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamedValue) GetOrdinal() int64 {
	if x != nil {
		return x.Ordinal
	}
	return 0
}

func (x *NamedValue) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*Value_Null
	//	*Value_Int64
	//	*Value_Float64
	//	*Value_Bool
	//	*Value_Bytes
	//	*Value_String_
	//	*Value_Time
	Value isValue_Value `protobuf_oneof:"value"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{4}
}

func (m *Value) GetValue() isValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Value) GetNull() bool {
	if x, ok := x.GetValue().(*Value_Null); ok {
		return x.Null
	}
	return false
}

func (x *Value) GetInt64() int64 {
	if x, ok := x.GetValue().(*Value_Int64); ok {
		return x.Int64
	}
	return 0
}

func (x *Value) GetFloat64() float64 {
	if x, ok := x.GetValue().(*Value_Float64); ok {
		return x.Float64
	}
	return 0
}

func (x *Value) GetBool() bool {
	if x, ok := x.GetValue().(*Value_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *Value) GetBytes() []byte {
	if x, ok := x.GetValue().(*Value_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (x *Value) GetString_() string {
	if x, ok := x.GetValue().(*Value_String_); ok {
		return x.String_
	}
	return ""
}

func (x *Value) GetTime() *timestamppb.Timestamp {
	if x, ok := x.GetValue().(*Value_Time); ok {
		return x.Time
	}
	return nil
}

type isValue_Value interface {
	isValue_Value()
}

type Value_Null struct {
	Null bool `protobuf:"varint,1,opt,name=null,proto3,oneof"`
}

type Value_Int64 struct {
	Int64 int64 `protobuf:"varint,2,opt,name=int64,proto3,oneof"`
}

type Value_Float64 struct {
	Float64 float64 `protobuf:"fixed64,3,opt,name=float64,proto3,oneof"`
}

type Value_Bool struct {
	Bool bool `protobuf:"varint,4,opt,name=bool,proto3,oneof"`
}

type Value_Bytes struct {
	Bytes []byte `protobuf:"bytes,5,opt,name=bytes,proto3,oneof"`
}

type Value_String_ struct {
	String_ string `protobuf:"bytes,6,opt,name=string,proto3,oneof"`
}

type Value_Time struct {
	Time *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=time,proto3,oneof"`
}

func (*Value_Null) isValue_Value() {}

func (*Value_Int64) isValue_Value() {}

func (*Value_Float64) isValue_Value() {}

func (*Value_Bool) isValue_Value() {}

func (*Value_Bytes) isValue_Value() {}

func (*Value_String_) isValue_Value() {}

func (*Value_Time) isValue_Value() {}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns []*Column `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *Row) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *Column) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Column) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_plugin_proto protoreflect.FileDescriptor

var file_plugin_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x73, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x23, 0x0a,
	0x03, 0x74, 0x78, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x78, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x74,
	0x78, 0x6e, 0x22, 0xb4, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x73,
	0x65, 0x72, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x61, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x6f,
	0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x12, 0x23, 0x0a, 0x03, 0x74, 0x78, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x78, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x52, 0x03, 0x74, 0x78, 0x6e, 0x22, 0xb7, 0x01, 0x0a, 0x0a, 0x54, 0x78,
	0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x6f, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x73, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f,
	0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f,
	0x6e, 0x6c, 0x79, 0x22, 0x5e, 0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12,
	0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a,
	0x04, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x75, 0x6c, 0x6c, 0x12, 0x16, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x07, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x07,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2e, 0x0a, 0x03, 0x52, 0x6f,
	0x77, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0x40, 0x0a, 0x06, 0x43, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x6c, 0x0a, 0x06,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x52, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x78, 0x6e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x78, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_proto_rawDescOnce sync.Once
	file_plugin_proto_rawDescData = file_plugin_proto_rawDesc
)

func file_plugin_proto_rawDescGZIP() []byte {
	file_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_proto_rawDescData)
	})
	return file_plugin_proto_rawDescData
}

var file_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_plugin_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: proto.Request
	(*Response)(nil),              // 1: proto.Response
	(*TxnContext)(nil),            // 2: proto.TxnContext
	(*NamedValue)(nil),            // 3: proto.NamedValue
	(*Value)(nil),                 // 4: proto.Value
	(*Row)(nil),                   // 5: proto.Row
	(*Column)(nil),                // 6: proto.Column
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_plugin_proto_depIdxs = []int32{
	3,  // 0: proto.Request.args:type_name -> proto.NamedValue
	2,  // 1: proto.Request.txn:type_name -> proto.TxnContext
	5,  // 2: proto.Response.rows:type_name -> proto.Row
	2,  // 3: proto.Response.txn:type_name -> proto.TxnContext
	4,  // 4: proto.NamedValue.value:type_name -> proto.Value
	7,  // 5: proto.Value.time:type_name -> google.protobuf.Timestamp
	6,  // 6: proto.Row.columns:type_name -> proto.Column
	4,  // 7: proto.Column.value:type_name -> proto.Value
	0,  // 8: proto.Driver.Query:input_type -> proto.Request
	2,  // 9: proto.Driver.CommitOrRollback:input_type -> proto.TxnContext
	1,  // 10: proto.Driver.Query:output_type -> proto.Response
	2,  // 11: proto.Driver.CommitOrRollback:output_type -> proto.TxnContext
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_plugin_proto_init() }
func file_plugin_proto_init() {
	if File_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxnContext); i {
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
		file_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamedValue); i {
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
		file_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_plugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_plugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Column); i {
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
	file_plugin_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Value_Null)(nil),
		(*Value_Int64)(nil),
		(*Value_Float64)(nil),
		(*Value_Bool)(nil),
		(*Value_Bytes)(nil),
		(*Value_String_)(nil),
		(*Value_Time)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_proto_depIdxs,
		MessageInfos:      file_plugin_proto_msgTypes,
	}.Build()
	File_plugin_proto = out.File
	file_plugin_proto_rawDesc = nil
	file_plugin_proto_goTypes = nil
	file_plugin_proto_depIdxs = nil
}
