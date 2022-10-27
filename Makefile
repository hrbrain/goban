
test:
	go test --cover -mod=mod ./... -count=1

testv:
	go test -v --cover -mod=mod ./... -count=1