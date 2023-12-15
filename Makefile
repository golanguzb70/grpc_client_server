copy_proto:
	rm -rf ./client/protos/* && cp -r ./proto/* ./client/protos/
	rm -rf ./server/protos/* && cp -r ./proto/* ./server/protos/