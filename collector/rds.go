package collector

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

/**
函数：GetRdsResourceID
功能：获取rds资源id列表
*/
func GetRdsResourceID(region string) []string {
	config := getConfig()
	client, err := rds.NewClientWithAccessKey(region, config.Accesskey, config.Accesssecrt)

	request := rds.CreateDescribeDBInstancesRequest()
	request.Scheme = "https"

	response, err := client.DescribeDBInstances(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	DBInstanceID := []string{}
	for _, item := range response.Items.DBInstance {
		DBInstanceID = append(DBInstanceID, item.DBInstanceId)
	}
	return DBInstanceID
}

/**
函数：GetRdsTagValue
功能：获取rds资源id列表
*/
func GetRdsTagValue(resourceID string, region string) [][]string {
	config := getConfig()
	client, err := rds.NewClientWithAccessKey(region, config.Accesskey, config.Accesssecrt)
	request := rds.CreateDescribeTagsRequest()
	request.Scheme = "https"
	request.DBInstanceId = resourceID
	response, err := client.DescribeTags(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	tagValues := make([][]string, 0)
	result := response.Items.TagInfos
	for _, item := range result {
		tagValue := []string{}
		tagValue = append(tagValue, item.TagKey)
		tagValue = append(tagValue, item.TagValue)
		tagValues = append(tagValues, tagValue)
	}
	//如果没有绑定标签，返回默认值，解决promethues panic错误
	if len(tagValues) == 0 {
		tagValues = [][]string{{"undefined", "undefined"}}
	}
	return tagValues
}
