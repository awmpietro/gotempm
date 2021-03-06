package globalUtils

import (
	"context"
	"encoding/base64"
	"fmt"
	//"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/micro/v3/service/context/metadata"
	"log"
	"strconv"
)

// AuthUtils contains methods to simplify authentication
type AuthUtils struct{}

// GetCurrentUserFromContext adds user  to the context during authentication. this function extracts it so
// that it can be used sending audit records to the broker
func (a *AuthUtils) GetCurrentUserFromContext(ctx context.Context) (int64, error) {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return 0, fmt.Errorf("unable to get user from metadata")
	}
	userId, err := euidToId(meta["Userid"])
	//userId, err := strconv.ParseInt(meta["Userid"], 10, 64)
	if err != nil {
		return 0, err
	}
	log.Printf("userid from metadata: %d\n", userId)
	return userId, nil
}

// euidToId converts  []byte based id to a regular int64 id
func euidToId(euid string) (int64, error) {
	decoded, err := base64.StdEncoding.DecodeString(euid)
	if err != nil {
		log.Printf("Unable to decode eUid. Error: %v\n", err)
		return 0, err
	}
	id, err := strconv.ParseInt(string(decoded), 10, 64)
	if err != nil {
		log.Printf("Unableto parse eUid. Error: %v\n", err)
		return 0, err
	}
	return id, nil
}
