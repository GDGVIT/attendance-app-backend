help:
	@echo ''
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo 'make dev: make dev for development work'
	@echo 'make production: docker production build'

dev:
	if [ ! -f .env ]; then cp .env.example .env; fi;
	docker build -t dev_go_attendance_server -f Dockerfile-dev .
	docker run -it --rm -p 8000:8000 -v "$(PWD):/app" -v app_logs:/app_logs --name attendance-container-dev --env-file .env dev_go_attendance_server

prod:
	if [ ! -f .env ]; then cp .env.example .env; fi;
	docker build -t prod_go_attendance_server -f Dockerfile .
	if docker ps -a --format "{{.Names}}" | grep -q '^attendance-container-prod$$'; then \
        docker stop attendance-container-prod; \
        docker rm attendance-container-prod; \
		echo "Removed old prod container"; \
    fi;
	docker run -d -p 8000:8000 -v app_logs:/app_logs --name attendance-container-prod --env-file .env prod_go_attendance_server

remove-prod:
	if docker ps -a --format "{{.Names}}" | grep -q '^attendance-container-prod$$'; then \
		docker stop attendance-container-prod; \
		docker rm attendance-container-prod; \
		echo "Production container has been removed."; \
	else \
		echo "No running production container found."; \
	fi;
