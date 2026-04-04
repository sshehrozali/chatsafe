# Retag HEAD and push → triggers .github/workflows/release.yml (GoReleaser).
TAG ?= v1.0.0

.PHONY: release-ci
release-ci:
	git tag -fa "$(TAG)" -m "$(TAG)"
	git push -f origin "$(TAG)"
