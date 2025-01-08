// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package base

import (
	"fmt"
	thrift "github.com/cloudwego/kitex/pkg/protocol/bthrift/apache"
	"strings"
)

type ResponseMetadata struct {
	RequestID string `thrift:"RequestID,1,required" frugal:"1,required,string" json:"RequestID"`
	AccountID string `thrift:"AccountID,2,required" frugal:"2,required,string" json:"AccountID"`
}

func NewResponseMetadata() *ResponseMetadata {
	return &ResponseMetadata{}
}

func (p *ResponseMetadata) InitDefault() {
}

func (p *ResponseMetadata) GetRequestID() (v string) {
	return p.RequestID
}

func (p *ResponseMetadata) GetAccountID() (v string) {
	return p.AccountID
}
func (p *ResponseMetadata) SetRequestID(val string) {
	p.RequestID = val
}
func (p *ResponseMetadata) SetAccountID(val string) {
	p.AccountID = val
}

var fieldIDToName_ResponseMetadata = map[int16]string{
	1: "RequestID",
	2: "AccountID",
}

func (p *ResponseMetadata) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetRequestID bool = false
	var issetAccountID bool = false

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
				issetRequestID = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetAccountID = true
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

	if !issetRequestID {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetAccountID {
		fieldId = 2
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_ResponseMetadata[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_ResponseMetadata[fieldId]))
}

func (p *ResponseMetadata) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.RequestID = _field
	return nil
}
func (p *ResponseMetadata) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.AccountID = _field
	return nil
}

func (p *ResponseMetadata) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("ResponseMetadata"); err != nil {
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

func (p *ResponseMetadata) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("RequestID", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.RequestID); err != nil {
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

func (p *ResponseMetadata) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("AccountID", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.AccountID); err != nil {
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

func (p *ResponseMetadata) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ResponseMetadata(%+v)", *p)

}

func (p *ResponseMetadata) DeepEqual(ano *ResponseMetadata) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.RequestID) {
		return false
	}
	if !p.Field2DeepEqual(ano.AccountID) {
		return false
	}
	return true
}

func (p *ResponseMetadata) Field1DeepEqual(src string) bool {

	if strings.Compare(p.RequestID, src) != 0 {
		return false
	}
	return true
}
func (p *ResponseMetadata) Field2DeepEqual(src string) bool {

	if strings.Compare(p.AccountID, src) != 0 {
		return false
	}
	return true
}
