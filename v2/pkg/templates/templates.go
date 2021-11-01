//go:generate dstdocgen -path "" -structure Template -output templates_doc.go -package templates
package templates

import (
	"github.com/projectdiscovery/nuclei/v2/pkg/model"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/dns"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/file"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/headless"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/http"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/network"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/others/ssl"
	"github.com/projectdiscovery/nuclei/v2/pkg/protocols/others/websocket"
	"github.com/projectdiscovery/nuclei/v2/pkg/workflows"
)

// Template is a YAML input file which defines all the requests and
// other metadata for a template.
type Template struct {
	// description: |
	//   ID is the unique id for the template.
	//
	//   #### Good IDs
	//
	//   A good ID uniquely identifies what the requests in the template
	//   are doing. Let's say you have a template that identifies a git-config
	//   file on the webservers, a good name would be `git-config-exposure`. Another
	//   example name is `azure-apps-nxdomain-takeover`.
	// examples:
	//   - name: ID Example
	//     value: "\"CVE-2021-19520\""
	ID string `yaml:"id" jsonschema:"title=id of the template,description=The Unique ID for the template,example=cve-2021-19520,pattern=^([a-zA-Z0-9]+[-_])*[a-zA-Z0-9]+$"`
	// description: |
	//   Info contains metadata information about the template.
	// examples:
	//   - value: exampleInfoStructure
	Info model.Info `yaml:"info" jsonschema:"title=info for the template,description=Info contains metadata for the template"`
	// description: |
	//   Requests contains the http request to make in the template.
	// examples:
	//   - value: exampleNormalHTTPRequest
	RequestsHTTP []*http.Request `yaml:"requests,omitempty" json:"requests,omitempty" jsonschema:"title=http requests to make,description=HTTP requests to make for the template"`
	// description: |
	//   DNS contains the dns request to make in the template
	// examples:
	//   - value: exampleNormalDNSRequest
	RequestsDNS []*dns.Request `yaml:"dns,omitempty" json:"dns,omitempty" jsonschema:"title=dns requests to make,description=DNS requests to make for the template"`
	// description: |
	//   File contains the file request to make in the template
	// examples:
	//   - value: exampleNormalFileRequest
	RequestsFile []*file.Request `yaml:"file,omitempty" json:"file,omitempty" jsonschema:"title=file requests to make,description=File requests to make for the template"`
	// description: |
	//   Network contains the network request to make in the template
	// examples:
	//   - value: exampleNormalNetworkRequest
	RequestsNetwork []*network.Request `yaml:"network,omitempty" json:"network,omitempty" jsonschema:"title=network requests to make,description=Network requests to make for the template"`
	// description: |
	//   Headless contains the headless request to make in the template.
	RequestsHeadless []*headless.Request `yaml:"headless,omitempty" json:"headless,omitempty" jsonschema:"title=headless requests to make,description=Headless requests to make for the template"`
	// description: |
	//   SSL contains the SSL request to make in the template.
	RequestsSSL []*ssl.Request `yaml:"ssl,omitempty" json:"ssl,omitempty" jsonschema:"title=ssl requests to make,description=SSL requests to make for the template"`
	// description: |
	//   Websocket contains the Websocket request to make in the template.
	RequestsWebsocket []*websocket.Request `yaml:"websocket,omitempty" json:"websocket,omitempty" jsonschema:"title=websocket requests to make,description=Websocket requests to make for the template"`

	// description: |
	//   Workflows is a yaml based workflow declaration code.
	workflows.Workflow `yaml:",inline,omitempty" jsonschema:"title=workflows to run,description=Workflows to run for the template"`
	CompiledWorkflow   *workflows.Workflow `yaml:"-" json:"-" jsonschema:"-"`

	// description: |
	//   Self Contained marks Requests for the template as self-contained
	SelfContained bool `yaml:"self-contained,omitempty" jsonschema:"title=mark requests as self-contained,description=Mark Requests for the template as self-contained"`

	// TotalRequests is the total number of requests for the template.
	TotalRequests int `yaml:"-" json:"-"`
	// Executer is the actual template executor for running template requests
	Executer protocols.Executer `yaml:"-" json:"-"`

	Path string `yaml:"-" json:"-"`
}

// TemplateTypes is a list of accepted template types
var TemplateTypes = []string{
	"dns",
	"file",
	"http",
	"headless",
	"network",
	"workflow",
	"ssl",
	"websocket",
}

// Type returns the type of the template
func (t *Template) Type() string {
	switch {
	case len(t.RequestsDNS) > 0:
		return "dns"
	case len(t.RequestsFile) > 0:
		return "file"
	case len(t.RequestsHTTP) > 0:
		return "http"
	case len(t.RequestsHeadless) > 0:
		return "headless"
	case len(t.RequestsNetwork) > 0:
		return "network"
	case len(t.Workflows) > 0:
		return "workflow"
	case len(t.RequestsSSL) > 0:
		return "ssl"
	case len(t.RequestsWebsocket) > 0:
		return "websocket"
	default:
		return ""
	}
}
