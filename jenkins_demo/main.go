package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/bndr/gojenkins"
)

const ConfigString = `<?xml version='1.1' encoding='UTF-8'?>
<flow-definition plugin="workflow-job@1292.v27d8cc3e2602">
<actions>
  <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@2.2131.vb_9788088fdb_5"/>
  <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@2.2131.vb_9788088fdb_5">
	<jobProperties/>
	<triggers/>
	<parameters/>
	<options/>
  </org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction>
</actions>
<description></description>
<keepDependencies>false</keepDependencies>
<properties/>
<definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@3659.v582dc37621d8">
  <script>node{
	  stage(&apos;Loading&apos;)
	  def rootDir = pwd()
	  println(rootDir)
	  def pipeline = load &apos;pipeline.groovy&apos;
	  pipeline(&apos;${app_name}&apos;,&apos;${app_group}&apos;)
}</script>
  <sandbox>true</sandbox>
</definition>
<triggers/>
<disabled>false</disabled>
</flow-definition>`

func parseString(name, group string) string {
	str := strings.Replace(ConfigString, "${app_name}", name, -1)
	str = strings.Replace(str, "${app_group}", group, -1)
	return str
}
func main() {
	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, "http://172.168.1.10:8088/", "admin", "123456")
	_, err := jenkins.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("jenkins 连接成功")
	nodes, _ := jenkins.GetAllNodes(ctx)
	for _, node := range nodes {
		node.Poll(ctx)
		if ok, _ := node.IsOnline(ctx); ok {
			nodeName := node.GetName()
			fmt.Printf("nodeName:%v\n", nodeName)
		}
	}
	job, _ := jenkins.GetJob(ctx, "ota")
	job.Poll(ctx)
	fmt.Println(job.GetName(), job.Base)
	build, _ := job.GetLastBuild(ctx)
	fmt.Println(build.GetBuildNumber())

	// fmt.Println("新建任务...")
	// config := parseString("ota", "company")
	// _, err = jenkins.CreateJob(ctx, config, "ota")
	// if err != nil {
	// 	fmt.Printf("err:%v\n", err)
	// 	return
	// }
	// fmt.Println("任务创建成功...")

	// jobs, _ := jenkins.GetAllJobs(ctx)
	// for _, job := range jobs {
	// 	fmt.Println(job.GetName(), job.Base)

	// }

	task, err := jenkins.GetQueue(ctx)
	if err != nil {
		fmt.Println(err)
	}
	task.Poll(ctx)
	// tasks := task.GetTasksForJob("ota")
	tasks := task.Tasks()
	fmt.Println(task, &tasks)
	for _, v := range tasks {
		fmt.Println(v.Raw.Task.URL, v.GetWhy(), v.Base)

		// job, _ := v.GetJob(ctx)
		// fmt.Println(job.GetName())
		_, err := v.Cancel(ctx)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(v.Raw.Task.URL, v.GetWhy(), v.Base)

		// if flag {
		// 	fmt.Println("success")
		// 	job, _ := v.GetJob(ctx)
		// 	fmt.Println(job.GetName())
		// }
	}
}
