* `collector`实现一个采集器，用于采集指标数据
* alicloud_tags_exporter.go 程序主逻辑
* config.json 存储阿里云秘钥信息
* 编译依赖
```
go get -u github.com/aliyun/alibaba-cloud-sdk-go/sdk
go get github.com/prometheus/client_golang/prometheus/promhttp
go get -u github.com/spf13/viper
```
执行`go build`编译运行，然后访问`http://localhost:19001/metrics`就可以看到采集到的指标数据。
```
* 使用帮助
./alicloud_tags_exporter -h
Usage of ./alicloud_tags_exporter:
  -listen-port string
        An port to listen on for web interface . (default "19001")
  -region string
        alicloud region id (default "cn-shanghai")
```
效果如下：
```
# HELP ecs_alicloud_tags descripe of ecs_alicloud_tags
# TYPE ecs_alicloud_tags gauge
ecs_alicloud_tags{region="cn-shanghai",resourceID="i-xxx",tagKey="Env",tagValue="development"} 1
ecs_alicloud_tags{region="cn-shanghai",resourceID="i-xxx",tagKey="Team",tagValue="cloud"} 1
ecs_alicloud_tags{region="cn-shanghai",resourceID="i-xxx",tagKey="undefined",tagValue="undefined"} 1
ecs_alicloud_tags{region="cn-shanghai",resourceID="i-xxx",tagKey="ros-aliyun-created",tagValue="k8s_nodes_config_stack1"} 1
ecs_alicloud_tags{region="cn-shanghai",resourceID="i-xxx",tagKey="undefined",tagValue="undefined"} 1
ecs_alicloud_tags{region="cn-shanghai",resourceID="i-xxxx",tagKey="ros-aliyun-created",tagValue="k8s_nodes_config_stack_7"} 1
ecs_alicloud_tags{region="cn-shanghai",resourceID="i-xxx",tagKey="ros-aliyun-created",tagValue="k8s_nodes_config_stack_2"} 1
# HELP rds_alicloud_tags descripe of rds_alicloud_tags
# TYPE rds_alicloud_tags gauge
rds_alicloud_tags{region="cn-shanghai",resourceID="pgm-xxx",tagKey="app",tagValue="mmmm"} 1
rds_alicloud_tags{region="cn-shanghai",resourceID="pgm-xxx",tagKey="env",tagValue="development"} 1
rds_alicloud_tags{region="cn-shanghai",resourceID="pgm-xxx",tagKey="name",tagValue="mmmm-pg"} 1
rds_alicloud_tags{region="cn-shanghai",resourceID="pgm-xxx",tagKey="team",tagValue="cloud"} 1
rds_alicloud_tags{region="cn-shanghai",resourceID="rm-xxx",tagKey="undefined",tagValue="undefined"} 1
rds_alicloud_tags{region="cn-shanghai",resourceID="rm-xxx",tagKey="undefined",tagValue="undefined"} 1
rds_alicloud_tags{region="cn-shanghai",resourceID="rm-xxxxx",tagKey="undefined",tagValue="undefined"} 1
rds_alicloud_tags{region="cn-shanghai",resourceID="rm-xxx",tagKey="undefined",tagValue="undefined"} 1
