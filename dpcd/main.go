package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

//https://open.feishu.cn/open-apis/bot/v2/hook/bf8bb912-bc2e-40ad-9533-fcb8068aa621
var (
	listenPort = flag.String("port", ":8090", "listen port")
	kubedir    = flag.String("kubedir", os.Getenv("HOME")+"/.kube", "kube config dir")
	// webhook地址
	url = "https://open.feishu.cn/open-apis/bot/v2/hook/f9c5d5ca-b83e-4bbb-af39-92dc7e868d4c"
)

type kubeClient struct {
	*exec.Cmd
}

func kubectl(args ...string) *kubeClient {
	return &kubeClient{
		&exec.Cmd{
			Path: "/usr/local/bin/kubectl",
			Args: append([]string{"kubectl"}, args...),
		},
	}
}

func (h *kubeClient) log() *kubeClient {
	h.Stdout = os.Stdout
	h.Stderr = os.Stderr
	return h
}

func (h *kubeClient) setEnv(kubeconf string) *kubeClient {
	h.Args = append(h.Args, "--kubeconfig", kubeconf)
	return h
}

// 执行任意Git命令的封装
func RunGitCommand(name string, arg ...string) (string, error) {
	gitpath := "/root/kube-conf-dev" // 从配置文件中获取当前git仓库的路径

	cmd := exec.Command(name, arg...)
	cmd.Dir = gitpath // 指定工作目录为git仓库目录
	//cmd.Stderr = os.Stderr
	msg, err := cmd.CombinedOutput() // 混合输出stdout+stderr
	cmd.Run()
	// 报错时 exit status 1
	return string(msg), err
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/status/:deploy", func(c *gin.Context) {
		d := c.Param("deploy")
		path, e := kubectl("-n", "test",
			"get",
			"deploy",
			d,
			"-o",
			"jsonpath={.spec.template.spec.containers[0].image}").CombinedOutput()
		if e != nil {
			fmt.Println(e.Error())
			return
		}

		c.String(http.StatusOK, string(path))

	})

	r.GET("/status", func(c *gin.Context) {
		result, e := kubectl("-n", "test",
			"get",
			"pod",
		).CombinedOutput()

		if e != nil {
			fmt.Println(e.Error())
			return
		}
		c.String(http.StatusOK, string(result))
		//fmt.Println(string(result))

	})

	r.POST("/:kubeConfig", func(c *gin.Context) {

		kubeConfig := c.Param("kubeConfig")

		var eventAliyum map[string]map[string]string
		c.BindJSON(&eventAliyum)

		repo := eventAliyum["repository"]
		tag := eventAliyum["push_data"]["tag"]
		name := eventAliyum["repository"]["name"] //stage_hello
		namespace := strings.Split(name, "_")[0]  //stage
		appname := strings.Split(name, "_")[1]    //hello
		fmt.Printf("tag:%s,name:%s,namespace:%s,appname:%s\n", tag, name, namespace, appname)
		go func(repo map[string]string, tag string) {
			kubeconf := *kubedir + "/" + kubeConfig
			kf := kubeconf + "-" + namespace

			if _, err := os.Stat(kf); !os.IsNotExist(err) {
				fmt.Println("INFO: current use kube config: " + kf)
				kubeconf = kf
			}
			//registry.cn-hangzhou.aliyuncs.com/unipal/stage_hello:be12e31b-117
			npath := "registry." + repo["region"] + ".aliyuncs.com/" + repo["namespace"] + "/" + repo["name"] + ":" + tag
			fmt.Println(npath)
			result, e := kubectl("-n", namespace,
				"set",
				"image",
				"deploy",
				appname,
				appname+"="+npath).setEnv(kubeconf).CombinedOutput()
			fmt.Println("INFO: "+"command is: \n", string(result))
			if e != nil {
				fmt.Println("ERROR: "+appname+" fail upgrade ", e)
				return
			}
			// if string(result) == "" {
			// 	fmt.Println("INFO: eid: " + appname + " upgrade complete (not change)")
			// 	return
			// }

			var msg string
			if string(result) == "" {
				if namespace == "dev" {
					//拉取代码
					_, err := RunGitCommand("git", "pull")
					if err != nil {
						fmt.Println("git pull error")
					}
					fmt.Println("INFO: " + "git pull complete")

					//更新配置文件
					configfile := "/root/kube-conf-dev/app-dev/" + appname + "/app_dev.yaml"
					_, e = kubectl("apply", "-f", configfile).setEnv(kubeconf).CombinedOutput()
					fmt.Println("INFO:" + appname + "upgrade config complete")
					if e != nil {
						fmt.Println("ERROR: "+appname+" config update fail ", e)
						msg = "ERROR: " + namespace + " " + appname + " config update fail"
						sendMsg(url, msg)
						return
					}
					//重启服务
					result, e = kubectl("-n", namespace, "rollout", "restart", "deploy", appname).setEnv(kubeconf).CombinedOutput()
					fmt.Println("INFO: "+"command is: ", string(result))
					if e != nil {
						fmt.Println("ERROR: "+appname+" fail restart deploy ", e)
						msg = "ERROR: " + appname + " fail restart deploy"
						sendMsg(url, msg)
						return
					}
					fmt.Println("INFO: eid: " + appname + " restart complete")
					msg = "INFO: " + appname + " restart complete"
					sendMsg(url, msg)
					return
				}
				if namespace == "test" {
					//重启服务
					result, e = kubectl("-n", namespace, "rollout", "restart", "deploy", appname).setEnv(kubeconf).CombinedOutput()
					fmt.Println("INFO: "+"command is: ", string(result))
					if e != nil {
						fmt.Println("ERROR: "+appname+" fail restart deploy ", e)
						return
					}
					fmt.Println("INFO: eid: "+appname+" upgrade complete ", string(result))
					msg = "INFO: " + namespace + "/" + appname + " upgrade complete"
					sendMsg(url, msg)
					return
				}
				fmt.Println("INFO: eid: " + appname + " upgrade complete (not change)")
				return
			}
			fmt.Println("INFO: eid: "+appname+" upgrade complete ", string(result))
			// msg = fmt.Sprintf(`INFO: name: %s`, appname)
			msg = "INFO: " + namespace + "/" + appname + " upgrade complete"
			sendMsg(url, msg)
		}(repo, tag)

	})

	return r
}

//飞书通知
func sendMsg(url, msg string) {
	// json
	contentType := "application/json"
	// data
	sendData := `{
		"msg_type": "text",
		"content": {
			"text": " ` + msg + `"
		}
	 }`
	// request
	result, err := http.Post(url, contentType, strings.NewReader(sendData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer result.Body.Close()

}

func main() {

	flag.Parse()
	fmt.Println("\n Info:\n  \tlisten: " + *listenPort + "\n\tkubedir: " + *kubedir)
	gin.SetMode(gin.ReleaseMode)
	r := setupRouter()

	r.Run(*listenPort)

}

//linux平台编译
//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dpcd
//http://120.55.54.179:18090/config
