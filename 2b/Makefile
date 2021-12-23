all: runTS runGo

containerTS:
		docker build -q -t 2b-ts ./ts

runTS: containerTS
		docker run -v `pwd`/input.txt:/usr/app/input.txt 2b-ts

containerGo:
		docker build -q -t 2b-go ./go

runGo: containerGo
		docker run -v `pwd`/input.txt:/app/input.txt 2b-go