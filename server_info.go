package jira

import "context"

// ServerInfoService handles staties for the Jira instance / API.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-server-info/#api-group-server-info
type ServerInfoService struct {
	client *Client
}

type ServerInfo struct {
	// The base URL of the Jira instance.
	BaseURL string `json:"baseUrl"`
	// The timestamp when the Jira version was built.
	BuildDate string `json:"buildDate"`
	// The build number of the Jira version.
	BuildNumber int `json:"buildNumber"`
	// The unique identifier of the Jira version.
	ScmInfo string `json:"scmInfo"`
	// The time in Jira when this request was responded to.
	ServerTime string `json:"serverTime"`
	// The name of the Jira instance.
	ServerTitle string `json:"serverTitle"`
	// The version of Jira.
	Version string `json:"version"`
	// The major, minor, and revision version numbers of the Jira version.
	VersionNumbers []int `json:"versionNumbers"`
}

// GetServerInfoWithContext returns information about the responding jira server
//
// https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-server-info/#api-group-server-info
func (s ServerInfoService) GetServerInfoWithContext(ctx context.Context) (*ServerInfo, *Response, error) {
	apiEndpoint := "rest/api/2/serverInfo"
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	serverInfo := ServerInfo{}
	resp, err := s.client.Do(req, &serverInfo)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}

	return &serverInfo, resp, nil
}

// GetServerInfo wraps GetServerInfoWithContext using the background context.
func (s *ServerInfoService) GetServerInfo() (*ServerInfo, *Response, error) {
	return s.GetServerInfoWithContext(context.Background())
}
