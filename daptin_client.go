package daptin_go_client

type JsonApiObject struct {
}

type DaptinQueryParameters struct {
}

type DaptinActionResponse struct {
}

type DaptinClient interface {
	FindOne(tableName string, referenceId string) (JsonApiObject, error)
	FindAll(...DaptinQueryParameters) ([]JsonApiObject, error)
	Create(object JsonApiObject) (JsonApiObject, error)
	Update(object JsonApiObject) error
	Delete(tableName string, referenceId string) error
	Execute(actionName string, tableName string, parameters map[string]interface{}) (DaptinActionResponse, error)
	Authorize(username string, password string) error
}

func NewDaptinClient(endpoint string) DaptinClient {

	return &daptinClientImpl{
		endpoint: endpoint,
	}
}

func NewDaptinClientWithAuthToken(endpoint string, authToken string) DaptinClient {

	return &daptinClientImpl{
		endpoint:  endpoint,
		authToken: authToken,
	}
}
