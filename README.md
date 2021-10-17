# Promred
This package provides [RED Method](https://grafana.com/blog/2018/08/02/the-red-method-how-to-instrument-your-services/) 
instrumentation via [Prometheus](https://prometheus.io/) for the following

- [AWS SNS](#aws-sns)
- [AWS SQS](#aws-sqs)
- [Doers](#doers)
- [Handlers](#handlers)
- [Redis](#redis)

## AWS SNS
Disclaimer: This makes use of [V2 of the AWS-SDK-Go](https://github.com/aws/aws-sdk-go-v2)

Available methods
- Publish

### How to use
```go
import (
	instrumentation "github.com/jamieaitken/promred/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

snsClient := sns.New(sns.Options{})

instr := instrumentation.New(snsClient)

out, err := instr.Publish(ctx, &sns.PublishInput{}, nil, "main")
if err != nil {
	return err
}
```


## AWS SQS
Disclaimer: This makes use of [V2 of the AWS-SDK-Go](https://github.com/aws/aws-sdk-go-v2)

Available methods
- ReceiveMessage
- SendMessage

### How to use
```go
import (
	instrumentation "github.com/jamieaitken/promred/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

sqsClient := sqs.New(sqs.Options{})

instr := instrumentation.New(sqsClient)

out, err := instr.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{}, nil, "main")
if err != nil {
	return err
}
```

## Doers
A doer performs an HTTP Request and returns either an HTTP Response or an error.

Available methods
- Do

### How to use
```go
import (
	instrumentation "github.com/jamieaitken/promred/doer"
	"net/http"
)

httpClient := &http.Client{}

instr := instrumentation.New(httpClient)

request := http.NewRequestWithContext(context.Background, http.MethodGet, "https://example.com", nil)

res, err := instr.Do(request)
if err != nil {
	return err
}
```

## Handlers
This provides instrumentation for http.Handlers, more specifically [HandlerFunc](https://pkg.go.dev/net/http#HandlerFunc)

Available methods
- HandleFor

### How to use
```go
import (
    instrumentation "github.com/jamieaitken/promred/handler"
    "net/http"
)

instr := instrumentation.New()

router := new(http.ServeMux)

router.HandleFunc("/v1/docs", h.HandleFor(handler.Get))


```

## Redis
This accepts a [go-redis](https://github.com/go-redis/redis) client and provides instrumentation for the following 
methods

Available methods
- Get
- HGet
- MGet
- MSet
- Set
- SetEX
- Ping

### How to use 
```go
import (
    "github.com/go-redis/redis/v8"
    instrumentation "github.com/jamieaitken/promred/redis"
)

redisClient := redis.NewClient(&redis.Options{})

instr := instrumentation.New(redisClient)

cmd := instr.Get(context.Background(), "key", "main")
if cmd.Err() != nil {
	return cmd.Err()
}
```