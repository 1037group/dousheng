.PHONY: all

SUBDIRS = ./cmd/comment ./cmd/favorite ./cmd/feed ./cmd/message ./cmd/publish ./cmd/relation ./cmd/user

update:
	kitex --thrift-plugin validator -module github.com/1037group/dousheng idl/comment.thrift # execute in the project root directory

	kitex --thrift-plugin validator -module github.com/1037group/dousheng idl/favorite.thrift # execute in the project root directory

	kitex --thrift-plugin validator -module github.com/1037group/dousheng idl/feed.thrift # execute in the project root directory

	kitex --thrift-plugin validator -module github.com/1037group/dousheng idl/message.thrift # execute in the project root directory

	kitex --thrift-plugin validator -module github.com/1037group/dousheng idl/publish.thrift # execute in the project root directory

	kitex --thrift-plugin validator -module github.com/1037group/dousheng idl/relation.thrift # execute in the project root directory

	kitex --thrift-plugin validator -module github.com/1037group/dousheng idl/user.thrift # execute in the project root directory

	@list='$(SUBDIRS)'; for subdir in $$list; do \
    		echo "make in $$subdir";\
    		cd $$subdir && $(MAKE) && cd ../..;\
    	done

	cd ./cmd/api &&	make update_api

go_build:
	@list='$(SUBDIRS)'; for subdir in $$list; do \
        		echo "go build in $$subdir";\
        		cd $$subdir && go build && cd ../..;\
		done
	cd ./cmd/api && go build
