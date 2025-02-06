// Code generated by thriftgo (0.3.18). DO NOT EDIT.

package comment

import (
	"context"
	"fmt"
	thrift "github.com/cloudwego/kitex/pkg/protocol/bthrift/apache"
	"strings"
)

type CreateCommentRequest struct {
	AuthorID  string  `thrift:"AuthorID,1,required" frugal:"1,required,string" json:"AuthorID"`
	ArticleID string  `thrift:"ArticleID,2,required" frugal:"2,required,string" json:"ArticleID"`
	ParentID  *string `thrift:"ParentID,3,optional" frugal:"3,optional,string" json:"ParentID,omitempty"`
	Content   string  `thrift:"Content,4,required" frugal:"4,required,string" json:"Content"`
}

func NewCreateCommentRequest() *CreateCommentRequest {
	return &CreateCommentRequest{}
}

func (p *CreateCommentRequest) InitDefault() {
}

func (p *CreateCommentRequest) GetAuthorID() (v string) {
	return p.AuthorID
}

func (p *CreateCommentRequest) GetArticleID() (v string) {
	return p.ArticleID
}

var CreateCommentRequest_ParentID_DEFAULT string

func (p *CreateCommentRequest) GetParentID() (v string) {
	if !p.IsSetParentID() {
		return CreateCommentRequest_ParentID_DEFAULT
	}
	return *p.ParentID
}

func (p *CreateCommentRequest) GetContent() (v string) {
	return p.Content
}
func (p *CreateCommentRequest) SetAuthorID(val string) {
	p.AuthorID = val
}
func (p *CreateCommentRequest) SetArticleID(val string) {
	p.ArticleID = val
}
func (p *CreateCommentRequest) SetParentID(val *string) {
	p.ParentID = val
}
func (p *CreateCommentRequest) SetContent(val string) {
	p.Content = val
}

var fieldIDToName_CreateCommentRequest = map[int16]string{
	1: "AuthorID",
	2: "ArticleID",
	3: "ParentID",
	4: "Content",
}

func (p *CreateCommentRequest) IsSetParentID() bool {
	return p.ParentID != nil
}

func (p *CreateCommentRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetAuthorID bool = false
	var issetArticleID bool = false
	var issetContent bool = false

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
				issetAuthorID = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
				issetArticleID = true
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
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
				issetContent = true
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

	if !issetAuthorID {
		fieldId = 1
		goto RequiredFieldNotSetError
	}

	if !issetArticleID {
		fieldId = 2
		goto RequiredFieldNotSetError
	}

	if !issetContent {
		fieldId = 4
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CreateCommentRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_CreateCommentRequest[fieldId]))
}

func (p *CreateCommentRequest) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.AuthorID = _field
	return nil
}
func (p *CreateCommentRequest) ReadField2(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.ArticleID = _field
	return nil
}
func (p *CreateCommentRequest) ReadField3(iprot thrift.TProtocol) error {

	var _field *string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = &v
	}
	p.ParentID = _field
	return nil
}
func (p *CreateCommentRequest) ReadField4(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.Content = _field
	return nil
}

func (p *CreateCommentRequest) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("CreateCommentRequest"); err != nil {
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

func (p *CreateCommentRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("AuthorID", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.AuthorID); err != nil {
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

func (p *CreateCommentRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("ArticleID", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.ArticleID); err != nil {
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

func (p *CreateCommentRequest) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetParentID() {
		if err = oprot.WriteFieldBegin("ParentID", thrift.STRING, 3); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.ParentID); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *CreateCommentRequest) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("Content", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Content); err != nil {
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

func (p *CreateCommentRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateCommentRequest(%+v)", *p)

}

func (p *CreateCommentRequest) DeepEqual(ano *CreateCommentRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.AuthorID) {
		return false
	}
	if !p.Field2DeepEqual(ano.ArticleID) {
		return false
	}
	if !p.Field3DeepEqual(ano.ParentID) {
		return false
	}
	if !p.Field4DeepEqual(ano.Content) {
		return false
	}
	return true
}

func (p *CreateCommentRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.AuthorID, src) != 0 {
		return false
	}
	return true
}
func (p *CreateCommentRequest) Field2DeepEqual(src string) bool {

	if strings.Compare(p.ArticleID, src) != 0 {
		return false
	}
	return true
}
func (p *CreateCommentRequest) Field3DeepEqual(src *string) bool {

	if p.ParentID == src {
		return true
	} else if p.ParentID == nil || src == nil {
		return false
	}
	if strings.Compare(*p.ParentID, *src) != 0 {
		return false
	}
	return true
}
func (p *CreateCommentRequest) Field4DeepEqual(src string) bool {

	if strings.Compare(p.Content, src) != 0 {
		return false
	}
	return true
}

type CreateCommentResponse struct {
	ID string `thrift:"ID,1,required" frugal:"1,required,string" json:"ID"`
}

func NewCreateCommentResponse() *CreateCommentResponse {
	return &CreateCommentResponse{}
}

func (p *CreateCommentResponse) InitDefault() {
}

func (p *CreateCommentResponse) GetID() (v string) {
	return p.ID
}
func (p *CreateCommentResponse) SetID(val string) {
	p.ID = val
}

var fieldIDToName_CreateCommentResponse = map[int16]string{
	1: "ID",
}

func (p *CreateCommentResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetID bool = false

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
				issetID = true
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

	if !issetID {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CreateCommentResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_CreateCommentResponse[fieldId]))
}

func (p *CreateCommentResponse) ReadField1(iprot thrift.TProtocol) error {

	var _field string
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		_field = v
	}
	p.ID = _field
	return nil
}

func (p *CreateCommentResponse) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("CreateCommentResponse"); err != nil {
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

func (p *CreateCommentResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("ID", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.ID); err != nil {
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

func (p *CreateCommentResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateCommentResponse(%+v)", *p)

}

func (p *CreateCommentResponse) DeepEqual(ano *CreateCommentResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.ID) {
		return false
	}
	return true
}

func (p *CreateCommentResponse) Field1DeepEqual(src string) bool {

	if strings.Compare(p.ID, src) != 0 {
		return false
	}
	return true
}

type CommonService interface {
	CreateComment(ctx context.Context, req *CreateCommentRequest) (r *CreateCommentResponse, err error)
}

type CommonServiceCreateCommentArgs struct {
	Req *CreateCommentRequest `thrift:"req,1" frugal:"1,default,CreateCommentRequest" json:"req"`
}

func NewCommonServiceCreateCommentArgs() *CommonServiceCreateCommentArgs {
	return &CommonServiceCreateCommentArgs{}
}

func (p *CommonServiceCreateCommentArgs) InitDefault() {
}

var CommonServiceCreateCommentArgs_Req_DEFAULT *CreateCommentRequest

func (p *CommonServiceCreateCommentArgs) GetReq() (v *CreateCommentRequest) {
	if !p.IsSetReq() {
		return CommonServiceCreateCommentArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *CommonServiceCreateCommentArgs) SetReq(val *CreateCommentRequest) {
	p.Req = val
}

var fieldIDToName_CommonServiceCreateCommentArgs = map[int16]string{
	1: "req",
}

func (p *CommonServiceCreateCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommonServiceCreateCommentArgs) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CommonServiceCreateCommentArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *CommonServiceCreateCommentArgs) ReadField1(iprot thrift.TProtocol) error {
	_field := NewCreateCommentRequest()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Req = _field
	return nil
}

func (p *CommonServiceCreateCommentArgs) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("CreateComment_args"); err != nil {
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

func (p *CommonServiceCreateCommentArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *CommonServiceCreateCommentArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommonServiceCreateCommentArgs(%+v)", *p)

}

func (p *CommonServiceCreateCommentArgs) DeepEqual(ano *CommonServiceCreateCommentArgs) bool {
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

func (p *CommonServiceCreateCommentArgs) Field1DeepEqual(src *CreateCommentRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type CommonServiceCreateCommentResult struct {
	Success *CreateCommentResponse `thrift:"success,0,optional" frugal:"0,optional,CreateCommentResponse" json:"success,omitempty"`
}

func NewCommonServiceCreateCommentResult() *CommonServiceCreateCommentResult {
	return &CommonServiceCreateCommentResult{}
}

func (p *CommonServiceCreateCommentResult) InitDefault() {
}

var CommonServiceCreateCommentResult_Success_DEFAULT *CreateCommentResponse

func (p *CommonServiceCreateCommentResult) GetSuccess() (v *CreateCommentResponse) {
	if !p.IsSetSuccess() {
		return CommonServiceCreateCommentResult_Success_DEFAULT
	}
	return p.Success
}
func (p *CommonServiceCreateCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*CreateCommentResponse)
}

var fieldIDToName_CommonServiceCreateCommentResult = map[int16]string{
	0: "success",
}

func (p *CommonServiceCreateCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommonServiceCreateCommentResult) Read(iprot thrift.TProtocol) (err error) {

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
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_CommonServiceCreateCommentResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *CommonServiceCreateCommentResult) ReadField0(iprot thrift.TProtocol) error {
	_field := NewCreateCommentResponse()
	if err := _field.Read(iprot); err != nil {
		return err
	}
	p.Success = _field
	return nil
}

func (p *CommonServiceCreateCommentResult) Write(oprot thrift.TProtocol) (err error) {

	var fieldId int16
	if err = oprot.WriteStructBegin("CreateComment_result"); err != nil {
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

func (p *CommonServiceCreateCommentResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *CommonServiceCreateCommentResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CommonServiceCreateCommentResult(%+v)", *p)

}

func (p *CommonServiceCreateCommentResult) DeepEqual(ano *CommonServiceCreateCommentResult) bool {
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

func (p *CommonServiceCreateCommentResult) Field0DeepEqual(src *CreateCommentResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
