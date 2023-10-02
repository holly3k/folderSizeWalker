# folderSizeWalkerSingalBinary

scan a folder and generate a JSON with detail size for every folder and files in it
can be compile to an executable file or a share lib for other application

## 1 To run as a executable app
```
go build
folderSizeWalker c:\\
```
## 2 To compile as a lib for used by other application
```
go build -o folderSizeWalker.dll -buildmode=c-shared folderSizeWalker.go
```
