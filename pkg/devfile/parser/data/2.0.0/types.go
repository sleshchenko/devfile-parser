package version200

import "github.com/cli-playground/devfile-parser/pkg/devfile/parser/data/common"

// CommandGroupType describes the kind of command group.
// +kubebuilder:validation:Enum=build;run;test;debug
type CommandGroupType string

const (
	BuildCommandGroupType CommandGroupType = "build"
	RunCommandGroupType   CommandGroupType = "run"
	TestCommandGroupType  CommandGroupType = "test"
	DebugCommandGroupType CommandGroupType = "debug"
)

// Devfile200 Devfile schema.
type Devfile200 struct {

	// Predefined, ready-to-use, workspace-related commands
	Commands []common.DevfileCommand `json:"commands,omitempty"`

	// List of the workspace components, such as editor and plugins, user-provided containers, or other types of components
	Components []common.DevfileComponent `json:"components,omitempty"`

	// Bindings of commands to events. Each command is referred-to by its name.
	Events common.DevfileEvents `json:"events,omitempty"`

	// Optional metadata
	Metadata common.DevfileMetadata `json:"metadata,omitempty"`

	// Parent workspace template
	Parent common.DevfileParent `json:"parent,omitempty"`

	// Projects worked on in the workspace, containing names and sources locations
	Projects []common.DevfileProject `json:"projects,omitempty"`

	// Devfile schema version
	SchemaVersion string `json:"schemaVersion"`
}

// CommandsItems
type Command struct {

	// Composite command that allows executing several sub-commands either sequentially or concurrently
	Composite *Composite `json:"composite,omitempty"`

	// CLI Command executed in a component container
	Exec *Exec `json:"exec,omitempty"`

	// Command providing the definition of a VsCode launch action
	VscodeLaunch *VscodeLaunch `json:"vscodeLaunch,omitempty"`

	// Command providing the definition of a VsCode Task
	VscodeTask *VscodeTask `json:"vscodeTask,omitempty"`
}

// ComponentsItems
type Component struct {

	// Allows adding and configuring workspace-related containers
	Container *Container `json:"container,omitempty"`

	// Allows importing into the workspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.
	Kubernetes *Kubernetes `json:"kubernetes,omitempty"`

	// Allows importing into the workspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.
	Openshift *Openshift `json:"openshift,omitempty"`

	// Allows importing a plugin. Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as `DevWorkspaceTemplate` Kubernetes Custom Resources
	Plugin *Plugin `json:"plugin,omitempty"`

	// Allows specifying the definition of a volume shared by several other components
	Volume *Volume `json:"volume,omitempty"`
}

// Composite Composite command that allows executing several sub-commands either sequentially or concurrently
type Composite struct {

	// Optional map of free-form additional command attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// The commands that comprise this composite command
	Commands []string `json:"commands,omitempty"`

	// Defines the group this command is part of
	Group *Group `json:"group,omitempty"`

	// Mandatory identifier that allows referencing this command in composite commands, or from a parent, or in events.
	Id string `json:"id"`

	// Optional label that provides a label for this command to be used in Editor UI menus for example
	Label string `json:"label,omitempty"`

	// Indicates if the sub-commands should be executed concurrently
	Parallel bool `json:"parallel,omitempty"`
}

// Configuration
type Configuration struct {
	CookiesAuthEnabled bool   `json:"cookiesAuthEnabled,omitempty"`
	Discoverable       bool   `json:"discoverable,omitempty"`
	Path               string `json:"path,omitempty"`

	// The is the low-level protocol of traffic coming through this endpoint. Default value is "tcp"
	Protocol string `json:"protocol,omitempty"`
	Public   bool   `json:"public,omitempty"`

	// The is the URL scheme to use when accessing the endpoint. Default value is "http"
	Scheme string `json:"scheme,omitempty"`
	Secure bool   `json:"secure,omitempty"`
	Type   string `json:"type,omitempty"`
}

// Container Allows adding and configuring workspace-related containers
type Container struct {

	// The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command. Defaults to an empty array, meaning use whatever is defined in the image.
	Args []string `json:"args,omitempty"`

	// The command to run in the dockerimage component instead of the default one provided in the image. Defaults to an empty array, meaning use whatever is defined in the image.
	Command []string `json:"command,omitempty"`

	Endpoints []*Endpoint `json:"endpoints,omitempty"`

	// Environment variables used in this container
	Env          []*Env `json:"env,omitempty"`
	Image        string `json:"image,omitempty"`
	MemoryLimit  string `json:"memoryLimit,omitempty"`
	MountSources bool   `json:"mountSources,omitempty"`
	Name         string `json:"name"`

	// Optional specification of the path in the container where project sources should be transferred/mounted when `mountSources` is `true`. When omitted, the value of the `PROJECTS_ROOT` environment variable is used.
	SourceMapping string `json:"sourceMapping,omitempty"`

	// List of volumes mounts that should be mounted is this container.
	VolumeMounts []*VolumeMount `json:"volumeMounts,omitempty"`
}

// Endpoint
type Endpoint struct {
	Attributes    map[string]string `json:"attributes,omitempty"`
	Configuration *Configuration    `json:"configuration,omitempty"`
	Name          string            `json:"name"`
	TargetPort    int32             `json:"targetPort"`
}

// Env
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Events Bindings of commands to events. Each command is referred-to by its name.
type Events struct {

	// Names of commands that should be executed after the workspace is completely started. In the case of Che-Theia, these commands should be executed after all plugins and extensions have started, including project cloning. This means that those commands are not triggered until the user opens the IDE in his browser.
	PostStart []string `json:"postStart,omitempty"`

	// Names of commands that should be executed after stopping the workspace.
	PostStop []string `json:"postStop,omitempty"`

	// Names of commands that should be executed before the workspace start. Kubernetes-wise, these commands would typically be executed in init containers of the workspace POD.
	PreStart []string `json:"preStart,omitempty"`

	// Names of commands that should be executed before stopping the workspace.
	PreStop []string `json:"preStop,omitempty"`
}

// Exec CLI Command executed in a component container
type Exec struct {

	// Optional map of free-form additional command attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// The actual command-line string
	CommandLine string `json:"commandLine,omitempty"`

	// Describes component to which given action relates
	Component string `json:"component,omitempty"`

	// Optional list of environment variables that have to be set before running the command
	Env []*Env `json:"env,omitempty"`

	// Defines the group this command is part of
	Group *Group `json:"group,omitempty"`

	// Mandatory identifier that allows referencing this command in composite commands, or from a parent, or in events.
	Id string `json:"id"`

	// Optional label that provides a label for this command to be used in Editor UI menus for example
	Label string `json:"label,omitempty"`

	// Working directory where the command should be executed
	WorkingDir string `json:"workingDir,omitempty"`
}

// Git Project's Git source
type Git struct {

	// The branch to check
	Branch string `json:"branch,omitempty"`

	// Project's source location address. Should be URL for git and github located projects, or; file:// for zip
	Location string `json:"location,omitempty"`

	// Part of project to populate in the working directory.
	SparseCheckoutDir string `json:"sparseCheckoutDir,omitempty"`

	// The tag or commit id to reset the checked out branch to
	StartPoint string `json:"startPoint,omitempty"`
}

// Github Project's GitHub source
type Github struct {

	// The branch to check
	Branch string `json:"branch,omitempty"`

	// Project's source location address. Should be URL for git and github located projects, or; file:// for zip
	Location string `json:"location,omitempty"`

	// Part of project to populate in the working directory.
	SparseCheckoutDir string `json:"sparseCheckoutDir,omitempty"`

	// The tag or commit id to reset the checked out branch to
	StartPoint string `json:"startPoint,omitempty"`
}

// Group Defines the group this command is part of
type Group struct {

	// Identifies the default command for a given group kind
	IsDefault bool `json:"isDefault,omitempty"`

	// Kind of group the command is part of
	Kind CommandGroupType `json:"kind"`
}

// Kubernetes Allows importing into the workspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.
type Kubernetes struct {

	// Inlined manifest
	Inlined string `json:"inlined,omitempty"`

	// Mandatory name that allows referencing the component in commands, or inside a parent
	Name string `json:"name"`

	// Location in a file fetched from a uri.
	Uri string `json:"uri,omitempty"`
}

// Metadata Optional metadata
type Metadata struct {

	// Optional devfile name
	Name string `json:"name,omitempty"`

	// Optional semver-compatible version
	Version string `json:"version,omitempty"`
}

// Openshift Configuration overriding for an OpenShift component
type Openshift struct {

	// Inlined manifest
	Inlined string `json:"inlined,omitempty"`

	// Mandatory name that allows referencing the component in commands, or inside a parent
	Name string `json:"name"`

	// Location in a file fetched from a uri.
	Uri string `json:"uri,omitempty"`
}

// Parent Parent workspace template
type Parent struct {

	// Predefined, ready-to-use, workspace-related commands
	Commands []*Command `json:"commands,omitempty"`

	// List of the workspace components, such as editor and plugins, user-provided containers, or other types of components
	Components []*Component `json:"components,omitempty"`

	// Bindings of commands to events. Each command is referred-to by its name.
	Events *Events `json:"events,omitempty"`

	// Id in a registry that contains a Devfile yaml file
	Id string `json:"id,omitempty"`

	// Reference to a Kubernetes CRD of type DevWorkspaceTemplate
	Kubernetes *Kubernetes `json:"kubernetes,omitempty"`

	// Projects worked on in the workspace, containing names and sources locations
	Projects []*Project `json:"projects,omitempty"`

	RegistryUrl string `json:"registryUrl,omitempty"`

	// Uri of a Devfile yaml file
	Uri string `json:"uri,omitempty"`
}

// Plugin Allows importing a plugin. Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as `DevWorkspaceTemplate` Kubernetes Custom Resources
type Plugin struct {

	// Overrides of commands encapsulated in a plugin. Overriding is done using a strategic merge
	Commands []*Command `json:"commands,omitempty"`

	// Overrides of components encapsulated in a plugin. Overriding is done using a strategic merge
	Components []*Component `json:"components,omitempty"`

	// Id in a registry that contains a Devfile yaml file
	Id string `json:"id,omitempty"`

	// Reference to a Kubernetes CRD of type DevWorkspaceTemplate
	Kubernetes *Kubernetes `json:"kubernetes,omitempty"`

	// Optional name that allows referencing the component in commands, or inside a parent If omitted it will be infered from the location (uri or registryEntry)
	Name        string `json:"name,omitempty"`
	RegistryUrl string `json:"registryUrl,omitempty"`

	// Uri of a Devfile yaml file
	Uri string `json:"uri,omitempty"`
}

// ProjectsItems
type Project struct {

	// Path relative to the root of the projects to which this project should be cloned into. This is a unix-style relative path (i.e. uses forward slashes). The path is invalid if it is absolute or tries to escape the project root through the usage of '..'. If not specified, defaults to the project name.
	ClonePath string `json:"clonePath,omitempty"`

	// Project's Git source
	Git *Git `json:"git,omitempty"`

	// Project's GitHub source
	Github *Github `json:"github,omitempty"`

	// Project name
	Name string `json:"name"`

	// Project's Zip source
	Zip *Zip `json:"zip,omitempty"`
}

// Volume Allows specifying the definition of a volume shared by several other components
type Volume struct {

	// Mandatory name that allows referencing the Volume component in Container volume mounts or inside a parent
	Name string `json:"name"`

	// Size of the volume
	Size string `json:"size,omitempty"`
}

// VolumeMountsItems Volume that should be mounted to a component container
type VolumeMount struct {

	// The volume mount name is the name of an existing `Volume` component. If no corresponding `Volume` component exist it is implicitly added. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.
	Name string `json:"name"`

	// The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is `/<name>`.
	Path string `json:"path,omitempty"`
}

// VscodeLaunch Command providing the definition of a VsCode launch action
type VscodeLaunch struct {

	// Optional map of free-form additional command attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// Defines the group this command is part of
	Group *Group `json:"group,omitempty"`

	// Mandatory identifier that allows referencing this command in composite commands, or from a parent, or in events.
	Id string `json:"id"`

	// Inlined content of the VsCode configuration
	Inlined string `json:"inlined,omitempty"`

	// Location as an absolute of relative URI the VsCode configuration will be fetched from
	Uri string `json:"uri,omitempty"`
}

// VscodeTask Command providing the definition of a VsCode Task
type VscodeTask struct {

	// Optional map of free-form additional command attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// Defines the group this command is part of
	Group *Group `json:"group,omitempty"`

	// Mandatory identifier that allows referencing this command in composite commands, or from a parent, or in events.
	Id string `json:"id"`

	// Inlined content of the VsCode configuration
	Inlined string `json:"inlined,omitempty"`

	// Location as an absolute of relative URI the VsCode configuration will be fetched from
	Uri string `json:"uri,omitempty"`
}

// Zip Project's Zip source
type Zip struct {

	// Project's source location address. Should be URL for git and github located projects, or; file:// for zip
	Location string `json:"location,omitempty"`

	// Part of project to populate in the working directory.
	SparseCheckoutDir string `json:"sparseCheckoutDir,omitempty"`
}
