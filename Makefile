start-frontend:
	cd old.frontend
	npm start

format-frontend:
	npx prettier --write old.frontend/src/

start-backend:
	cd old.backend
	python3 app.py

format-backend:
	black ./old.backend/

build-server:
	go build ./cmd/server

build-ui:
	go build ./cmd/ui

build-server-release:
	go build -ldflags "-s -w" ./cmd/server

build-ui-release:
	go build -ldflags "-s -w" ./cmd/ui
