build-fe:
	cd web && npm install && npm run build && rm -rf ../static && mv dist ../static

build-be:
	go build .

install:
	sudo cp -r static /opt/gober/ && sudo cp gober /opt/gober/ && sudo chmod +x /opt/gober/gober

build-gober: build-fe build-be

install-gober: build-gober install
