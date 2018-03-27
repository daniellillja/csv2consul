package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	cfg *api.Config
}

func NewConsulClient() ConsulClient {
	consulConfig := api.DefaultConfig()
	return ConsulClient{
		cfg: consulConfig,
	}
}

func (c ConsulClient) EnvironmentKvpPut(kvpEnv EnvironmentKvp) error {
	log.Printf("Connecting to consul: address=%s\n", c.cfg.Address)
	client, err := api.NewClient(c.cfg)
	if err != nil {
		return err
	}

	kv := api.KVPair{
		Key:   fmt.Sprintf("%s/%s", kvpEnv.Environment, kvpEnv.Key),
		Value: []byte(kvpEnv.Value),
	}
	wOpts := api.WriteOptions{}
	log.Printf("Saving kv pair: key=%s, value=%s\n", kv.Key, string(kv.Value))
	_, err = client.KV().Put(&kv, &wOpts)
	if err != nil {
		return err
	}

	return nil
}
