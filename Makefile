# Go parameters
GOCMD=go
GOINSTALL=go build
GOARM = env GOOS=linux GOARCH=arm go build
GOLINUX86 = env GOOS=linux GOARCH=386 go build
GOLINUX64 = env GOOS=linux GOARCH=amd64 go build
GODARWIN32 = env GOOS=darwin GOARCH=386 go build
GODARWIN64 = env GOOS=darwin GOARCH=amd64 go build
GOWINDOWS32 = env GOOS=windows GOARCH=386 go build
GOWINDOWS64 = env GOOS=windows GOARCH=amd64 go build
# Avoid problem with the gopath adding
# temporarely ourself to it
GOPATH := $(CURDIR)

# Create a stand-alone repository
# in order to avoid to download
# untested 3° party dependences

GOPATH := $(CURDIR)/src/_vendor:$(GOPATH)

all:
	cd ./src && $(GOINSTALL) 
	cd ./src && mv src mDota
deploy:
	cd src && $(GOARM) && mv src mDota_arm
	cd src && $(GOLINUX86) && mv src mDota_linux_amd64
	cd src && $(GOLINUX64) && mv src mDota_linux_386
	cd src && $(GODARWIN32) && mv src mDota_darwin_386
	cd src && $(GODARWIN64) && mv src mDota_darwin_amd64
	cd src && $(GOWINDOWS32) && mv src.exe mDota_windows_386.exe
	cd src && $(GOWINDOWS64) && mv src.exe mDota_windows_amd64.exe
	mkdir windows_amd64 && cp -r data windows_amd64 && mv src/mDota_windows_amd64.exe windows_amd64
	mkdir windows_386 && cp -r data windows_386 && mv src/mDota_windows_386.exe windows_386
	mkdir darwin_amd64 && cp -r data darwin_amd64 && mv src/mDota_darwin_amd64 darwin_amd64
	mkdir darwin_386 && cp -r data darwin_386 && mv src/mDota_darwin_386 darwin_386
	mkdir linux_amd64 && cp -r data linux_amd64 && mv src/mDota_linux_amd64 linux_amd64
	mkdir linux_386 && cp -r data linux_386 && mv src/mDota_linux_386 linux_386
	mkdir linux_arm && cp -r data linux_arm && mv src/mDota_arm linux_arm
	zip -r windows_amd64 windows_amd64
	zip -r windows_386 windows_386
	zip -r darwin_amd64 darwin_amd64
	zip -r darwin_386 darwin_386
	zip -r linux_amd64 linux_amd64  
	zip -r linux_386 linux_386
	zip -r linux_arm linux_arm
	rm -rf windows_amd64
	rm -rf windows_386
	rm -rf darwin_amd64
	rm -rf darwin_386
	rm -rf linux_amd64
	rm -rf linux_386
	rm -rf linux_arm