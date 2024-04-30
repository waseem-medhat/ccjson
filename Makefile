debug:
	@echo
	@echo "Original JSON"
	@echo "-------------"
	@cat ${file}
	@echo
	@echo
	@echo "Checker Output"
	@echo "-------------"
	@DEBUG=true go run . ${file}
	@echo
