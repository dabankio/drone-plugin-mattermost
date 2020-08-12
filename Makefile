dockerImage:
	GOOS=linux GOARCH=amd64 go build -o plugin.run
	docker build --tag drone-plugin-mattermost:1 . 
	#test: docker run --rm -e "maxRetry=3" -e "host=http://" -e "token=" -e "channel=" drone-plugin-mattermost:1
	echo docker tag drone-plugin-mattermost:1 dabankio/drone-plugin-mattermost:1
	echo push: docker push dabankio/drone-plugin-mattermost:1