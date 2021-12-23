all: runTS runGo

containerTS:
		docker build -q -t 2a-ts ./ts

runTS: containerTS
		docker run -v `pwd`/input.txt:/usr/app/input.txt 2a-ts

containerGo:
		docker build -q -t 2a-go ./go

runGo: containerGo
		docker run -v `pwd`/input.txt:/app/input.txt 2a-go