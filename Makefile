echo: 
	pwd
	echo "build main.go!"

build: echo
	rm -f -r main
	go build main.go

docker: build
	docker build -t ygdxdyj/httpserver:v1.0 .
push: docker
	docker push ygdxdyj/httpserver:v1.0
