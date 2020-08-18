dockerImage:
	#CGO_ENABLED is needed for linux build, see https://stackoverflow.com/questions/51508150/standard-init-linux-go190-exec-user-process-caused-no-such-file-or-directory
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o plugin.run 
	docker build --tag drone-plugin-mattermost:4 . 
	#test: docker run --rm -e "maxRetry=4" -e "host=http://" -e "token=" -e "channel=" drone-plugin-mattermost:1
	@echo docker tag drone-plugin-mattermost:4 dabankio/drone-plugin-mattermost:4
	@echo push: docker push dabankio/drone-plugin-mattermost:4