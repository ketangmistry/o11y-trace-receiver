package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsublite/pscompat"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func traces(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	topicPath := os.Getenv("PUBSUB_TOPIC")

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		// TODO: Handle error - should continue but show error
	}

	publisher, err := pscompat.NewPublisherClient(ctx, topicPath)
	if err != nil {
		// TODO: Handle error  - should continue but show error
	}

	result := publisher.Publish(ctx, &pubsub.Message{Data: []byte(b)})

	id, err := result.Get(ctx)
	if err != nil {
		// TODO: Handle error  - should continue but show error
	}

	publisher.Stop()

	fmt.Printf("%s", id)
}

func main() {

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("o11y-trace-receiver"),
		newrelic.ConfigLicense(os.Getenv("NR_INGEST_LICENSE_KEY")),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		// TODO: Handle error -  - should continue but show error
	}

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/v1/traces", traces))

	http.ListenAndServe(":8090", nil)
}
