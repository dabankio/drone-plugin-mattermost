package main

//copied from https://github.com/drone/drone-plugin-starter

type env struct {
	Name   string
	Usage  string
	EnvVar string
	Value  string
}

func (e env) String() string {
	return e.Value
}

var envList = []env{
	{
		Name:   "repo.fullname",
		Usage:  "repository full name",
		EnvVar: "DRONE_REPO",
	},
	{
		Name:   "repo.owner",
		Usage:  "repository owner",
		EnvVar: "DRONE_REPO_OWNER",
	},
	{
		Name:   "repo.name",
		Usage:  "repository name",
		EnvVar: "DRONE_REPO_NAME",
	},
	{
		Name:   "repo.link",
		Usage:  "repository link",
		EnvVar: "DRONE_REPO_LINK",
	},
	{
		Name:   "repo.avatar",
		Usage:  "repository avatar",
		EnvVar: "DRONE_REPO_AVATAR",
	},
	{
		Name:   "repo.branch",
		Usage:  "repository default branch",
		EnvVar: "DRONE_REPO_BRANCH",
	},
	{
		Name:   "repo.private",
		Usage:  "repository is private",
		EnvVar: "DRONE_REPO_PRIVATE",
	},
	{
		Name:   "repo.trusted",
		Usage:  "repository is trusted",
		EnvVar: "DRONE_REPO_TRUSTED",
	},

	// commit args

	{
		Name:   "remote.url",
		Usage:  "git remote url",
		EnvVar: "DRONE_REMOTE_URL",
	},
	{
		Name:   "commit.sha",
		Usage:  "git commit sha",
		EnvVar: "DRONE_COMMIT_SHA",
	},
	{
		Name:   "commit.ref",
		Value:  "refs/heads/master",
		Usage:  "git commit ref",
		EnvVar: "DRONE_COMMIT_REF",
	},
	{
		Name:   "commit.branch",
		Value:  "master",
		Usage:  "git commit branch",
		EnvVar: "DRONE_COMMIT_BRANCH",
	},
	{
		Name:   "commit.message",
		Usage:  "git commit message",
		EnvVar: "DRONE_COMMIT_MESSAGE",
	},
	{
		Name:   "commit.link",
		Usage:  "git commit link",
		EnvVar: "DRONE_COMMIT_LINK",
	},
	{
		Name:   "commit.author.name",
		Usage:  "git author name",
		EnvVar: "DRONE_COMMIT_AUTHOR",
	},
	{
		Name:   "commit.author.email",
		Usage:  "git author email",
		EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
	},
	{
		Name:   "commit.author.avatar",
		Usage:  "git author avatar",
		EnvVar: "DRONE_COMMIT_AUTHOR_AVATAR",
	},

	// build args
	{
		Name:   "build.event",
		Value:  "push",
		Usage:  "build event",
		EnvVar: "DRONE_BUILD_EVENT",
	},
	{
		Name:   "build.number",
		Usage:  "build number",
		EnvVar: "DRONE_BUILD_NUMBER",
	},
	{
		Name:   "build.created",
		Usage:  "build created",
		EnvVar: "DRONE_BUILD_CREATED",
	},
	{
		Name:   "build.started",
		Usage:  "build started",
		EnvVar: "DRONE_BUILD_STARTED",
	},
	{
		Name:   "build.finished",
		Usage:  "build finished",
		EnvVar: "DRONE_BUILD_FINISHED",
	},
	{
		Name:   "build.status",
		Usage:  "build status",
		Value:  "success",
		EnvVar: "DRONE_BUILD_STATUS",
	},
	{
		Name:   "build.link",
		Usage:  "build link",
		EnvVar: "DRONE_BUILD_LINK",
	},
	{
		Name:   "build.deploy",
		Usage:  "build deployment target",
		EnvVar: "DRONE_DEPLOY_TO",
	},
	{
		Name:   "yaml.verified",
		Usage:  "build yaml is verified",
		EnvVar: "DRONE_YAML_VERIFIED",
	},
	{
		Name:   "yaml.signed",
		Usage:  "build yaml is signed",
		EnvVar: "DRONE_YAML_SIGNED",
	},

	// prev build args
	{
		Name:   "prev.build.number",
		Usage:  "previous build number",
		EnvVar: "DRONE_PREV_BUILD_NUMBER",
	},
	{
		Name:   "prev.build.status",
		Usage:  "previous build status",
		EnvVar: "DRONE_PREV_BUILD_STATUS",
	},
	{
		Name:   "prev.commit.sha",
		Usage:  "previous build sha",
		EnvVar: "DRONE_PREV_COMMIT_SHA",
	},
}
