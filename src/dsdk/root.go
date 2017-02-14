package dsdk

type Client struct {
}

var Cpool *ConnectionPool

func NewClient(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*Client, error) {
	var err error
	//Initialize global connection object
	Cpool, err = NewConnPool(hostname, port, username, password, apiVersion, tenant, timeout, headers, secure)
	if err != nil {
		return nil, err
	}
	conn := Cpool.GetConn()
	defer Cpool.ReleaseConn(conn)
	return &Client{}, nil

}

func (ep Client) GetEp(path string) IEndpoint {
	return NewEp("", path)
}
