// Code generated by protoc-gen-go.
// source: snapshot_response.proto
// DO NOT EDIT!

package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ProtoSnapshotResponse struct {
	Success          *bool  `protobuf:"varint,1,req" json:"Success,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ProtoSnapshotResponse) Reset()         { *m = ProtoSnapshotResponse{} }
func (m *ProtoSnapshotResponse) String() string { return proto.CompactTextString(m) }
func (*ProtoSnapshotResponse) ProtoMessage()    {}

func (m *ProtoSnapshotResponse) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func init() {
}
