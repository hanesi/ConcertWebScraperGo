package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Initialize variables and an AWS session with configured credentials
	var key string
	// var bandMap map[string]string

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		panic(err)
	}

	// Define client and bucket name (Use environment variables)
	s3Client := s3.New(sess)
	bucket := "ian-test-bucket-go-python"

	// List objects in bucket, retrieve the key from the returned result
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(bucket),
		MaxKeys: aws.Int64(3),
	}
	result, err := s3Client.ListObjectsV2(input)
	if err != nil {
		fmt.Println(err)
	}

	for i := range result.Contents {
		if strings.Contains(*result.Contents[i].Key, "Bands") == true {
			key = *result.Contents[i].Key
		}
	}

	// Using the key, get the object from the bucket
	obj, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Println(err)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, obj.Body); err != nil {
		fmt.Println(err)
	}
	test, _ := json.Marshal(buf.Bytes())
	fmt.Println(test)
}
