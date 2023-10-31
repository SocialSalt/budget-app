start-frontend:
	cd src/frontend
	npm start

format-frontend:
	npx prettier --write src/frontend/src/

start-backend:
	cd src/backend
	python3 app.py

format-backend:
	black ./src/backend/
