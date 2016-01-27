package main

type Params struct {
	AccessKey         string `json:"access_key"`
	SecretKey         string `json:"secret_key"`
	Region            string `json:"region"`
	Application       string `json:"application"`
	VersionLabel      string `json:"version_label"`
	Description       string `json:"description"`
	AutoCreate        bool   `json:"auto_create"`
	Process           bool   `json:"process"`
	BucketName        string `json:"bucket_name"`
	BucketKey         string `json:"bucket_key"`
	EnvironmentName   string `json:"environment_name"`
	UpdateEnvironment bool   `json:"update_environment"`
}
