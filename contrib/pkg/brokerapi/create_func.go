package brokerapi

type CreateFunc func(name, url, username, password string) BrokerClient
