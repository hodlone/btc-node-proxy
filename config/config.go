package config

import (
	"log"
	"os"
)

//EnvVar ...
type EnvVar struct {
	Key   string
	Value string
}

func (envVar *EnvVar) expected() string {
	if envVar.Value == "" {
		log.Panicf("Could not get %v environment variable", envVar.Key)
	}
	return envVar.Value
}

func (envVar *EnvVar) optional() string {
	return envVar.Value
}

func getEnv(name string) *EnvVar {
	return &EnvVar{
		Key:   name,
		Value: os.Getenv(name),
	}
}

var (
	// ServiceName ...
	ServiceName string = os.Getenv("SERVICE_NAME")
	// ServerPort ...
	ServerPort string = os.Getenv("PORT")
	// NatsAddr ...
	NatsAddr string = os.Getenv("NATS_ADDR")
	//NatsName ...
	NatsName = os.Getenv("NATS_NAME")
	// GrpcServerPort ...
	GrpcServerPort string = os.Getenv("GRPC_SERVER_PORT")
	// StanClusterID is the STAN cluster id
	StanClusterID string = os.Getenv("STAN_CLUSTER_ID")

	// Redis cache configuration
	Redis = struct {
		Addr string
		Pass string
	}{
		getEnv("REDIS_ADDR").optional(),
		getEnv("REDIS_PASS").optional(),
	}
)
