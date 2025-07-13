package cloud

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (c *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (c *CloudDb) Write(data []byte) {
}
