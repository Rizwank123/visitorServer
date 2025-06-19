# Setup envoironment
setup: 
	@echo "Setting up the Platform"
	@echo "Generating the .env file"
	cp sample.env .env
	@echo "set up the envoironment values"

# Setup docker
docker-setup:
	@echo "Building the docker image"
	docker-compose down 
	docker-compose up --build -d
	@echo "Docker image build and started successfully"


# test the app
test: 
	@echo "Running the unit and integration tests"
	go test ./... -v
	
# run the app
run:
	@echo "Start the Application"
	go run github.com/Rizwank123/visitorServer
		