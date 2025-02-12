package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

// 定义全局变量
var (
	region     string = "cn-shenzhen" // 存储区域
	bucketName string = "mi-oss-demo" // 存储空间名称
)

func main() {
	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewEnvironmentVariableCredentialsProvider()).
		WithRegion(region)

	// 创建OSS客户端
	client := oss.NewClient(cfg)

	// 创建列出存储空间的请求
	ProcessObjectRequest(client)
	// ListBucketsRequest(client)

}

func ProcessObjectRequest(client *oss.Client) {

	destObjName := "a1/3.jpg"
	process := fmt.Sprintf("image/resize,w_100|sys/saveas,o_%v", base64.URLEncoding.EncodeToString([]byte(destObjName)))
	request := &oss.ProcessObjectRequest{
		Bucket:  oss.Ptr(bucketName),
		Key:     oss.Ptr(destObjName),
		Process: oss.Ptr(process),
	}

	fmt.Println(request)
	xx, err := client.ProcessObject(context.TODO(), request)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())

		return
	}
	fmt.Println(xx)

}

func ListBucketsRequest(client *oss.Client) {
	// 创建列出对象的请求
	request := &oss.ListObjectsV2Request{
		Bucket:    oss.Ptr(bucketName),
		Prefix:    oss.Ptr("a1/"),
		Delimiter: oss.Ptr("/"),
	}
	// 创建分页器
	p := client.NewListObjectsV2Paginator(request)
	// 初始化页码计数器
	var i int
	log.Println("Objects:")
	// 遍历分页器中的每一页
	for p.HasNext() {
		i++
		// 获取下一页的数据
		page, err := p.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("failed to get page %v, %v", i, err)
		}

		// 打印该页中的每个对象的信息
		for _, obj := range page.Contents {
			log.Printf("Object:%v, %v, %v\n", oss.ToString(obj.Key), obj.Size, oss.ToTime(obj.LastModified))
		}
	}
}
