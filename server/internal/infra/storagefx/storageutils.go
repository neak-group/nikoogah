package storagefx

import (
	// 	"context"
	"fmt"
	// 	"time"
	// 	"github.com/neak-group/nikoogah/utils/uuid"
)

// func GetUploadLink(ctx context.Context, userID string) (preSignedLink string, imageKey string, err error) {

// 	randomID := uuid.New().String()

// 	uploadKey := fmt.Sprintf("%s%d", randomID, time.Now().Unix())
// 	uploadPath := PrepareUploadPath(userID, uploadKey)

// 	expiry := time.Second * 2 * 60 * 60 // 2 hours
// 	preSignedLink, err = GetUploadSignedURL(ctx, UploadBucket, uploadPath, expiry)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	return preSignedLink, uploadKey, nil
// }

// func PrepareUploadPath(userID string, uploadKey string) string {
// 	return fmt.Sprintf("upload/%s/%s", userID, uploadKey)
// }

func PrepareUploadPath(userID string, uploadKey string) string {
	return fmt.Sprintf("upload/%s/%s", userID, uploadKey)
}
