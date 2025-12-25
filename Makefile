#DEV

build-dev: 
	docker build -t videochat -f containers/Dockerfile . && docker build -t tunel -f containers/images/Dockerfile.turn .

clean-dev:
	docker-compose -f containers/compose/dc.dev.yml.down

run-dev:
	docker-compose -f containers/composes/dc.dev.yml up