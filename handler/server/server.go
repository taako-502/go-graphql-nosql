package server

type server struct {
	corsAllowedOrigins []string
	dynamoEndpoint     string
	graphqlServerPort  string
	authServerURL      string
	awsConfig          *awsConfig
}

type awsConfig struct {
	region   string
	s3Bucket string
}

func NewServer(awsRegion, s3Bucket string, corsAllowedOrigins []string, dynamoEndpoint, graphqlServerPort, authServerURL string) *server {
	return &server{
		corsAllowedOrigins: corsAllowedOrigins,
		dynamoEndpoint:     dynamoEndpoint,
		graphqlServerPort:  graphqlServerPort,
		authServerURL:      authServerURL,
		awsConfig: &awsConfig{
			region:   awsRegion,
			s3Bucket: s3Bucket,
		},
	}
}
