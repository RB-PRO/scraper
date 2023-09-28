all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/PhotoTemaParser.git

pull:
	git pull git@github.com:RB-PRO/PhotoTemaParser.git

pushW:
	git push https://github.com/RB-PRO/PhotoTemaParser.git

pullW:
	git pull https://github.com/RB-PRO/PhotoTemaParser.git

doc:
	godoc -http :8080