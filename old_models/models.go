package models

// ----------------------------
// --- Models v 0.9.0

// InputModel ...
type InputModel struct {
	MappedTo          string   `json:"mapped_to,omitempty" yaml:"mapped_to,omitempty"`
	Title             string   `json:"title,omitempty" yaml:"title,omitempty"`
	Description       string   `json:"description,omitempty" yaml:"description,omitempty"`
	Value             string   `json:"value,omitempty" yaml:"value,omitempty"`
	ValueOptions      []string `json:"value_options,omitempty" yaml:"value_options,omitempty"`
	IsRequired        bool     `json:"is_required,omitempty" yaml:"is_required,omitempty"`
	IsExpand          bool     `json:"is_expand,omitempty" yaml:"is_expand,omitempty"`
	IsDontChangeValue bool     `json:"is_dont_change_value,omitempty" yaml:"is_dont_change_value,omitempty"`
}

// OutputModel ...
type OutputModel struct {
	MappedTo    string `json:"mapped_to,omitempty" yaml:"mapped_to,omitempty"`
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

// StepModel ...
type StepModel struct {
	ID                  string            `json:"id" yaml:"id,omitempty"`
	SteplibSource       string            `json:"steplib_source" yaml:"steplib_source,omitempty"`
	VersionTag          string            `json:"version_tag" yaml:"version_tag,omitempty"`
	Name                string            `json:"name" yaml:"name" yaml:"name,omitempty"`
	Description         string            `json:"description,omitempty" yaml:"description,omitempty"`
	Website             string            `json:"website" yaml:"website"`
	ForkURL             string            `json:"fork_url,omitempty" yaml:"fork_url,omitempty"`
	Source              map[string]string `json:"source" yaml:"source"`
	HostOsTags          []string          `json:"host_os_tags,omitempty" yaml:"host_os_tags,omitempty"`
	ProjectTypeTags     []string          `json:"project_type_tags,omitempty" yaml:"project_type_tags,omitempty"`
	TypeTags            []string          `json:"type_tags,omitempty" yaml:"type_tags,omitempty"`
	IsRequiresAdminUser bool              `json:"is_requires_admin_user,omitempty" yaml:"is_requires_admin_user,omitempty"`
	IsAlwaysRun         bool              `json:"is_always_run,omitempty" yaml:"is_always_run,omitempty"`
	Inputs              []InputModel      `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Outputs             []OutputModel     `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}

// WorkflowModel ...
type WorkflowModel struct {
	FormatVersion string       `json:"format_version" yaml:"format_version,omitempty"`
	Environments  []InputModel `json:"environments" yaml:"environments,omitempty"`
	Steps         []StepModel  `json:"steps" yaml:"steps,omitempty"`
}

// StepGroupModel ...
type StepGroupModel struct {
	ID       string      `json:"id"`
	Versions []StepModel `json:"versions"`
	Latest   StepModel   `json:"latest"`
}

// StepHash ...
type StepHash map[string]StepGroupModel

// StepCollectionModel ...
type StepCollectionModel struct {
	FormatVersion        string   `json:"format_version"`
	GeneratedAtTimeStamp int64    `json:"generated_at_timestamp"`
	Steps                StepHash `json:"steps"`
	SteplibSource        string   `json:"steplib_source"`
}
