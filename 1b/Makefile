all: runTS runGo

containerTS:
		docker build -q -t 1b-ts ./ts

runTS: containerTS
		docker run -v `pwd`/input.txt:/usr/app/input.txt 1b-ts

containerGo:
		docker build -q -t 1b-go ./go

runGo: containerGo
		docker run -v `pwd`/input.txt:/app/input.txt 1b-go