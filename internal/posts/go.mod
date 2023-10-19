module github.com/OrIX219/SomethingSocial/internal/posts

go 1.21.0

require (
	github.com/OrIX219/SomethingSocial/internal/common v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/render v1.0.3
	github.com/google/uuid v1.3.1
	github.com/oapi-codegen/runtime v1.0.0
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
	google.golang.org/appengine v1.6.7
	google.golang.org/grpc v1.59.0
)

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231009173412-8bfb1ae86b6c // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace github.com/OrIX219/SomethingSocial/internal/common => ../common
