// Copyright 2021 The GoLand Authors.
// Author: Spume
// Modified: 2022/4/22

package miniocli

type MinioCli struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	WebPrefix       string
}

func NewMinio(endpoint string, accessKeyID string, secretAccessKey string) *MinioCli {
	return &MinioCli{
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		WebPrefix:       "",
	}
}
