package daptin_go_client

type daptinClientImpl struct {
	endpoint string
	authToken string
}

func (d daptinClientImpl) FindOne(tableName string, referenceId string) (JsonApiObject, error) {
	panic("implement me")
}

func (d daptinClientImpl) FindAll(parameters ...DaptinQueryParameters) ([]JsonApiObject, error) {
	panic("implement me")
}

func (d daptinClientImpl) Create(object JsonApiObject) (JsonApiObject, error) {
	panic("implement me")
}

func (d daptinClientImpl) Update(object JsonApiObject) error {
	panic("implement me")
}

func (d daptinClientImpl) Delete(tableName string, referenceId string) error {
	panic("implement me")
}

func (d daptinClientImpl) Execute(actionName string, tableName string, parameters map[string]interface{}) (DaptinActionResponse, error) {
	panic("implement me")
}

func (d daptinClientImpl) Authorize(username string, password string) error {
	panic("implement me")
}
