package main

import (
	"fmt"
	"os"

	"github.com/bndr/gojenkins"
)

func main() {
	jenkinsURL := os.Getenv("JENKINS_URL")
	jenkinsUser := os.Getenv("JENKINS_USER")
	jenkinsToken := os.Getenv("JENKINS_TOKEN")
	jenkinsJob := os.Getenv("JENKINS_JOB")

	jenkins := gojenkins.CreateJenkins(nil, jenkinsURL, jenkinsUser, jenkinsToken)

	_, err := jenkins.Init()
	if err != nil {
		panic(err)
	}

	job, err := jenkins.GetJob(jenkinsJob)
	if err != nil {
		panic(err)
	}

	var buildsResp struct {
		Builds []*gojenkins.BuildResponse `json:"allBuilds"`
	}
	_, err = jenkins.Requester.GetJSON(job.Base, &buildsResp, map[string]string{"tree": "allBuilds[number,id,timestamp,result,duration]"})
	if err != nil {
		panic(err)
	}

	for _, build := range buildsResp.Builds {
		if build.Result == "" {
			continue
		}
		// fmt.Printf("%+v\n", build)
		// fmt.Printf("%v\t%v\t%v\n", build.Number, time.Unix(build.Timestamp/1000, 0), time.Duration(build.Duration)*time.Millisecond)
		fmt.Printf("%v\t%v\t%v\n", build.Number, build.Timestamp, build.Duration)
	}
}
