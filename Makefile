build-fe:
	cd web && npm install && npm run build && rm -rf ../static && mv dist ../static

run:
	go run main.go

serve: build-fe run
