package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// CONF - the config path location
const CONF = "github.com/crunchyroll/s3-sidecar/config/config.yml"

func main() {
	api := App{}

	confPath := flag.String("config", "", "config file to use")
	flag.Parse()

	/*
		confPath, err := filepath.Abs("../" + CONF)
		if err != nil {
			fmt.Printf("invalid config file location. error: %v. \n", err)
			os.Exit(1)
		}
	*/
	api.Initialize(*confPath)
	api.Run()
}

/*
	support for 405 if not head or get
	support for 403 if not 127.0.0.1
	support for json formatted logger
	support for range
	support for error interface, a wrapper around aws error type
	handle timeout with the sdk
*/

func GetObject(key string) (error, string) {
	ctx := context.Background()
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(endpoints.UsWest2RegionID)}))
	svc := s3.New(sess, &aws.Config{Region: aws.String(endpoints.UsWest2RegionID)})

	result, err := svc.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String("ellation-cx-proto0-vod-media/"),
		Key:    aws.String(key),
	})

	if err != nil {
		// Cast err to awserr.Error to handle specific error codes.
		aerr, ok := err.(awserr.Error)
		if ok && aerr.Code() == s3.ErrCodeNoSuchKey {
			// Specific error code handling
		}
		fmt.Println("error:", err)
		return err, ""
	}

	// Make sure to close the body when done with it for S3 GetObject APIs or
	// will leak connections.
	defer result.Body.Close()

	fmt.Println("Object Size:", aws.Int64Value(result.ContentLength))
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	newStr := buf.String()
	fmt.Println("Object body:", newStr)
	return nil, newStr
}
