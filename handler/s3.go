package handler

import (
	"context"

	"github.com/crunchyroll/s3-sidecar/logging"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Handler - contains the specific to store data into kms
type S3Handler struct {
	bucket   *string
	context  context.Context
	logger   logging.Logger
	s3Client *s3.S3
}

// NewS3Handler - returns a S3 handler for a given bucket and region
func NewS3Handler(region, bucket string) *S3Handler {
	// ctx := context.Background()
	sess := session.Must(session.NewSession(&aws.Config{Region: &region}))
	svc := s3.New(sess, &aws.Config{Region: &region})

	return &S3Handler{s3Client: svc, bucket: &bucket}
}

// GetItemContentWithCustomInput - supports byte-range request and other various customization
func (sh *S3Handler) GetItemContentWithCustomInput(key string, custom CustomInput) (error, string) {
	result, err := sh.s3Client.GetObject(&s3.GetObjectInput{
		Bucket:                     sh.bucket,
		Key:                        &key,
		IfMatch:                    custom.IfMatch,
		IfModifiedSince:            custom.IfModifiedSince,
		IfNoneMatch:                custom.IfNoneMatch,
		IfUnmodifiedSince:          custom.IfUnmodifiedSince,
		PartNumber:                 custom.PartNumber,
		Range:                      custom.Range,
		ResponseCacheControl:       custom.ResponseCacheControl,
		ResponseContentDisposition: custom.ResponseContentDisposition,
		ResponseContentEncoding:    custom.ResponseContentEncoding,
		ResponseContentLanguage:    custom.ResponseContentLanguage,
		ResponseContentType:        custom.ResponseContentType,
		ResponseExpires:            custom.ResponseExpires,
		VersionId:                  custom.VersionId,
	})

	// TODO: check for connection close
	if err != nil {
		// Cast err to awserr.Error to handle specific error codes.
		aerr, ok := err.(awserr.Error)
		if ok && aerr.Code() == s3.ErrCodeNoSuchKey {
			// Specific error code handling
		}
		return err, ""
	}
	return nil, result.GoString()
}

// GetItemContent - returns the entire content of an item from the input bucket
func (sh *S3Handler) GetItemContent(key string) (error, string) {
	return sh.GetItemContentWithCustomInput(key, CustomInput{})
}

// GetItemContentWithContext - ** this is NOT READY
func (sh *S3Handler) GetItemContentWithContext(key string) (error, string) {
	/*
		result, err := svc.GetObjectWithContext(ctx, &s3.GetObjectInput{
			Bucket: sh.bucket,
			Key:    key,
		})

		// TODO: check for connection close
		if err != nil {
			// Cast err to awserr.Error to handle specific error codes.
			aerr, ok := err.(awserr.Error)
			if ok && aerr.Code() == s3.ErrCodeNoSuchKey {
				// Specific error code handling
			}
			fmt.Println("error:", err)
			return err, ""
		}
	*/
	return nil, ""
}
