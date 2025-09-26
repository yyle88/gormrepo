COVERAGE_DIR ?= .coverage.out

# cp from: https://github.com/yyle88/gormcngen/blob/5f75d814c71ec306b276a804c134de6655913951/Makefile#L4
test:
	@if [ -d $(COVERAGE_DIR) ]; then rm -r $(COVERAGE_DIR); fi
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
