package main

import (
	"log"
	"os"
	"text/template"
	"fmt"
	"io/ioutil"
)

type Project struct {
	GithubOwner			string
	RepoName 			string
	LocalRepoPath 		string
}

type LocalConfig struct {
	GoPath				string
	GithubToken			string
}

type ProjectFile struct {
	TemplateFilePath	string
	RepoFilePath 		string
	Render				bool
}

func main() {
	// Check for needed environment variables
	localConfig, err := envVarCheck()
	if err != nil {
		log.Fatal(err)
	}
	
	// Set the project
	x := Project{GithubOwner: "codingdiaz", RepoName: "testing1234"}
	// Set the repo path
	setRepoPath(&x, &localConfig)
	// Create the github repo (local folder for now)
	err = createGitRepo(&x, &localConfig)
	if err != nil {log.Fatal(err)}
	err = createFiles(&x, &localConfig)
	if err != nil {log.Fatal(err)}
}


func envVarCheck() (localConfig LocalConfig, err error) {
	// check for a gopath
	localConfig.GoPath = os.Getenv("GOPATH")
	if localConfig.GoPath == "" {
		return localConfig, fmt.Errorf("No GOPATH environment variable set")
	}

	// check for a github token
	localConfig.GithubToken = os.Getenv("GITHUB_TOKEN")
	if localConfig.GithubToken == "" {
		return localConfig, fmt.Errorf("No GITHUB_TOKEN environment variable set")
	}

	return
}

func setRepoPath(project *Project, localConfig *LocalConfig) () {
	// set the repo path based on inputs
	project.LocalRepoPath = fmt.Sprintf("%s/src/github.com/%s/%s", 
		localConfig.GoPath, 
		project.GithubOwner, 
		project.RepoName,
	)
}

func createGitRepo(project *Project, localConfig *LocalConfig) (err error) {
	err = os.Mkdir(project.LocalRepoPath,0755)
	if err != nil {
		return err
	}
	return 
}

func createFiles(project *Project, localConfig *LocalConfig) (err error) {
	fileSlice := []ProjectFile{}
	fileSlice = append(fileSlice, ProjectFile{
		"templates/Dockerfile",
		"Dockerfile", 
		true,
	})

	for _, file := range fileSlice {
		if file.Render == true {
			err = renderFile(&file, project)
			if err != nil {
				return
			}
		} else {
			err = saveFile(&file, project)
			if err != nil {
				return
			}
		}
	}
	return nil
}


func renderFile (file *ProjectFile, project *Project) (err error) {
	f, err := os.Create(fmt.Sprintf("%s/%s",project.LocalRepoPath,file.RepoFilePath))
	if err != nil {
		return err
	}
	tmpl, err := template.ParseFiles(file.TemplateFilePath)
	if err != nil { 
		return err
	}
	err = tmpl.Execute(f, project)
	if err != nil { 
		return err
	 }
	f.Close()
	return nil
}

func saveFile (file *ProjectFile, project *Project) (err error) {
	contents, err := ioutil.ReadFile(file.TemplateFilePath)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s",project.LocalRepoPath,file.RepoFilePath),contents,0644)
	return nil
}