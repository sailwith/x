package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"hash"
	"io"
	"log"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var expireTime int64 = 3600

type PostPolicy struct {
	Expiration string          `json:"expiration"`
	Conditions [][]interface{} `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string
	Host        string
	Expire      int64
	Signature   string
	Policy      string
	Directory   string
	Callback    string
}

type PutInfo struct {
	SignedURL string
}

type Config struct {
	Endpoint        string
	BucketName      string
	AccessKeyID     string
	AccessKeySecret string
}

type OSS struct {
	endpoint        string
	bucketName      string
	host            string
	accessKeyID     string
	accessKeySecret string
	bucket          *oss.Bucket
}

func New(c Config) *OSS {
	host := "https://" + c.BucketName + "." + c.Endpoint
	client, err := oss.New(c.Endpoint, c.AccessKeyID, c.AccessKeySecret)
	if err != nil {
		log.Fatal(err)
	}

	bucket, err := client.Bucket(c.BucketName)
	if err != nil {
		log.Fatal(err)
	}

	return &OSS{
		endpoint:        c.Endpoint,
		bucketName:      c.BucketName,
		host:            host,
		accessKeyID:     c.AccessKeyID,
		accessKeySecret: c.AccessKeySecret,
		bucket:          bucket,
	}
}

func (o *OSS) PutInfo(key, contentType, callback, callbackVar string) (*PutInfo, error) {
	opts := []oss.Option{oss.ForbidOverWrite(true)}
	if contentType != "" {
		opts = append(opts, oss.ContentType(contentType))
	}
	if callback != "" {
		opts = append(opts, oss.Callback(callback))
	}
	if callbackVar != "" {
		opts = append(opts, oss.CallbackVar(callbackVar))
	}
	signedURL, err := o.bucket.SignURL(key, oss.HTTPPut, expireTime, opts...)
	if err != nil {
		return nil, err
	}

	var putInfo PutInfo
	putInfo.SignedURL = signedURL
	return &putInfo, nil
}

func (o *OSS) PostInfo(dir string) (*PolicyToken, error) {
	now := time.Now().Unix()
	expireEnd := now + expireTime
	tokenExpire := gmtISO8601(expireEnd)

	//create post policy json
	var policy PostPolicy
	policy.Expiration = tokenExpire
	condition1 := []interface{}{"starts-with", "$key", dir}
	condition2 := []interface{}{"content-length-range", 1, 10485760}
	policy.Conditions = append(policy.Conditions, condition1)
	policy.Conditions = append(policy.Conditions, condition2)

	//calucate signature
	result, err := json.Marshal(policy)
	if err != nil {
		return nil, err
	}
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(o.accessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var policyToken PolicyToken
	policyToken.AccessKeyId = o.accessKeyID
	policyToken.Host = o.host
	policyToken.Expire = expireEnd
	policyToken.Signature = signedStr
	policyToken.Directory = dir
	policyToken.Policy = debyte
	return &policyToken, nil
}

func gmtISO8601(expireEnd int64) string {
	return time.Unix(expireEnd, 0).UTC().Format("2006-01-02T15:04:05Z")
}
