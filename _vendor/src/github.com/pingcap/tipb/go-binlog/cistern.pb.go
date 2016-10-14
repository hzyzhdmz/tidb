// Code generated by protoc-gen-gogo.
// source: cistern.proto
// DO NOT EDIT!

/*
	Package binlog is a generated protocol buffer package.

	It is generated from these files:
		cistern.proto

	It has these top-level messages:
		DumpBinlogReq
		DumpBinlogResp
*/
package binlog

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	io "io"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DumpBinlogReq struct {
	// beginCommitTS speicifies the position from which begin to dump binlogs.
	// note that actually the result of dump starts from the one next to beginCommitTS
	// it should be zero in case of the first request.
	BeginCommitTS int64 `protobuf:"varint,1,opt,name=beginCommitTS,proto3" json:"beginCommitTS,omitempty"`
	// limit defines the maximum number of binlogs requested once call.
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *DumpBinlogReq) Reset()                    { *m = DumpBinlogReq{} }
func (m *DumpBinlogReq) String() string            { return proto.CompactTextString(m) }
func (*DumpBinlogReq) ProtoMessage()               {}
func (*DumpBinlogReq) Descriptor() ([]byte, []int) { return fileDescriptorCistern, []int{0} }

type DumpBinlogResp struct {
	// An empty errmsg means that the successful acquisition of binlogs.
	Errmsg string `protobuf:"bytes,1,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	// payloads is an array of binlog bytecodes returned.
	Payloads [][]byte `protobuf:"bytes,2,rep,name=payloads" json:"payloads,omitempty"`
	// endCommitTS is the commitTS of the last one in binlog payloads
	// client can use it as beginCommitTS of the next request of dump.
	EndCommitTS int64 `protobuf:"varint,3,opt,name=endCommitTS,proto3" json:"endCommitTS,omitempty"`
}

func (m *DumpBinlogResp) Reset()                    { *m = DumpBinlogResp{} }
func (m *DumpBinlogResp) String() string            { return proto.CompactTextString(m) }
func (*DumpBinlogResp) ProtoMessage()               {}
func (*DumpBinlogResp) Descriptor() ([]byte, []int) { return fileDescriptorCistern, []int{1} }

func init() {
	proto.RegisterType((*DumpBinlogReq)(nil), "binlog.DumpBinlogReq")
	proto.RegisterType((*DumpBinlogResp)(nil), "binlog.DumpBinlogResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Cistern service

type CisternClient interface {
	// DumpBinlog dumps a continuous binlogs from a given position in binlog-server.
	DumpBinlog(ctx context.Context, in *DumpBinlogReq, opts ...grpc.CallOption) (*DumpBinlogResp, error)
}

type cisternClient struct {
	cc *grpc.ClientConn
}

func NewCisternClient(cc *grpc.ClientConn) CisternClient {
	return &cisternClient{cc}
}

func (c *cisternClient) DumpBinlog(ctx context.Context, in *DumpBinlogReq, opts ...grpc.CallOption) (*DumpBinlogResp, error) {
	out := new(DumpBinlogResp)
	err := grpc.Invoke(ctx, "/binlog.Cistern/DumpBinlog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cistern service

type CisternServer interface {
	// DumpBinlog dumps a continuous binlogs from a given position in binlog-server.
	DumpBinlog(context.Context, *DumpBinlogReq) (*DumpBinlogResp, error)
}

func RegisterCisternServer(s *grpc.Server, srv CisternServer) {
	s.RegisterService(&_Cistern_serviceDesc, srv)
}

func _Cistern_DumpBinlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpBinlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CisternServer).DumpBinlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binlog.Cistern/DumpBinlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CisternServer).DumpBinlog(ctx, req.(*DumpBinlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cistern_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binlog.Cistern",
	HandlerType: (*CisternServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DumpBinlog",
			Handler:    _Cistern_DumpBinlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptorCistern,
}

func (m *DumpBinlogReq) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DumpBinlogReq) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BeginCommitTS != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintCistern(data, i, uint64(m.BeginCommitTS))
	}
	if m.Limit != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintCistern(data, i, uint64(m.Limit))
	}
	return i, nil
}

func (m *DumpBinlogResp) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *DumpBinlogResp) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Errmsg) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintCistern(data, i, uint64(len(m.Errmsg)))
		i += copy(data[i:], m.Errmsg)
	}
	if len(m.Payloads) > 0 {
		for _, b := range m.Payloads {
			data[i] = 0x12
			i++
			i = encodeVarintCistern(data, i, uint64(len(b)))
			i += copy(data[i:], b)
		}
	}
	if m.EndCommitTS != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintCistern(data, i, uint64(m.EndCommitTS))
	}
	return i, nil
}

func encodeFixed64Cistern(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Cistern(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCistern(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *DumpBinlogReq) Size() (n int) {
	var l int
	_ = l
	if m.BeginCommitTS != 0 {
		n += 1 + sovCistern(uint64(m.BeginCommitTS))
	}
	if m.Limit != 0 {
		n += 1 + sovCistern(uint64(m.Limit))
	}
	return n
}

func (m *DumpBinlogResp) Size() (n int) {
	var l int
	_ = l
	l = len(m.Errmsg)
	if l > 0 {
		n += 1 + l + sovCistern(uint64(l))
	}
	if len(m.Payloads) > 0 {
		for _, b := range m.Payloads {
			l = len(b)
			n += 1 + l + sovCistern(uint64(l))
		}
	}
	if m.EndCommitTS != 0 {
		n += 1 + sovCistern(uint64(m.EndCommitTS))
	}
	return n
}

func sovCistern(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCistern(x uint64) (n int) {
	return sovCistern(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DumpBinlogReq) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCistern
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DumpBinlogReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DumpBinlogReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeginCommitTS", wireType)
			}
			m.BeginCommitTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.BeginCommitTS |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Limit |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCistern(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCistern
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
func (m *DumpBinlogResp) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCistern
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DumpBinlogResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DumpBinlogResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errmsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCistern
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errmsg = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payloads", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCistern
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payloads = append(m.Payloads, make([]byte, postIndex-iNdEx))
			copy(m.Payloads[len(m.Payloads)-1], data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndCommitTS", wireType)
			}
			m.EndCommitTS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.EndCommitTS |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCistern(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCistern
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
func skipCistern(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCistern
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowCistern
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCistern
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCistern
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipCistern(data[start:])
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
	ErrInvalidLengthCistern = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCistern   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cistern.proto", fileDescriptorCistern) }

var fileDescriptorCistern = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x2c, 0x2e,
	0x49, 0x2d, 0xca, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xca, 0xcc, 0xcb, 0xc9,
	0x4f, 0x97, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x0b, 0xe9, 0x83, 0x58, 0x10, 0x59, 0x25, 0x6f,
	0x2e, 0x5e, 0x97, 0xd2, 0xdc, 0x02, 0x27, 0xb0, 0x9a, 0xa0, 0xd4, 0x42, 0x21, 0x15, 0x2e, 0xde,
	0xa4, 0xd4, 0xf4, 0xcc, 0x3c, 0xe7, 0xfc, 0xdc, 0xdc, 0xcc, 0x92, 0x90, 0x60, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xe6, 0x20, 0x54, 0x41, 0x21, 0x11, 0x2e, 0xd6, 0x9c, 0x4c, 0x20, 0x53, 0x82, 0x09,
	0x28, 0xcb, 0x1a, 0x04, 0xe1, 0x28, 0xa5, 0x71, 0xf1, 0x21, 0x1b, 0x56, 0x5c, 0x20, 0x24, 0xc6,
	0xc5, 0x96, 0x5a, 0x54, 0x94, 0x5b, 0x9c, 0x0e, 0x36, 0x86, 0x33, 0x08, 0xca, 0x13, 0x92, 0xe2,
	0xe2, 0x28, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x29, 0x06, 0x1a, 0xc1, 0xac, 0xc1, 0x13, 0x04,
	0xe7, 0x0b, 0x29, 0x70, 0x71, 0xa7, 0xe6, 0xa5, 0xc0, 0xed, 0x67, 0x06, 0xdb, 0x8f, 0x2c, 0x64,
	0xe4, 0xc1, 0xc5, 0xee, 0x0c, 0xf1, 0xa3, 0x90, 0x2d, 0x17, 0x17, 0xc2, 0x4a, 0x21, 0x51, 0x3d,
	0x88, 0x67, 0xf5, 0x50, 0xfc, 0x24, 0x25, 0x86, 0x4d, 0xb8, 0xb8, 0x40, 0x89, 0xc1, 0x49, 0xe0,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x40, 0xfc, 0x00, 0x88, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03,
	0x87, 0x8b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x28, 0xb0, 0xc9, 0x70, 0x46, 0x01, 0x00, 0x00,
}
