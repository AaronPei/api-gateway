// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/api-gateway/example/echo/service/echo.proto

/*
	Package echo is a generated protocol buffer package.

	It is generated from these files:
		github.com/api-gateway/example/echo/service/echo.proto

	It has these top-level messages:
		EchoRequest
		EchoResponse
*/
package echo

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import http "net/http"
import strings "strings"
import math "math"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EchoRequest struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptorEcho, []int{0} }

func (m *EchoRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type EchoResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptorEcho, []int{1} }

func (m *EchoResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "echo.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "echo.EchoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	Ping(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf2.Timestamp, error)
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Ping(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf2.Timestamp, error) {
	out := new(google_protobuf2.Timestamp)
	err := grpc.Invoke(ctx, "/echo.Echo/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := grpc.Invoke(ctx, "/echo.Echo/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

type EchoServer interface {
	Ping(context.Context, *google_protobuf1.Empty) (*google_protobuf2.Timestamp, error)
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Ping(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Echo_Ping_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _Echo_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/api-gateway/example/echo/service/echo.proto",
}

func (m *EchoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEcho(dAtA, i, uint64(len(m.Text)))
		i += copy(dAtA[i:], m.Text)
	}
	return i, nil
}

func (m *EchoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEcho(dAtA, i, uint64(len(m.Text)))
		i += copy(dAtA[i:], m.Text)
	}
	return i, nil
}

func encodeVarintEcho(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EchoRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovEcho(uint64(l))
	}
	return n
}

func (m *EchoResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovEcho(uint64(l))
	}
	return n
}

func sovEcho(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEcho(x uint64) (n int) {
	return sovEcho(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EchoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEcho
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EchoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEcho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEcho
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEcho(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEcho
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EchoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEcho
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EchoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEcho
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEcho
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEcho(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEcho
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEcho(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEcho
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEcho
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEcho
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEcho
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEcho
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEcho(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEcho = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEcho   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/api-gateway/example/echo/service/echo.proto", fileDescriptorEcho)
}

var fileDescriptorEcho = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x89, 0x14, 0xd1, 0x38, 0x82, 0x46, 0x10, 0xa9, 0x52, 0x35, 0x2b, 0x11, 0x4c, 0x50,
	0xc1, 0x85, 0xcb, 0x81, 0xd9, 0xb9, 0x90, 0xc1, 0x0b, 0xa4, 0xe5, 0xd9, 0x06, 0xa6, 0x4d, 0x9c,
	0xbc, 0x8e, 0x33, 0x5b, 0xaf, 0xe0, 0xc6, 0x23, 0xb9, 0x14, 0xbc, 0x80, 0x14, 0x0f, 0x22, 0x49,
	0x3a, 0x30, 0xa8, 0xbb, 0xf7, 0xde, 0xff, 0xe5, 0xff, 0x7f, 0x42, 0x6f, 0x4a, 0x8d, 0x55, 0x9b,
	0x8b, 0xc2, 0xd4, 0x52, 0x59, 0x7d, 0x51, 0x2a, 0x84, 0x67, 0xb5, 0x90, 0x30, 0x57, 0xb5, 0x9d,
	0x80, 0x84, 0xa2, 0x32, 0xd2, 0xc1, 0x74, 0xa6, 0x8b, 0xb8, 0x08, 0x3b, 0x35, 0x68, 0x58, 0xe2,
	0xe7, 0xf4, 0xa8, 0x34, 0xa6, 0x9c, 0x80, 0x7f, 0x29, 0x55, 0xd3, 0x18, 0x54, 0xa8, 0x4d, 0xe3,
	0x22, 0x93, 0x1e, 0xf6, 0x6a, 0xd8, 0xf2, 0xf6, 0x51, 0x42, 0x6d, 0x71, 0xd1, 0x8b, 0xc7, 0xbf,
	0x45, 0xd4, 0x35, 0x38, 0x54, 0xb5, 0x8d, 0x00, 0x3f, 0xa5, 0x5b, 0xa3, 0xa2, 0x32, 0x63, 0x78,
	0x6a, 0xc1, 0x21, 0x63, 0x34, 0x41, 0x98, 0xe3, 0x01, 0x39, 0x21, 0x67, 0x9b, 0xe3, 0x30, 0x73,
	0x4e, 0x07, 0x11, 0x71, 0xd6, 0x34, 0x0e, 0xfe, 0x63, 0xae, 0xde, 0x08, 0x4d, 0x3c, 0xc4, 0xee,
	0x68, 0x72, 0xaf, 0x9b, 0x92, 0xed, 0x8b, 0x98, 0x2c, 0x96, 0xc9, 0x62, 0xe4, 0x6b, 0xa5, 0xe9,
	0x9f, 0xfb, 0xc3, 0xb2, 0x11, 0xdf, 0x79, 0xf9, 0xfc, 0x7e, 0x5d, 0xa3, 0x6c, 0x43, 0xce, 0x2e,
	0xa5, 0xf5, 0x2e, 0xc3, 0xde, 0x75, 0x57, 0x84, 0x4f, 0x59, 0x69, 0x9a, 0xb2, 0xd5, 0x53, 0x6c,
	0xc6, 0xf7, 0x82, 0xc1, 0x36, 0x0f, 0x06, 0x5e, 0xbe, 0x25, 0xe7, 0xc3, 0xc1, 0x7b, 0x97, 0x91,
	0x8f, 0x2e, 0x23, 0x5f, 0x5d, 0x46, 0xf2, 0xf5, 0x90, 0x7b, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xff, 0xe1, 0xe8, 0x69, 0x92, 0x01, 0x00, 0x00,
}

const PROTO_JSON ="[{\"Package\":\"echo\",\"Service\":\"Echo\",\"Method\":{\"name\":\"Ping\",\"input_type\":\".google.protobuf.Empty\",\"output_type\":\".google.protobuf.Timestamp\",\"options\":{}},\"Pattern\":{\"Verb\":\"GET\",\"Path\":\"/v1/ping\",\"Body\":\"\"},\"Options\":{}},{\"Package\":\"echo\",\"Service\":\"Echo\",\"Method\":{\"name\":\"Echo\",\"input_type\":\".echo.EchoRequest\",\"output_type\":\".echo.EchoResponse\",\"options\":{}},\"Pattern\":{\"Verb\":\"POST\",\"Path\":\"/v1/echo\",\"Body\":\"*\"},\"Options\":{}}]"

func init() {
	 if _, err := (&http.Client{}).Post("http://api-gateway:8080/loader", "", strings.NewReader(PROTO_JSON)); err != nil {
			fmt.Println(err)
	 }
}