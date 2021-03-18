package proxyplane

type ProxyPlane struct {
	HTTPProxy *HTTPProxy
}

func MustNewProxyPlane() *ProxyPlane {
	return &ProxyPlane{
		HTTPProxy: MustNewHTTPProxy(),
	}
}

func (pp *ProxyPlane) Stop() {
	pp.HTTPProxy.Stop()
}

func (pp *ProxyPlane) Start() {
	go pp.HTTPProxy.Start()
}
