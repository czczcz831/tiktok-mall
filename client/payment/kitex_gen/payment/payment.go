// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package payment

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type CreditCard struct {
	Uuid               string `thrift:"uuid,1" frugal:"1,default,string" json:"uuid"`
	UserUuid           string `thrift:"user_uuid,2" frugal:"2,default,string" json:"user_uuid"`
	CreditCardNumber   string `thrift:"credit_card_number,3" frugal:"3,default,string" json:"credit_card_number"`
	CreditCardCvv      int64  `thrift:"credit_card_cvv,4" frugal:"4,default,i64" json:"credit_card_cvv"`
	CreditCardExpMonth int64  `thrift:"credit_card_exp_month,5" frugal:"5,default,i64" json:"credit_card_exp_month"`
	CreditCardExpYear  int64  `thrift:"credit_card_exp_year,6" frugal:"6,default,i64" json:"credit_card_exp_year"`
}

func NewCreditCard() *CreditCard {
	return &CreditCard{}
}

func (p *CreditCard) InitDefault() {
}

func (p *CreditCard) GetUuid() (v string) {
	return p.Uuid
}

func (p *CreditCard) GetUserUuid() (v string) {
	return p.UserUuid
}

func (p *CreditCard) GetCreditCardNumber() (v string) {
	return p.CreditCardNumber
}

func (p *CreditCard) GetCreditCardCvv() (v int64) {
	return p.CreditCardCvv
}

func (p *CreditCard) GetCreditCardExpMonth() (v int64) {
	return p.CreditCardExpMonth
}

func (p *CreditCard) GetCreditCardExpYear() (v int64) {
	return p.CreditCardExpYear
}
func (p *CreditCard) SetUuid(val string) {
	p.Uuid = val
}
func (p *CreditCard) SetUserUuid(val string) {
	p.UserUuid = val
}
func (p *CreditCard) SetCreditCardNumber(val string) {
	p.CreditCardNumber = val
}
func (p *CreditCard) SetCreditCardCvv(val int64) {
	p.CreditCardCvv = val
}
func (p *CreditCard) SetCreditCardExpMonth(val int64) {
	p.CreditCardExpMonth = val
}
func (p *CreditCard) SetCreditCardExpYear(val int64) {
	p.CreditCardExpYear = val
}

var fieldIDToName_CreditCard = map[int16]string{
	1: "uuid",
	2: "user_uuid",
	3: "credit_card_number",
	4: "credit_card_cvv",
	5: "credit_card_exp_month",
	6: "credit_card_exp_year",
}

func (p *CreditCard) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 5:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField6(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CreditCard[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *CreditCard) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Uuid = _field
	return nil
}
func (p *CreditCard) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.UserUuid = _field
	return nil
}
func (p *CreditCard) ReadField3(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.CreditCardNumber = _field
	return nil
}
func (p *CreditCard) ReadField4(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.CreditCardCvv = _field
	return nil
}
func (p *CreditCard) ReadField5(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.CreditCardExpMonth = _field
	return nil
}
func (p *CreditCard) ReadField6(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.CreditCardExpYear = _field
	return nil
}

func (p *CreditCard) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("CreditCard"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}
		if err = p.writeField6(oprot); err != nil {
			fieldId = 6
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *CreditCard) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("uuid", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Uuid); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *CreditCard) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user_uuid", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UserUuid); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *CreditCard) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("credit_card_number", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.CreditCardNumber); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *CreditCard) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("credit_card_cvv", thrift.I64, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.CreditCardCvv); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *CreditCard) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("credit_card_exp_month", thrift.I64, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.CreditCardExpMonth); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *CreditCard) writeField6(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("credit_card_exp_year", thrift.I64, 6); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.CreditCardExpYear); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 6 end error: ", p), err)
}

func (p *CreditCard) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreditCard(%+v)", *p)

}

func (p *CreditCard) DeepEqual(ano *CreditCard) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Uuid) {
		return false
	}
	if !p.Field2DeepEqual(ano.UserUuid) {
		return false
	}
	if !p.Field3DeepEqual(ano.CreditCardNumber) {
		return false
	}
	if !p.Field4DeepEqual(ano.CreditCardCvv) {
		return false
	}
	if !p.Field5DeepEqual(ano.CreditCardExpMonth) {
		return false
	}
	if !p.Field6DeepEqual(ano.CreditCardExpYear) {
		return false
	}
	return true
}

func (p *CreditCard) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Uuid, src) != 0 {
		return false
	}
	return true
}
func (p *CreditCard) Field2DeepEqual(src string) bool {

	if strings.Compare(p.UserUuid, src) != 0 {
		return false
	}
	return true
}
func (p *CreditCard) Field3DeepEqual(src string) bool {

	if strings.Compare(p.CreditCardNumber, src) != 0 {
		return false
	}
	return true
}
func (p *CreditCard) Field4DeepEqual(src int64) bool {

	if p.CreditCardCvv != src {
		return false
	}
	return true
}
func (p *CreditCard) Field5DeepEqual(src int64) bool {

	if p.CreditCardExpMonth != src {
		return false
	}
	return true
}
func (p *CreditCard) Field6DeepEqual(src int64) bool {

	if p.CreditCardExpYear != src {
		return false
	}
	return true
}

type ChargeReq struct {
	UserUuid   string      `thrift:"user_uuid,1" frugal:"1,default,string" json:"user_uuid"`
	OrderUuid  string      `thrift:"order_uuid,2" frugal:"2,default,string" json:"order_uuid"`
	Amount     int64       `thrift:"amount,3" frugal:"3,default,i64" json:"amount"`
	CreditCard *CreditCard `thrift:"credit_card,4" frugal:"4,default,CreditCard" json:"credit_card"`
}

func NewChargeReq() *ChargeReq {
	return &ChargeReq{}
}

func (p *ChargeReq) InitDefault() {
}

func (p *ChargeReq) GetUserUuid() (v string) {
	return p.UserUuid
}

func (p *ChargeReq) GetOrderUuid() (v string) {
	return p.OrderUuid
}

func (p *ChargeReq) GetAmount() (v int64) {
	return p.Amount
}

var ChargeReq_CreditCard_DEFAULT *CreditCard

func (p *ChargeReq) GetCreditCard() (v *CreditCard) {
	if !p.IsSetCreditCard() {
		return ChargeReq_CreditCard_DEFAULT
	}
	return p.CreditCard
}
func (p *ChargeReq) SetUserUuid(val string) {
	p.UserUuid = val
}
func (p *ChargeReq) SetOrderUuid(val string) {
	p.OrderUuid = val
}
func (p *ChargeReq) SetAmount(val int64) {
	p.Amount = val
}
func (p *ChargeReq) SetCreditCard(val *CreditCard) {
	p.CreditCard = val
}

var fieldIDToName_ChargeReq = map[int16]string{
	1: "user_uuid",
	2: "order_uuid",
	3: "amount",
	4: "credit_card",
}

func (p *ChargeReq) IsSetCreditCard() bool {
	return p.CreditCard != nil
}

func (p *ChargeReq) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 4:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ChargeReq[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ChargeReq) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.UserUuid = _field
	return nil
}
func (p *ChargeReq) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.OrderUuid = _field
	return nil
}
func (p *ChargeReq) ReadField3(iprot thrift.TProtocol) error {

	var _field int64
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Amount = _field
	return nil
}
func (p *ChargeReq) ReadField4(iprot thrift.TProtocol) error {
	_field := NewCreditCard()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.CreditCard = _field
	return nil
}

func (p *ChargeReq) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("ChargeReq"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ChargeReq) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user_uuid", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UserUuid); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ChargeReq) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("order_uuid", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.OrderUuid); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *ChargeReq) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("amount", thrift.I64, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.Amount); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *ChargeReq) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("credit_card", thrift.STRUCT, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.CreditCard.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *ChargeReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ChargeReq(%+v)", *p)

}

func (p *ChargeReq) DeepEqual(ano *ChargeReq) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.UserUuid) {
		return false
	}
	if !p.Field2DeepEqual(ano.OrderUuid) {
		return false
	}
	if !p.Field3DeepEqual(ano.Amount) {
		return false
	}
	if !p.Field4DeepEqual(ano.CreditCard) {
		return false
	}
	return true
}

func (p *ChargeReq) Field1DeepEqual(src string) bool {

	if strings.Compare(p.UserUuid, src) != 0 {
		return false
	}
	return true
}
func (p *ChargeReq) Field2DeepEqual(src string) bool {

	if strings.Compare(p.OrderUuid, src) != 0 {
		return false
	}
	return true
}
func (p *ChargeReq) Field3DeepEqual(src int64) bool {

	if p.Amount != src {
		return false
	}
	return true
}
func (p *ChargeReq) Field4DeepEqual(src *CreditCard) bool {

	if !p.CreditCard.DeepEqual(src) {
		return false
	}
	return true
}

type ChargeResp struct {
	TransactionUuid string `thrift:"transaction_uuid,1" frugal:"1,default,string" json:"transaction_uuid"`
	Success         bool   `thrift:"success,2" frugal:"2,default,bool" json:"success"`
}

func NewChargeResp() *ChargeResp {
	return &ChargeResp{}
}

func (p *ChargeResp) InitDefault() {
}

func (p *ChargeResp) GetTransactionUuid() (v string) {
	return p.TransactionUuid
}

func (p *ChargeResp) GetSuccess() (v bool) {
	return p.Success
}
func (p *ChargeResp) SetTransactionUuid(val string) {
	p.TransactionUuid = val
}
func (p *ChargeResp) SetSuccess(val bool) {
	p.Success = val
}

var fieldIDToName_ChargeResp = map[int16]string{
	1: "transaction_uuid",
	2: "success",
}

func (p *ChargeResp) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.BOOL {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ChargeResp[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *ChargeResp) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.TransactionUuid = _field
	return nil
}
func (p *ChargeResp) ReadField2(iprot thrift.TProtocol) error {

	var _field bool
	if v, err := iprot.ReadBool(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Success = _field
	return nil
}

func (p *ChargeResp) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("ChargeResp"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *ChargeResp) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("transaction_uuid", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.TransactionUuid); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *ChargeResp) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("success", thrift.BOOL, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteBool(p.Success); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *ChargeResp) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ChargeResp(%+v)", *p)

}

func (p *ChargeResp) DeepEqual(ano *ChargeResp) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.TransactionUuid) {
		return false
	}
	if !p.Field2DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *ChargeResp) Field1DeepEqual(src string) bool {

	if strings.Compare(p.TransactionUuid, src) != 0 {
		return false
	}
	return true
}
func (p *ChargeResp) Field2DeepEqual(src bool) bool {

	if p.Success != src {
		return false
	}
	return true
}

type PaymentService interface {
	Charge(ctx context.Context, req *ChargeReq) (r *ChargeResp, err error)
}

type PaymentServiceClient struct {
	c thrift.TClient
}

func NewPaymentServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PaymentServiceClient {
	return &PaymentServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewPaymentServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PaymentServiceClient {
	return &PaymentServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewPaymentServiceClient(c thrift.TClient) *PaymentServiceClient {
	return &PaymentServiceClient{
		c: c,
	}
}

func (p *PaymentServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *PaymentServiceClient) Charge(ctx context.Context, req *ChargeReq) (r *ChargeResp, err error) {
	var _args PaymentServiceChargeArgs
	_args.Req = req
	var _result PaymentServiceChargeResult
	if err = p.Client_().Call(ctx, "Charge", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type PaymentServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      PaymentService
}

func (p *PaymentServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *PaymentServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *PaymentServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewPaymentServiceProcessor(handler PaymentService) *PaymentServiceProcessor {
	self := &PaymentServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("Charge", &paymentServiceProcessorCharge{handler: handler})
	return self
}
func (p *PaymentServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type paymentServiceProcessorCharge struct {
	handler PaymentService
}

func (p *paymentServiceProcessorCharge) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := PaymentServiceChargeArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Charge", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := PaymentServiceChargeResult{}
	var retval *ChargeResp
	if retval, err2 = p.handler.Charge(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Charge: "+err2.Error())
		oprot.WriteMessageBegin("Charge", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Charge", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type PaymentServiceChargeArgs struct {
	Req *ChargeReq `thrift:"req,1" frugal:"1,default,ChargeReq" json:"req"`
}

func NewPaymentServiceChargeArgs() *PaymentServiceChargeArgs {
	return &PaymentServiceChargeArgs{}
}

func (p *PaymentServiceChargeArgs) InitDefault() {
}

var PaymentServiceChargeArgs_Req_DEFAULT *ChargeReq

func (p *PaymentServiceChargeArgs) GetReq() (v *ChargeReq) {
	if !p.IsSetReq() {
		return PaymentServiceChargeArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *PaymentServiceChargeArgs) SetReq(val *ChargeReq) {
	p.Req = val
}

var fieldIDToName_PaymentServiceChargeArgs = map[int16]string{
	1: "req",
}

func (p *PaymentServiceChargeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PaymentServiceChargeArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PaymentServiceChargeArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PaymentServiceChargeArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewChargeReq()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *PaymentServiceChargeArgs) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("Charge_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PaymentServiceChargeArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *PaymentServiceChargeArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceChargeArgs(%+v)", *p)

}

func (p *PaymentServiceChargeArgs) DeepEqual(ano *PaymentServiceChargeArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *PaymentServiceChargeArgs) Field1DeepEqual(src *ChargeReq) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type PaymentServiceChargeResult struct {
	Success *ChargeResp `thrift:"success,0,optional" frugal:"0,optional,ChargeResp" json:"success,omitempty"`
}

func NewPaymentServiceChargeResult() *PaymentServiceChargeResult {
	return &PaymentServiceChargeResult{}
}

func (p *PaymentServiceChargeResult) InitDefault() {
}

var PaymentServiceChargeResult_Success_DEFAULT *ChargeResp

func (p *PaymentServiceChargeResult) GetSuccess() (v *ChargeResp) {
	if !p.IsSetSuccess() {
		return PaymentServiceChargeResult_Success_DEFAULT
	}
	return p.Success
}
func (p *PaymentServiceChargeResult) SetSuccess(x interface{}) {
	p.Success = x.(*ChargeResp)
}

var fieldIDToName_PaymentServiceChargeResult = map[int16]string{
	0: "success",
}

func (p *PaymentServiceChargeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PaymentServiceChargeResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_PaymentServiceChargeResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *PaymentServiceChargeResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewChargeResp()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *PaymentServiceChargeResult) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("Charge_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *PaymentServiceChargeResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *PaymentServiceChargeResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PaymentServiceChargeResult(%+v)", *p)

}

func (p *PaymentServiceChargeResult) DeepEqual(ano *PaymentServiceChargeResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *PaymentServiceChargeResult) Field0DeepEqual(src *ChargeResp) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
