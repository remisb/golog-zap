IMAGE = remisb/golog-zap
VERSION = 0.1.1

run:
	go run main.go

# based on https://github.com/labstack/armor/blob/master/Makefile
publish
	git tag v$(VERSION)
	git push origin --tags
	# goreleaser --rm-dist
	# docker build -t $(IMAGE):$(VERSION) -t $(IMAGE) .
    # docker push $(IMAGE):$(VERSION)
    # docker push $(IMAGE):latest

.PHONY: publish
