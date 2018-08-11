package main

import (
	"crypto/rsa"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"os"
	"time"
)

var privateKeyPath string
var jaegerHostPort string
var secretPrivateKey *rsa.PrivateKey

func main() {
	privateKeyPath = "/secrets/key/private.pem"
	jaegerHostPort = os.Getenv("JAEGER_HOST_PORT")
	if jaegerHostPort == "" {
		jaegerHostPort = "localhost:6832"
	}
	var err error
	secretPrivateKey, err = loadPrivateKey(privateKeyPath)
	if err != nil {
		panic("unable to load private key")
	}

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			LocalAgentHostPort:  jaegerHostPort,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	tracer, closer, err := cfg.New(
		"login",
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		panic(err)
	}

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	s := newServer()
	s.start()
}
