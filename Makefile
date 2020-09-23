## Create container
container:
	docker build --rm --tag smol:go .

## Handle program
run:
	docker run --name smol-go --publish 80:8080 --rm smol:go ./main
start:
	docker run --name smol-go --publish 80:8080 --rm --detach smol:go ./main
stop:
	docker stop smol-go

## Get container size
size: start
	docker ps --size | grep "smol:go" | sed -e "s/.*virtual \(.*\)).*/\1/"