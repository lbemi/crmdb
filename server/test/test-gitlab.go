package main

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"log"
	"time"
)

func main() {
	git, err := gitlab.NewClient("sQn-9rr8tbeGj-sJy6FH", gitlab.WithBaseURL("http://git.yunling.org"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
	for _, user := range users {
		fmt.Println(user.Name, user.Email)
	}
	projects, _, err := git.Projects.ListProjects(&gitlab.ListProjectsOptions{Search: gitlab.Ptr("yunling")})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	for _, project := range projects {
		fmt.Println(project.PathWithNamespace, project.Name, project.ID)
	}
	opt := &gitlab.ListProjectPipelinesOptions{
		Scope:         gitlab.Ptr("branches"),
		Status:        gitlab.Ptr(gitlab.Running),
		Ref:           gitlab.Ptr("master"),
		YamlErrors:    gitlab.Ptr(true),
		Name:          gitlab.Ptr("name"),
		Username:      gitlab.Ptr("username"),
		UpdatedAfter:  gitlab.Ptr(time.Now().Add(-24 * 365 * time.Hour)),
		UpdatedBefore: gitlab.Ptr(time.Now().Add(-7 * 24 * time.Hour)),
		OrderBy:       gitlab.Ptr("status"),
		Sort:          gitlab.Ptr("asc"),
	}
	pipelines, _, err := git.Pipelines.ListProjectPipelines(177, opt)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	for _, pipeline := range pipelines {
		fmt.Println(pipeline.ID, pipeline.Status, pipeline.Ref, pipeline.Source)
	}

}
