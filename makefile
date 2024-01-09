# Variables
IMAGE_NAME := inputobject2/random-picker
TAGS := latest 0.0.1 0.0 0

# Build the Docker image and tag it
build:
	@docker build -t $(IMAGE_NAME):latest .

# Tag the image with additional tags
tag: build
	@$(foreach tag,$(TAGS),docker tag $(IMAGE_NAME):latest $(IMAGE_NAME):$(tag);)

# Push the Docker image with all tags
push: tag
	@$(foreach tag,$(TAGS),docker push $(IMAGE_NAME):$(tag);)

# All-in-one command to build, tag, and push
all: push

.PHONY: build tag push all
