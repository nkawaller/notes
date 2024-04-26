package main

import (
	"os"
)

func generateStaticSite() error {
	err := os.MkdirAll("deploy/static", 0755)
	if err != nil {
		return err
	}
	return nil
}
