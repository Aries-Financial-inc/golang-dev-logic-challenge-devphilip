# Options Contract Analysis
## This is an options analysis Backend written in Go.

### To run the application

#### 1. Clone the repository
git clone https://github.com/philipakpeki/golang-dev-logic-challenge-devphilip.git

#### 2. Move into the cloned directory
cd golang-dev-logic-challenge-devphilip

#### 3. get dependencies
go mod download

#### 4. run the application
go run .

#### 5. test the endpoint using curl
curl -X POST http://localhost:8080/analyze -d @testdata/testdata.json

### you can visit the swagger docs at:
http://localhost:8080/swagger-ui/index.html

### To run the tests
go test tests/analysis_test.go -v