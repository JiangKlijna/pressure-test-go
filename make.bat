@echo off

if "%1%" == "" (
	go build -ldflags "-s -w"
) else (
	if "%1%" == "run" (
		go build -o out\\tmp.exe && out\\tmp.exe
	) else (
		if "%1%" == "test" (
			go test
		) else (
			if "%1%" == "clean" (
				rd /s /q out
			) else (
				if "%1%" == "publish" (
					call:publish
				) else (
					echo Example:
					echo 	make
					echo 	make run
					echo 	make test
					echo 	make clean
					echo 	make publish
				)
			)
		)
	)
)
goto :eof

:publish
	call:cross_compile darwin 386
	call:cross_compile darwin amd64

	call:cross_compile linux 386
	call:cross_compile linux amd64
	call:cross_compile linux arm
	call:cross_compile linux arm64
	call:cross_compile linux mips
	call:cross_compile linux mips64

	call:cross_compile windows 386
	call:cross_compile windows amd64
goto :eof

:cross_compile
	set GOOS=%~1
	set GOARCH=%~2
	set TARGET_FILE=rainx_%~1_%~2
	if %~1==windows set TARGET_FILE=%TARGET_FILE%.exe
	go build -i -o out/%TARGET_FILE% -ldflags "-s -w"
	echo compile %TARGET_FILE% done
goto :eof
