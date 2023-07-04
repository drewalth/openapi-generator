setup: install_go_dependencies install_git_hooks

install_go_dependencies:
	@go get

install_git_hooks:
	@git config core.hooksPath hooks
