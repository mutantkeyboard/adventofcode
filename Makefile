.PHONY: run test clean

# Default day value
day ?= 1

# Run solution for specific day
run:
	@if [ -d "$(day)" ]; then \
		cd $(day) && go run solution.go; \
	else \
		echo "Day $(day) directory does not exist"; \
		exit 1; \
	fi


# Clean build artifacts
clean:
	@find . -type f -name "*.exe" -delete
	@find . -type f -name "*.test" -delete
	@find . -type f -name "*.out" -delete


# Generate new day directory and files
generate_day:
	@if [ -d "$(day)" ]; then \
		echo "Directory $(day) already exists"; \
		exit 1; \
	else \
		mkdir -p $(day); \
		echo "// Solution for Advent of Code 2024 - Day $(day)\n\npackage main\n\nfunc main() {\n\t// Your solution here\n}" > $(day)/solution.go; \
		echo "# Day $(day): [Problem Name]\n\n## Part One\n\n[Problem description]\n\n## Part Two\n\n[Problem description]" > $(day)/README.md; \
		touch $(day)/input.txt; \
		echo "Created directory and files for Day $(day)"; \
	fi

# Generate multiple days at once
generate_days:
	@for i in $$(seq $(start_day) $(end_day)); do \
		$(MAKE) generate_day day=$$i; \
	done
	@echo "Generated days $(start_day) through $(end_day)"