TAG := $(shell git rev-parse --short HEAD)
IMAGE := johnwesonga/go-http-server

build:
	docker build -t $(IMAGE):$(TAG) .

load:
	kind load docker-image $(IMAGE):$(TAG)

deploy:
	cd k8s/overlays/dev && \
	kustomize edit set image $(IMAGE)=$(IMAGE):$(TAG) && \
	kubectl apply -k .

all: build load deploy