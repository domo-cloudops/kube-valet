CUR_TAG ?= $(shell git describe --abbrev=0 --tags 2>/dev/null)
PREV_TAG ?= $(shell git describe --abbrev=0 --tags $(CUR_TAG)^)

VERSION ?= $(shell echo $${BRANCH_NAME:-local} | sed s/[^a-zA-Z0-9_-]/_/)_$(shell git describe --always --dirty)
IMAGE ?= domoinc/kube-valet

.PHONY: all kube-valet valetctl test release

all: customresources build

build: kube-valet valetctl

kube-valet:
	mkdir build || true
	CGO_ENABLED=0 GOOS=linux go build -v -i -pkgdir $(PWD)/build/gopkgs --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo -o build/kube-valet

valetctl:
	mkdir build || true
	CGO_ENABLED=0 GOOS=linux go build -v -i -pkgdir $(PWD)/build/gopkgs --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo -o build/valetctl bin/valetctl.go

clean:
	rm build/* || true

test: test-customresources test-pkgs

test-pkgs:
	# client-go is huge, install deps so future tests are faster
	go test -i ./pkg/...

	# run tests
	go test -v ./pkg/...

docker-image:
	docker build -t $(IMAGE):$(VERSION) .

release: docker-image
	git tag -f $(VERSION) -m "Kube-valet release $(VERSION)"

push-release: release
	git push github $(VERSION)

	# Push the versioned tag
	docker push $(IMAGE):$(VERSION)

	# Push the latest tag
	docker tag $(IMAGE):$(VERSION) $(IMAGE):latest
	docker push $(IMAGE):latest

# Targets to build custom resources and clients

customresources: clean-customresources gen-customresources test-customresources

gen-customresources: clean-customresources
    # Install vendor files from modules
	go mod vendor

	cat ./vendor/k8s.io/code-generator/generate-groups.sh ./_openapi/openapi-gen.sh > ./vendor/k8s.io/code-generator/generate-groups-custom.sh
	chmod +x ./vendor/k8s.io/code-generator/generate-groups-custom.sh

	mkdir -p ./pkg/client/openapi

	echo "package openapi" > ./pkg/client/openapi/doc.go
	cp ./_openapi/path_template.tmpl ./pkg/client/openapi
	cp ./_openapi/print_test.go ./pkg/client/openapi

	# Generate client and deepcopy
	./vendor/k8s.io/code-generator/generate-groups-custom.sh deepcopy,client,informer,lister,openapi \
	github.com/domoinc/kube-valet/pkg/client \
	github.com/domoinc/kube-valet/pkg/apis \
	"assignments:v1alpha1" \
	--output-base ./build \
	--go-header-file "$(PWD)/boilerplate/boilerplate.go.txt"

	# Move generated files
	mv build/github.com/domoinc/kube-valet/pkg/apis/assignments/v1alpha1/zz_generated.deepcopy.go pkg/apis/assignments/v1alpha1/
	mv \
		build/github.com/domoinc/kube-valet/pkg/client/clientset \
		build/github.com/domoinc/kube-valet/pkg/client/informers \
		build/github.com/domoinc/kube-valet/pkg/client/listers \
		pkg/client
	mv \
		build/github.com/domoinc/kube-valet/pkg/client/openapi/* \
		pkg/client/openapi

	# Cleanup gen dir
	rm -rf build/github.com

	go test ./pkg/client/openapi/*.go -test.run=TestWriteOpenAPISpec

	rm ./pkg/client/openapi/path_template.tmpl
	rm ./pkg/client/openapi/print_test.go
	rm ./pkg/client/openapi/openapi_generated.go

clean-customresources:
	# Delete all generated code.
	rm -rf pkg/client
	rm -f pkg/apis/*/*/zz_generated.deepcopy.go

# This is a basic smoke-test to make sure the types compile
test-customresources:
	go build -o build/crud -i _examples/clients/crud.go
	go build -o build/list -i _examples/clients/list.go

	@echo "All custom resource client test binaries compiled!"

release-notes:
	@echo "## Changes since $(PREV_TAG)"
	@git log --no-merges --format='%s'  $(PREV_TAG)..$(CUR_TAG) | sed 's/^/ - /'

