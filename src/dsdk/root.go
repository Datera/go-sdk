package dsdk

type RootEp struct {
	Path string
}

var Cpool *ConnectionPool

func NewRootEp(hostname, port, username, password, apiVersion, tenant, timeout string, headers map[string]string, secure bool) (*RootEp, error) {
	var err error
	//Initialize global connection object
	Cpool, err = NewConnPool(hostname, port, username, password, apiVersion, tenant, timeout, headers, secure)
	if err != nil {
		return nil, err
	}
	conn := Cpool.GetConn()
	defer Cpool.ReleaseConn(conn)
	// err = conn.Login()
	// if err != nil {
	// 	return nil, err
	// }
	return &RootEp{
		Path: "",
	}, nil

}

func (ep RootEp) GetEp(path string) IEndpoint {
	return NewEp(ep.Path, path)
}
