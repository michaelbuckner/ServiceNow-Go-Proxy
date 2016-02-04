package models

type SNOWUserData struct {
	Result []struct {
		Active              string `json:"active"`
		AgentStatus         string `json:"agent_status"`
		Building            string `json:"building"`
		CalendarIntegration string `json:"calendar_integration"`
		City                string `json:"city"`
		Company             struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"company"`
		CostCenter struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"cost_center"`
		Country            string `json:"country"`
		DateFormat         string `json:"date_format"`
		DefaultPerspective string `json:"default_perspective"`
		Department         struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"department"`
		EduStatus               string `json:"edu_status"`
		Email                   string `json:"email"`
		EmployeeNumber          string `json:"employee_number"`
		FailedAttempts          string `json:"failed_attempts"`
		FirstName               string `json:"first_name"`
		Gender                  string `json:"gender"`
		GeolocationTracked      string `json:"geolocation_tracked"`
		HomePhone               string `json:"home_phone"`
		InternalIntegrationUser string `json:"internal_integration_user"`
		Introduction            string `json:"introduction"`
		LastLogin               string `json:"last_login"`
		LastLoginDevice         string `json:"last_login_device"`
		LastLoginTime           string `json:"last_login_time"`
		LastName                string `json:"last_name"`
		LastPassword            string `json:"last_password"`
		LastPositionUpdate      string `json:"last_position_update"`
		Latitude                string `json:"latitude"`
		LdapServer              string `json:"ldap_server"`
		Location                struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"location"`
		LockedOut string `json:"locked_out"`
		Longitude string `json:"longitude"`
		Manager   struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"manager"`
		MiddleName         string `json:"middle_name"`
		MobilePhone        string `json:"mobile_phone"`
		Name               string `json:"name"`
		Notification       string `json:"notification"`
		OnSchedule         string `json:"on_schedule"`
		PasswordNeedsReset string `json:"password_needs_reset"`
		Phone              string `json:"phone"`
		Photo              string `json:"photo"`
		PreferredLanguage  string `json:"preferred_language"`
		Roles              string `json:"roles"`
		Schedule           string `json:"schedule"`
		Source             string `json:"source"`
		State              string `json:"state"`
		Street             string `json:"street"`
		SysClassName       string `json:"sys_class_name"`
		SysCreatedBy       string `json:"sys_created_by"`
		SysCreatedOn       string `json:"sys_created_on"`
		SysDomain          struct {
			Link  string `json:"link"`
			Value string `json:"value"`
		} `json:"sys_domain"`
		SysDomainPath        string `json:"sys_domain_path"`
		SysID                string `json:"sys_id"`
		SysModCount          string `json:"sys_mod_count"`
		SysTags              string `json:"sys_tags"`
		SysUpdatedBy         string `json:"sys_updated_by"`
		SysUpdatedOn         string `json:"sys_updated_on"`
		TimeFormat           string `json:"time_format"`
		TimeZone             string `json:"time_zone"`
		Title                string `json:"title"`
		UserName             string `json:"user_name"`
		UserPassword         string `json:"user_password"`
		Vip                  string `json:"vip"`
		WebServiceAccessOnly string `json:"web_service_access_only"`
		Zip                  string `json:"zip"`
	} `json:"result"`
}
