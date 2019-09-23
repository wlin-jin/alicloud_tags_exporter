package collector

import (
	"fmt"
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

/**
函数：GetEcsTagValues
功能：根据资源id获取ecs tags value
*/
func GetEcsTagValue(resourceID string, region string) [][]string {
	config := getConfig()
	ecsClient, err := ecs.NewClientWithAccessKey(
		region,
		config.Accesskey,
		config.Accesssecrt)
	if err != nil {
		// 异常处理
		fmt.Println(err)
	}
	// 创建API请求并设置参数
	request := ecs.CreateListTagResourcesRequest()
	request.ResourceType = "instance"
	request.ResourceId = &[]string{resourceID}
	// 发起请求并处理异常
	response, err := ecsClient.ListTagResources(request)
	if err != nil {
		// 异常处理
		log.Printf("function:%s 请求%s异常:%s", "GetEcsTagValue", "ListTagResources", err)
	}
	tagValues := [][]string{}
	for _, v := range response.TagResources.TagResource {
		tagValue := []string{}
		tagValue = append(tagValue, v.TagKey)
		tagValue = append(tagValue, v.TagValue)
		tagValues = append(tagValues, tagValue)
	}
	//如果没有绑定标签，返回默认值，解决promethues panic错误
	if len(tagValues) == 0 {
		tagValues = [][]string{{"undefined", "undefined"}}
	}
	return tagValues
}

/**
函数：GetEcsResourceID
功能：获取ecs资源id列表
*/
func GetEcsResourceID(region string) []string {
	config := getConfig()
	client, err := ecs.NewClientWithAccessKey(region, config.Accesskey, config.Accesssecrt)
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"
	response, err := client.DescribeInstances(request)
	if err != nil {
		log.Printf(err.Error())
	}
	resourceID := []string{}
	for _, v := range response.Instances.Instance {
		resourceID = append(resourceID, v.InstanceId)
	}
	return resourceID
}
