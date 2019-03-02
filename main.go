package main

import (
	"log"
	"os"
	"text/template"
	"fmt"
)


type Project struct {
	GithubOwner		string
	RepoName 		string
	LocalRepoPath 	string
}

func main() {
	// check for a gopath
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		log.Fatal("Please set a $GOPATH")
	}
	fmt.Println(goPath)

	// check for a github token
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == ""{
		log.Fatal("Please set a $GITHUB_TOKEN")
	}

	

	// Set the project
	x := Project{GithubOwner: "codingdiaz", RepoName: "testing1234"}
	x.LocalRepoPath = fmt.Sprintf("%s/src/github.com/%s/%s", goPath, x.GithubOwner, x.RepoName)
	fmt.Println(x.LocalRepoPath)
	err := os.Mkdir(x.LocalRepoPath,0755)
	if err != nil {log.Fatal(err)}
	f, err := os.Create(fmt.Sprintf("%s/%s",x.LocalRepoPath,"Dockerfile"))
	if err != nil {log.Fatal(err)}
	tmpl, err := template.ParseFiles("templates/Dockerfile")
	if err != nil { panic(err) }
	err = tmpl.Execute(f, x)
	if err != nil { panic(err) }
	f.Close()
}