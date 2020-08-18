package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
)

func init() {
	for _, x := range []*env{
		&host,
		&token,
		&channel,
		&maxRetry,

		&repoFullname,
		&remoteURL,
		&commitSha,
		&commitBranch,
		&commitMessage,
		&commitLink,
		&commitAuthorName,
		&buildEvent,
		&buildStatus,
		&buildNumber,
	} {
		x.Value = os.Getenv(x.EnvVar)
	}
}

// report things: repo(name),time,commit(author/hash/msg),branch,build result(event/)
func main() {
	c := model.NewAPIv4Client(host.Value)
	c.SetOAuthToken(token.Value)

	maxRetryCount, err := strconv.Atoi(maxRetry.Value)
	if err != nil {
		panic(err)
	}
	if maxRetryCount == 0 {
		maxRetryCount = 3
	}
	emoji := " ✅"
	if strings.Contains(buildStatus.Value, "failure") {
		emoji = "⚠️"
	}
	msg := fmt.Sprintf(`CI build: %s (branch: %s)
commit: %s(%s:%s)
event: %s:%s
%s%s`,
		repoFullname, commitBranch,
		commitLink, commitAuthorName, commitMessage,
		buildEvent, buildNumber,
		buildStatus, emoji)
	success := false
	for i := 0; i < maxRetryCount; i++ {
		_, resp := c.CreatePost(&model.Post{
			ChannelId: channel.Value,
			Message:   msg,
		})
		if resp.Error == nil {
			break
		}
		log.Printf("create post failed (i:%d), %v", i, resp.Error)
		time.Sleep(time.Duration(i) * time.Second)
		success = true
		continue
	}
	if !success {
		os.Exit(1) //failed
	}
}

var (
	//mattermost env
	host = env{
		EnvVar: "host",
	}
	token = env{
		EnvVar: "token",
	}
	channel = env{
		EnvVar: "channel",
	}
	maxRetry = env{
		Usage:  "max retry when post message failed",
		EnvVar: "maxRetry",
	}

	//drone env
	repoFullname = env{
		Name:   "repo.fullname",
		Usage:  "repository full name",
		EnvVar: "DRONE_REPO",
	}
	remoteURL = env{
		Name:   "remote.url",
		Usage:  "git remote url",
		EnvVar: "DRONE_REMOTE_URL",
	}
	commitSha = env{
		Name:   "commit.sha",
		Usage:  "git commit sha",
		EnvVar: "DRONE_COMMIT_SHA",
	}
	commitBranch = env{
		Name:   "commit.branch",
		Value:  "master",
		Usage:  "git commit branch",
		EnvVar: "DRONE_COMMIT_BRANCH",
	}
	commitMessage = env{
		Name:   "commit.message",
		Usage:  "git commit message",
		EnvVar: "DRONE_COMMIT_MESSAGE",
	}
	commitLink = env{
		Name:   "commit.link",
		Usage:  "git commit link",
		EnvVar: "DRONE_COMMIT_LINK",
	}
	commitAuthorName = env{
		Name:   "commit.author.name",
		Usage:  "git author name",
		EnvVar: "DRONE_COMMIT_AUTHOR",
	}
	buildEvent = env{
		Name:   "build.event",
		Value:  "push",
		Usage:  "build event",
		EnvVar: "DRONE_BUILD_EVENT",
	}
	buildStatus = env{
		Name:   "build.status",
		Usage:  "build status",
		Value:  "success",
		EnvVar: "DRONE_BUILD_STATUS",
	}
	buildNumber = env{
		Name:   "build.number",
		Usage:  "build number",
		EnvVar: "DRONE_BUILD_NUMBER",
	}
)
