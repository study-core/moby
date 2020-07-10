package client

import "net/http"

// NewClient initializes a new API client for the given host and API version.
// It uses the given http client as transport.
// It also initializes the custom http headers to add to each request.
//
// It won't send any version information if the version number is empty. It is
// highly recommended that you set a version or your client may break if the
// server is upgraded.
// Deprecated: use NewClientWithOpts
//
//
// todo NewClient:
//		为给定的主机和API版本初始化一个新的API客户端。
//		它使用给定的http客户端作为传输。
//		它还初始化要添加到每个请求的自定义http标头。
//
//		如果版本号为空，则不会发送任何版本信息。 强烈建议您设置版本，否则如果升级服务器，客户端可能会损坏。
//		不推荐使用：使用NewClientWithOpts
func NewClient(host string, version string, client *http.Client, httpHeaders map[string]string) (*Client, error) {
	return NewClientWithOpts(WithHost(host), WithVersion(version), WithHTTPClient(client), WithHTTPHeaders(httpHeaders))
}

// NewEnvClient initializes a new API client based on environment variables.
// See FromEnv for a list of support environment variables.
//
// Deprecated: use NewClientWithOpts(FromEnv)
func NewEnvClient() (*Client, error) {
	return NewClientWithOpts(FromEnv)
}
