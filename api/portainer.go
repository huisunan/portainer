package portainer

import (
	"context"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/volume"
	"github.com/portainer/portainer/api/database/models"
	gittypes "github.com/portainer/portainer/api/git/types"
	v1 "k8s.io/api/core/v1"
)

type (

	// AgentPlatform represents a platform type for an Agent
	AgentPlatform int

	// AzureCredentials represents the credentials used to connect to an Azure
	// environment(endpoint).
	AzureCredentials struct {
		// Azure application ID
		ApplicationID string `json:"ApplicationID" example:"eag7cdo9-o09l-9i83-9dO9-f0b23oe78db4"`
		// Azure tenant ID
		TenantID string `json:"TenantID" example:"34ddc78d-4fel-2358-8cc1-df84c8o839f5"`
		// Azure authentication key
		AuthenticationKey string `json:"AuthenticationKey" example:"cOrXoK/1D35w8YQ8nH1/8ZGwzz45JIYD5jxHKXEQknk="`
	}

	// CLIFlags represents the available flags on the CLI
	CLIFlags struct {
		Addr                      *string
		AddrHTTPS                 *string
		TunnelAddr                *string
		TunnelPort                *string
		AdminPassword             *string
		AdminPasswordFile         *string
		Assets                    *string
		Data                      *string
		FeatureFlags              *[]models.Pair
		DemoEnvironment           *bool
		EnableEdgeComputeFeatures *bool
		EndpointURL               *string
		Labels                    *[]models.Pair
		Logo                      *string
		NoAnalytics               *bool
		Templates                 *string
		TLS                       *bool
		TLSSkipVerify             *bool
		TLSCacert                 *string
		TLSCert                   *string
		TLSKey                    *string
		HTTPDisabled              *bool
		HTTPEnabled               *bool
		SSL                       *bool
		SSLCert                   *string
		SSLKey                    *string
		Rollback                  *bool
		SnapshotInterval          *string
		BaseURL                   *string
		InitialMmapSize           *int
		MaxBatchSize              *int
		MaxBatchDelay             *time.Duration
		SecretKeyName             *string
		LogLevel                  *string
	}

	// CustomTemplateVariableDefinition
	CustomTemplateVariableDefinition struct {
		Name         string `json:"name" example:"MY_VAR"`
		Label        string `json:"label" example:"My Variable"`
		DefaultValue string `json:"defaultValue" example:"default value"`
		Description  string `json:"description" example:"Description"`
	}

	// CustomTemplate represents a custom template
	CustomTemplate struct {
		// CustomTemplate Identifier
		ID CustomTemplateID `json:"Id" example:"1"`
		// Title of the template
		Title string `json:"Title" example:"Nginx"`
		// Description of the template
		Description string `json:"Description" example:"High performance web server"`
		// Path on disk to the repository hosting the Stack file
		ProjectPath string `json:"ProjectPath" example:"/data/custom_template/3"`
		// Path to the Stack file
		EntryPoint string `json:"EntryPoint" example:"docker-compose.yml"`
		// User identifier who created this template
		CreatedByUserID models.UserID `json:"CreatedByUserId" example:"3"`
		// A note that will be displayed in the UI. Supports HTML content
		Note string `json:"Note" example:"This is my <b>custom</b> template"`
		// Platform associated to the template.
		// Valid values are: 1 - 'linux', 2 - 'windows'
		Platform CustomTemplatePlatform `json:"Platform" example:"1" enums:"1,2"`
		// URL of the template's logo
		Logo string `json:"Logo" example:"https://cloudinovasi.id/assets/img/logos/nginx.png"`
		// Type of created stack (1 - swarm, 2 - compose)
		Type            StackType        `json:"Type" example:"1"`
		ResourceControl *ResourceControl `json:"ResourceControl"`
		Variables       []CustomTemplateVariableDefinition
	}

	// CustomTemplateID represents a custom template identifier
	CustomTemplateID int

	// CustomTemplatePlatform represents a custom template platform
	CustomTemplatePlatform int

	// DockerHub represents all the required information to connect and use the
	// Docker Hub
	DockerHub struct {
		// Is authentication against DockerHub enabled
		Authentication bool `json:"Authentication" example:"true"`
		// Username used to authenticate against the DockerHub
		Username string `json:"Username" example:"user"`
		// Password used to authenticate against the DockerHub
		Password string `json:"Password,omitempty" example:"passwd"`
	}

	// DockerSnapshot represents a snapshot of a specific Docker environment(endpoint) at a specific time
	DockerSnapshot struct {
		Time                    int64             `json:"Time"`
		DockerVersion           string            `json:"DockerVersion"`
		Swarm                   bool              `json:"Swarm"`
		TotalCPU                int               `json:"TotalCPU"`
		TotalMemory             int64             `json:"TotalMemory"`
		RunningContainerCount   int               `json:"RunningContainerCount"`
		StoppedContainerCount   int               `json:"StoppedContainerCount"`
		HealthyContainerCount   int               `json:"HealthyContainerCount"`
		UnhealthyContainerCount int               `json:"UnhealthyContainerCount"`
		VolumeCount             int               `json:"VolumeCount"`
		ImageCount              int               `json:"ImageCount"`
		ServiceCount            int               `json:"ServiceCount"`
		StackCount              int               `json:"StackCount"`
		SnapshotRaw             DockerSnapshotRaw `json:"DockerSnapshotRaw"`
		NodeCount               int               `json:"NodeCount"`
		GpuUseAll               bool              `json:"GpuUseAll"`
		GpuUseList              []string          `json:"GpuUseList"`
	}

	// DockerSnapshotRaw represents all the information related to a snapshot as returned by the Docker API

	DockerSnapshotRaw struct {
		Containers []types.Container       `json:"Containers" swaggerignore:"true"`
		Volumes    volume.VolumeListOKBody `json:"Volumes" swaggerignore:"true"`
		Networks   []types.NetworkResource `json:"Networks" swaggerignore:"true"`
		Images     []types.ImageSummary    `json:"Images" swaggerignore:"true"`
		Info       types.Info              `json:"Info" swaggerignore:"true"`
		Version    types.Version           `json:"Version" swaggerignore:"true"`
	}

	// EdgeGroup represents an Edge group
	EdgeGroup struct {
		// EdgeGroup Identifier
		ID           EdgeGroupID  `json:"Id" example:"1"`
		Name         string       `json:"Name"`
		Dynamic      bool         `json:"Dynamic"`
		TagIDs       []TagID      `json:"TagIds"`
		Endpoints    []EndpointID `json:"Endpoints"`
		PartialMatch bool         `json:"PartialMatch"`
	}

	// EdgeGroupID represents an Edge group identifier
	EdgeGroupID int

	// EdgeJob represents a job that can run on Edge environments(endpoints).
	EdgeJob struct {
		// EdgeJob Identifier
		ID             EdgeJobID                          `json:"Id" example:"1"`
		Created        int64                              `json:"Created"`
		CronExpression string                             `json:"CronExpression"`
		Endpoints      map[EndpointID]EdgeJobEndpointMeta `json:"Endpoints"`
		Name           string                             `json:"Name"`
		ScriptPath     string                             `json:"ScriptPath"`
		Recurring      bool                               `json:"Recurring"`
		Version        int                                `json:"Version"`
	}

	// EdgeJobEndpointMeta represents a meta data object for an Edge job and Environment(Endpoint) relation
	EdgeJobEndpointMeta struct {
		LogsStatus  EdgeJobLogsStatus
		CollectLogs bool
	}

	// EdgeJobID represents an Edge job identifier
	EdgeJobID int

	// EdgeJobLogsStatus represent status of logs collection job
	EdgeJobLogsStatus int

	// EdgeSchedule represents a scheduled job that can run on Edge environments(endpoints).
	//
	// Deprecated: in favor of EdgeJob
	EdgeSchedule struct {
		// EdgeSchedule Identifier
		ID             ScheduleID   `json:"Id" example:"1"`
		CronExpression string       `json:"CronExpression"`
		Script         string       `json:"Script"`
		Version        int          `json:"Version"`
		Endpoints      []EndpointID `json:"Endpoints"`
	}

	//EdgeStack represents an edge stack
	EdgeStack struct {
		// EdgeStack Identifier
		ID             EdgeStackID                    `json:"Id" example:"1"`
		Name           string                         `json:"Name"`
		Status         map[EndpointID]EdgeStackStatus `json:"Status"`
		CreationDate   int64                          `json:"CreationDate"`
		EdgeGroups     []EdgeGroupID                  `json:"EdgeGroups"`
		ProjectPath    string                         `json:"ProjectPath"`
		EntryPoint     string                         `json:"EntryPoint"`
		Version        int                            `json:"Version"`
		ManifestPath   string
		DeploymentType EdgeStackDeploymentType

		// Deprecated
		Prune bool `json:"Prune"`
	}

	EdgeStackDeploymentType int

	//EdgeStackID represents an edge stack id
	EdgeStackID int

	//EdgeStackStatus represents an edge stack status
	EdgeStackStatus struct {
		Type       EdgeStackStatusType `json:"Type"`
		Error      string              `json:"Error"`
		EndpointID EndpointID          `json:"EndpointID"`
	}

	//EdgeStackStatusType represents an edge stack status type
	EdgeStackStatusType int

	// Environment(Endpoint) represents a Docker environment(endpoint) with all the info required
	// to connect to it
	Endpoint struct {
		// Environment(Endpoint) Identifier
		ID EndpointID `json:"Id" example:"1"`
		// Environment(Endpoint) name
		Name string `json:"Name" example:"my-environment"`
		// Environment(Endpoint) environment(endpoint) type. 1 for a Docker environment(endpoint), 2 for an agent on Docker environment(endpoint) or 3 for an Azure environment(endpoint).
		Type EndpointType `json:"Type" example:"1"`
		// URL or IP address of the Docker host associated to this environment(endpoint)
		URL string `json:"URL" example:"docker.mydomain.tld:2375"`
		// Environment(Endpoint) group identifier
		GroupID EndpointGroupID `json:"GroupId" example:"1"`
		// URL or IP address where exposed containers will be reachable
		PublicURL        string                  `json:"PublicURL" example:"docker.mydomain.tld:2375"`
		Gpus             []models.Pair           `json:"Gpus"`
		TLSConfig        models.TLSConfiguration `json:"TLSConfig"`
		AzureCredentials AzureCredentials        `json:"AzureCredentials,omitempty" example:""`
		// List of tag identifiers to which this environment(endpoint) is associated
		TagIDs []TagID `json:"TagIds"`
		// The status of the environment(endpoint) (1 - up, 2 - down)
		Status EndpointStatus `json:"Status" example:"1"`
		// List of snapshots
		Snapshots []DockerSnapshot `json:"Snapshots" example:""`
		// List of user identifiers authorized to connect to this environment(endpoint)
		UserAccessPolicies UserAccessPolicies `json:"UserAccessPolicies"`
		// List of team identifiers authorized to connect to this environment(endpoint)
		TeamAccessPolicies models.TeamAccessPolicies `json:"TeamAccessPolicies" example:""`
		// The identifier of the edge agent associated with this environment(endpoint)
		EdgeID string `json:"EdgeID,omitempty" example:""`
		// The key which is used to map the agent to Portainer
		EdgeKey string `json:"EdgeKey" example:""`
		// The check in interval for edge agent (in seconds)
		EdgeCheckinInterval int `json:"EdgeCheckinInterval" example:"5"`
		// Associated Kubernetes data
		Kubernetes KubernetesData `json:"Kubernetes" example:""`
		// Maximum version of docker-compose
		ComposeSyntaxMaxVersion string `json:"ComposeSyntaxMaxVersion" example:"3.8"`
		// Environment(Endpoint) specific security settings
		SecuritySettings EndpointSecuritySettings
		// The identifier of the AMT Device associated with this environment(endpoint)
		AMTDeviceGUID string `json:"AMTDeviceGUID,omitempty" example:"4c4c4544-004b-3910-8037-b6c04f504633"`
		// LastCheckInDate mark last check-in date on checkin
		LastCheckInDate int64
		// QueryDate of each query with the endpoints list
		QueryDate int64
		// IsEdgeDevice marks if the environment was created as an EdgeDevice
		IsEdgeDevice bool
		// Whether the device has been trusted or not by the user
		UserTrusted bool

		Edge struct {
			// Whether the device has been started in edge async mode
			AsyncMode bool
			// The ping interval for edge agent - used in edge async mode [seconds]
			PingInterval int `json:"PingInterval" example:"60"`
			// The snapshot interval for edge agent - used in edge async mode [seconds]
			SnapshotInterval int `json:"SnapshotInterval" example:"60"`
			// The command list interval for edge agent - used in edge async mode [seconds]
			CommandInterval int `json:"CommandInterval" example:"60"`
		}

		Agent struct {
			Version string `example:"1.0.0"`
		}

		// Deprecated fields
		// Deprecated in DBVersion == 4
		TLS           bool   `json:"TLS,omitempty"`
		TLSCACertPath string `json:"TLSCACert,omitempty"`
		TLSCertPath   string `json:"TLSCert,omitempty"`
		TLSKeyPath    string `json:"TLSKey,omitempty"`

		// Deprecated in DBVersion == 18
		AuthorizedUsers []models.UserID `json:"AuthorizedUsers"`
		AuthorizedTeams []models.TeamID `json:"AuthorizedTeams"`

		// Deprecated in DBVersion == 22
		Tags []string `json:"Tags"`
	}

	// EndpointAuthorizations represents the authorizations associated to a set of environments(endpoints)
	EndpointAuthorizations map[EndpointID]models.Authorizations

	// EndpointGroup represents a group of environments(endpoints)
	EndpointGroup struct {
		// Environment(Endpoint) group Identifier
		ID EndpointGroupID `json:"Id" example:"1"`
		// Environment(Endpoint) group name
		Name string `json:"Name" example:"my-environment-group"`
		// Description associated to the environment(endpoint) group
		Description        string                    `json:"Description" example:"Environment(Endpoint) group description"`
		UserAccessPolicies UserAccessPolicies        `json:"UserAccessPolicies" example:""`
		TeamAccessPolicies models.TeamAccessPolicies `json:"TeamAccessPolicies" example:""`
		// List of tags associated to this environment(endpoint) group
		TagIDs []TagID `json:"TagIds"`

		// Deprecated fields
		Labels []models.Pair `json:"Labels"`

		// Deprecated in DBVersion == 18
		AuthorizedUsers []models.UserID `json:"AuthorizedUsers"`
		AuthorizedTeams []models.TeamID `json:"AuthorizedTeams"`

		// Deprecated in DBVersion == 22
		Tags []string `json:"Tags"`
	}

	// EndpointGroupID represents an environment(endpoint) group identifier
	EndpointGroupID int

	// EndpointID represents an environment(endpoint) identifier
	EndpointID int

	// EndpointStatus represents the status of an environment(endpoint)
	EndpointStatus int

	// EndpointSyncJob represents a scheduled job that synchronize environments(endpoints) based on an external file
	// Deprecated
	EndpointSyncJob struct{}

	// EndpointSecuritySettings represents settings for an environment(endpoint)
	EndpointSecuritySettings struct {
		// Whether non-administrator should be able to use bind mounts when creating containers
		AllowBindMountsForRegularUsers bool `json:"allowBindMountsForRegularUsers" example:"false"`
		// Whether non-administrator should be able to use privileged mode when creating containers
		AllowPrivilegedModeForRegularUsers bool `json:"allowPrivilegedModeForRegularUsers" example:"false"`
		// Whether non-administrator should be able to browse volumes
		AllowVolumeBrowserForRegularUsers bool `json:"allowVolumeBrowserForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use the host pid
		AllowHostNamespaceForRegularUsers bool `json:"allowHostNamespaceForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use device mapping
		AllowDeviceMappingForRegularUsers bool `json:"allowDeviceMappingForRegularUsers" example:"true"`
		// Whether non-administrator should be able to manage stacks
		AllowStackManagementForRegularUsers bool `json:"allowStackManagementForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use container capabilities
		AllowContainerCapabilitiesForRegularUsers bool `json:"allowContainerCapabilitiesForRegularUsers" example:"true"`
		// Whether non-administrator should be able to use sysctl settings
		AllowSysctlSettingForRegularUsers bool `json:"allowSysctlSettingForRegularUsers" example:"true"`
		// Whether host management features are enabled
		EnableHostManagementFeatures bool `json:"enableHostManagementFeatures" example:"true"`
	}

	// EndpointType represents the type of an environment(endpoint)
	EndpointType int

	// EndpointRelation represents a environment(endpoint) relation object
	EndpointRelation struct {
		EndpointID EndpointID
		EdgeStacks map[EdgeStackID]bool
	}

	// Extension represents a deprecated Portainer extension
	Extension struct {
		// Extension Identifier
		ID               ExtensionID        `json:"Id" example:"1"`
		Enabled          bool               `json:"Enabled"`
		Name             string             `json:"Name,omitempty"`
		ShortDescription string             `json:"ShortDescription,omitempty"`
		Description      string             `json:"Description,omitempty"`
		DescriptionURL   string             `json:"DescriptionURL,omitempty"`
		Price            string             `json:"Price,omitempty"`
		PriceDescription string             `json:"PriceDescription,omitempty"`
		Deal             bool               `json:"Deal,omitempty"`
		Available        bool               `json:"Available,omitempty"`
		License          LicenseInformation `json:"License,omitempty"`
		Version          string             `json:"Version"`
		UpdateAvailable  bool               `json:"UpdateAvailable"`
		ShopURL          string             `json:"ShopURL,omitempty"`
		Images           []string           `json:"Images,omitempty"`
		Logo             string             `json:"Logo,omitempty"`
	}

	// ExtensionID represents a extension identifier
	ExtensionID int

	// GitlabRegistryData represents data required for gitlab registry to work
	GitlabRegistryData struct {
		ProjectID   int    `json:"ProjectId"`
		InstanceURL string `json:"InstanceURL"`
		ProjectPath string `json:"ProjectPath"`
	}

	HelmUserRepositoryID int

	// HelmUserRepositories stores a Helm repository URL for the given user
	HelmUserRepository struct {
		// Membership Identifier
		ID HelmUserRepositoryID `json:"Id" example:"1"`
		// User identifier
		UserID models.UserID `json:"UserId" example:"1"`
		// Helm repository URL
		URL string `json:"URL" example:"https://charts.bitnami.com/bitnami"`
	}

	// QuayRegistryData represents data required for Quay registry to work
	QuayRegistryData struct {
		UseOrganisation  bool   `json:"UseOrganisation"`
		OrganisationName string `json:"OrganisationName"`
	}

	// EcrData represents data required for ECR registry
	EcrData struct {
		Region string `json:"Region" example:"ap-southeast-2"`
	}

	// JobType represents a job type
	JobType int

	K8sNamespaceInfo struct {
		IsSystem  bool `json:"IsSystem"`
		IsDefault bool `json:"IsDefault"`
	}

	K8sNodeLimits struct {
		CPU    int64 `json:"CPU"`
		Memory int64 `json:"Memory"`
	}

	K8sNodesLimits map[string]*K8sNodeLimits

	K8sNamespaceAccessPolicy struct {
		UserAccessPolicies UserAccessPolicies        `json:"UserAccessPolicies"`
		TeamAccessPolicies models.TeamAccessPolicies `json:"TeamAccessPolicies"`
	}

	// KubernetesData contains all the Kubernetes related environment(endpoint) information
	KubernetesData struct {
		Snapshots     []KubernetesSnapshot    `json:"Snapshots"`
		Configuration KubernetesConfiguration `json:"Configuration"`
	}

	// KubernetesSnapshot represents a snapshot of a specific Kubernetes environment(endpoint) at a specific time
	KubernetesSnapshot struct {
		Time              int64  `json:"Time"`
		KubernetesVersion string `json:"KubernetesVersion"`
		NodeCount         int    `json:"NodeCount"`
		TotalCPU          int64  `json:"TotalCPU"`
		TotalMemory       int64  `json:"TotalMemory"`
	}

	// KubernetesConfiguration represents the configuration of a Kubernetes environment(endpoint)
	KubernetesConfiguration struct {
		UseLoadBalancer                 bool                           `json:"UseLoadBalancer"`
		UseServerMetrics                bool                           `json:"UseServerMetrics"`
		EnableResourceOverCommit        bool                           `json:"EnableResourceOverCommit"`
		ResourceOverCommitPercentage    int                            `json:"ResourceOverCommitPercentage"`
		StorageClasses                  []KubernetesStorageClassConfig `json:"StorageClasses"`
		IngressClasses                  []KubernetesIngressClassConfig `json:"IngressClasses"`
		RestrictDefaultNamespace        bool                           `json:"RestrictDefaultNamespace"`
		IngressAvailabilityPerNamespace bool                           `json:"IngressAvailabilityPerNamespace"`
	}

	// KubernetesStorageClassConfig represents a Kubernetes Storage Class configuration
	KubernetesStorageClassConfig struct {
		Name                 string   `json:"Name"`
		AccessModes          []string `json:"AccessModes"`
		Provisioner          string   `json:"Provisioner"`
		AllowVolumeExpansion bool     `json:"AllowVolumeExpansion"`
	}

	// KubernetesIngressClassConfig represents a Kubernetes Ingress Class configuration
	KubernetesIngressClassConfig struct {
		Name              string   `json:"Name"`
		Type              string   `json:"Type"`
		GloballyBlocked   bool     `json:"Blocked"`
		BlockedNamespaces []string `json:"BlockedNamespaces"`
	}

	// KubernetesShellPod represents a Kubectl Shell details to facilitate pod exec functionality
	KubernetesShellPod struct {
		Namespace        string
		PodName          string
		ContainerName    string
		ShellExecCommand string
	}

	// LicenseInformation represents information about an extension license
	LicenseInformation struct {
		LicenseKey string `json:"LicenseKey,omitempty"`
		Company    string `json:"Company,omitempty"`
		Expiration string `json:"Expiration,omitempty"`
		Valid      bool   `json:"Valid,omitempty"`
	}

	// Registry represents a Docker registry with all the info required
	// to connect to it
	Registry struct {
		// Registry Identifier
		ID RegistryID `json:"Id" example:"1"`
		// Registry Type (1 - Quay, 2 - Azure, 3 - Custom, 4 - Gitlab, 5 - ProGet, 6 - DockerHub, 7 - ECR)
		Type RegistryType `json:"Type" enums:"1,2,3,4,5,6,7"`
		// Registry Name
		Name string `json:"Name" example:"my-registry"`
		// URL or IP address of the Docker registry
		URL string `json:"URL" example:"registry.mydomain.tld:2375"`
		// Base URL, introduced for ProGet registry
		BaseURL string `json:"BaseURL" example:"registry.mydomain.tld:2375"`
		// Is authentication against this registry enabled
		Authentication bool `json:"Authentication" example:"true"`
		// Username or AccessKeyID used to authenticate against this registry
		Username string `json:"Username" example:"registry user"`
		// Password or SecretAccessKey used to authenticate against this registry
		Password                string                           `json:"Password,omitempty" example:"registry_password"`
		ManagementConfiguration *RegistryManagementConfiguration `json:"ManagementConfiguration"`
		Gitlab                  GitlabRegistryData               `json:"Gitlab"`
		Quay                    QuayRegistryData                 `json:"Quay"`
		Ecr                     EcrData                          `json:"Ecr"`
		RegistryAccesses        RegistryAccesses                 `json:"RegistryAccesses"`

		// Deprecated fields
		// Deprecated in DBVersion == 31
		UserAccessPolicies UserAccessPolicies `json:"UserAccessPolicies"`
		// Deprecated in DBVersion == 31
		TeamAccessPolicies models.TeamAccessPolicies `json:"TeamAccessPolicies"`

		// Deprecated in DBVersion == 18
		AuthorizedUsers []models.UserID `json:"AuthorizedUsers"`
		// Deprecated in DBVersion == 18
		AuthorizedTeams []models.TeamID `json:"AuthorizedTeams"`

		// Stores temporary access token
		AccessToken       string `json:"AccessToken,omitempty"`
		AccessTokenExpiry int64  `json:"AccessTokenExpiry,omitempty"`
	}

	RegistryAccesses map[EndpointID]RegistryAccessPolicies

	RegistryAccessPolicies struct {
		UserAccessPolicies UserAccessPolicies        `json:"UserAccessPolicies"`
		TeamAccessPolicies models.TeamAccessPolicies `json:"TeamAccessPolicies"`
		Namespaces         []string                  `json:"Namespaces"`
	}

	// RegistryID represents a registry identifier
	RegistryID int

	// RegistryManagementConfiguration represents a configuration that can be used to query
	// the registry API via the registry management extension.
	RegistryManagementConfiguration struct {
		Type              RegistryType            `json:"Type"`
		Authentication    bool                    `json:"Authentication"`
		Username          string                  `json:"Username"`
		Password          string                  `json:"Password"`
		TLSConfig         models.TLSConfiguration `json:"TLSConfig"`
		Ecr               EcrData                 `json:"Ecr"`
		AccessToken       string                  `json:"AccessToken,omitempty"`
		AccessTokenExpiry int64                   `json:"AccessTokenExpiry,omitempty"`
	}

	// RegistryType represents a type of registry
	RegistryType int

	// ResourceControl represent a reference to a Docker resource with specific access controls
	ResourceControl struct {
		// ResourceControl Identifier
		ID ResourceControlID `json:"Id" example:"1"`
		// Docker resource identifier on which access control will be applied.\
		// In the case of a resource control applied to a stack, use the stack name as identifier
		ResourceID string `json:"ResourceId" example:"617c5f22bb9b023d6daab7cba43a57576f83492867bc767d1c59416b065e5f08"`
		// List of Docker resources that will inherit this access control
		SubResourceIDs []string `json:"SubResourceIds" example:"617c5f22bb9b023d6daab7cba43a57576f83492867bc767d1c59416b065e5f08"`
		// Type of Docker resource. Valid values are: 1- container, 2 -service
		// 3 - volume, 4 - secret, 5 - stack, 6 - config or 7 - custom template
		Type         ResourceControlType         `json:"Type" example:"1"`
		UserAccesses []UserResourceAccess        `json:"UserAccesses" example:""`
		TeamAccesses []models.TeamResourceAccess `json:"TeamAccesses" example:""`
		// Permit access to the associated resource to any user
		Public bool `json:"Public" example:"true"`
		// Permit access to resource only to admins
		AdministratorsOnly bool `json:"AdministratorsOnly" example:"true"`
		System             bool `json:"System" example:""`

		// Deprecated fields
		// Deprecated in DBVersion == 2
		OwnerID     models.UserID              `json:"OwnerId,omitempty"`
		AccessLevel models.ResourceAccessLevel `json:"AccessLevel,omitempty"`
	}

	// ResourceControlID represents a resource control identifier
	ResourceControlID int

	// ResourceControlType represents the type of resource associated to the resource control (volume, container, service...)
	ResourceControlType int

	// APIKeyID represents an API key identifier
	APIKeyID int

	// APIKey represents an API key
	APIKey struct {
		ID          APIKeyID      `json:"id" example:"1"`
		UserID      models.UserID `json:"userId" example:"1"`
		Description string        `json:"description" example:"portainer-api-key"`
		Prefix      string        `json:"prefix"`           // API key identifier (7 char prefix)
		DateCreated int64         `json:"dateCreated"`      // Unix timestamp (UTC) when the API key was created
		LastUsed    int64         `json:"lastUsed"`         // Unix timestamp (UTC) when the API key was last used
		Digest      []byte        `json:"digest,omitempty"` // Digest represents SHA256 hash of the raw API key
	}

	// Schedule represents a scheduled job.
	// It only contains a pointer to one of the JobRunner implementations
	// based on the JobType.
	// NOTE: The Recurring option is only used by ScriptExecutionJob at the moment
	// Deprecated in favor of EdgeJob
	Schedule struct {
		// Schedule Identifier
		ID             ScheduleID `json:"Id" example:"1"`
		Name           string
		CronExpression string
		Recurring      bool
		Created        int64
		JobType        JobType
		EdgeSchedule   *EdgeSchedule
	}

	// ScheduleID represents a schedule identifier.
	// Deprecated in favor of EdgeJob
	ScheduleID int

	// ScriptExecutionJob represents a scheduled job that can execute a script via a privileged container
	ScriptExecutionJob struct {
		Endpoints     []EndpointID
		Image         string
		ScriptPath    string
		RetryCount    int
		RetryInterval int
	}

	// SnapshotJob represents a scheduled job that can create environment(endpoint) snapshots
	SnapshotJob struct{}

	// SoftwareEdition represents an edition of Portainer
	SoftwareEdition int

	// SSLSettings represents a pair of SSL certificate and key
	SSLSettings struct {
		CertPath    string `json:"certPath"`
		KeyPath     string `json:"keyPath"`
		SelfSigned  bool   `json:"selfSigned"`
		HTTPEnabled bool   `json:"httpEnabled"`
	}

	// Stack represents a Docker stack created via docker stack deploy
	Stack struct {
		// Stack Identifier
		ID StackID `json:"Id" example:"1"`
		// Stack name
		Name string `json:"Name" example:"myStack"`
		// Stack type. 1 for a Swarm stack, 2 for a Compose stack
		Type StackType `json:"Type" example:"2"`
		// Environment(Endpoint) identifier. Reference the environment(endpoint) that will be used for deployment
		EndpointID EndpointID `json:"EndpointId" example:"1"`
		// Cluster identifier of the Swarm cluster where the stack is deployed
		SwarmID string `json:"SwarmId" example:"jpofkc0i9uo9wtx1zesuk649w"`
		// Path to the Stack file
		EntryPoint string `json:"EntryPoint" example:"docker-compose.yml"`
		// A list of environment(endpoint) variables used during stack deployment
		Env []models.Pair `json:"Env" example:""`
		//
		ResourceControl *ResourceControl `json:"ResourceControl" example:""`
		// Stack status (1 - active, 2 - inactive)
		Status StackStatus `json:"Status" example:"1"`
		// Path on disk to the repository hosting the Stack file
		ProjectPath string `example:"/data/compose/myStack_jpofkc0i9uo9wtx1zesuk649w"`
		// The date in unix time when stack was created
		CreationDate int64 `example:"1587399600"`
		// The username which created this stack
		CreatedBy string `example:"admin"`
		// The date in unix time when stack was last updated
		UpdateDate int64 `example:"1587399600"`
		// The username which last updated this stack
		UpdatedBy string `example:"bob"`
		// Only applies when deploying stack with multiple files
		AdditionalFiles []string `json:"AdditionalFiles"`
		// The auto update settings of a git stack
		AutoUpdate *StackAutoUpdate `json:"AutoUpdate"`
		// The stack deployment option
		Option *StackOption `json:"Option"`
		// The git config of this stack
		GitConfig *gittypes.RepoConfig
		// Whether the stack is from a app template
		FromAppTemplate bool `example:"false"`
		// Kubernetes namespace if stack is a kube application
		Namespace string `example:"default"`
		// IsComposeFormat indicates if the Kubernetes stack is created from a Docker Compose file
		IsComposeFormat bool `example:"false"`
	}

	//StackAutoUpdate represents the git auto sync config for stack deployment
	StackAutoUpdate struct {
		// Auto update interval
		Interval string `example:"1m30s"`
		// A UUID generated from client
		Webhook string `example:"05de31a2-79fa-4644-9c12-faa67e5c49f0"`
		// Autoupdate job id
		JobID string `example:"15"`
	}

	// StackOption represents the options for stack deployment
	StackOption struct {
		// Prune services that are no longer referenced
		Prune bool `example:"false"`
	}

	// StackID represents a stack identifier (it must be composed of Name + "_" + SwarmID to create a unique identifier)
	StackID int

	// StackStatus represent a status for a stack
	StackStatus int

	// StackType represents the type of the stack (compose v2, stack deploy v3)
	StackType int

	// Status represents the application status
	Status struct {
		// Portainer API version
		Version string `json:"Version" example:"2.0.0"`
		// Server Instance ID
		InstanceID string `example:"299ab403-70a8-4c05-92f7-bf7a994d50df"`
	}

	// Tag represents a tag that can be associated to a resource
	Tag struct {
		// Tag identifier
		ID TagID `example:"1"`
		// Tag name
		Name string `json:"Name" example:"org/acme"`
		// A set of environment(endpoint) ids that have this tag
		Endpoints map[EndpointID]bool `json:"Endpoints"`
		// A set of environment(endpoint) group ids that have this tag
		EndpointGroups map[EndpointGroupID]bool `json:"EndpointGroups"`
	}

	// TagID represents a tag identifier
	TagID int

	// Template represents an application template that can be used as an App Template
	// or an Edge template
	Template struct {
		// Mandatory container/stack fields
		// Template Identifier
		ID TemplateID `json:"Id" example:"1"`
		// Template type. Valid values are: 1 (container), 2 (Swarm stack) or 3 (Compose stack)
		Type TemplateType `json:"type" example:"1"`
		// Title of the template
		Title string `json:"title" example:"Nginx"`
		// Description of the template
		Description string `json:"description" example:"High performance web server"`
		// Whether the template should be available to administrators only
		AdministratorOnly bool `json:"administrator_only" example:"true"`

		// Mandatory container fields
		// Image associated to a container template. Mandatory for a container template
		Image string `json:"image" example:"nginx:latest"`

		// Mandatory stack fields
		Repository TemplateRepository `json:"repository"`

		// Mandatory Edge stack fields
		// Stack file used for this template
		StackFile string `json:"stackFile"`

		// Optional stack/container fields
		// Default name for the stack/container to be used on deployment
		Name string `json:"name,omitempty" example:"mystackname"`
		// URL of the template's logo
		Logo string `json:"logo,omitempty" example:"https://cloudinovasi.id/assets/img/logos/nginx.png"`
		// A list of environment(endpoint) variables used during the template deployment
		Env []TemplateEnv `json:"env,omitempty"`
		// A note that will be displayed in the UI. Supports HTML content
		Note string `json:"note,omitempty" example:"This is my <b>custom</b> template"`
		// Platform associated to the template.
		// Valid values are: 'linux', 'windows' or leave empty for multi-platform
		Platform string `json:"platform,omitempty" example:"linux"`
		// A list of categories associated to the template
		Categories []string `json:"categories,omitempty" example:"database"`

		// Optional container fields
		// The URL of a registry associated to the image for a container template
		Registry string `json:"registry,omitempty" example:"quay.io"`
		// The command that will be executed in a container template
		Command string `json:"command,omitempty" example:"ls -lah"`
		// Name of a network that will be used on container deployment if it exists inside the environment(endpoint)
		Network string `json:"network,omitempty" example:"mynet"`
		// A list of volumes used during the container template deployment
		Volumes []TemplateVolume `json:"volumes,omitempty"`
		// A list of ports exposed by the container
		Ports []string `json:"ports,omitempty" example:"8080:80/tcp"`
		// Container labels
		Labels []models.Pair `json:"labels,omitempty" example:""`
		// Whether the container should be started in privileged mode
		Privileged bool `json:"privileged,omitempty" example:"true"`
		// Whether the container should be started in
		// interactive mode (-i -t equivalent on the CLI)
		Interactive bool `json:"interactive,omitempty" example:"true"`
		// Container restart policy
		RestartPolicy string `json:"restart_policy,omitempty" example:"on-failure"`
		// Container hostname
		Hostname string `json:"hostname,omitempty" example:"mycontainer"`
	}

	// TemplateEnv represents a template environment(endpoint) variable configuration
	TemplateEnv struct {
		// name of the environment(endpoint) variable
		Name string `json:"name" example:"MYSQL_ROOT_PASSWORD"`
		// Text for the label that will be generated in the UI
		Label string `json:"label,omitempty" example:"Root password"`
		// Content of the tooltip that will be generated in the UI
		Description string `json:"description,omitempty" example:"MySQL root account password"`
		// Default value that will be set for the variable
		Default string `json:"default,omitempty" example:"default_value"`
		// If set to true, will not generate any input for this variable in the UI
		Preset bool `json:"preset,omitempty" example:"false"`
		// A list of name/value that will be used to generate a dropdown in the UI
		Select []TemplateEnvSelect `json:"select,omitempty"`
	}

	// TemplateEnvSelect represents text/value pair that will be displayed as a choice for the
	// template user
	TemplateEnvSelect struct {
		// Some text that will displayed as a choice
		Text string `json:"text" example:"text value"`
		// A value that will be associated to the choice
		Value string `json:"value" example:"value"`
		// Will set this choice as the default choice
		Default bool `json:"default" example:"false"`
	}

	// TemplateID represents a template identifier
	TemplateID int

	// TemplateRepository represents the git repository configuration for a template
	TemplateRepository struct {
		// URL of a git repository used to deploy a stack template. Mandatory for a Swarm/Compose stack template
		URL string `json:"url" example:"https://github.com/portainer/portainer-compose"`
		// Path to the stack file inside the git repository
		StackFile string `json:"stackfile" example:"./subfolder/docker-compose.yml"`
	}

	// TemplateType represents the type of a template
	TemplateType int

	// TemplateVolume represents a template volume configuration
	TemplateVolume struct {
		// Path inside the container
		Container string `json:"container" example:"/data"`
		// Path on the host
		Bind string `json:"bind,omitempty" example:"/tmp"`
		// Whether the volume used should be readonly
		ReadOnly bool `json:"readonly,omitempty" example:"true"`
	}

	// TLSFileType represents a type of TLS file required to connect to a Docker environment(endpoint).
	// It can be either a TLS CA file, a TLS certificate file or a TLS key file
	TLSFileType int

	// TokenData represents the data embedded in a JWT token
	TokenData struct {
		ID                  models.UserID
		Username            string
		Role                UserRole
		ForceChangePassword bool
	}

	// TunnelDetails represents information associated to a tunnel
	TunnelDetails struct {
		Status       string
		LastActivity time.Time
		Port         int
		Jobs         []EdgeJob
		Credentials  string
	}

	// TunnelServerInfo represents information associated to the tunnel server
	TunnelServerInfo struct {
		PrivateKeySeed string `json:"PrivateKeySeed"`
	}

	// User represents a user account
	User struct {
		// User Identifier
		ID       models.UserID `json:"Id" example:"1"`
		Username string        `json:"Username" example:"bob"`
		Password string        `json:"Password,omitempty" swaggerignore:"true"`
		// User Theme
		UserTheme string `example:"dark"`
		// User role (1 for administrator account and 2 for regular account)
		Role         UserRole `json:"Role" example:"1"`
		TokenIssueAt int64    `json:"TokenIssueAt" example:"1"`

		// Deprecated fields
		// Deprecated in DBVersion == 25
		PortainerAuthorizations models.Authorizations  `json:"PortainerAuthorizations"`
		EndpointAuthorizations  EndpointAuthorizations `json:"EndpointAuthorizations"`
	}

	// UserAccessPolicies represent the association of an access policy and a user
	UserAccessPolicies map[models.UserID]models.AccessPolicy

	// UserResourceAccess represents the level of control on a resource for a specific user
	UserResourceAccess struct {
		UserID      models.UserID              `json:"UserId"`
		AccessLevel models.ResourceAccessLevel `json:"AccessLevel"`
	}

	// UserRole represents the role of a user. It can be either an administrator
	// or a regular user
	UserRole int

	// Webhook represents a url webhook that can be used to update a service
	Webhook struct {
		// Webhook Identifier
		ID          WebhookID   `json:"Id" example:"1"`
		Token       string      `json:"Token"`
		ResourceID  string      `json:"ResourceId"`
		EndpointID  EndpointID  `json:"EndpointId"`
		RegistryID  RegistryID  `json:"RegistryId"`
		WebhookType WebhookType `json:"Type"`
	}

	// WebhookID represents a webhook identifier.
	WebhookID int

	// WebhookType represents the type of resource a webhook is related to
	WebhookType int

	Snapshot struct {
		EndpointID EndpointID          `json:"EndpointId"`
		Docker     *DockerSnapshot     `json:"Docker"`
		Kubernetes *KubernetesSnapshot `json:"Kubernetes"`
	}

	// CLIService represents a service for managing CLI
	CLIService interface {
		ParseFlags(version string) (*CLIFlags, error)
		ValidateFlags(flags *CLIFlags) error
	}

	// ComposeStackManager represents a service to manage Compose stacks
	ComposeStackManager interface {
		ComposeSyntaxMaxVersion() string
		NormalizeStackName(name string) string
		Up(ctx context.Context, stack *Stack, endpoint *Endpoint, forceRereate bool) error
		Down(ctx context.Context, stack *Stack, endpoint *Endpoint) error
		Pull(ctx context.Context, stack *Stack, endpoint *Endpoint) error
	}

	// CryptoService represents a service for encrypting/hashing data
	CryptoService interface {
		Hash(data string) (string, error)
		CompareHashAndData(hash string, data string) error
	}

	// DigitalSignatureService represents a service to manage digital signatures
	DigitalSignatureService interface {
		ParseKeyPair(private, public []byte) error
		GenerateKeyPair() ([]byte, []byte, error)
		EncodedPublicKey() string
		PEMHeaders() (string, string)
		CreateSignature(message string) (string, error)
	}

	// DockerSnapshotter represents a service used to create Docker environment(endpoint) snapshots
	DockerSnapshotter interface {
		CreateSnapshot(endpoint *Endpoint) (*DockerSnapshot, error)
	}

	// FileService represents a service for managing files
	FileService interface {
		GetDockerConfigPath() string
		GetFileContent(trustedRootPath, filePath string) ([]byte, error)
		Copy(fromFilePath string, toFilePath string, deleteIfExists bool) error
		Rename(oldPath, newPath string) error
		RemoveDirectory(directoryPath string) error
		StoreTLSFileFromBytes(folder string, fileType TLSFileType, data []byte) (string, error)
		GetPathForTLSFile(folder string, fileType TLSFileType) (string, error)
		DeleteTLSFile(folder string, fileType TLSFileType) error
		DeleteTLSFiles(folder string) error
		GetStackProjectPath(stackIdentifier string) string
		StoreStackFileFromBytes(stackIdentifier, fileName string, data []byte) (string, error)
		UpdateStoreStackFileFromBytes(stackIdentifier, fileName string, data []byte) (string, error)
		RemoveStackFileBackup(stackIdentifier, fileName string) error
		RollbackStackFile(stackIdentifier, fileName string) error
		GetEdgeStackProjectPath(edgeStackIdentifier string) string
		StoreEdgeStackFileFromBytes(edgeStackIdentifier, fileName string, data []byte) (string, error)
		StoreRegistryManagementFileFromBytes(folder, fileName string, data []byte) (string, error)
		KeyPairFilesExist() (bool, error)
		StoreKeyPair(private, public []byte, privatePEMHeader, publicPEMHeader string) error
		LoadKeyPair() ([]byte, []byte, error)
		WriteJSONToFile(path string, content interface{}) error
		FileExists(path string) (bool, error)
		StoreEdgeJobFileFromBytes(identifier string, data []byte) (string, error)
		GetEdgeJobFolder(identifier string) string
		ClearEdgeJobTaskLogs(edgeJobID, taskID string) error
		GetEdgeJobTaskLogFileContent(edgeJobID, taskID string) (string, error)
		StoreEdgeJobTaskLogFileFromBytes(edgeJobID, taskID string, data []byte) error
		GetBinaryFolder() string
		StoreCustomTemplateFileFromBytes(identifier, fileName string, data []byte) (string, error)
		GetCustomTemplateProjectPath(identifier string) string
		GetTemporaryPath() (string, error)
		GetDatastorePath() string
		GetDefaultSSLCertsPath() (string, string)
		StoreSSLCertPair(cert, key []byte) (string, string, error)
		CopySSLCertPair(certPath, keyPath string) (string, string, error)
		CopySSLCACert(caCertPath string) (string, error)
		StoreFDOProfileFileFromBytes(fdoProfileIdentifier string, data []byte) (string, error)
	}

	// GitService represents a service for managing Git
	GitService interface {
		CloneRepository(destination string, repositoryURL, referenceName, username, password string) error
		LatestCommitID(repositoryURL, referenceName, username, password string) (string, error)
		ListRefs(repositoryURL, username, password string, hardRefresh bool) ([]string, error)
		ListFiles(repositoryURL, referenceName, username, password string, hardRefresh bool, includeExts []string) ([]string, error)
	}

	// OpenAMTService represents a service for managing OpenAMT
	OpenAMTService interface {
		Configure(configuration models.OpenAMTConfiguration) error
		DeviceInformation(configuration models.OpenAMTConfiguration, deviceGUID string) (*models.OpenAMTDeviceInformation, error)
		EnableDeviceFeatures(configuration models.OpenAMTConfiguration, deviceGUID string, features models.OpenAMTDeviceEnabledFeatures) (string, error)
		ExecuteDeviceAction(configuration models.OpenAMTConfiguration, deviceGUID string, action string) error
	}

	// KubeClient represents a service used to query a Kubernetes environment(endpoint)
	KubeClient interface {
		SetupUserServiceAccount(userID int, teamIDs []int, restrictDefaultNamespace bool) error
		GetServiceAccount(tokendata *TokenData) (*v1.ServiceAccount, error)
		GetServiceAccountBearerToken(userID int) (string, error)
		CreateUserShellPod(ctx context.Context, serviceAccountName, shellPodImage string) (*KubernetesShellPod, error)
		StartExecProcess(token string, useAdminToken bool, namespace, podName, containerName string, command []string, stdin io.Reader, stdout io.Writer, errChan chan error)

		HasStackName(namespace string, stackName string) (bool, error)
		NamespaceAccessPoliciesDeleteNamespace(namespace string) error
		CreateNamespace(info models.K8sNamespaceDetails) error
		UpdateNamespace(info models.K8sNamespaceDetails) error
		GetNamespaces() (map[string]K8sNamespaceInfo, error)
		DeleteNamespace(namespace string) error
		GetConfigMapsAndSecrets(namespace string) ([]models.K8sConfigMapOrSecret, error)
		GetIngressControllers() models.K8sIngressControllers
		CreateIngress(namespace string, info models.K8sIngressInfo) error
		UpdateIngress(namespace string, info models.K8sIngressInfo) error
		GetIngresses(namespace string) ([]models.K8sIngressInfo, error)
		DeleteIngresses(reqs models.K8sIngressDeleteRequests) error
		CreateService(namespace string, service models.K8sServiceInfo) error
		UpdateService(namespace string, service models.K8sServiceInfo) error
		GetServices(namespace string) ([]models.K8sServiceInfo, error)
		DeleteServices(reqs models.K8sServiceDeleteRequests) error
		GetNodesLimits() (K8sNodesLimits, error)
		GetNamespaceAccessPolicies() (map[string]K8sNamespaceAccessPolicy, error)
		UpdateNamespaceAccessPolicies(accessPolicies map[string]K8sNamespaceAccessPolicy) error
		DeleteRegistrySecret(registry *Registry, namespace string) error
		CreateRegistrySecret(registry *Registry, namespace string) error
		IsRegistrySecret(namespace, secretName string) (bool, error)
		ToggleSystemState(namespace string, isSystem bool) error
	}

	// KubernetesDeployer represents a service to deploy a manifest inside a Kubernetes environment(endpoint)
	KubernetesDeployer interface {
		Deploy(userID models.UserID, endpoint *Endpoint, manifestFiles []string, namespace string) (string, error)
		Remove(userID models.UserID, endpoint *Endpoint, manifestFiles []string, namespace string) (string, error)
		ConvertCompose(data []byte) ([]byte, error)
	}

	// KubernetesSnapshotter represents a service used to create Kubernetes environment(endpoint) snapshots
	KubernetesSnapshotter interface {
		CreateSnapshot(endpoint *Endpoint) (*KubernetesSnapshot, error)
	}

	// LDAPService represents a service used to authenticate users against a LDAP/AD
	LDAPService interface {
		AuthenticateUser(username, password string, settings *models.LDAPSettings) error
		TestConnectivity(settings *models.LDAPSettings) error
		GetUserGroups(username string, settings *models.LDAPSettings) ([]string, error)
		SearchGroups(settings *models.LDAPSettings) ([]models.LDAPUser, error)
		SearchUsers(settings *models.LDAPSettings) ([]string, error)
	}

	// OAuthService represents a service used to authenticate users using OAuth
	OAuthService interface {
		Authenticate(code string, configuration *models.OAuthSettings) (string, error)
	}

	// ReverseTunnelService represents a service used to manage reverse tunnel connections.
	ReverseTunnelService interface {
		StartTunnelServer(addr, port string, snapshotService SnapshotService) error
		StopTunnelServer() error
		GenerateEdgeKey(url, host string, endpointIdentifier int) string
		SetTunnelStatusToActive(endpointID EndpointID)
		SetTunnelStatusToRequired(endpointID EndpointID) error
		SetTunnelStatusToIdle(endpointID EndpointID)
		KeepTunnelAlive(endpointID EndpointID, ctx context.Context, maxKeepAlive time.Duration)
		GetTunnelDetails(endpointID EndpointID) TunnelDetails
		GetActiveTunnel(endpoint *Endpoint) (TunnelDetails, error)
		AddEdgeJob(endpointID EndpointID, edgeJob *EdgeJob)
		RemoveEdgeJob(edgeJobID EdgeJobID)
	}

	// Server defines the interface to serve the API
	Server interface {
		Start() error
	}

	// SnapshotService represents a service for managing environment(endpoint) snapshots
	SnapshotService interface {
		Start()
		SetSnapshotInterval(snapshotInterval string) error
		SnapshotEndpoint(endpoint *Endpoint) error
		FillSnapshotData(endpoint *Endpoint) error
	}

	// SwarmStackManager represents a service to manage Swarm stacks
	SwarmStackManager interface {
		Login(registries []Registry, endpoint *Endpoint) error
		Logout(endpoint *Endpoint) error
		Deploy(stack *Stack, prune bool, pullImage bool, endpoint *Endpoint) error
		Remove(stack *Stack, endpoint *Endpoint) error
		NormalizeStackName(name string) string
	}
)

const (
	// APIVersion is the version number of the Portainer API
	APIVersion = "2.17.0"
	// DBVersion is the version number of the Portainer database
	DBVersion = 80
	// ComposeSyntaxMaxVersion is a maximum supported version of the docker compose syntax
	ComposeSyntaxMaxVersion = "3.9"
	// AssetsServerURL represents the URL of the Portainer asset server
	AssetsServerURL = "https://portainer-io-assets.sfo2.digitaloceanspaces.com"
	// MessageOfTheDayURL represents the URL where Portainer MOTD message can be retrieved
	MessageOfTheDayURL = AssetsServerURL + "/motd.json"
	// VersionCheckURL represents the URL used to retrieve the latest version of Portainer
	VersionCheckURL = "https://api.github.com/repos/portainer/portainer/releases/latest"
	// PortainerAgentHeader represents the name of the header available in any agent response
	PortainerAgentHeader = "Portainer-Agent"
	// PortainerAgentEdgeIDHeader represent the name of the header containing the Edge ID associated to an agent/agent cluster
	PortainerAgentEdgeIDHeader = "X-PortainerAgent-EdgeID"
	// HTTPResponseAgentPlatform represents the name of the header containing the Agent platform
	HTTPResponseAgentPlatform = "Portainer-Agent-Platform"
	// PortainerAgentTargetHeader represent the name of the header containing the target node name
	PortainerAgentTargetHeader = "X-PortainerAgent-Target"
	// PortainerAgentSignatureHeader represent the name of the header containing the digital signature
	PortainerAgentSignatureHeader = "X-PortainerAgent-Signature"
	// PortainerAgentPublicKeyHeader represent the name of the header containing the public key
	PortainerAgentPublicKeyHeader = "X-PortainerAgent-PublicKey"
	// PortainerAgentKubernetesSATokenHeader represent the name of the header containing a Kubernetes SA token
	PortainerAgentKubernetesSATokenHeader = "X-PortainerAgent-SA-Token"
	// PortainerAgentSignatureMessage represents the message used to create a digital signature
	// to be used when communicating with an agent
	PortainerAgentSignatureMessage = "Portainer-App"
	// DefaultSnapshotInterval represents the default interval between each environment snapshot job
	DefaultSnapshotInterval = "5m"
	// DefaultEdgeAgentCheckinIntervalInSeconds represents the default interval (in seconds) used by Edge agents to checkin with the Portainer instance
	DefaultEdgeAgentCheckinIntervalInSeconds = 5
	// DefaultTemplatesURL represents the URL to the official templates supported by Portainer
	DefaultTemplatesURL = "https://raw.githubusercontent.com/portainer/templates/master/templates-2.0.json"
	// DefaultHelmrepositoryURL represents the URL to the official templates supported by Bitnami
	DefaultHelmRepositoryURL = "https://charts.bitnami.com/bitnami"
	// DefaultUserSessionTimeout represents the default timeout after which the user session is cleared
	DefaultUserSessionTimeout = "8h"
	// DefaultUserSessionTimeout represents the default timeout after which the user session is cleared
	DefaultKubeconfigExpiry = "0"
	// DefaultKubectlShellImage represents the default image and tag for the kubectl shell
	DefaultKubectlShellImage = "portainer/kubectl-shell"
	// WebSocketKeepAlive web socket keep alive for edge environments
	WebSocketKeepAlive = 1 * time.Hour
)

const FeatureFlagEdgeRemoteUpdate models.Feature = "edgeRemoteUpdate"

// List of supported features
var SupportedFeatureFlags = []models.Feature{
	FeatureFlagEdgeRemoteUpdate,
}

const (
	_ AgentPlatform = iota
	// AgentPlatformDocker represent the Docker platform (Standalone/Swarm)
	AgentPlatformDocker
	// AgentPlatformKubernetes represent the Kubernetes platform
	AgentPlatformKubernetes
)

const (
	_ EdgeJobLogsStatus = iota
	// EdgeJobLogsStatusIdle represents an idle log collection job
	EdgeJobLogsStatusIdle
	// EdgeJobLogsStatusPending represents a pending log collection job
	EdgeJobLogsStatusPending
	// EdgeJobLogsStatusCollected represents a completed log collection job
	EdgeJobLogsStatusCollected
)

const (
	_ CustomTemplatePlatform = iota
	// CustomTemplatePlatformLinux represents a custom template for linux
	CustomTemplatePlatformLinux
	// CustomTemplatePlatformWindows represents a custom template for windows
	CustomTemplatePlatformWindows
)

const (
	// EdgeStackDeploymentCompose represent an edge stack deployed using a compose file
	EdgeStackDeploymentCompose EdgeStackDeploymentType = iota
	// EdgeStackDeploymentKubernetes represent an edge stack deployed using a kubernetes manifest file
	EdgeStackDeploymentKubernetes
)

const (
	_ EdgeStackStatusType = iota
	//StatusOk represents a successfully deployed edge stack
	StatusOk
	//StatusError represents an edge environment(endpoint) which failed to deploy its edge stack
	StatusError
	//StatusAcknowledged represents an acknowledged edge stack
	StatusAcknowledged
)

const (
	_ EndpointStatus = iota
	// EndpointStatusUp is used to represent an available environment(endpoint)
	EndpointStatusUp
	// EndpointStatusDown is used to represent an unavailable environment(endpoint)
	EndpointStatusDown
)

const (
	_ EndpointType = iota
	// DockerEnvironment represents an environment(endpoint) connected to a Docker environment(endpoint)
	DockerEnvironment
	// AgentOnDockerEnvironment represents an environment(endpoint) connected to a Portainer agent deployed on a Docker environment(endpoint)
	AgentOnDockerEnvironment
	// AzureEnvironment represents an environment(endpoint) connected to an Azure environment(endpoint)
	AzureEnvironment
	// EdgeAgentOnDockerEnvironment represents an environment(endpoint) connected to an Edge agent deployed on a Docker environment(endpoint)
	EdgeAgentOnDockerEnvironment
	// KubernetesLocalEnvironment represents an environment(endpoint) connected to a local Kubernetes environment(endpoint)
	KubernetesLocalEnvironment
	// AgentOnKubernetesEnvironment represents an environment(endpoint) connected to a Portainer agent deployed on a Kubernetes environment(endpoint)
	AgentOnKubernetesEnvironment
	// EdgeAgentOnKubernetesEnvironment represents an environment(endpoint) connected to an Edge agent deployed on a Kubernetes environment(endpoint)
	EdgeAgentOnKubernetesEnvironment
)

const (
	_ JobType = iota
	// SnapshotJobType is a system job used to create environment(endpoint) snapshots
	SnapshotJobType = 2
)

const (
	_ models.MembershipRole = iota
	// TeamLeader represents a leader role inside a team
	TeamLeader
	// TeamMember represents a member role inside a team
	TeamMember
)

const (
	_ SoftwareEdition = iota
	// PortainerCE represents the community edition of Portainer
	PortainerCE
	// PortainerBE represents the business edition of Portainer
	PortainerBE
	// PortainerEE represents the business edition of Portainer
	PortainerEE
)

const (
	_ RegistryType = iota
	// QuayRegistry represents a Quay.io registry
	QuayRegistry
	// AzureRegistry represents an ACR registry
	AzureRegistry
	// CustomRegistry represents a custom registry
	CustomRegistry
	// GitlabRegistry represents a gitlab registry
	GitlabRegistry
	// ProGetRegistry represents a proget registry
	ProGetRegistry
	// DockerHubRegistry represents a dockerhub registry
	DockerHubRegistry
	// EcrRegistry represents an ECR registry
	EcrRegistry
)

const (
	_ models.ResourceAccessLevel = iota
	// ReadWriteAccessLevel represents an access level with read-write permissions on a resource
	ReadWriteAccessLevel
)

const (
	_ ResourceControlType = iota
	// ContainerResourceControl represents a resource control associated to a Docker container
	ContainerResourceControl
	// ServiceResourceControl represents a resource control associated to a Docker service
	ServiceResourceControl
	// VolumeResourceControl represents a resource control associated to a Docker volume
	VolumeResourceControl
	// NetworkResourceControl represents a resource control associated to a Docker network
	NetworkResourceControl
	// SecretResourceControl represents a resource control associated to a Docker secret
	SecretResourceControl
	// StackResourceControl represents a resource control associated to a stack composed of Docker services
	StackResourceControl
	// ConfigResourceControl represents a resource control associated to a Docker config
	ConfigResourceControl
	// CustomTemplateResourceControl represents a resource control associated to a custom template
	CustomTemplateResourceControl
	// ContainerGroupResourceControl represents a resource control associated to an Azure container group
	ContainerGroupResourceControl
)

const (
	_ StackType = iota
	// DockerSwarmStack represents a stack managed via docker stack
	DockerSwarmStack
	// DockerComposeStack represents a stack managed via docker-compose
	DockerComposeStack
	// KubernetesStack represents a stack managed via kubectl
	KubernetesStack
)

// StackStatus represents a status for a stack
const (
	_ StackStatus = iota
	StackStatusActive
	StackStatusInactive
)

const (
	_ TemplateType = iota
	// ContainerTemplate represents a container template
	ContainerTemplate
	// SwarmStackTemplate represents a template used to deploy a Swarm stack
	SwarmStackTemplate
	// ComposeStackTemplate represents a template used to deploy a Compose stack
	ComposeStackTemplate
	// EdgeStackTemplate represents a template used to deploy an Edge stack
	EdgeStackTemplate
)

const (
	// TLSFileCA represents a TLS CA certificate file
	TLSFileCA TLSFileType = iota
	// TLSFileCert represents a TLS certificate file
	TLSFileCert
	// TLSFileKey represents a TLS key file
	TLSFileKey
)

const (
	_ UserRole = iota
	// AdministratorRole represents an administrator user role
	AdministratorRole
	// StandardUserRole represents a regular user role
	StandardUserRole
)

const (
	_ WebhookType = iota
	// ServiceWebhook is a webhook for restarting a docker service
	ServiceWebhook
)

const (
	// EdgeAgentIdle represents an idle state for a tunnel connected to an Edge environment(endpoint).
	EdgeAgentIdle string = "IDLE"
	// EdgeAgentManagementRequired represents a required state for a tunnel connected to an Edge environment(endpoint)
	EdgeAgentManagementRequired string = "REQUIRED"
	// EdgeAgentActive represents an active state for a tunnel connected to an Edge environment(endpoint)
	EdgeAgentActive string = "ACTIVE"
)

// represents an authorization type
const (
	OperationDockerContainerArchiveInfo         models.Authorization = "DockerContainerArchiveInfo"
	OperationDockerContainerList                models.Authorization = "DockerContainerList"
	OperationDockerContainerExport              models.Authorization = "DockerContainerExport"
	OperationDockerContainerChanges             models.Authorization = "DockerContainerChanges"
	OperationDockerContainerInspect             models.Authorization = "DockerContainerInspect"
	OperationDockerContainerTop                 models.Authorization = "DockerContainerTop"
	OperationDockerContainerLogs                models.Authorization = "DockerContainerLogs"
	OperationDockerContainerStats               models.Authorization = "DockerContainerStats"
	OperationDockerContainerAttachWebsocket     models.Authorization = "DockerContainerAttachWebsocket"
	OperationDockerContainerArchive             models.Authorization = "DockerContainerArchive"
	OperationDockerContainerCreate              models.Authorization = "DockerContainerCreate"
	OperationDockerContainerPrune               models.Authorization = "DockerContainerPrune"
	OperationDockerContainerKill                models.Authorization = "DockerContainerKill"
	OperationDockerContainerPause               models.Authorization = "DockerContainerPause"
	OperationDockerContainerUnpause             models.Authorization = "DockerContainerUnpause"
	OperationDockerContainerRestart             models.Authorization = "DockerContainerRestart"
	OperationDockerContainerStart               models.Authorization = "DockerContainerStart"
	OperationDockerContainerStop                models.Authorization = "DockerContainerStop"
	OperationDockerContainerWait                models.Authorization = "DockerContainerWait"
	OperationDockerContainerResize              models.Authorization = "DockerContainerResize"
	OperationDockerContainerAttach              models.Authorization = "DockerContainerAttach"
	OperationDockerContainerExec                models.Authorization = "DockerContainerExec"
	OperationDockerContainerRename              models.Authorization = "DockerContainerRename"
	OperationDockerContainerUpdate              models.Authorization = "DockerContainerUpdate"
	OperationDockerContainerPutContainerArchive models.Authorization = "DockerContainerPutContainerArchive"
	OperationDockerContainerDelete              models.Authorization = "DockerContainerDelete"
	OperationDockerImageList                    models.Authorization = "DockerImageList"
	OperationDockerImageSearch                  models.Authorization = "DockerImageSearch"
	OperationDockerImageGetAll                  models.Authorization = "DockerImageGetAll"
	OperationDockerImageGet                     models.Authorization = "DockerImageGet"
	OperationDockerImageHistory                 models.Authorization = "DockerImageHistory"
	OperationDockerImageInspect                 models.Authorization = "DockerImageInspect"
	OperationDockerImageLoad                    models.Authorization = "DockerImageLoad"
	OperationDockerImageCreate                  models.Authorization = "DockerImageCreate"
	OperationDockerImagePrune                   models.Authorization = "DockerImagePrune"
	OperationDockerImagePush                    models.Authorization = "DockerImagePush"
	OperationDockerImageTag                     models.Authorization = "DockerImageTag"
	OperationDockerImageDelete                  models.Authorization = "DockerImageDelete"
	OperationDockerImageCommit                  models.Authorization = "DockerImageCommit"
	OperationDockerImageBuild                   models.Authorization = "DockerImageBuild"
	OperationDockerNetworkList                  models.Authorization = "DockerNetworkList"
	OperationDockerNetworkInspect               models.Authorization = "DockerNetworkInspect"
	OperationDockerNetworkCreate                models.Authorization = "DockerNetworkCreate"
	OperationDockerNetworkConnect               models.Authorization = "DockerNetworkConnect"
	OperationDockerNetworkDisconnect            models.Authorization = "DockerNetworkDisconnect"
	OperationDockerNetworkPrune                 models.Authorization = "DockerNetworkPrune"
	OperationDockerNetworkDelete                models.Authorization = "DockerNetworkDelete"
	OperationDockerVolumeList                   models.Authorization = "DockerVolumeList"
	OperationDockerVolumeInspect                models.Authorization = "DockerVolumeInspect"
	OperationDockerVolumeCreate                 models.Authorization = "DockerVolumeCreate"
	OperationDockerVolumePrune                  models.Authorization = "DockerVolumePrune"
	OperationDockerVolumeDelete                 models.Authorization = "DockerVolumeDelete"
	OperationDockerExecInspect                  models.Authorization = "DockerExecInspect"
	OperationDockerExecStart                    models.Authorization = "DockerExecStart"
	OperationDockerExecResize                   models.Authorization = "DockerExecResize"
	OperationDockerSwarmInspect                 models.Authorization = "DockerSwarmInspect"
	OperationDockerSwarmUnlockKey               models.Authorization = "DockerSwarmUnlockKey"
	OperationDockerSwarmInit                    models.Authorization = "DockerSwarmInit"
	OperationDockerSwarmJoin                    models.Authorization = "DockerSwarmJoin"
	OperationDockerSwarmLeave                   models.Authorization = "DockerSwarmLeave"
	OperationDockerSwarmUpdate                  models.Authorization = "DockerSwarmUpdate"
	OperationDockerSwarmUnlock                  models.Authorization = "DockerSwarmUnlock"
	OperationDockerNodeList                     models.Authorization = "DockerNodeList"
	OperationDockerNodeInspect                  models.Authorization = "DockerNodeInspect"
	OperationDockerNodeUpdate                   models.Authorization = "DockerNodeUpdate"
	OperationDockerNodeDelete                   models.Authorization = "DockerNodeDelete"
	OperationDockerServiceList                  models.Authorization = "DockerServiceList"
	OperationDockerServiceInspect               models.Authorization = "DockerServiceInspect"
	OperationDockerServiceLogs                  models.Authorization = "DockerServiceLogs"
	OperationDockerServiceCreate                models.Authorization = "DockerServiceCreate"
	OperationDockerServiceUpdate                models.Authorization = "DockerServiceUpdate"
	OperationDockerServiceDelete                models.Authorization = "DockerServiceDelete"
	OperationDockerSecretList                   models.Authorization = "DockerSecretList"
	OperationDockerSecretInspect                models.Authorization = "DockerSecretInspect"
	OperationDockerSecretCreate                 models.Authorization = "DockerSecretCreate"
	OperationDockerSecretUpdate                 models.Authorization = "DockerSecretUpdate"
	OperationDockerSecretDelete                 models.Authorization = "DockerSecretDelete"
	OperationDockerConfigList                   models.Authorization = "DockerConfigList"
	OperationDockerConfigInspect                models.Authorization = "DockerConfigInspect"
	OperationDockerConfigCreate                 models.Authorization = "DockerConfigCreate"
	OperationDockerConfigUpdate                 models.Authorization = "DockerConfigUpdate"
	OperationDockerConfigDelete                 models.Authorization = "DockerConfigDelete"
	OperationDockerTaskList                     models.Authorization = "DockerTaskList"
	OperationDockerTaskInspect                  models.Authorization = "DockerTaskInspect"
	OperationDockerTaskLogs                     models.Authorization = "DockerTaskLogs"
	OperationDockerPluginList                   models.Authorization = "DockerPluginList"
	OperationDockerPluginPrivileges             models.Authorization = "DockerPluginPrivileges"
	OperationDockerPluginInspect                models.Authorization = "DockerPluginInspect"
	OperationDockerPluginPull                   models.Authorization = "DockerPluginPull"
	OperationDockerPluginCreate                 models.Authorization = "DockerPluginCreate"
	OperationDockerPluginEnable                 models.Authorization = "DockerPluginEnable"
	OperationDockerPluginDisable                models.Authorization = "DockerPluginDisable"
	OperationDockerPluginPush                   models.Authorization = "DockerPluginPush"
	OperationDockerPluginUpgrade                models.Authorization = "DockerPluginUpgrade"
	OperationDockerPluginSet                    models.Authorization = "DockerPluginSet"
	OperationDockerPluginDelete                 models.Authorization = "DockerPluginDelete"
	OperationDockerSessionStart                 models.Authorization = "DockerSessionStart"
	OperationDockerDistributionInspect          models.Authorization = "DockerDistributionInspect"
	OperationDockerBuildPrune                   models.Authorization = "DockerBuildPrune"
	OperationDockerBuildCancel                  models.Authorization = "DockerBuildCancel"
	OperationDockerPing                         models.Authorization = "DockerPing"
	OperationDockerInfo                         models.Authorization = "DockerInfo"
	OperationDockerEvents                       models.Authorization = "DockerEvents"
	OperationDockerSystem                       models.Authorization = "DockerSystem"
	OperationDockerVersion                      models.Authorization = "DockerVersion"

	OperationDockerAgentPing         models.Authorization = "DockerAgentPing"
	OperationDockerAgentList         models.Authorization = "DockerAgentList"
	OperationDockerAgentHostInfo     models.Authorization = "DockerAgentHostInfo"
	OperationDockerAgentBrowseDelete models.Authorization = "DockerAgentBrowseDelete"
	OperationDockerAgentBrowseGet    models.Authorization = "DockerAgentBrowseGet"
	OperationDockerAgentBrowseList   models.Authorization = "DockerAgentBrowseList"
	OperationDockerAgentBrowsePut    models.Authorization = "DockerAgentBrowsePut"
	OperationDockerAgentBrowseRename models.Authorization = "DockerAgentBrowseRename"

	OperationPortainerDockerHubInspect      models.Authorization = "PortainerDockerHubInspect"
	OperationPortainerDockerHubUpdate       models.Authorization = "PortainerDockerHubUpdate"
	OperationPortainerEndpointGroupCreate   models.Authorization = "PortainerEndpointGroupCreate"
	OperationPortainerEndpointGroupList     models.Authorization = "PortainerEndpointGroupList"
	OperationPortainerEndpointGroupDelete   models.Authorization = "PortainerEndpointGroupDelete"
	OperationPortainerEndpointGroupInspect  models.Authorization = "PortainerEndpointGroupInspect"
	OperationPortainerEndpointGroupUpdate   models.Authorization = "PortainerEndpointGroupEdit"
	OperationPortainerEndpointGroupAccess   models.Authorization = "PortainerEndpointGroupAccess "
	OperationPortainerEndpointList          models.Authorization = "PortainerEndpointList"
	OperationPortainerEndpointInspect       models.Authorization = "PortainerEndpointInspect"
	OperationPortainerEndpointCreate        models.Authorization = "PortainerEndpointCreate"
	OperationPortainerEndpointJob           models.Authorization = "PortainerEndpointJob"
	OperationPortainerEndpointSnapshots     models.Authorization = "PortainerEndpointSnapshots"
	OperationPortainerEndpointSnapshot      models.Authorization = "PortainerEndpointSnapshot"
	OperationPortainerEndpointUpdate        models.Authorization = "PortainerEndpointUpdate"
	OperationPortainerEndpointUpdateAccess  models.Authorization = "PortainerEndpointUpdateAccess"
	OperationPortainerEndpointDelete        models.Authorization = "PortainerEndpointDelete"
	OperationPortainerExtensionList         models.Authorization = "PortainerExtensionList"
	OperationPortainerExtensionInspect      models.Authorization = "PortainerExtensionInspect"
	OperationPortainerExtensionCreate       models.Authorization = "PortainerExtensionCreate"
	OperationPortainerExtensionUpdate       models.Authorization = "PortainerExtensionUpdate"
	OperationPortainerExtensionDelete       models.Authorization = "PortainerExtensionDelete"
	OperationPortainerMOTD                  models.Authorization = "PortainerMOTD"
	OperationPortainerRegistryList          models.Authorization = "PortainerRegistryList"
	OperationPortainerRegistryInspect       models.Authorization = "PortainerRegistryInspect"
	OperationPortainerRegistryCreate        models.Authorization = "PortainerRegistryCreate"
	OperationPortainerRegistryConfigure     models.Authorization = "PortainerRegistryConfigure"
	OperationPortainerRegistryUpdate        models.Authorization = "PortainerRegistryUpdate"
	OperationPortainerRegistryUpdateAccess  models.Authorization = "PortainerRegistryUpdateAccess"
	OperationPortainerRegistryDelete        models.Authorization = "PortainerRegistryDelete"
	OperationPortainerResourceControlCreate models.Authorization = "PortainerResourceControlCreate"
	OperationPortainerResourceControlUpdate models.Authorization = "PortainerResourceControlUpdate"
	OperationPortainerResourceControlDelete models.Authorization = "PortainerResourceControlDelete"
	OperationPortainerRoleList              models.Authorization = "PortainerRoleList"
	OperationPortainerRoleInspect           models.Authorization = "PortainerRoleInspect"
	OperationPortainerRoleCreate            models.Authorization = "PortainerRoleCreate"
	OperationPortainerRoleUpdate            models.Authorization = "PortainerRoleUpdate"
	OperationPortainerRoleDelete            models.Authorization = "PortainerRoleDelete"
	OperationPortainerScheduleList          models.Authorization = "PortainerScheduleList"
	OperationPortainerScheduleInspect       models.Authorization = "PortainerScheduleInspect"
	OperationPortainerScheduleFile          models.Authorization = "PortainerScheduleFile"
	OperationPortainerScheduleTasks         models.Authorization = "PortainerScheduleTasks"
	OperationPortainerScheduleCreate        models.Authorization = "PortainerScheduleCreate"
	OperationPortainerScheduleUpdate        models.Authorization = "PortainerScheduleUpdate"
	OperationPortainerScheduleDelete        models.Authorization = "PortainerScheduleDelete"
	OperationPortainerSettingsInspect       models.Authorization = "PortainerSettingsInspect"
	OperationPortainerSettingsUpdate        models.Authorization = "PortainerSettingsUpdate"
	OperationPortainerSettingsLDAPCheck     models.Authorization = "PortainerSettingsLDAPCheck"
	OperationPortainerStackList             models.Authorization = "PortainerStackList"
	OperationPortainerStackInspect          models.Authorization = "PortainerStackInspect"
	OperationPortainerStackFile             models.Authorization = "PortainerStackFile"
	OperationPortainerStackCreate           models.Authorization = "PortainerStackCreate"
	OperationPortainerStackMigrate          models.Authorization = "PortainerStackMigrate"
	OperationPortainerStackUpdate           models.Authorization = "PortainerStackUpdate"
	OperationPortainerStackDelete           models.Authorization = "PortainerStackDelete"
	OperationPortainerTagList               models.Authorization = "PortainerTagList"
	OperationPortainerTagCreate             models.Authorization = "PortainerTagCreate"
	OperationPortainerTagDelete             models.Authorization = "PortainerTagDelete"
	OperationPortainerTeamMembershipList    models.Authorization = "PortainerTeamMembershipList"
	OperationPortainerTeamMembershipCreate  models.Authorization = "PortainerTeamMembershipCreate"
	OperationPortainerTeamMembershipUpdate  models.Authorization = "PortainerTeamMembershipUpdate"
	OperationPortainerTeamMembershipDelete  models.Authorization = "PortainerTeamMembershipDelete"
	OperationPortainerTeamList              models.Authorization = "PortainerTeamList"
	OperationPortainerTeamInspect           models.Authorization = "PortainerTeamInspect"
	OperationPortainerTeamMemberships       models.Authorization = "PortainerTeamMemberships"
	OperationPortainerTeamCreate            models.Authorization = "PortainerTeamCreate"
	OperationPortainerTeamUpdate            models.Authorization = "PortainerTeamUpdate"
	OperationPortainerTeamDelete            models.Authorization = "PortainerTeamDelete"
	OperationPortainerTemplateList          models.Authorization = "PortainerTemplateList"
	OperationPortainerTemplateInspect       models.Authorization = "PortainerTemplateInspect"
	OperationPortainerTemplateCreate        models.Authorization = "PortainerTemplateCreate"
	OperationPortainerTemplateUpdate        models.Authorization = "PortainerTemplateUpdate"
	OperationPortainerTemplateDelete        models.Authorization = "PortainerTemplateDelete"
	OperationPortainerUploadTLS             models.Authorization = "PortainerUploadTLS"
	OperationPortainerUserList              models.Authorization = "PortainerUserList"
	OperationPortainerUserInspect           models.Authorization = "PortainerUserInspect"
	OperationPortainerUserMemberships       models.Authorization = "PortainerUserMemberships"
	OperationPortainerUserCreate            models.Authorization = "PortainerUserCreate"
	OperationPortainerUserListToken         models.Authorization = "PortainerUserListToken"
	OperationPortainerUserCreateToken       models.Authorization = "PortainerUserCreateToken"
	OperationPortainerUserRevokeToken       models.Authorization = "PortainerUserRevokeToken"
	OperationPortainerUserUpdate            models.Authorization = "PortainerUserUpdate"
	OperationPortainerUserUpdatePassword    models.Authorization = "PortainerUserUpdatePassword"
	OperationPortainerUserDelete            models.Authorization = "PortainerUserDelete"
	OperationPortainerWebsocketExec         models.Authorization = "PortainerWebsocketExec"
	OperationPortainerWebhookList           models.Authorization = "PortainerWebhookList"
	OperationPortainerWebhookCreate         models.Authorization = "PortainerWebhookCreate"
	OperationPortainerWebhookDelete         models.Authorization = "PortainerWebhookDelete"

	OperationDockerUndefined      models.Authorization = "DockerUndefined"
	OperationDockerAgentUndefined models.Authorization = "DockerAgentUndefined"
	OperationPortainerUndefined   models.Authorization = "PortainerUndefined"

	EndpointResourcesAccess models.Authorization = "EndpointResourcesAccess"

	// Deprecated operations
	OperationPortainerEndpointExtensionAdd    models.Authorization = "PortainerEndpointExtensionAdd"
	OperationPortainerEndpointExtensionRemove models.Authorization = "PortainerEndpointExtensionRemove"
	OperationIntegrationStoridgeAdmin         models.Authorization = "IntegrationStoridgeAdmin"
)

const (
	AzurePathContainerGroups = "/subscriptions/*/providers/Microsoft.ContainerInstance/containerGroups"
	AzurePathContainerGroup  = "/subscriptions/*/resourceGroups/*/providers/Microsoft.ContainerInstance/containerGroups/*"
)
