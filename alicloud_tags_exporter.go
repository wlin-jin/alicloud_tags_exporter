package main

import (
	"alicloud_tags_exporter/collector"
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 初始化，管理promethues tag
type TagManager struct {
	//用来做唯一标识，获取对应资源下的tag
	ResourceID string
	Region     string
	ResType    string
	// promethues变量
	TagCountDesc *prometheus.Desc
}

/**
 * 接口：Describe
 * 功能：传递结构体中的指标描述符到channel
 */
func (c *TagManager) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.TagCountDesc
}

/**
 * 接口：Collect
 * 功能：抓取最新的数据，传递给channel
 */
func (c *TagManager) Collect(ch chan<- prometheus.Metric) {
	//根据资源id获取对应的tagkey，tagvalue，返回类型[][]string
	metricValues := getTagValues(c.ResourceID, c.Region, c.ResType)
	//这里定义成了gaugevalue类型，值为1
	for _, metricValue := range metricValues {
		ch <- prometheus.MustNewConstMetric(
			c.TagCountDesc,
			prometheus.GaugeValue,
			1,
			metricValue...,
		)
	}
}

/**
 * 接口：NewTagManager
 * 功能：创建promethues监控指标，
 * 命名规范：指标描述 ecs_alicloud_tags 2个常量label：resourceID,region variableLabels: tagKey,tagValue
 */
func NewTagManager(resourceID string, namespace string, region string, resType string) *TagManager {
	return &TagManager{
		ResourceID: resourceID,
		Region:     region,
		ResType:    resType,
		TagCountDesc: prometheus.NewDesc(
			namespace,
			"descripe of "+namespace,
			[]string{"tagKey", "tagValue"},
			prometheus.Labels{"resourceID": resourceID, "region": region},
		),
	}
}

/**
函数：getTagValue
功能：获取不同类型资源下的tag 资源
*/
func getTagValues(resourceID string, region string, resType string) [][]string {
	result := make([][]string, 50)
	switch resType {
	case "ecs":
		result = collector.GetEcsTagValue(resourceID, region)
	case "rds":
		result = collector.GetRdsTagValue(resourceID, region)
	}
	return result
}

/**
函数：getResourceIDs
功能：获取不同类型资源下的resouceID 资源，用来注册metric
*/
func getResourceIDs(region string, resType string) []string {
	result := make([]string, 50)
	switch resType {
	case "ecs":
		result = collector.GetEcsResourceID(region)
	case "rds":
		result = collector.GetRdsResourceID(region)
	}
	return result
}

func main() {
	// 命令行参数
	listenAddr := flag.String("listen-port", "19001", "An port to listen on for web interface .")
	region := flag.String("region", "cn-shanghai", "alicloud region id")
	flag.Parse()
	//定义获取tag的资源类型
	ResTypeList := []string{"ecs", "rds"}
	reg := prometheus.NewPedanticRegistry()
	//获取ecs资源id，用于注册指标
	for _, resType := range ResTypeList {
		ResourceIDs := getResourceIDs(*region, resType)
		log.Printf("%s资源列表如下:", resType)
		log.Println(ResourceIDs)
		for _, resourceID := range ResourceIDs {
			worker := NewTagManager(resourceID, resType+"_alicloud_tags", *region, resType)
			reg.MustRegister(worker)
		}
	}
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Printf("Starting Server at http://localhost:%s%s", *listenAddr, "/metrics")
	log.Fatal(http.ListenAndServe(":"+*listenAddr, nil))
}
