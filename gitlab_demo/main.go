package main

import (
	"fmt"
	"log"
	"source/go-learning/gitlab_demo/database"

	"github.com/xanzy/go-gitlab"
)

func main() {
	var item = database.Item{}
	var items = []database.Item{}
	token := "glpat-57WwyFMeYAsszMzyzjsn"
	url := "http://172.168.1.10/api/v4"
	git, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
	// if err != nil {
	// 	log.Fatalf("Failed to get users err: %v", err)
	// }
	// // fmt.Println(users)
	// for _, v := range users {
	// 	fmt.Println(v.Name, v.Username, v.ID)
	// }
	var gid int
	groupOption := &gitlab.ListGroupsOptions{Search: gitlab.String("server")}
	groups, _, err := git.Groups.ListGroups(groupOption)
	if err != nil {
		log.Fatalf("Failed to get groups err: %v", err)
	}
	for _, group := range groups {
		fmt.Println(group.ID, group.Name)
		gid = group.ID
	}
	opt := &gitlab.ListGroupProjectsOptions{ListOptions: gitlab.ListOptions{Page: 1, PerPage: 50}}
	projects, _, err := git.Groups.ListGroupProjects(gid, opt)
	// listOption := &gitlab.ListProjectsOptions{ListOptions: gitlab.ListOptions{Page: 1, PerPage: 50}}
	// // listOption.ListOptions.Page++
	// projects, _, err := git.Projects.ListProjects(listOption)
	if err != nil {
		log.Fatalf("Failed to get projects err: %v", err)
	}
	for _, v := range projects {
		fmt.Println(v.ID, v.Name, v.HTTPURLToRepo, v.SSHURLToRepo)
		item.CodeID = v.ID
		item.AppName = v.Name
		item.AppGroup = "server"
		item.AppType = "go"
		item.HTTPURLToRepo = v.HTTPURLToRepo
		item.SSHURLToRepo = v.SSHURLToRepo
		items = append(items, item)

	}
	fmt.Printf("####%#v\n", items)
	err = database.InitMysql()
	if err != nil {
		log.Fatalf("连接数据库错误: %v", err)
	}
	// for _, item := range items {
	// 	n, err := database.AddItem(item)
	// 	if err != nil {
	// 		log.Fatalf("插入数据错误: %v", err)
	// 	}
	// 	fmt.Printf("insert success,affected rows%v\n", n)
	// }

	b, _ := database.GetItemFromName()
	fmt.Println(b)
}
