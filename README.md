# HTTP Test Caller

This project is a wrapper for testing HTTP handler, it will pass the request to the handler and record the response using `httptest.ResponseRecorder`. If the response type is set, it will try to unmarshal the response body to the type.

## Usage

### Importing

```go
import "github.com/AccelByte/http-test-caller"
```

### Make a request

```go
resp, body, err :=
    caller.Call(handler).
        To(gorequest.New().    // using gorequest to construct the request
            Post("/example").  // we are testing handler locally, doesn't need full URL
            Type(restful.MIME_JSON).
            SendString(`{"Example": "test"}`).
            MakeRequest()).
        Read(&models.ExampleResponse{}).
        Execute()
    
// we can assert the resp.Code & unmarshaled body
```