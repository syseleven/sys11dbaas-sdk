package sys11dbaassdk

import "time"

// requests

type CreatePostgreSQLRequest struct {
	Organization      string                        `json:"-"`
	Project           string                        `json:"-"`
	ServiceConfig     *PSQLServiceConfigRequest     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigRequest `json:"application_config,omitempty"`
	Name              string                        `json:"name,omitempty"`
	Description       string                        `json:"description,omitempty"`
}

type GetPostgreSQLRequest struct {
	Organization string `json:"-"`
	Project      string `json:"-"`
	UUID         string `json:"-"`
}

type GetPostgreSQLsRequest struct {
	Organization string `json:"-"`
	Project      string `json:"-"`
}

type DeletePostgreSQLRequest struct {
	Organization string `json:"-"`
	Project      string `json:"-"`
	UUID         string `json:"-"`
}

type UpdatePostgreSQLRequest struct {
	Organization      string                              `json:"-"`
	Project           string                              `json:"-"`
	ServiceConfig     *PSQLServiceConfigUpdateRequest     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigUpdateRequest `json:"application_config,omitempty"`
	Name              string                              `json:"name,omitempty"`
	Description       string                              `json:"description,omitempty"`
	UUID              string                              `json:"-"`
}

// responses
type CreatePostgreSQLResponse struct {
	ServiceConfig     *PSQLServiceConfigResponse     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigResponse `json:"application_config,omitempty"`
	Status            string                         `json:"status,omitempty"`
	Name              string                         `json:"name,omitempty"`
	Description       string                         `json:"description,omitempty"`
	UUID              string                         `json:"uuid,omitempty"`
	CreatedBy         string                         `json:"created_by,omitempty"`
	CreatedAt         string                         `json:"created_at,omitempty"`
	LastModifiedBy    string                         `json:"last_modified_by,omitempty"`
	LastModifiedAt    string                         `json:"last_modified_at,omitempty"`
}

type GetPostgreSQLResponse struct {
	ServiceConfig     *PSQLServiceConfigResponse     `json:"service_config,omitempty"`
	ApplicationConfig *PSQLApplicationConfigResponse `json:"application_config,omitempty"`
	Status            string                         `json:"status,omitempty"`
	Name              string                         `json:"name,omitempty"`
	Description       string                         `json:"description,omitempty"`
	UUID              string                         `json:"uuid,omitempty"`
	CreatedBy         string                         `json:"created_by,omitempty"`
	CreatedAt         string                         `json:"created_at,omitempty"`
	LastModifiedBy    string                         `json:"last_modified_by,omitempty"`
	LastModifiedAt    string                         `json:"last_modified_at,omitempty"`
}

type GetPostgreSQLsResponse []*GetPostgreSQLResponse

type DeletePostgreSQLResponse struct {
}

type UpdatePostgreSQLResponse struct {
}

// data structs

type PSQLServiceConfigRequest struct {
	Disksize          *int               `json:"disksize,omitempty"`
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty"`
	RemoteIPs         []string           `json:"remote_ips,omitempty"`
	Type              string             `json:"type,omitempty"`
	Flavor            string             `json:"flavor,omitempty"`
	Region            string             `json:"region,omitempty"`
}

type PSQLServiceConfigUpdateRequest struct {
	Disksize          *int               `json:"disksize,omitempty"`
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty"`
	RemoteIPs         []string           `json:"remote_ips,omitempty"`
	Type              string             `json:"type,omitempty"`
	Flavor            string             `json:"flavor,omitempty"`
}

type PSQLServiceConfigResponse struct {
	Disksize          *int               `json:"disksize,omitempty"`
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window,omitempty"`
	RemoteIPs         []string           `json:"remote_ips,omitempty"`
	Type              string             `json:"type,omitempty"`
	Flavor            string             `json:"flavor,omitempty"`
	Region            string             `json:"region,omitempty"`
}

type PSQLApplicationConfigRequest struct {
	Type             string                `json:"type,omitempty"`
	Password         string                `json:"password,omitempty"`
	Instances        *int                  `json:"instances,omitempty"`
	Version          string                `json:"version,omitempty"`
	ScheduledBackups *PSQLScheduledBackups `json:"scheduled_backups,omitempty"`
	Recovery         *PSQLRecovery         `json:"recovery,omitempty"`
}

type PSQLApplicationConfigUpdateRequest struct {
	Password         string                `json:"password,omitempty"`
	Instances        *int                  `json:"instances,omitempty"`
	Version          string                `json:"version,omitempty"`
	ScheduledBackups *PSQLScheduledBackups `json:"scheduled_backups,omitempty"`
}

type PSQLApplicationConfigResponse struct {
	Type             string                `json:"type,omitempty"`
	Password         string                `json:"password,omitempty"`
	Instances        *int                  `json:"instances,omitempty"`
	Version          string                `json:"version,omitempty"`
	Hostname         string                `json:"hostname,omitempty"`
	IPAddress        string                `json:"ip_address,omitempty"`
	ScheduledBackups *PSQLScheduledBackups `json:"scheduled_backups,omitempty"`
}

type PSQLScheduledBackups struct {
	Schedule  *PSQLScheduledBackupsSchedule `json:"schedule,omitempty"`
	Retention *int                          `json:"retention,omitempty"`
}

type PSQLScheduledBackupsSchedule struct {
	Hour   *int `json:"hour,omitempty"`
	Minute *int `json:"minute,omitempty"`
}

type MaintenanceWindow struct {
	DayOfWeek   *int `json:"day_of_week,omitempty"`
	StartHour   *int `json:"start_hour,omitempty"`
	StartMinute *int `json:"start_minute,omitempty"`
}

type PSQLRecovery struct {
	Source     string     `json:"source,omitempty"`
	Exclusive  bool       `json:"exclusive,omitempty"`
	TargetName string     `json:"targetName,omitempty"`
	TargetLSN  string     `json:"targetLSN,omitempty"`
	TargetXID  string     `json:"targetXID,omitempty"`
	TargetTime *time.Time `json:"targetTime,omitempty"`
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
)
