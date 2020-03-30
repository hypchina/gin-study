set d=%date:~0,4%%date:~5,2%%date:~8,2%%time:~0,2%%time:~3,2%%time:~6,2%

set GO111MODULE=off
::set GOROOT=D:\ProgramFiles\Go
set GOPATH=E:\data\storage\jianguoyun\data\project\go
set GOOS=windows
set GOARCH=386
set CGO_ENABLED=1
go build -o ./exec/monitor_386_%d%.exe

set GO111MODULE=off
::set GOROOT=D:\ProgramFiles\Go
set GOPATH=E:\data\storage\jianguoyun\data\project\go
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
go build -o ./exec/monitor_amd64_%d%.exe

pause