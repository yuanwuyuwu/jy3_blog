blog_server : blog_server.go
	go build blog_server.go

.PHONY: run_blog

run_blog:
	sudo nohup ./blog_server > blog.log &

.PHONY: kill_blog

kill_blog:
	echo $$(ps -ef|grep blog_server|grep -v grep|awk '{print $$2}')
	sudo kill -9 $$(ps -ef|grep blog_server|grep -v grep|awk '{print $$2}')
