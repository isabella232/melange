// Code generated by protoc-gen-go.
// source: source/dap.proto
// DO NOT EDIT!

package wire

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Register struct {
	Keys             []*Data `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}

func (m *Register) GetKeys() []*Data {
	if m != nil {
		return m.Keys
	}
	return nil
}

type Unregister struct {
	Keys             []*Data `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Unregister) Reset()         { *m = Unregister{} }
func (m *Unregister) String() string { return proto.CompactTextString(m) }
func (*Unregister) ProtoMessage()    {}

func (m *Unregister) GetKeys() []*Data {
	if m != nil {
		return m.Keys
	}
	return nil
}

type DownloadMessages struct {
	Since            *uint64 `protobuf:"varint,1,opt,name=since" json:"since,omitempty"`
	Context          *bool   `protobuf:"varint,2,opt,name=context" json:"context,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DownloadMessages) Reset()         { *m = DownloadMessages{} }
func (m *DownloadMessages) String() string { return proto.CompactTextString(m) }
func (*DownloadMessages) ProtoMessage()    {}

func (m *DownloadMessages) GetSince() uint64 {
	if m != nil && m.Since != nil {
		return *m.Since
	}
	return 0
}

func (m *DownloadMessages) GetContext() bool {
	if m != nil && m.Context != nil {
		return *m.Context
	}
	return false
}

type PublishMessage struct {
	Data             []byte   `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	To               []string `protobuf:"bytes,2,rep,name=to" json:"to,omitempty"`
	Name             *string  `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	Alert            *bool    `protobuf:"varint,4,opt,name=alert" json:"alert,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PublishMessage) Reset()         { *m = PublishMessage{} }
func (m *PublishMessage) String() string { return proto.CompactTextString(m) }
func (*PublishMessage) ProtoMessage()    {}

func (m *PublishMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PublishMessage) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PublishMessage) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PublishMessage) GetAlert() bool {
	if m != nil && m.Alert != nil {
		return *m.Alert
	}
	return false
}

type UpdateMessage struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Data             []byte  `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateMessage) Reset()         { *m = UpdateMessage{} }
func (m *UpdateMessage) String() string { return proto.CompactTextString(m) }
func (*UpdateMessage) ProtoMessage()    {}

func (m *UpdateMessage) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UpdateMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PublishDataMessage struct {
	Header           []byte   `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	To               []string `protobuf:"bytes,2,rep,name=to" json:"to,omitempty"`
	Name             *string  `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	Length           *uint64  `protobuf:"varint,4,req,name=length" json:"length,omitempty"`
	Hash             []byte   `protobuf:"bytes,5,req,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PublishDataMessage) Reset()         { *m = PublishDataMessage{} }
func (m *PublishDataMessage) String() string { return proto.CompactTextString(m) }
func (*PublishDataMessage) ProtoMessage()    {}

func (m *PublishDataMessage) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PublishDataMessage) GetTo() []string {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PublishDataMessage) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PublishDataMessage) GetLength() uint64 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *PublishDataMessage) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Data struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Data             []byte  `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}

func (m *Data) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetData struct {
	Key              *string `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetData) Reset()         { *m = GetData{} }
func (m *GetData) String() string { return proto.CompactTextString(m) }
func (*GetData) ProtoMessage()    {}

func (m *GetData) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

type Response struct {
	Code             *uint32 `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Message          *string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Length           *uint64 `protobuf:"varint,3,opt,name=length" json:"length,omitempty"`
	Data             []byte  `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetCode() uint32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Response) GetLength() uint64 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ResponseMessage struct {
	Data             []byte  `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	Context          []*Data `protobuf:"bytes,2,rep,name=context" json:"context,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseMessage) Reset()         { *m = ResponseMessage{} }
func (m *ResponseMessage) String() string { return proto.CompactTextString(m) }
func (*ResponseMessage) ProtoMessage()    {}

func (m *ResponseMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ResponseMessage) GetContext() []*Data {
	if m != nil {
		return m.Context
	}
	return nil
}

func init() {
}