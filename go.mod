module github.com/mkyc/go-use-versioned-structs-tests

go 1.15

//for local development in parallel to  because of https://github.com/golang/go/issues/31365#issuecomment-481713692
//replace github.com/mkyc/go-structs-versioning-tests => ../go-structs-versioning-tests

require github.com/mkyc/go-structs-versioning-tests v0.0.1-alpha.2
