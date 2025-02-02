package consul

import (
	fmt "fmt"
	net "net"

	consul "github.com/hashicorp/consul/api"
)

type Client interface {
	// Get a Service from consul
	Service(string, string) ([]*consul.ServiceEntry, *consul.QueryMeta, error)
	// Register a service with local agent
	Register(string, int) error
	// Deregister a service with local agent
	DeRegister(string) error
}

type client struct {
	consul *consul.Client
}

//NewConsul returns a Client interface for given consul address
func NewConsulClient(addr string) (Client, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &client{consul: c}, nil
}

// Register a service with consul local agent
func (c *client) Register(name string, port int) error {
	HOST, err := getLocalIP()
	if err != nil {
		return err
	}
	script := "curl http://10.225.241.195:8082/health"
	reg := &consul.AgentServiceRegistration{
		ID:      name,
		Name:    name,
		Port:    port,
		Address: HOST,
		Check: &consul.AgentServiceCheck{
			Script:   script,
			Interval: "10s",
		},
	}
	err = c.consul.Agent().ServiceRegister(reg)
	if err != nil {
		return err
	}
	return nil
}

// DeRegister a service with consul local agent
func (c *client) DeRegister(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}

// Service return a service
func (c *client) Service(service string, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	passingOnly := true
	addrs, meta, err := c.consul.Health().Service(service, tag, passingOnly, nil)
	if len(addrs) == 0 && err == nil {
		return nil, nil, fmt.Errorf("service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, nil, err
	}
	return addrs, meta, nil
}

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", err
}
