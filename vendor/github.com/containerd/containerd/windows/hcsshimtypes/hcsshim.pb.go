// Code generated by protoc-gen-gogo.
// source: github.com/containerd/containerd/windows/hcsshimtypes/hcsshim.proto
// DO NOT EDIT!

/*
	Package hcsshimtypes is a generated protocol buffer package.

	It is generated from these files:
		github.com/containerd/containerd/windows/hcsshimtypes/hcsshim.proto

	It has these top-level messages:
		CreateOptions
		ProcessDetails
*/
package hcsshimtypes

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/types"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CreateOptions struct {
	TerminateDuration time.Duration `protobuf:"bytes,1,opt,name=terminate_duration,json=terminateDuration,stdduration" json:"terminate_duration"`
}

func (m *CreateOptions) Reset()                    { *m = CreateOptions{} }
func (*CreateOptions) ProtoMessage()               {}
func (*CreateOptions) Descriptor() ([]byte, []int) { return fileDescriptorHcsshim, []int{0} }

// ProcessDetails contains additional information about a process
// ProcessDetails is made of the same fields as found in hcsshim.ProcessListItem
type ProcessDetails struct {
	ImageName                    string    `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	CreatedAt                    time.Time `protobuf:"bytes,2,opt,name=created_at,json=createdAt,stdtime" json:"created_at"`
	KernelTime_100Ns             uint64    `protobuf:"varint,3,opt,name=kernel_time_100_ns,json=kernelTime100Ns,proto3" json:"kernel_time_100_ns,omitempty"`
	MemoryCommitBytes            uint64    `protobuf:"varint,4,opt,name=memory_commit_bytes,json=memoryCommitBytes,proto3" json:"memory_commit_bytes,omitempty"`
	MemoryWorkingSetPrivateBytes uint64    `protobuf:"varint,5,opt,name=memory_working_set_private_bytes,json=memoryWorkingSetPrivateBytes,proto3" json:"memory_working_set_private_bytes,omitempty"`
	MemoryWorkingSetSharedBytes  uint64    `protobuf:"varint,6,opt,name=memory_working_set_shared_bytes,json=memoryWorkingSetSharedBytes,proto3" json:"memory_working_set_shared_bytes,omitempty"`
	ProcessID                    uint32    `protobuf:"varint,7,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	UserTime_100Ns               uint64    `protobuf:"varint,8,opt,name=user_time_100_ns,json=userTime100Ns,proto3" json:"user_time_100_ns,omitempty"`
}

func (m *ProcessDetails) Reset()                    { *m = ProcessDetails{} }
func (*ProcessDetails) ProtoMessage()               {}
func (*ProcessDetails) Descriptor() ([]byte, []int) { return fileDescriptorHcsshim, []int{1} }

func init() {
	proto.RegisterType((*CreateOptions)(nil), "containerd.windows.hcsshim.CreateOptions")
	proto.RegisterType((*ProcessDetails)(nil), "containerd.windows.hcsshim.ProcessDetails")
}
func (m *CreateOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintHcsshim(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(m.TerminateDuration)))
	n1, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.TerminateDuration, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *ProcessDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessDetails) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ImageName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(len(m.ImageName)))
		i += copy(dAtA[i:], m.ImageName)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintHcsshim(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)))
	n2, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.KernelTime_100Ns != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.KernelTime_100Ns))
	}
	if m.MemoryCommitBytes != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.MemoryCommitBytes))
	}
	if m.MemoryWorkingSetPrivateBytes != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.MemoryWorkingSetPrivateBytes))
	}
	if m.MemoryWorkingSetSharedBytes != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.MemoryWorkingSetSharedBytes))
	}
	if m.ProcessID != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.ProcessID))
	}
	if m.UserTime_100Ns != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintHcsshim(dAtA, i, uint64(m.UserTime_100Ns))
	}
	return i, nil
}

func encodeFixed64Hcsshim(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Hcsshim(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintHcsshim(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CreateOptions) Size() (n int) {
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.TerminateDuration)
	n += 1 + l + sovHcsshim(uint64(l))
	return n
}

func (m *ProcessDetails) Size() (n int) {
	var l int
	_ = l
	l = len(m.ImageName)
	if l > 0 {
		n += 1 + l + sovHcsshim(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovHcsshim(uint64(l))
	if m.KernelTime_100Ns != 0 {
		n += 1 + sovHcsshim(uint64(m.KernelTime_100Ns))
	}
	if m.MemoryCommitBytes != 0 {
		n += 1 + sovHcsshim(uint64(m.MemoryCommitBytes))
	}
	if m.MemoryWorkingSetPrivateBytes != 0 {
		n += 1 + sovHcsshim(uint64(m.MemoryWorkingSetPrivateBytes))
	}
	if m.MemoryWorkingSetSharedBytes != 0 {
		n += 1 + sovHcsshim(uint64(m.MemoryWorkingSetSharedBytes))
	}
	if m.ProcessID != 0 {
		n += 1 + sovHcsshim(uint64(m.ProcessID))
	}
	if m.UserTime_100Ns != 0 {
		n += 1 + sovHcsshim(uint64(m.UserTime_100Ns))
	}
	return n
}

func sovHcsshim(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHcsshim(x uint64) (n int) {
	return sovHcsshim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CreateOptions) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateOptions{`,
		`TerminateDuration:` + strings.Replace(strings.Replace(this.TerminateDuration.String(), "Duration", "google_protobuf1.Duration", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessDetails) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessDetails{`,
		`ImageName:` + fmt.Sprintf("%v", this.ImageName) + `,`,
		`CreatedAt:` + strings.Replace(strings.Replace(this.CreatedAt.String(), "Timestamp", "google_protobuf2.Timestamp", 1), `&`, ``, 1) + `,`,
		`KernelTime_100Ns:` + fmt.Sprintf("%v", this.KernelTime_100Ns) + `,`,
		`MemoryCommitBytes:` + fmt.Sprintf("%v", this.MemoryCommitBytes) + `,`,
		`MemoryWorkingSetPrivateBytes:` + fmt.Sprintf("%v", this.MemoryWorkingSetPrivateBytes) + `,`,
		`MemoryWorkingSetSharedBytes:` + fmt.Sprintf("%v", this.MemoryWorkingSetSharedBytes) + `,`,
		`ProcessID:` + fmt.Sprintf("%v", this.ProcessID) + `,`,
		`UserTime_100Ns:` + fmt.Sprintf("%v", this.UserTime_100Ns) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHcsshim(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CreateOptions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHcsshim
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
			return fmt.Errorf("proto: CreateOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TerminateDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHcsshim
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.TerminateDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHcsshim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHcsshim
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
func (m *ProcessDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHcsshim
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
			return fmt.Errorf("proto: ProcessDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImageName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
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
				return ErrInvalidLengthHcsshim
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImageName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHcsshim
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KernelTime_100Ns", wireType)
			}
			m.KernelTime_100Ns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KernelTime_100Ns |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryCommitBytes", wireType)
			}
			m.MemoryCommitBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryCommitBytes |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryWorkingSetPrivateBytes", wireType)
			}
			m.MemoryWorkingSetPrivateBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryWorkingSetPrivateBytes |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryWorkingSetSharedBytes", wireType)
			}
			m.MemoryWorkingSetSharedBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryWorkingSetSharedBytes |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessID", wireType)
			}
			m.ProcessID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProcessID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserTime_100Ns", wireType)
			}
			m.UserTime_100Ns = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHcsshim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserTime_100Ns |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHcsshim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHcsshim
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
func skipHcsshim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHcsshim
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
					return 0, ErrIntOverflowHcsshim
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
					return 0, ErrIntOverflowHcsshim
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
				return 0, ErrInvalidLengthHcsshim
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHcsshim
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
				next, err := skipHcsshim(dAtA[start:])
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
	ErrInvalidLengthHcsshim = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHcsshim   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/windows/hcsshimtypes/hcsshim.proto", fileDescriptorHcsshim)
}

var fileDescriptorHcsshim = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x1b, 0x36, 0xc6, 0x62, 0x54, 0x60, 0x86, 0x43, 0x28, 0x90, 0x54, 0xbb, 0x50, 0x09,
	0x94, 0x74, 0x70, 0xe4, 0x44, 0x5a, 0x21, 0xed, 0x32, 0xa6, 0x0c, 0x09, 0x09, 0x21, 0x59, 0x6e,
	0xf2, 0x48, 0xad, 0xd5, 0x71, 0x64, 0xbb, 0x54, 0xbd, 0xf1, 0x11, 0x38, 0xf2, 0x49, 0xf8, 0x0c,
	0x3d, 0x72, 0xe4, 0x34, 0x58, 0x3e, 0x09, 0x8a, 0xed, 0x96, 0x51, 0x38, 0x71, 0xf3, 0xf3, 0xff,
	0xf7, 0x7e, 0xaf, 0x7e, 0x0d, 0x1a, 0x95, 0x4c, 0x4f, 0xe7, 0x93, 0x38, 0x17, 0x3c, 0xc9, 0x45,
	0xa5, 0x29, 0xab, 0x40, 0x16, 0x57, 0x8f, 0x0b, 0x56, 0x15, 0x62, 0xa1, 0x92, 0x69, 0xae, 0xd4,
	0x94, 0x71, 0xbd, 0xac, 0x61, 0x53, 0xc4, 0xb5, 0x14, 0x5a, 0xe0, 0xde, 0x6f, 0x3c, 0x76, 0x78,
	0xec, 0x88, 0xde, 0xbd, 0x52, 0x94, 0xc2, 0x60, 0x49, 0x7b, 0xb2, 0x1d, 0xbd, 0xb0, 0x14, 0xa2,
	0x9c, 0x41, 0x62, 0xaa, 0xc9, 0xfc, 0x43, 0x52, 0xcc, 0x25, 0xd5, 0x4c, 0x54, 0x2e, 0x8f, 0xb6,
	0x73, 0xcd, 0x38, 0x28, 0x4d, 0x79, 0x6d, 0x81, 0xc3, 0x1c, 0x75, 0x47, 0x12, 0xa8, 0x86, 0xd7,
	0x75, 0xdb, 0xa6, 0x70, 0x86, 0xb0, 0x06, 0xc9, 0x59, 0x45, 0x35, 0x90, 0xb5, 0x2d, 0xf0, 0xfa,
	0xde, 0xe0, 0xe6, 0xb3, 0xfb, 0xb1, 0xd5, 0xc5, 0x6b, 0x5d, 0x3c, 0x76, 0x40, 0xba, 0xbf, 0xba,
	0x88, 0x3a, 0x5f, 0x7e, 0x44, 0x5e, 0x76, 0xb0, 0x69, 0x5f, 0x87, 0x87, 0x5f, 0x77, 0xd0, 0xad,
	0x53, 0x29, 0x72, 0x50, 0x6a, 0x0c, 0x9a, 0xb2, 0x99, 0xc2, 0x8f, 0x10, 0x62, 0x9c, 0x96, 0x40,
	0x2a, 0xca, 0xc1, 0xe8, 0xfd, 0xcc, 0x37, 0x37, 0x27, 0x94, 0x03, 0x1e, 0x21, 0x94, 0x9b, 0x9f,
	0x55, 0x10, 0xaa, 0x83, 0x6b, 0x66, 0x7a, 0xef, 0xaf, 0xe9, 0x6f, 0xd6, 0x8f, 0xb1, 0xe3, 0x3f,
	0xb7, 0xe3, 0x7d, 0xd7, 0xf7, 0x52, 0xe3, 0x27, 0x08, 0x9f, 0x83, 0xac, 0x60, 0x46, 0xda, 0x57,
	0x93, 0xa3, 0xe1, 0x90, 0x54, 0x2a, 0xd8, 0xe9, 0x7b, 0x83, 0xdd, 0xec, 0xb6, 0x4d, 0x5a, 0xc3,
	0xd1, 0x70, 0x78, 0xa2, 0x70, 0x8c, 0xee, 0x72, 0xe0, 0x42, 0x2e, 0x49, 0x2e, 0x38, 0x67, 0x9a,
	0x4c, 0x96, 0x1a, 0x54, 0xb0, 0x6b, 0xe8, 0x03, 0x1b, 0x8d, 0x4c, 0x92, 0xb6, 0x01, 0x7e, 0x85,
	0xfa, 0x8e, 0x5f, 0x08, 0x79, 0xce, 0xaa, 0x92, 0x28, 0xd0, 0xa4, 0x96, 0xec, 0x63, 0xbb, 0x38,
	0xdb, 0x7c, 0xdd, 0x34, 0x3f, 0xb4, 0xdc, 0x5b, 0x8b, 0x9d, 0x81, 0x3e, 0xb5, 0x90, 0xf5, 0x8c,
	0x51, 0xf4, 0x0f, 0x8f, 0x9a, 0x52, 0x09, 0x85, 0xd3, 0xec, 0x19, 0xcd, 0x83, 0x6d, 0xcd, 0x99,
	0x61, 0xac, 0xe5, 0x29, 0x42, 0xb5, 0x5d, 0x30, 0x61, 0x45, 0x70, 0xa3, 0xef, 0x0d, 0xba, 0x69,
	0xb7, 0xb9, 0x88, 0x7c, 0xb7, 0xf6, 0xe3, 0x71, 0xe6, 0x3b, 0xe0, 0xb8, 0xc0, 0x8f, 0xd1, 0x9d,
	0xb9, 0x02, 0xf9, 0xc7, 0x5a, 0xf6, 0xcd, 0x90, 0x6e, 0x7b, 0xbf, 0x59, 0x4a, 0xfa, 0x7e, 0x75,
	0x19, 0x76, 0xbe, 0x5f, 0x86, 0x9d, 0x4f, 0x4d, 0xe8, 0xad, 0x9a, 0xd0, 0xfb, 0xd6, 0x84, 0xde,
	0xcf, 0x26, 0xf4, 0xde, 0xa5, 0xff, 0xf5, 0xbd, 0xbf, 0xb8, 0x5a, 0x4c, 0xf6, 0xcc, 0x1f, 0xf9,
	0xfc, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xad, 0xdf, 0xf8, 0x3c, 0x03, 0x00, 0x00,
}
