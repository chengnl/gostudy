// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package timestamp

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"vrv/im/service/vrv"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = vrv.GoUnusedProtection__
var GoUnusedProtection__ int

// ********************************页码时间戳类型***************************
// 121: 好友页码
// 122: 表情页码
// 123: 群成员页码
// 124: 群文件页码
//
// Attributes:
//  - Timestamps
//  - Buddytimestamps
type TimestampResult_ struct {
	Timestamps      []int64 `thrift:"timestamps,1" json:"timestamps,omitempty"`
	Buddytimestamps []int64 `thrift:"buddytimestamps,2" json:"buddytimestamps,omitempty"`
}

func NewTimestampResult_() *TimestampResult_ {
	return &TimestampResult_{}
}

var TimestampResult__Timestamps_DEFAULT []int64

func (p *TimestampResult_) GetTimestamps() []int64 {
	return p.Timestamps
}

var TimestampResult__Buddytimestamps_DEFAULT []int64

func (p *TimestampResult_) GetBuddytimestamps() []int64 {
	return p.Buddytimestamps
}
func (p *TimestampResult_) IsSetTimestamps() bool {
	return p.Timestamps != nil
}

func (p *TimestampResult_) IsSetBuddytimestamps() bool {
	return p.Buddytimestamps != nil
}

func (p *TimestampResult_) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TimestampResult_) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]int64, 0, size)
	p.Timestamps = tSlice
	for i := 0; i < size; i++ {
		var _elem0 int64
		if v, err := iprot.ReadI64(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.Timestamps = append(p.Timestamps, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TimestampResult_) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]int64, 0, size)
	p.Buddytimestamps = tSlice
	for i := 0; i < size; i++ {
		var _elem1 int64
		if v, err := iprot.ReadI64(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem1 = v
		}
		p.Buddytimestamps = append(p.Buddytimestamps, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *TimestampResult_) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("TimestampResult"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TimestampResult_) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetTimestamps() {
		if err := oprot.WriteFieldBegin("timestamps", thrift.LIST, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:timestamps: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.I64, len(p.Timestamps)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Timestamps {
			if err := oprot.WriteI64(int64(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:timestamps: ", p), err)
		}
	}
	return err
}

func (p *TimestampResult_) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetBuddytimestamps() {
		if err := oprot.WriteFieldBegin("buddytimestamps", thrift.LIST, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:buddytimestamps: ", p), err)
		}
		if err := oprot.WriteListBegin(thrift.I64, len(p.Buddytimestamps)); err != nil {
			return thrift.PrependError("error writing list begin: ", err)
		}
		for _, v := range p.Buddytimestamps {
			if err := oprot.WriteI64(int64(v)); err != nil {
				return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return thrift.PrependError("error writing list end: ", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:buddytimestamps: ", p), err)
		}
	}
	return err
}

func (p *TimestampResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TimestampResult_(%+v)", *p)
}