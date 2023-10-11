build_on_windows:
	@echo "----------"
	@echo "Building Demo Executable (Win64)"
	@echo "----------"
	set GOOS=windows&& set GOARCH=amd64&& go build -o ../../bin/win64/demo.exe .
	@echo "----------"
	@echo "Building Demo Executable (Linux64)"
	@echo "----------"
	set GOOS=linux&& set GOARCH=amd64&& go build -o ../../bin/linux64/demo .	

build_on_linux:
	@echo "----------"
	@echo "Building Demo Executable (Win64)"
	@echo "----------"
	GOOS=windows GOARCH=amd64 go build -o ../../bin/win64/demo.exe .
	@echo "----------"
	@echo "Building Demo Executable (Linux64)"
	@echo "----------"
	GOOS=linux GOARCH=amd6 go build -o ../../bin/linux64/demo .