package handler

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"golang.org/x/net/http2"
)

// snippet-start:[s3.go.customHttpClient_struct]
type S3ClientSettings struct {
	ConnKeepAlive    time.Duration
	Connect          time.Duration
	ExpectContinue   time.Duration
	IdleConn         time.Duration
	MaxAllIdleConns  int
	MaxHostIdleConns int
	ResponseHeader   time.Duration
	TLSHandshake     time.Duration
}

func NewS3Client(httpSettings S3ClientSettings) *http.Client {
	tr := &http.Transport{
		ResponseHeaderTimeout: httpSettings.ResponseHeader,
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: httpSettings.ConnKeepAlive,
			DualStack: true,
			Timeout:   httpSettings.Connect,
		}).DialContext,
		MaxIdleConns:          httpSettings.MaxAllIdleConns,
		IdleConnTimeout:       httpSettings.IdleConn,
		TLSHandshakeTimeout:   httpSettings.TLSHandshake,
		MaxIdleConnsPerHost:   httpSettings.MaxHostIdleConns,
		ExpectContinueTimeout: httpSettings.ExpectContinue,
	}

	// So client makes HTTP/2 requests
	http2.ConfigureTransport(tr)

	return &http.Client{
		Transport: tr,
	}
}

// client - to test out the S3Client
func client(bucket, itemKey, region string) {
	if bucket == "" || itemKey == "" {
		fmt.Println("You must supply the name of the bucket and item")
		fmt.Println("Usage: go run customHttpClient -b bucket-name -i item-name [-s] (show the bucket item as a string)")
		os.Exit(1)
	}

	fmt.Println("Getting item " + itemKey + " from bucket " + bucket + " in " + region)

	sess := session.Must(session.NewSession(&aws.Config{
		Region: &region,
		HTTPClient: NewS3Client(S3ClientSettings{
			Connect:          5 * time.Second,
			ExpectContinue:   1 * time.Second,
			IdleConn:         90 * time.Second,
			ConnKeepAlive:    30 * time.Second,
			MaxAllIdleConns:  100,
			MaxHostIdleConns: 10,
			ResponseHeader:   5 * time.Second,
			TLSHandshake:     5 * time.Second,
		}),
	}))

	client := s3.New(sess)

	obj, err := client.GetObject(&s3.GetObjectInput{Bucket: &bucket, Key: &itemKey})
	if err != nil {
		fmt.Println("Got error calling GetObject:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Convert body from IO.ReadCloser to string:
	buf := new(bytes.Buffer)
	buf.ReadFrom(obj.Body)
	newBytes := buf.String()
	s := string(newBytes)

	fmt.Println("Bucket item as string:")
	fmt.Println(s)
}
