package utils

type Configurations struct {
	Version                     string `required:"true"`
	AgentDeployments            string `required:"true" split_words:"true"`
	DbServer                    string `required:"true" split_words:"true"`
	JwtSecret                   string `required:"true" split_words:"true"`
	SelfAgent                   string `required:"true" split_words:"true"`
	AgentScope                  string `required:"true" split_words:"true"`
	AgentNamespace              string `required:"true" split_words:"true"`
	LitmusPortalNamespace       string `required:"true" split_words:"true"`
	DbUser                      string `required:"true" split_words:"true"`
	DbPassword                  string `required:"true" split_words:"true"`
	ChaosCenterScope            string `required:"true" split_words:"true"`
	SubscriberImage             string `required:"true" split_words:"true"`
	EventTrackerImage           string `required:"true" split_words:"true"`
	ArgoWorkflowControllerImage string `required:"true" split_words:"true"`
	ArgoWorkflowExecutorImage   string `required:"true" split_words:"true"`
	LitmusChaosOperatorImage    string `required:"true" split_words:"true"`
	LitmusChaosRunnerImage      string `required:"true" split_words:"true"`
	LitmusChaosExporterImage    string `required:"true" split_words:"true"`
	ContainerRuntimeExecutor    string `required:"true" split_words:"true"`
	HubBranchName               string `required:"true" split_words:"true"`
	WorkflowHelperImageVersion  string `required:"true" split_words:"true"`
	ServerServiceName           string `split_words:"true"`
	NodeName                    string `split_words:"true"`
	Ingress                     string `split_words:"true"`
	IngressName                 string `split_words:"true"`
	ChaosCenterUiEndpoint       string `split_words:"true"`
	TlsCertB64                  string `split_words:"true"`
	TlsSecretName               string `split_words:"true"`
	LitmusAuthGrpcEndpoint      string `split_words:"true" default:"localhost"`
	LitmusAuthGrpcPort          string `split_words:"true" default:":3030"`
	KubeConfigFilePath          string `split_words:"true"`
	RemoteHubMaxSize            string `split_words:"true"`
	SkipSslVerify               string `split_words:"true"`
	SelfAgentNodeSelector       string `split_words:"true"`
	SelfAgentTolerations        string `split_words:"true"`
	HttpPort                    string `split_words:"true" default:"8080"`
	RpcPort                     string `split_words:"true" default:"8000"`
}

var Config Configurations
