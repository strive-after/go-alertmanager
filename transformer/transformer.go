package transformer

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	model "webhook/module"
)

// TransformToMarkdown transform alertmanager notification to dingtalk markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.DingTalkMarkdown, err error) {
	var Alarm_level string
	switch notification.CommonLabels["severity"] {
	case "critical":
		Alarm_level = "严重"
	case "warning":
		Alarm_level = "警报"
	case "info":
		Alarm_level = "信息"
	}
	//groupKey := notification.GroupLabels
	status := notification.Status

	//annotations := notification.CommonAnnotations

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf(" **当前状态:**%s \n\n",status))
	buffer.WriteString(fmt.Sprintf("**出现的问题:**%s \n\n",notification.CommonLabels["alertname"]))
	buffer.WriteString(fmt.Sprintf("**集群名称:**%s \n\n",notification.CommonLabels["cluster_name"]))
	buffer.WriteString(fmt.Sprintf("**指标收集方式:** %s \n\n",notification.CommonLabels["alert_type"]))
	buffer.WriteString(fmt.Sprintf("**报警级别:**%s \n\n",Alarm_level))

 	for _, alert := range notification.Alerts {
		buffer.WriteString(fmt.Sprint(strings.Repeat("-",10)))
		buffer.WriteString(fmt.Sprintf("\n\n"))
		buffer.WriteString(fmt.Sprintf("**警报线:**  %v \n\n",alert.Labels["threshold_value"]))

		//buffer.WriteString(fmt.Sprintf("\n\n **属性:**   %s \n\n",alert.Labels["instance"]))
		//针对node报警会有node信息
		//针对pod报警会有pod信息
		//如果有容器这个标签那么根据容器的标签写入  如果没有按照node的标签来写入
		if alert.Labels["container"] != "" {
			buffer.WriteString(fmt.Sprintf("**容器:**   %s \n\n",alert.Labels["container"]))
			buffer.WriteString(fmt.Sprintf("**容器名称:** %s \n\n",alert.Labels["pod_name"]))
			buffer.WriteString(fmt.Sprintf("**命名空间:**   %s \n\n",alert.Labels["namespace"]))
			buffer.WriteString(fmt.Sprintf("**容器所在节点:**   %s \n\n",alert.Labels["node"]))
		}else {
			buffer.WriteString(fmt.Sprintf("**节点:**   %s \n\n",alert.Labels["host_ip"]))
		}



		//共用
		buffer.WriteString(fmt.Sprintf(" **开始时间:**%s\n\n", alert.StartsAt.Add(8*time.Hour).Format("2006-01-02 15:04:05")))
		buffer.WriteString(fmt.Sprintf("**当前值:**   %s \n\n",alert.Annotations["current_value"]))
	}


	markdown = &model.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Title: fmt.Sprintf("####alter名称 %s\n\n", notification.CommonLabels["alertname"]),
			Text:  buffer.String(),
		},
		At: &model.At{
			IsAtAll: false,
		},
	}
	return
}