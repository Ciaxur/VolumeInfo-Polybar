APP = volume-info

# Build Source
# Copy Over Default Notifcation Image
build:
	go build -o ./bin/$(APP) ./src

# Clean up Dev Bin Directory
clean:
	rm -rf ./bin/*
