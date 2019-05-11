package apiserver

type Config interface{ Complete() CompletedConfig }
type CompletedConfig interface {
	NewServer(stopCh <-chan struct{}) (*ServiceCatalogAPIServer, error)
}
