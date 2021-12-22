all: runTS runGo

containerTS:
		docker build -t 1-ts ./ts

runTS: containerTS
		docker run -v `pwd`/input.txt:/usr/app/input.txt 1-ts

containerGo:
		docker build -t 1-go ./go

runGo: containerGo
		docker run -v `pwd`/input.txt:/app/input.txt 1-go