package response

import (
	"github.com/lianzhilu/chat-paper/cp-core/kitex/kitex_gen/base"
)

type ChatPaperResponse struct {
	ResponseMetadata base.ResponseMetadata
	Result           interface{} `json:",omitempty"`
}

func ConstructSuccessResponse(result interface{}) ChatPaperResponse {
	return ChatPaperResponse{
		ResponseMetadata: base.ResponseMetadata{
			RequestID: "",
			Action:    "",
		},
		Result: result,
	}
}

func ConstructErrorResponse(code string, message string) ChatPaperResponse {
	return ChatPaperResponse{
		ResponseMetadata: base.ResponseMetadata{
			RequestID: "",
			Action:    "",
			Error:     &base.Error{Code: code, Message: message},
		},
	}
}
