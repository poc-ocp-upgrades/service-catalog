package apiserver

type RunnableServer interface{ Run(<-chan struct{}) error }
