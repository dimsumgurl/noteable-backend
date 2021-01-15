package mongo

type User struct {
	Email        string
	PasswordHash string
}

func CreateCollection() error {
	return nil
}

func (c *MongoClient) Create() error {
	return nil
}

func (c *MongoClient) RetrieveByID() error {
	return nil
}

func (c *MongoClient) RetrieveByEmail() error {
	return nil
}

func (c *MongoClient) Update() error {
	return nil
}

func (c *MongoClient) Delete() error {
	return nil
}
