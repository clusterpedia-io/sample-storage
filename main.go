package main

import (
	plugin "github.com/clusterpedia-io/sample-storage/storage"
)

func init() {
	plugin.RegisterStorageLayer()
}
