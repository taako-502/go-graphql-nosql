package server

type server struct {
	corsAllowedOrigins []string
	dynamoEndpoint     string
	graphqlServerPort  string
	awsConfig          *awsConfig
}

type awsConfig struct {
	region string
}

func NewServerForLocal(awsRegion string, corsAllowedOrigins []string, dynamoEndpoint, graphqlServerPort string) *server {
	return &server{
		corsAllowedOrigins: corsAllowedOrigins,
		dynamoEndpoint:     dynamoEndpoint,
		graphqlServerPort:  graphqlServerPort,
		awsConfig: &awsConfig{
			region: awsRegion,
		},
	}
}

func NewServerForLambda(awsRegion string, corsAllowedOrigins []string) *server {
	return &server{
		corsAllowedOrigins: corsAllowedOrigins,
		awsConfig: &awsConfig{
			region: awsRegion,
		},
	}
}
