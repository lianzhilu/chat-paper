package generator

import "github.com/rs/xid"

func GenerateUserID() string {
	id := xid.New()
	return id.String()
}
