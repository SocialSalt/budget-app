start-frontend:
	cd frontend
	npm start

format-frontend:
	npx prettier --write frontend/src/

start-backend:
	cd backend
	python3 app.py

format-backend:
	black ./backend/
