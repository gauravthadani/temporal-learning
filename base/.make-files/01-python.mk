
## Python:

.ONESHELL:
ENV_PREFIX=$(shell python -c "if __import__('pathlib').Path('.venv/bin/pip').exists(): print('.venv/bin/')")


.PHONY: show
show: ## Show the current environment.
	@echo "Current environment:"
	@echo "Running using $(ENV_PREFIX)"
	@$(ENV_PREFIX)python -V
	@$(ENV_PREFIX)python -m site

.PHONY: clean
clean: ## Clean unused files.
	@find ./ -name '*.pyc' -exec rm -f {} \;
	@find ./ -name '__pycache__' -exec rm -rf {} \;
	@find ./ -name 'Thumbs.db' -exec rm -f {} \;
	@find ./ -name '*~' -exec rm -f {} \;
	@rm -rf .cache
	@rm -rf .pytest_cache
	@rm -rf .mypy_cache
	@rm -rf build
	@rm -rf dist
	@rm -rf *.egg-info
	@rm -rf htmlcov
	@rm -rf .tox/
	@rm -rf docs/_build	


.PHONY: fmt
fmt: ## Format code using black & isort.
	$(ENV_PREFIX)isort $(PACKAGE_NAME)/
	$(ENV_PREFIX)black -l 79 $(PACKAGE_NAME)/
	$(ENV_PREFIX)black -l 79 tests/

.PHONY: lint
lint: ## Run pep8, black, mypy linters.
	$(ENV_PREFIX)flake8 $(PACKAGE_NAME)/
	$(ENV_PREFIX)black -l 79 --check $(PACKAGE_NAME)/
	$(ENV_PREFIX)black -l 79 --check tests/
	$(ENV_PREFIX)mypy --ignore-missing-imports $(PACKAGE_NAME)/

.PHONY: test
test: ## Run all tests.
	@#$(ENV_PREFIX)coverage run -m pytest
	@$(ENV_PREFIX)pytest --cov

.PHONY: virtualenv
virtualenv: ## Create a virtual environment.
	@echo "creating virtualenv ..."
	@rm -rf .venv
	@python3 -m venv .venv
	@./.venv/bin/pip install -U pip
	@./.venv/bin/pip install -e .[tests]
	@echo
	@echo "!!! Please run 'source .venv/bin/activate' to enable the environment !!!"


.PHONY: package
package: ## packages the disrtibution
	@echo "packaging distribution ..."
	@zip .venv.zip -r .venv
