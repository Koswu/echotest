ENV_CGO = CGO_ENABLED=0
ENV_WINDOWS = GOOS=windows
ENV_LINUX = GOOS=linux
ENV_386 = GOARCH=386
ENV_AMD64 = GOARCH=amd64
ENV_ARM = GOARCH=arm
ENV_ARM64 = GOARCH=arm64

windows-386 = echotest-windows-386.exe
windows-amd64 = echotest-windows-amd64.exe
linux-386 = echotest-linux-386
linux-amd64 = echotest-linux-amd64
linux-arm = echotest-linux-arm
linux-arm64 = echotest-linux-arm64


main: all
	echo "Done"

build:
	mkdir build -p

$(windows-386):build
	$(ENV_CGO) $(ENV_WINDOWS) $(ENV_386) go build -o build/$@ main.go

$(windows-amd64):build
	$(ENV_CGO) $(ENV_WINDOWS) $(ENV_AMD64) go build -o build/$@ main.go

$(linux-386):build
	$(ENV_CGO) $(ENV_LINUX) $(ENV_386) go build -o build/$@ main.go

$(linux-amd64):build
	$(ENV_CGO) $(ENV_LINUX) $(ENV_AMD64) go build -o build/$@ main.go

$(linux-arm):build
	$(ENV_CGO) $(ENV_LINUX) $(ENV_ARM) go build -o build/$@ main.go

$(linux-arm64):build
	$(ENV_CGO) $(ENV_LINUX) $(ENV_ARM64) go build -o build/$@ main.go

all:$(windows-386) $(windows-amd64) $(linux-386) $(linux-amd64) $(linux-arm) $(linux-arm64)



