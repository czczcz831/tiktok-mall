// Code generated by Kitex v0.9.1. DO NOT EDIT.

package order

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/cloudwego/kitex/pkg/protocol/bthrift"

	"github.com/czczcz831/tiktok-mall/app/order/kitex_gen/checkout"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
	_ = bthrift.BinaryWriter(nil)
	_ = checkout.KitexUnusedProtection
)

func (p *Order) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.BOOL {
				l, err = p.FastReadField4(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				l, err = p.FastReadField5(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 6:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField6(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_Order[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *Order) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Uuid = v

	}
	return offset, nil
}

func (p *Order) FastReadField2(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.UserUuid = v

	}
	return offset, nil
}

func (p *Order) FastReadField3(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Total = v

	}
	return offset, nil
}

func (p *Order) FastReadField4(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadBool(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.IsPaid = v

	}
	return offset, nil
}

func (p *Order) FastReadField5(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadI64(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.CreatedAt = v

	}
	return offset, nil
}

func (p *Order) FastReadField6(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.Items = make([]*checkout.OrderItem, 0, size)
	for i := 0; i < size; i++ {
		_elem := checkout.NewOrderItem()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.Items = append(p.Items, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

// for compatibility
func (p *Order) FastWrite(buf []byte) int {
	return 0
}

func (p *Order) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "Order")
	if p != nil {
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
		offset += p.fastWriteField4(buf[offset:], binaryWriter)
		offset += p.fastWriteField5(buf[offset:], binaryWriter)
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField6(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *Order) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("Order")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
		l += p.field4Length()
		l += p.field5Length()
		l += p.field6Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *Order) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "uuid", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Uuid)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Order) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "user_uuid", thrift.STRING, 2)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.UserUuid)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Order) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "total", thrift.I64, 3)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.Total)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Order) fastWriteField4(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "is_paid", thrift.BOOL, 4)
	offset += bthrift.Binary.WriteBool(buf[offset:], p.IsPaid)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Order) fastWriteField5(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "created_at", thrift.I64, 5)
	offset += bthrift.Binary.WriteI64(buf[offset:], p.CreatedAt)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Order) fastWriteField6(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "items", thrift.LIST, 6)
	listBeginOffset := offset
	offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
	var length int
	for _, v := range p.Items {
		length++
		offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
	}
	bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
	offset += bthrift.Binary.WriteListEnd(buf[offset:])
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *Order) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("uuid", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.Uuid)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Order) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("user_uuid", thrift.STRING, 2)
	l += bthrift.Binary.StringLengthNocopy(p.UserUuid)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Order) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("total", thrift.I64, 3)
	l += bthrift.Binary.I64Length(p.Total)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Order) field4Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("is_paid", thrift.BOOL, 4)
	l += bthrift.Binary.BoolLength(p.IsPaid)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Order) field5Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("created_at", thrift.I64, 5)
	l += bthrift.Binary.I64Length(p.CreatedAt)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *Order) field6Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("items", thrift.LIST, 6)
	l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.Items))
	for _, v := range p.Items {
		l += v.BLength()
	}
	l += bthrift.Binary.ListEndLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *CreateOrderReq) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField2(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.LIST {
				l, err = p.FastReadField3(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CreateOrderReq[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *CreateOrderReq) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.UserUuid = v

	}
	return offset, nil
}

func (p *CreateOrderReq) FastReadField2(buf []byte) (int, error) {
	offset := 0

	tmp := checkout.NewAddress()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Address = tmp
	return offset, nil
}

func (p *CreateOrderReq) FastReadField3(buf []byte) (int, error) {
	offset := 0

	_, size, l, err := bthrift.Binary.ReadListBegin(buf[offset:])
	offset += l
	if err != nil {
		return offset, err
	}
	p.Items = make([]*checkout.OrderItem, 0, size)
	for i := 0; i < size; i++ {
		_elem := checkout.NewOrderItem()
		if l, err := _elem.FastRead(buf[offset:]); err != nil {
			return offset, err
		} else {
			offset += l
		}

		p.Items = append(p.Items, _elem)
	}
	if l, err := bthrift.Binary.ReadListEnd(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	return offset, nil
}

// for compatibility
func (p *CreateOrderReq) FastWrite(buf []byte) int {
	return 0
}

func (p *CreateOrderReq) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "CreateOrderReq")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
		offset += p.fastWriteField2(buf[offset:], binaryWriter)
		offset += p.fastWriteField3(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *CreateOrderReq) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("CreateOrderReq")
	if p != nil {
		l += p.field1Length()
		l += p.field2Length()
		l += p.field3Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *CreateOrderReq) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "user_uuid", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.UserUuid)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *CreateOrderReq) fastWriteField2(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "address", thrift.STRUCT, 2)
	offset += p.Address.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *CreateOrderReq) fastWriteField3(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "items", thrift.LIST, 3)
	listBeginOffset := offset
	offset += bthrift.Binary.ListBeginLength(thrift.STRUCT, 0)
	var length int
	for _, v := range p.Items {
		length++
		offset += v.FastWriteNocopy(buf[offset:], binaryWriter)
	}
	bthrift.Binary.WriteListBegin(buf[listBeginOffset:], thrift.STRUCT, length)
	offset += bthrift.Binary.WriteListEnd(buf[offset:])
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *CreateOrderReq) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("user_uuid", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.UserUuid)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *CreateOrderReq) field2Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("address", thrift.STRUCT, 2)
	l += p.Address.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *CreateOrderReq) field3Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("items", thrift.LIST, 3)
	l += bthrift.Binary.ListBeginLength(thrift.STRUCT, len(p.Items))
	for _, v := range p.Items {
		l += v.BLength()
	}
	l += bthrift.Binary.ListEndLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *CreateOrderResp) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CreateOrderResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *CreateOrderResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := NewOrder()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Order = tmp
	return offset, nil
}

// for compatibility
func (p *CreateOrderResp) FastWrite(buf []byte) int {
	return 0
}

func (p *CreateOrderResp) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "CreateOrderResp")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *CreateOrderResp) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("CreateOrderResp")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *CreateOrderResp) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "order", thrift.STRUCT, 1)
	offset += p.Order.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *CreateOrderResp) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("order", thrift.STRUCT, 1)
	l += p.Order.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *MarkOrderPaidReq) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_MarkOrderPaidReq[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *MarkOrderPaidReq) FastReadField1(buf []byte) (int, error) {
	offset := 0

	if v, l, err := bthrift.Binary.ReadString(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l

		p.Uuid = v

	}
	return offset, nil
}

// for compatibility
func (p *MarkOrderPaidReq) FastWrite(buf []byte) int {
	return 0
}

func (p *MarkOrderPaidReq) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "MarkOrderPaidReq")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *MarkOrderPaidReq) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("MarkOrderPaidReq")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *MarkOrderPaidReq) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "uuid", thrift.STRING, 1)
	offset += bthrift.Binary.WriteStringNocopy(buf[offset:], binaryWriter, p.Uuid)

	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *MarkOrderPaidReq) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("uuid", thrift.STRING, 1)
	l += bthrift.Binary.StringLengthNocopy(p.Uuid)

	l += bthrift.Binary.FieldEndLength()
	return l
}

func (p *MarkOrderPaidResp) FastRead(buf []byte) (int, error) {
	var err error
	var offset int
	var l int
	var fieldTypeId thrift.TType
	var fieldId int16
	_, l, err = bthrift.Binary.ReadStructBegin(buf)
	offset += l
	if err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, l, err = bthrift.Binary.ReadFieldBegin(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				l, err = p.FastReadField1(buf[offset:])
				offset += l
				if err != nil {
					goto ReadFieldError
				}
			} else {
				l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
				offset += l
				if err != nil {
					goto SkipFieldError
				}
			}
		default:
			l, err = bthrift.Binary.Skip(buf[offset:], fieldTypeId)
			offset += l
			if err != nil {
				goto SkipFieldError
			}
		}

		l, err = bthrift.Binary.ReadFieldEnd(buf[offset:])
		offset += l
		if err != nil {
			goto ReadFieldEndError
		}
	}
	l, err = bthrift.Binary.ReadStructEnd(buf[offset:])
	offset += l
	if err != nil {
		goto ReadStructEndError
	}

	return offset, nil
ReadStructBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_MarkOrderPaidResp[fieldId]), err)
SkipFieldError:
	return offset, thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)
ReadFieldEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return offset, thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *MarkOrderPaidResp) FastReadField1(buf []byte) (int, error) {
	offset := 0

	tmp := NewOrder()
	if l, err := tmp.FastRead(buf[offset:]); err != nil {
		return offset, err
	} else {
		offset += l
	}
	p.Order = tmp
	return offset, nil
}

// for compatibility
func (p *MarkOrderPaidResp) FastWrite(buf []byte) int {
	return 0
}

func (p *MarkOrderPaidResp) FastWriteNocopy(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteStructBegin(buf[offset:], "MarkOrderPaidResp")
	if p != nil {
		offset += p.fastWriteField1(buf[offset:], binaryWriter)
	}
	offset += bthrift.Binary.WriteFieldStop(buf[offset:])
	offset += bthrift.Binary.WriteStructEnd(buf[offset:])
	return offset
}

func (p *MarkOrderPaidResp) BLength() int {
	l := 0
	l += bthrift.Binary.StructBeginLength("MarkOrderPaidResp")
	if p != nil {
		l += p.field1Length()
	}
	l += bthrift.Binary.FieldStopLength()
	l += bthrift.Binary.StructEndLength()
	return l
}

func (p *MarkOrderPaidResp) fastWriteField1(buf []byte, binaryWriter bthrift.BinaryWriter) int {
	offset := 0
	offset += bthrift.Binary.WriteFieldBegin(buf[offset:], "order", thrift.STRUCT, 1)
	offset += p.Order.FastWriteNocopy(buf[offset:], binaryWriter)
	offset += bthrift.Binary.WriteFieldEnd(buf[offset:])
	return offset
}

func (p *MarkOrderPaidResp) field1Length() int {
	l := 0
	l += bthrift.Binary.FieldBeginLength("order", thrift.STRUCT, 1)
	l += p.Order.BLength()
	l += bthrift.Binary.FieldEndLength()
	return l
}
