package sys11dbaassdk

import "time"

// v1

// requests

type CreatePostgreSQLRequestV1 struct {
	Organization      string                          `json:"-"`
	Project           string                          `json:"-"`
	ServiceConfig     *PSQLServiceConfigRequestV1     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigRequestV1 `json:"application_config,omitempty"`
	Name              string                          `json:"name,omitempty"`
	Description       string                          `json:"description,omitempty"`
	Verbose           bool                            `json:"-"`
}

type GetPostgreSQLRequestV1 struct {
	Organization string `json:"-"`
	Project      string `json:"-"`
	UUID         string `json:"-"`
	Verbose      bool   `json:"-"`
}

type GetPostgreSQLsRequestV1 struct {
	Organization string `json:"-"`
	Project      string `json:"-"`
	Verbose      bool   `json:"-"`
}

type DeletePostgreSQLRequestV1 struct {
	Organization string `json:"-"`
	Project      string `json:"-"`
	UUID         string `json:"-"`
	Verbose      bool   `json:"-"`
}

type UpdatePostgreSQLRequestV1 struct {
	Organization      string                                `json:"-"`
	Project           string                                `json:"-"`
	ServiceConfig     *PSQLServiceConfigUpdateRequestV1     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigUpdateRequestV1 `json:"application_config,omitempty"`
	Name              string                                `json:"name,omitempty"`
	Description       string                                `json:"description,omitempty"`
	UUID              string                                `json:"-"`
	Verbose           bool                                  `json:"-"`
}

// responses
type CreatePostgreSQLResponseV1 struct {
	ServiceConfig     *PSQLServiceConfigResponseV1     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigResponseV1 `json:"application_config,omitempty"`
	Status            string                           `json:"status,omitempty"`
	Phase             string                           `json:"phase,omitempty"`
	ResourceStatus    string                           `json:"resource_status,omitempty"`
	Name              string                           `json:"name,omitempty"`
	Description       string                           `json:"description,omitempty"`
	UUID              string                           `json:"uuid,omitempty"`
	CreatedBy         string                           `json:"created_by,omitempty"`
	CreatedAt         string                           `json:"created_at,omitempty"`
	LastModifiedBy    string                           `json:"last_modified_by,omitempty"`
	LastModifiedAt    string                           `json:"last_modified_at,omitempty"`
}

type GetPostgreSQLResponseV1 struct {
	ServiceConfig     *PSQLServiceConfigResponseV1     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigResponseV1 `json:"application_config,omitempty"`
	Status            string                           `json:"status,omitempty"`
	Phase             string                           `json:"phase,omitempty"`
	ResourceStatus    string                           `json:"resource_status,omitempty"`
	Name              string                           `json:"name,omitempty"`
	Description       string                           `json:"description,omitempty"`
	UUID              string                           `json:"uuid,omitempty"`
	CreatedBy         string                           `json:"created_by,omitempty"`
	CreatedAt         string                           `json:"created_at,omitempty"`
	LastModifiedBy    string                           `json:"last_modified_by,omitempty"`
	LastModifiedAt    string                           `json:"last_modified_at,omitempty"`
}

type GetPostgreSQLsResponseV1 []*GetPostgreSQLResponseV1

type DeletePostgreSQLResponseV1 struct {
}

type UpdatePostgreSQLResponseV1 struct {
}

// data structs

type PSQLServiceConfigRequestV1 struct {
	Disksize          *int                 `json:"disksize,omitempty"`
	MaintenanceWindow *MaintenanceWindowV1 `json:"maintenance_window,omitempty"`
	RemoteIPs         []string             `json:"remote_ips,omitempty"`
	Type              string               `json:"type,omitempty"`
	Flavor            string               `json:"flavor,omitempty"`
	Region            string               `json:"region,omitempty"`
}

type PSQLServiceConfigUpdateRequestV1 struct {
	Disksize          *int                 `json:"disksize,omitempty"`
	MaintenanceWindow *MaintenanceWindowV1 `json:"maintenance_window,omitempty"`
	RemoteIPs         []string             `json:"remote_ips,omitempty"`
	Type              string               `json:"type,omitempty"`
	Flavor            string               `json:"flavor,omitempty"`
}

type PSQLServiceConfigResponseV1 struct {
	Disksize          *int                 `json:"disksize,omitempty"`
	MaintenanceWindow *MaintenanceWindowV1 `json:"maintenance_window,omitempty"`
	RemoteIPs         []string             `json:"remote_ips,omitempty"`
	Type              string               `json:"type,omitempty"`
	Flavor            string               `json:"flavor,omitempty"`
	Region            string               `json:"region,omitempty"`
}

type PSQLApplicationConfigRequestV1 struct {
	Type             string                  `json:"type,omitempty"`
	Password         string                  `json:"password,omitempty"`
	Instances        *int                    `json:"instances,omitempty"`
	Version          string                  `json:"version,omitempty"`
	ScheduledBackups *PSQLScheduledBackupsV1 `json:"scheduled_backups,omitempty"`
	Recovery         *PSQLRecoveryV1         `json:"recovery,omitempty"`
}

type PSQLApplicationConfigUpdateRequestV1 struct {
	Password         string                  `json:"password,omitempty"`
	Instances        *int                    `json:"instances,omitempty"`
	Version          string                  `json:"version,omitempty"`
	ScheduledBackups *PSQLScheduledBackupsV1 `json:"scheduled_backups,omitempty"`
}

type PSQLApplicationConfigResponseV1 struct {
	Type             string                  `json:"type,omitempty"`
	Password         string                  `json:"password,omitempty"`
	Instances        *int                    `json:"instances,omitempty"`
	Version          string                  `json:"version,omitempty"`
	Hostname         string                  `json:"hostname,omitempty"`
	IPAddress        string                  `json:"ip_address,omitempty"`
	ScheduledBackups *PSQLScheduledBackupsV1 `json:"scheduled_backups,omitempty"`
	Recovery         *PSQLRecoveryV1         `json:"recovery,omitempty"`
}

type PSQLScheduledBackupsV1 struct {
	Schedule  *PSQLScheduledBackupsScheduleV1 `json:"schedule,omitempty"`
	Retention *int                            `json:"retention,omitempty"`
}

type PSQLScheduledBackupsScheduleV1 struct {
	Hour   *int `json:"hour,omitempty"`
	Minute *int `json:"minute,omitempty"`
}

type MaintenanceWindowV1 struct {
	DayOfWeek   *int `json:"day_of_week,omitempty"`
	StartHour   *int `json:"start_hour,omitempty"`
	StartMinute *int `json:"start_minute,omitempty"`
}

type PSQLRecoveryV1 struct {
	Source     string     `json:"source,omitempty"`
	Exclusive  bool       `json:"exclusive,omitempty"`
	TargetName string     `json:"targetName,omitempty"`
	TargetLSN  string     `json:"targetLSN,omitempty"`
	TargetXID  string     `json:"targetXID,omitempty"`
	TargetTime *time.Time `json:"targetTime,omitempty"`
}

// v2

// requests

type CreatePostgreSQLRequestV2 struct {
	Organization      string                          `json:"-"`
	Project           string                          `json:"-"`
	ServiceConfig     *PSQLServiceConfigRequestV2     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigRequestV2 `json:"application_config,omitempty"`
	Name              string                          `json:"name,omitempty"`
	Description       string                          `json:"description,omitempty"`
	Verbose           bool                            `json:"-"`
}

type GetPostgreSQLRequestV2 struct {
	Organization string `json:"-"`
	Project      string `json:"-"`
	UUID         string `json:"-"`
	Verbose      bool   `json:"-"`
}

type GetPostgreSQLsRequestV2 struct {
	Organization string `json:"-"`
	Project      string `json:"-"`
	Verbose      bool   `json:"-"`
}

type DeletePostgreSQLRequestV2 struct {
	Organization string `json:"-"`
	Project      string `json:"-"`
	UUID         string `json:"-"`
	Verbose      bool   `json:"-"`
}

type UpdatePostgreSQLRequestV2 struct {
	Organization      string                                `json:"-"`
	Project           string                                `json:"-"`
	ServiceConfig     *PSQLServiceConfigUpdateRequestV2     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigUpdateRequestV2 `json:"application_config,omitempty"`
	Name              string                                `json:"name,omitempty"`
	Description       string                                `json:"description,omitempty"`
	UUID              string                                `json:"-"`
	Verbose           bool                                  `json:"-"`
}

// responses
type CreatePostgreSQLResponseV2 struct {
	ServiceConfig     *PSQLServiceConfigResponseV2     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigResponseV2 `json:"application_config,omitempty"`
	Status            string                           `json:"status,omitempty"`
	Phase             string                           `json:"phase,omitempty"`
	ResourceStatus    string                           `json:"resource_status,omitempty"`
	Name              string                           `json:"name,omitempty"`
	Description       string                           `json:"description,omitempty"`
	UUID              string                           `json:"uuid,omitempty"`
	CreatedBy         string                           `json:"created_by,omitempty"`
	CreatedAt         string                           `json:"created_at,omitempty"`
	LastModifiedBy    string                           `json:"last_modified_by,omitempty"`
	LastModifiedAt    string                           `json:"last_modified_at,omitempty"`
}

type GetPostgreSQLResponseV2 struct {
	ServiceConfig     *PSQLServiceConfigResponseV2     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigResponseV2 `json:"application_config,omitempty"`
	Status            string                           `json:"status,omitempty"`
	Phase             string                           `json:"phase,omitempty"`
	ResourceStatus    string                           `json:"resource_status,omitempty"`
	Name              string                           `json:"name,omitempty"`
	Description       string                           `json:"description,omitempty"`
	UUID              string                           `json:"uuid,omitempty"`
	CreatedBy         string                           `json:"created_by,omitempty"`
	CreatedAt         string                           `json:"created_at,omitempty"`
	LastModifiedBy    string                           `json:"last_modified_by,omitempty"`
	LastModifiedAt    string                           `json:"last_modified_at,omitempty"`
}

type GetPostgreSQLsResponseV2 []*GetPostgreSQLResponseV2

type DeletePostgreSQLResponseV2 struct {
}

type UpdatePostgreSQLResponseV2 struct {
}

// data structs

type PSQLServiceConfigRequestV2 struct {
	Disksize          *int                 `json:"disksize,omitempty"`
	MaintenanceWindow *MaintenanceWindowV2 `json:"maintenance_window,omitempty"`
	Type              string               `json:"type,omitempty"`
	Flavor            string               `json:"flavor,omitempty"`
	Region            string               `json:"region,omitempty"`
}

type PSQLServiceConfigUpdateRequestV2 struct {
	Disksize          *int                 `json:"disksize,omitempty"`
	MaintenanceWindow *MaintenanceWindowV2 `json:"maintenance_window,omitempty"`
	Type              string               `json:"type,omitempty"`
	Flavor            string               `json:"flavor,omitempty"`
}

type PSQLServiceConfigResponseV2 struct {
	Disksize          *int                 `json:"disksize,omitempty"`
	MaintenanceWindow *MaintenanceWindowV2 `json:"maintenance_window,omitempty"`
	Type              string               `json:"type,omitempty"`
	Flavor            string               `json:"flavor,omitempty"`
	Region            string               `json:"region,omitempty"`
}

type PSQLApplicationConfigRequestV2 struct {
	Type                 string                             `json:"type,omitempty"`
	Password             string                             `json:"password,omitempty"`
	Instances            *int                               `json:"instances,omitempty"`
	Version              string                             `json:"version,omitempty"`
	ScheduledBackups     *PSQLScheduledBackupsV2            `json:"scheduled_backups,omitempty"`
	Recovery             *PSQLRecoveryV2                    `json:"recovery,omitempty"`
	PrivateNetworkConfig *PSQLPrivateNetworkConfigRequestV2 `json:"private_networking,omitempty"`
	PublicNetworkConfig  *PSQLPublicNetworkConfigRequestV2  `json:"public_networking,omitempty"`
}

type PSQLApplicationConfigUpdateRequestV2 struct {
	Password             string                             `json:"password,omitempty"`
	Instances            *int                               `json:"instances,omitempty"`
	Version              string                             `json:"version,omitempty"`
	ScheduledBackups     *PSQLScheduledBackupsV2            `json:"scheduled_backups,omitempty"`
	PrivateNetworkConfig *PSQLPrivateNetworkConfigRequestV2 `json:"private_networking,omitempty"`
	PublicNetworkConfig  *PSQLPublicNetworkConfigRequestV2  `json:"public_networking,omitempty"`
}

type PSQLApplicationConfigResponseV2 struct {
	Type                 string                              `json:"type,omitempty"`
	Password             string                              `json:"password,omitempty"`
	Instances            *int                                `json:"instances,omitempty"`
	Version              string                              `json:"version,omitempty"`
	ScheduledBackups     *PSQLScheduledBackupsV2             `json:"scheduled_backups,omitempty"`
	Recovery             *PSQLRecoveryV2                     `json:"recovery,omitempty"`
	PrivateNetworkConfig *PSQLPrivateNetworkConfigResponseV2 `json:"private_networking,omitempty"`
	PublicNetworkConfig  *PSQLPublicNetworkConfigResponseV2  `json:"public_networking,omitempty"`
}

type PSQLScheduledBackupsV2 struct {
	Schedule  *PSQLScheduledBackupsScheduleV2 `json:"schedule,omitempty"`
	Retention *int                            `json:"retention,omitempty"`
}

type PSQLScheduledBackupsScheduleV2 struct {
	Hour   *int `json:"hour,omitempty"`
	Minute *int `json:"minute,omitempty"`
}

type MaintenanceWindowV2 struct {
	DayOfWeek   *int `json:"day_of_week,omitempty"`
	StartHour   *int `json:"start_hour,omitempty"`
	StartMinute *int `json:"start_minute,omitempty"`
}

type PSQLRecoveryV2 struct {
	Source     string     `json:"source,omitempty"`
	Exclusive  bool       `json:"exclusive,omitempty"`
	TargetName string     `json:"targetName,omitempty"`
	TargetLSN  string     `json:"targetLSN,omitempty"`
	TargetXID  string     `json:"targetXID,omitempty"`
	TargetTime *time.Time `json:"targetTime,omitempty"`
}

type PSQLPrivateNetworkConfigRequestV2 struct {
	Enabled          bool      `json:"enabled"`
	AllowedCIDRs     *[]string `json:"allowed_cidrs,omitempty"`
	SharedSubnetCIDR *string   `json:"shared_subnet_cidr,omitempty"`
}

type PSQLPrivateNetworkConfigResponseV2 struct {
	Enabled          bool      `json:"enabled"`
	AllowedCIDRs     *[]string `json:"allowed_cidrs,omitempty"`
	SharedSubnetCIDR *string   `json:"shared_subnet_cidr,omitempty"`
	Hostname         string    `json:"hostname,omitempty"`
	IPAddress        string    `json:"ip_address,omitempty"`
	SharedSubnetID   string    `json:"shared_subnet_id,omitempty"`
	SharedNetworkID  string    `json:"shared_network_id,omitempty"`
}

type PSQLPublicNetworkConfigRequestV2 struct {
	Enabled      bool      `json:"enabled"`
	AllowedCIDRs *[]string `json:"allowed_cidrs,omitempty"`
}

type PSQLPublicNetworkConfigResponseV2 struct {
	Enabled      bool      `json:"enabled"`
	AllowedCIDRs *[]string `json:"allowed_cidrs,omitempty"`
	Hostname     string    `json:"hostname,omitempty"`
	IPAddress    string    `json:"ip_address,omitempty"`
}

// pointer

func StrPtr(s string) *string {
	sPtr := s
	return &sPtr
}

func IntPtr(i int) *int {
	iPtr := i
	return &iPtr
}

func Int64ToIntPtr(i int64) *int {
	iPtr := int(i)
	return &iPtr
}

func TimePtr(t time.Time) *time.Time {
	tPtr := t
	return &tPtr
}

// consts

const (
	SERVICE_TYPE_DB = "database"

	APP_TYPE_PSQL = "postgresql"

	STATE_NOT_READY = "ClusterIsNotReady"
	STATE_READY     = "ClusterIsReady"
)
