package livestatus

import "time"

// AcknowledgeHostProblem is generated from the nagios external command definition:
// Desfinition: ACKNOWLEDGE_HOST_PROBLEM;<host_name>;<sticky>;<notify>;<persistent>;<author>;<comment>
// Description:
//  Allows you to acknowledge the current problem for the specified host.  By acknowledging the current problem, future notifications (for the same host state) are disabled.  If the "sticky" option is set to two (2), the acknowledgement will remain until the host returns to an UP state.  Otherwise the acknowledgement will automatically be removed when the host changes state.  If the "notify" option is set to one (1), a notification will be sent out to contacts indicating that the current host problem has been acknowledged.  If the "persistent" option is set to one (1), the comment associated with the acknowledgement will survive across restarts of the Nagios process.  If not, the comment will be deleted the next time Nagios restarts.
func (c *Command) AcknowledgeHostProblem(host_name string, sticky bool, notify bool, persistent bool, author string, comment string) {
	c.Raw("ACKNOWLEDGE_HOST_PROBLEM")
	c.Hostname(host_name)
	c.Sticky(sticky)
	c.Notify(notify)
	c.Persistent(persistent)
	c.Author(author)
	c.Comment(comment)
}

// AcknowledgeSvcProblem is generated from the nagios external command definition:
// Desfinition: ACKNOWLEDGE_SVC_PROBLEM;<host_name>;<service_description>;<sticky>;<notify>;<persistent>;<author>;<comment>
// Description:
//  Allows you to acknowledge the current problem for the specified service.  By acknowledging the current problem, future notifications (for the same servicestate) are disabled.  If the "sticky" option is set to two (2), the acknowledgement will remain until the service returns to an OK state.  Otherwise the acknowledgement will automatically be removed when the service changes state.  If the "notify" option is set to one (1), a notification will be sent out to contacts indicating that the current service problem has been acknowledged.  If the "persistent" option is set to one (1), the comment associated with the acknowledgement will survive across restarts of the Nagios process.  If not, the comment will be deleted the next time Nagios restarts.
func (c *Command) AcknowledgeSvcProblem(host_name string, service_description string, sticky bool, notify bool, persistent bool, author string, comment string) {
	c.Raw("ACKNOWLEDGE_SVC_PROBLEM")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.Sticky(sticky)
	c.Notify(notify)
	c.Persistent(persistent)
	c.Author(author)
	c.Comment(comment)
}

// AddHostComment is generated from the nagios external command definition:
// Desfinition: ADD_HOST_COMMENT;<host_name>;<persistent>;<author>;<comment>
// Description:
//  Adds a comment to a particular host.  If the "persistent" field is set to zero (0), the comment will be deleted the next time Nagios is restarted.  Otherwise, the comment will persist across program restarts until it is deleted manually.
func (c *Command) AddHostComment(host_name string, persistent bool, author string, comment string) {
	c.Raw("ADD_HOST_COMMENT")
	c.Hostname(host_name)
	c.Persistent(persistent)
	c.Author(author)
	c.Comment(comment)
}

// AddSvcComment is generated from the nagios external command definition:
// Desfinition: ADD_SVC_COMMENT;<host_name>;<service_description>;<persistent>;<author>;<comment>
// Description:
//  Adds a comment to a particular service.  If the "persistent" field is set to zero (0), the comment will be deleted the next time Nagios is restarted.  Otherwise, the comment will persist across program restarts until it is deleted manually.
func (c *Command) AddSvcComment(host_name string, service_description string, persistent bool, author string, comment string) {
	c.Raw("ADD_SVC_COMMENT")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.Persistent(persistent)
	c.Author(author)
	c.Comment(comment)
}

// ChangeContactHostNotificationTimeperiod is generated from the nagios external command definition:
// Desfinition: CHANGE_CONTACT_HOST_NOTIFICATION_TIMEPERIOD;<contact_name>;<notification_timeperiod>
// Description:
//  Changes the host notification timeperiod for a particular contact to what is specified by the "notification_timeperiod" option.  The "notification_timeperiod" option should be the short name of the timeperiod that is to be used as the contact's host notification timeperiod.  The timeperiod must have been configured in Nagios before it was last (re)started.
func (c *Command) ChangeContactHostNotificationTimeperiod(contact_name string, notification_timeperiod string) {
	c.Raw("CHANGE_CONTACT_HOST_NOTIFICATION_TIMEPERIOD")
	c.ContactName(contact_name)
	c.NotificationTimePeriod(notification_timeperiod)
}

// ChangeContactModattr is generated from the nagios external command definition:
// Desfinition: CHANGE_CONTACT_MODATTR;<contact_name>;<value>
// Description:
//  This command changes the modified attributes value for the specified contact.  Modified attributes values are used by Nagios to determine which object properties should be retained across program restarts.  Thus, modifying the value of the attributes can affect data retention.  This is an advanced option and should only be used by people who are intimately familiar with the data retention logic in Nagios.
func (c *Command) ChangeContactModattr(contact_name string, value string) {
	c.Raw("CHANGE_CONTACT_MODATTR")
	c.ContactName(contact_name)
	c.Value(value)
}

// ChangeContactModhattr is generated from the nagios external command definition:
// Desfinition: CHANGE_CONTACT_MODHATTR;<contact_name>;<value>
// Description:
//  This command changes the modified host attributes value for the specified contact.  Modified attributes values are used by Nagios to determine which object properties should be retained across program restarts.  Thus, modifying the value of the attributes can affect data retention.  This is an advanced option and should only be used by people who are intimately familiar with the data retention logic in Nagios.
func (c *Command) ChangeContactModhattr(contact_name string, value string) {
	c.Raw("CHANGE_CONTACT_MODHATTR")
	c.ContactName(contact_name)
	c.Value(value)
}

// ChangeContactModsattr is generated from the nagios external command definition:
// Desfinition: CHANGE_CONTACT_MODSATTR;<contact_name>;<value>
// Description:
//  This command changes the modified service attributes value for the specified contact.  Modified attributes values are used by Nagios to determine which object properties should be retained across program restarts.  Thus, modifying the value of the attributes can affect data retention.  This is an advanced option and should only be used by people who are intimately familiar with the data retention logic in Nagios.
func (c *Command) ChangeContactModsattr(contact_name string, value string) {
	c.Raw("CHANGE_CONTACT_MODSATTR")
	c.ContactName(contact_name)
	c.Value(value)
}

// ChangeContactSvcNotificationTimeperiod is generated from the nagios external command definition:
// Desfinition: CHANGE_CONTACT_SVC_NOTIFICATION_TIMEPERIOD;<contact_name>;<notification_timeperiod>
// Description:
//  Changes the service notification timeperiod for a particular contact to what is specified by the "notification_timeperiod" option.  The "notification_timeperiod" option should be the short name of the timeperiod that is to be used as the contact's service notification timeperiod.  The timeperiod must have been configured in Nagios before it was last (re)started.
func (c *Command) ChangeContactSvcNotificationTimeperiod(contact_name string, notification_timeperiod string) {
	c.Raw("CHANGE_CONTACT_SVC_NOTIFICATION_TIMEPERIOD")
	c.ContactName(contact_name)
	c.NotificationTimePeriod(notification_timeperiod)
}

// ChangeCustomContactVar is generated from the nagios external command definition:
// Desfinition: CHANGE_CUSTOM_CONTACT_VAR;<contact_name>;<varname>;<varvalue>
// Description:
//  Changes the value of a custom contact variable.
func (c *Command) ChangeCustomContactVar(contact_name string, varname string, varvalue string) {
	c.Raw("CHANGE_CUSTOM_CONTACT_VAR")
	c.ContactName(contact_name)
	c.VarName(varname)
	c.VarValue(varvalue)
}

// ChangeCustomHostVar is generated from the nagios external command definition:
// Desfinition: CHANGE_CUSTOM_HOST_VAR;<host_name>;<varname>;<varvalue>
// Description:
//  Changes the value of a custom host variable.
func (c *Command) ChangeCustomHostVar(host_name string, varname string, varvalue string) {
	c.Raw("CHANGE_CUSTOM_HOST_VAR")
	c.Hostname(host_name)
	c.VarName(varname)
	c.VarValue(varvalue)
}

// ChangeCustomSvcVar is generated from the nagios external command definition:
// Desfinition: CHANGE_CUSTOM_SVC_VAR;<host_name>;<service_description>;<varname>;<varvalue>
// Description:
//  Changes the value of a custom service variable.
func (c *Command) ChangeCustomSvcVar(host_name string, service_description string, varname string, varvalue string) {
	c.Raw("CHANGE_CUSTOM_SVC_VAR")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.VarName(varname)
	c.VarValue(varvalue)
}

// ChangeGlobalHostEventHandler is generated from the nagios external command definition:
// Desfinition: CHANGE_GLOBAL_HOST_EVENT_HANDLER;<event_handler_command>
// Description:
//  Changes the global host event handler command to be that specified by the "event_handler_command" option.  The "event_handler_command" option specifies the short name of the command that should be used as the new host event handler.  The command must have been configured in Nagios before it was last (re)started.
func (c *Command) ChangeGlobalHostEventHandler(event_handler_command string) {
	c.Raw("CHANGE_GLOBAL_HOST_EVENT_HANDLER")
	c.EventHandlerCommand(event_handler_command)
}

// ChangeGlobalSvcEventHandler is generated from the nagios external command definition:
// Desfinition: CHANGE_GLOBAL_SVC_EVENT_HANDLER;<event_handler_command>
// Description:
//  Changes the global service event handler command to be that specified by the "event_handler_command" option.  The "event_handler_command" option specifies the short name of the command that should be used as the new service event handler.  The command must have been configured in Nagios before it was last (re)started.
func (c *Command) ChangeGlobalSvcEventHandler(event_handler_command string) {
	c.Raw("CHANGE_GLOBAL_SVC_EVENT_HANDLER")
	c.EventHandlerCommand(event_handler_command)
}

// ChangeHostCheckCommand is generated from the nagios external command definition:
// Desfinition: CHANGE_HOST_CHECK_COMMAND;<host_name>;<check_command>
// Description:
//  Changes the check command for a particular host to be that specified by the "check_command" option.  The "check_command" option specifies the short name of the command that should be used as the new host check command.  The command must have been configured in Nagios before it was last (re)started.
func (c *Command) ChangeHostCheckCommand(host_name string, check_command string) {
	c.Raw("CHANGE_HOST_CHECK_COMMAND")
	c.Hostname(host_name)
	c.CheckCommand(check_command)
}

// ChangeHostCheckTimeperiod is generated from the nagios external command definition:
// Desfinition: CHANGE_HOST_CHECK_TIMEPERIOD;<host_name>;<timeperiod>
// Description:
//  Changes the valid check period for the specified host.
func (c *Command) ChangeHostCheckTimeperiod(host_name string, timeperiod string) {
	c.Raw("CHANGE_HOST_CHECK_TIMEPERIOD")
	c.Hostname(host_name)
	c.TimePeriod(timeperiod)
}

// ChangeHostEventHandler is generated from the nagios external command definition:
// Desfinition: CHANGE_HOST_EVENT_HANDLER;<host_name>;<event_handler_command>
// Description:
//  Changes the event handler command for a particular host to be that specified by the "event_handler_command" option.  The "event_handler_command" option specifies the short name of the command that should be used as the new host event handler.  The command must have been configured in Nagios before it was last (re)started.
func (c *Command) ChangeHostEventHandler(host_name string, event_handler_command string) {
	c.Raw("CHANGE_HOST_EVENT_HANDLER")
	c.Hostname(host_name)
	c.EventHandlerCommand(event_handler_command)
}

// ChangeHostModattr is generated from the nagios external command definition:
// Desfinition: CHANGE_HOST_MODATTR;<host_name>;<value>
// Description:
//  This command changes the modified attributes value for the specified host.  Modified attributes values are used by Nagios to determine which object properties should be retained across program restarts.  Thus, modifying the value of the attributes can affect data retention.  This is an advanced option and should only be used by people who are intimately familiar with the data retention logic in Nagios.
func (c *Command) ChangeHostModattr(host_name string, value string) {
	c.Raw("CHANGE_HOST_MODATTR")
	c.Hostname(host_name)
	c.Value(value)
}

// ChangeMaxHostCheckAttempts is generated from the nagios external command definition:
// Desfinition: CHANGE_MAX_HOST_CHECK_ATTEMPTS;<host_name>;<check_attempts>
// Description:
//  Changes the maximum number of check attempts (retries) for a particular host.
func (c *Command) ChangeMaxHostCheckAttempts(host_name string, check_attempts int) {
	c.Raw("CHANGE_MAX_HOST_CHECK_ATTEMPTS")
	c.Hostname(host_name)
	c.CheckAttempts(check_attempts)
}

// ChangeMaxSvcCheckAttempts is generated from the nagios external command definition:
// Desfinition: CHANGE_MAX_SVC_CHECK_ATTEMPTS;<host_name>;<service_description>;<check_attempts>
// Description:
//  Changes the maximum number of check attempts (retries) for a particular service.
func (c *Command) ChangeMaxSvcCheckAttempts(host_name string, service_description string, check_attempts int) {
	c.Raw("CHANGE_MAX_SVC_CHECK_ATTEMPTS")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.CheckAttempts(check_attempts)
}

// ChangeNormalHostCheckInterval is generated from the nagios external command definition:
// Desfinition: CHANGE_NORMAL_HOST_CHECK_INTERVAL;<host_name>;<check_interval>
// Description:
//  Changes the normal (regularly scheduled) check interval for a particular host.
func (c *Command) ChangeNormalHostCheckInterval(host_name string, check_interval time.Duration) {
	c.Raw("CHANGE_NORMAL_HOST_CHECK_INTERVAL")
	c.Hostname(host_name)
	c.Duration(check_interval)
}

// ChangeNormalSvcCheckInterval is generated from the nagios external command definition:
// Desfinition: CHANGE_NORMAL_SVC_CHECK_INTERVAL;<host_name>;<service_description>;<check_interval>
// Description:
//  Changes the normal (regularly scheduled) check interval for a particular service
func (c *Command) ChangeNormalSvcCheckInterval(host_name string, service_description string, check_interval time.Duration) {
	c.Raw("CHANGE_NORMAL_SVC_CHECK_INTERVAL")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.Duration(check_interval)
}

// ChangeRetryHostCheckInterval is generated from the nagios external command definition:
// Desfinition: CHANGE_RETRY_HOST_CHECK_INTERVAL;<host_name>;<service_description>;<check_interval>
// Description:
//  Changes the retry check interval for a particular host.
func (c *Command) ChangeRetryHostCheckInterval(host_name string, service_description string, check_interval time.Duration) {
	c.Raw("CHANGE_RETRY_HOST_CHECK_INTERVAL")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.Duration(check_interval)
}

// ChangeRetrySvcCheckInterval is generated from the nagios external command definition:
// Desfinition: CHANGE_RETRY_SVC_CHECK_INTERVAL;<host_name>;<service_description>;<check_interval>
// Description:
//  Changes the retry check interval for a particular service.
func (c *Command) ChangeRetrySvcCheckInterval(host_name string, service_description string, check_interval time.Duration) {
	c.Raw("CHANGE_RETRY_SVC_CHECK_INTERVAL")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.Duration(check_interval)
}

// ChangeSvcCheckCommand is generated from the nagios external command definition:
// Desfinition: CHANGE_SVC_CHECK_COMMAND;<host_name>;<service_description>;<check_command>
// Description:
//  Changes the check command for a particular service to be that specified by the "check_command" option.  The "check_command" option specifies the short name of the command that should be used as the new service check command.  The command must have been configured in Nagios before it was last (re)started.
func (c *Command) ChangeSvcCheckCommand(host_name string, service_description string, check_command string) {
	c.Raw("CHANGE_SVC_CHECK_COMMAND")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.CheckCommand(check_command)
}

// ChangeSvcCheckTimeperiod is generated from the nagios external command definition:
// Desfinition: CHANGE_SVC_CHECK_TIMEPERIOD;<host_name>;<service_description>;<check_timeperiod>
// Description:
//  Changes the check timeperiod for a particular service to what is specified by the "check_timeperiod" option.  The "check_timeperiod" option should be the short name of the timeperod that is to be used as the service check timeperiod.  The timeperiod must have been configured in Nagios before it was last (re)started.
func (c *Command) ChangeSvcCheckTimeperiod(host_name string, service_description string, check_timeperiod string) {
	c.Raw("CHANGE_SVC_CHECK_TIMEPERIOD")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.CheckTimePeriod(check_timeperiod)
}

// ChangeSvcEventHandler is generated from the nagios external command definition:
// Desfinition: CHANGE_SVC_EVENT_HANDLER;<host_name>;<service_description>;<event_handler_command>
// Description:
//  Changes the event handler command for a particular service to be that specified by the "event_handler_command" option.  The "event_handler_command" option specifies the short name of the command that should be used as the new service event handler.  The command must have been configured in Nagios before it was last (re)started.
func (c *Command) ChangeSvcEventHandler(host_name string, service_description string, event_handler_command string) {
	c.Raw("CHANGE_SVC_EVENT_HANDLER")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.EventHandlerCommand(event_handler_command)
}

// ChangeSvcModattr is generated from the nagios external command definition:
// Desfinition: CHANGE_SVC_MODATTR;<host_name>;<service_description>;<value>
// Description:
//  This command changes the modified attributes value for the specified service.  Modified attributes values are used by Nagios to determine which object properties should be retained across program restarts.  Thus, modifying the value of the attributes can affect data retention.  This is an advanced option and should only be used by people who are intimately familiar with the data retention logic in Nagios.
func (c *Command) ChangeSvcModattr(host_name string, service_description string, value string) {
	c.Raw("CHANGE_SVC_MODATTR")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.Value(value)
}

// ChangeSvcNotificationTimeperiod is generated from the nagios external command definition:
// Desfinition: CHANGE_SVC_NOTIFICATION_TIMEPERIOD;<host_name>;<service_description>;<notification_timeperiod>
// Description:
//  Changes the notification timeperiod for a particular service to what is specified by the "notification_timeperiod" option.  The "notification_timeperiod" option should be the short name of the timeperiod that is to be used as the service notification timeperiod.  The timeperiod must have been configured in Nagios before it was last (re)started.
func (c *Command) ChangeSvcNotificationTimeperiod(host_name string, service_description string, notification_timeperiod string) {
	c.Raw("CHANGE_SVC_NOTIFICATION_TIMEPERIOD")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.NotificationTimePeriod(notification_timeperiod)
}

// DelayHostNotification is generated from the nagios external command definition:
// Desfinition: DELAY_HOST_NOTIFICATION;<host_name>;<notification_time>
// Description:
//  Delays the next notification for a parciular service until "notification_time".  The "notification_time" argument is specified in time_t format (seconds since the UNIX epoch).  Note that this will only have an affect if the service stays in the same problem state that it is currently in.  If the service changes to another state, a new notification may go out before the time you specify in the "notification_time" argument.
func (c *Command) DelayHostNotification(host_name string, notification_time time.Time) {
	c.Raw("DELAY_HOST_NOTIFICATION")
	c.Hostname(host_name)
	c.NotificationTime(notification_time)
}

// DelaySvcNotification is generated from the nagios external command definition:
// Desfinition: DELAY_SVC_NOTIFICATION;<host_name>;<service_description>;<notification_time>
// Description:
//  Delays the next notification for a parciular service until "notification_time".  The "notification_time" argument is specified in time_t format (seconds since the UNIX epoch).  Note that this will only have an affect if the service stays in the same problem state that it is currently in.  If the service changes to another state, a new notification may go out before the time you specify in the "notification_time" argument.
func (c *Command) DelaySvcNotification(host_name string, service_description string, notification_time time.Time) {
	c.Raw("DELAY_SVC_NOTIFICATION")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.NotificationTime(notification_time)
}

// DelAllHostComments is generated from the nagios external command definition:
// Desfinition: DEL_ALL_HOST_COMMENTS;<host_name>
// Description:
//  Deletes all comments assocated with a particular host.
func (c *Command) DelAllHostComments(host_name string) {
	c.Raw("DEL_ALL_HOST_COMMENTS")
	c.Hostname(host_name)
}

// DelAllSvcComments is generated from the nagios external command definition:
// Desfinition: DEL_ALL_SVC_COMMENTS;<host_name>;<service_description>
// Description:
//  Deletes all comments associated with a particular service.
func (c *Command) DelAllSvcComments(host_name string, service_description string) {
	c.Raw("DEL_ALL_SVC_COMMENTS")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// DelHostComment is generated from the nagios external command definition:
// Desfinition: DEL_HOST_COMMENT;<comment_id>
// Description:
//  Deletes a host comment.  The id number of the comment that is to be deleted must be specified.
func (c *Command) DelHostComment(comment_id int) {
	c.Raw("DEL_HOST_COMMENT")
	c.CommentID(comment_id)
}

// DelHostDowntime is generated from the nagios external command definition:
// Desfinition: DEL_HOST_DOWNTIME;<downtime_id>
// Description:
//  Deletes the host downtime entry that has an ID number matching the "downtime_id" argument.  If the downtime is currently in effect, the host will come out of scheduled downtime (as long as there are no other overlapping active downtime entries).
func (c *Command) DelHostDowntime(downtime_id int) {
	c.Raw("DEL_HOST_DOWNTIME")
	c.DowntimeID(downtime_id)
}

// DelSvcComment is generated from the nagios external command definition:
// Desfinition: DEL_SVC_COMMENT;<comment_id>
// Description:
//  Deletes a service comment.  The id number of the comment that is to be deleted must be specified.
func (c *Command) DelSvcComment(comment_id int) {
	c.Raw("DEL_SVC_COMMENT")
	c.CommentID(comment_id)
}

// DelSvcDowntime is generated from the nagios external command definition:
// Desfinition: DEL_SVC_DOWNTIME;<downtime_id>
// Description:
//  Deletes the service downtime entry that has an ID number matching the "downtime_id" argument.  If the downtime is currently in effect, the service will come out of scheduled downtime (as long as there are no other overlapping active downtime entries).
func (c *Command) DelSvcDowntime(downtime_id int) {
	c.Raw("DEL_SVC_DOWNTIME")
	c.DowntimeID(downtime_id)
}

// DisableAllNotificationsBeyondHost is generated from the nagios external command definition:
// Desfinition: DISABLE_ALL_NOTIFICATIONS_BEYOND_HOST;<host_name>
// Description:
//  Disables notifications for all hosts and services "beyond" (e.g. on all child hosts of) the specified host.  The current notification setting for the specified host is not affected.
func (c *Command) DisableAllNotificationsBeyondHost(host_name string) {
	c.Raw("DISABLE_ALL_NOTIFICATIONS_BEYOND_HOST")
	c.Hostname(host_name)
}

// DisableContactgroupHostNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_CONTACTGROUP_HOST_NOTIFICATIONS;<contactgroup_name>
// Description:
//  Disables host notifications for all contacts in a particular contactgroup.
func (c *Command) DisableContactgroupHostNotifications(contactgroup_name string) {
	c.Raw("DISABLE_CONTACTGROUP_HOST_NOTIFICATIONS")
	c.ContactGroupName(contactgroup_name)
}

// DisableContactgroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_CONTACTGROUP_SVC_NOTIFICATIONS;<contactgroup_name>
// Description:
//  Disables service notifications for all contacts in a particular contactgroup.
func (c *Command) DisableContactgroupSvcNotifications(contactgroup_name string) {
	c.Raw("DISABLE_CONTACTGROUP_SVC_NOTIFICATIONS")
	c.ContactGroupName(contactgroup_name)
}

// DisableContactHostNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_CONTACT_HOST_NOTIFICATIONS;<contact_name>
// Description:
//  Disables host notifications for a particular contact.
func (c *Command) DisableContactHostNotifications(contact_name string) {
	c.Raw("DISABLE_CONTACT_HOST_NOTIFICATIONS")
	c.ContactName(contact_name)
}

// DisableContactSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_CONTACT_SVC_NOTIFICATIONS;<contact_name>
// Description:
//  Disables service notifications for a particular contact.
func (c *Command) DisableContactSvcNotifications(contact_name string) {
	c.Raw("DISABLE_CONTACT_SVC_NOTIFICATIONS")
	c.ContactName(contact_name)
}

// DisableEventHandlers is generated from the nagios external command definition:
// Desfinition: DISABLE_EVENT_HANDLERS
// Description:
//  Disables host and service event handlers on a program-wide basis.
func (c *Command) DisableEventHandlers() {
	c.Raw("DISABLE_EVENT_HANDLERS")
}

// DisableFailurePrediction is generated from the nagios external command definition:
// Desfinition: DISABLE_FAILURE_PREDICTION
// Description:
//  Disables failure prediction on a program-wide basis.  This feature is not currently implemented in Nagios.
func (c *Command) DisableFailurePrediction() {
	c.Raw("DISABLE_FAILURE_PREDICTION")
}

// DisableFlapDetection is generated from the nagios external command definition:
// Desfinition: DISABLE_FLAP_DETECTION
// Description:
//  Disables host and service flap detection on a program-wide basis.
func (c *Command) DisableFlapDetection() {
	c.Raw("DISABLE_FLAP_DETECTION")
}

// DisableHostgroupHostChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_HOST_CHECKS;<hostgroup_name>
// Description:
//  Disables active checks for all hosts in a particular hostgroup.
func (c *Command) DisableHostgroupHostChecks(hostgroup_name string) {
	c.Raw("DISABLE_HOSTGROUP_HOST_CHECKS")
	c.HostGroupName(hostgroup_name)
}

// DisableHostgroupHostNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_HOST_NOTIFICATIONS;<hostgroup_name>
// Description:
//  Disables notifications for all hosts in a particular hostgroup.  This does not disable notifications for the services associated with the hosts in the hostgroup - see the DISABLE_HOSTGROUP_SVC_NOTIFICATIONS command for that.
func (c *Command) DisableHostgroupHostNotifications(hostgroup_name string) {
	c.Raw("DISABLE_HOSTGROUP_HOST_NOTIFICATIONS")
	c.HostGroupName(hostgroup_name)
}

// DisableHostgroupPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_PASSIVE_HOST_CHECKS;<hostgroup_name>
// Description:
//  Disables passive checks for all hosts in a particular hostgroup.
func (c *Command) DisableHostgroupPassiveHostChecks(hostgroup_name string) {
	c.Raw("DISABLE_HOSTGROUP_PASSIVE_HOST_CHECKS")
	c.HostGroupName(hostgroup_name)
}

// DisableHostgroupPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_PASSIVE_SVC_CHECKS;<hostgroup_name>
// Description:
//  Disables passive checks for all services associated with hosts in a particular hostgroup.
func (c *Command) DisableHostgroupPassiveSvcChecks(hostgroup_name string) {
	c.Raw("DISABLE_HOSTGROUP_PASSIVE_SVC_CHECKS")
	c.HostGroupName(hostgroup_name)
}

// DisableHostgroupSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_SVC_CHECKS;<hostgroup_name>
// Description:
//  Disables active checks for all services associated with hosts in a particular hostgroup.
func (c *Command) DisableHostgroupSvcChecks(hostgroup_name string) {
	c.Raw("DISABLE_HOSTGROUP_SVC_CHECKS")
	c.HostGroupName(hostgroup_name)
}

// DisableHostgroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_HOSTGROUP_SVC_NOTIFICATIONS;<hostgroup_name>
// Description:
//  Disables notifications for all services associated with hosts in a particular hostgroup.  This does not disable notifications for the hosts in the hostgroup - see the DISABLE_HOSTGROUP_HOST_NOTIFICATIONS command for that.
func (c *Command) DisableHostgroupSvcNotifications(hostgroup_name string) {
	c.Raw("DISABLE_HOSTGROUP_SVC_NOTIFICATIONS")
	c.HostGroupName(hostgroup_name)
}

// DisableHostAndChildNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_AND_CHILD_NOTIFICATIONS;<host_name>
// Description:
//  Disables notifications for the specified host, as well as all hosts "beyond" (e.g. on all child hosts of) the specified host.
func (c *Command) DisableHostAndChildNotifications(host_name string) {
	c.Raw("DISABLE_HOST_AND_CHILD_NOTIFICATIONS")
	c.Hostname(host_name)
}

// DisableHostCheck is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_CHECK;<host_name>
// Description:
//  Disables (regularly scheduled and on-demand) active checks of the specified host.
func (c *Command) DisableHostCheck(host_name string) {
	c.Raw("DISABLE_HOST_CHECK")
	c.Hostname(host_name)
}

// DisableHostEventHandler is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_EVENT_HANDLER;<host_name>
// Description:
//  Disables the event handler for the specified host.
func (c *Command) DisableHostEventHandler(host_name string) {
	c.Raw("DISABLE_HOST_EVENT_HANDLER")
	c.Hostname(host_name)
}

// DisableHostFlapDetection is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_FLAP_DETECTION;<host_name>
// Description:
//  Disables flap detection for the specified host.
func (c *Command) DisableHostFlapDetection(host_name string) {
	c.Raw("DISABLE_HOST_FLAP_DETECTION")
	c.Hostname(host_name)
}

// DisableHostFreshnessChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_FRESHNESS_CHECKS
// Description:
//  Disables freshness checks of all hosts on a program-wide basis.
func (c *Command) DisableHostFreshnessChecks() {
	c.Raw("DISABLE_HOST_FRESHNESS_CHECKS")
}

// DisableHostNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_NOTIFICATIONS;<host_name>
// Description:
//  Disables notifications for a particular host.
func (c *Command) DisableHostNotifications(host_name string) {
	c.Raw("DISABLE_HOST_NOTIFICATIONS")
	c.Hostname(host_name)
}

// DisableHostSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_SVC_CHECKS;<host_name>
// Description:
//  Enables active checks of all services on the specified host.
func (c *Command) DisableHostSvcChecks(host_name string) {
	c.Raw("DISABLE_HOST_SVC_CHECKS")
	c.Hostname(host_name)
}

// DisableHostSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_HOST_SVC_NOTIFICATIONS;<host_name>
// Description:
//  Disables notifications for all services on the specified host.
func (c *Command) DisableHostSvcNotifications(host_name string) {
	c.Raw("DISABLE_HOST_SVC_NOTIFICATIONS")
	c.Hostname(host_name)
}

// DisableNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_NOTIFICATIONS
// Description:
//  Disables host and service notifications on a program-wide basis.
func (c *Command) DisableNotifications() {
	c.Raw("DISABLE_NOTIFICATIONS")
}

// DisablePassiveHostChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_PASSIVE_HOST_CHECKS;<host_name>
// Description:
//  Disables acceptance and processing of passive host checks for the specified host.
func (c *Command) DisablePassiveHostChecks(host_name string) {
	c.Raw("DISABLE_PASSIVE_HOST_CHECKS")
	c.Hostname(host_name)
}

// DisablePassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_PASSIVE_SVC_CHECKS;<host_name>;<service_description>
// Description:
//  Disables passive checks for the specified service.
func (c *Command) DisablePassiveSvcChecks(host_name string, service_description string) {
	c.Raw("DISABLE_PASSIVE_SVC_CHECKS")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// DisablePerformanceData is generated from the nagios external command definition:
// Desfinition: DISABLE_PERFORMANCE_DATA
// Description:
//  Disables the processing of host and service performance data on a program-wide basis.
func (c *Command) DisablePerformanceData() {
	c.Raw("DISABLE_PERFORMANCE_DATA")
}

// DisableServicegroupHostChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_HOST_CHECKS;<servicegroup_name>
// Description:
//  Disables active checks for all hosts that have services that are members of a particular hostgroup.
func (c *Command) DisableServicegroupHostChecks(servicegroup_name string) {
	c.Raw("DISABLE_SERVICEGROUP_HOST_CHECKS")
	c.ServiceGroupName(servicegroup_name)
}

// DisableServicegroupHostNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_HOST_NOTIFICATIONS;<servicegroup_name>
// Description:
//  Disables notifications for all hosts that have services that are members of a particular servicegroup.
func (c *Command) DisableServicegroupHostNotifications(servicegroup_name string) {
	c.Raw("DISABLE_SERVICEGROUP_HOST_NOTIFICATIONS")
	c.ServiceGroupName(servicegroup_name)
}

// DisableServicegroupPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_PASSIVE_HOST_CHECKS;<servicegroup_name>
// Description:
//  Disables the acceptance and processing of passive checks for all hosts that have services that are members of a particular service group.
func (c *Command) DisableServicegroupPassiveHostChecks(servicegroup_name string) {
	c.Raw("DISABLE_SERVICEGROUP_PASSIVE_HOST_CHECKS")
	c.ServiceGroupName(servicegroup_name)
}

// DisableServicegroupPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_PASSIVE_SVC_CHECKS;<servicegroup_name>
// Description:
//  Disables the acceptance and processing of passive checks for all services in a particular servicegroup.
func (c *Command) DisableServicegroupPassiveSvcChecks(servicegroup_name string) {
	c.Raw("DISABLE_SERVICEGROUP_PASSIVE_SVC_CHECKS")
	c.ServiceGroupName(servicegroup_name)
}

// DisableServicegroupSvcChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_SVC_CHECKS;<servicegroup_name>
// Description:
//  Disables active checks for all services in a particular servicegroup.
func (c *Command) DisableServicegroupSvcChecks(servicegroup_name string) {
	c.Raw("DISABLE_SERVICEGROUP_SVC_CHECKS")
	c.ServiceGroupName(servicegroup_name)
}

// DisableServicegroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICEGROUP_SVC_NOTIFICATIONS;<servicegroup_name>
// Description:
//  Disables notifications for all services that are members of a particular servicegroup.
func (c *Command) DisableServicegroupSvcNotifications(servicegroup_name string) {
	c.Raw("DISABLE_SERVICEGROUP_SVC_NOTIFICATIONS")
	c.ServiceGroupName(servicegroup_name)
}

// DisableServiceFlapDetection is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICE_FLAP_DETECTION;<host_name>;<service_description>
// Description:
//  Disables flap detection for the specified service.
func (c *Command) DisableServiceFlapDetection(host_name string, service_description string) {
	c.Raw("DISABLE_SERVICE_FLAP_DETECTION")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// DisableServiceFreshnessChecks is generated from the nagios external command definition:
// Desfinition: DISABLE_SERVICE_FRESHNESS_CHECKS
// Description:
//  Disables freshness checks of all services on a program-wide basis.
func (c *Command) DisableServiceFreshnessChecks() {
	c.Raw("DISABLE_SERVICE_FRESHNESS_CHECKS")
}

// DisableSvcCheck is generated from the nagios external command definition:
// Desfinition: DISABLE_SVC_CHECK;<host_name>;<service_description>
// Description:
//  Disables active checks for a particular service.
func (c *Command) DisableSvcCheck(host_name string, service_description string) {
	c.Raw("DISABLE_SVC_CHECK")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// DisableSvcEventHandler is generated from the nagios external command definition:
// Desfinition: DISABLE_SVC_EVENT_HANDLER;<host_name>;<service_description>
// Description:
//  Disables the event handler for the specified service.
func (c *Command) DisableSvcEventHandler(host_name string, service_description string) {
	c.Raw("DISABLE_SVC_EVENT_HANDLER")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// DisableSvcFlapDetection is generated from the nagios external command definition:
// Desfinition: DISABLE_SVC_FLAP_DETECTION;<host_name>;<service_description>
// Description:
//  Disables flap detection for the specified service.
func (c *Command) DisableSvcFlapDetection(host_name string, service_description string) {
	c.Raw("DISABLE_SVC_FLAP_DETECTION")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// DisableSvcNotifications is generated from the nagios external command definition:
// Desfinition: DISABLE_SVC_NOTIFICATIONS;<host_name>;<service_description>
// Description:
//  Disables notifications for a particular service.
func (c *Command) DisableSvcNotifications(host_name string, service_description string) {
	c.Raw("DISABLE_SVC_NOTIFICATIONS")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// EnableAllNotificationsBeyondHost is generated from the nagios external command definition:
// Desfinition: ENABLE_ALL_NOTIFICATIONS_BEYOND_HOST;<host_name>
// Description:
//  Enables notifications for all hosts and services "beyond" (e.g. on all child hosts of) the specified host.  The current notification setting for the specified host is not affected.  Notifications will only be sent out for these hosts and services if notifications are also enabled on a program-wide basis.
func (c *Command) EnableAllNotificationsBeyondHost(host_name string) {
	c.Raw("ENABLE_ALL_NOTIFICATIONS_BEYOND_HOST")
	c.Hostname(host_name)
}

// EnableContactgroupHostNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_CONTACTGROUP_HOST_NOTIFICATIONS;<contactgroup_name>
// Description:
//  Enables host notifications for all contacts in a particular contactgroup.
func (c *Command) EnableContactgroupHostNotifications(contactgroup_name string) {
	c.Raw("ENABLE_CONTACTGROUP_HOST_NOTIFICATIONS")
	c.ContactGroupName(contactgroup_name)
}

// EnableContactgroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_CONTACTGROUP_SVC_NOTIFICATIONS;<contactgroup_name>
// Description:
//  Enables service notifications for all contacts in a particular contactgroup.
func (c *Command) EnableContactgroupSvcNotifications(contactgroup_name string) {
	c.Raw("ENABLE_CONTACTGROUP_SVC_NOTIFICATIONS")
	c.ContactGroupName(contactgroup_name)
}

// EnableContactHostNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_CONTACT_HOST_NOTIFICATIONS;<contact_name>
// Description:
//  Enables host notifications for a particular contact.
func (c *Command) EnableContactHostNotifications(contact_name string) {
	c.Raw("ENABLE_CONTACT_HOST_NOTIFICATIONS")
	c.ContactName(contact_name)
}

// EnableContactSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_CONTACT_SVC_NOTIFICATIONS;<contact_name>
// Description:
//  Disables service notifications for a particular contact.
func (c *Command) EnableContactSvcNotifications(contact_name string) {
	c.Raw("ENABLE_CONTACT_SVC_NOTIFICATIONS")
	c.ContactName(contact_name)
}

// EnableEventHandlers is generated from the nagios external command definition:
// Desfinition: ENABLE_EVENT_HANDLERS
// Description:
//  Enables host and service event handlers on a program-wide basis.
func (c *Command) EnableEventHandlers() {
	c.Raw("ENABLE_EVENT_HANDLERS")
}

// EnableFailurePrediction is generated from the nagios external command definition:
// Desfinition: ENABLE_FAILURE_PREDICTION
// Description:
//  Enables failure prediction on a program-wide basis.  This feature is not currently implemented in Nagios.
func (c *Command) EnableFailurePrediction() {
	c.Raw("ENABLE_FAILURE_PREDICTION")
}

// EnableFlapDetection is generated from the nagios external command definition:
// Desfinition: ENABLE_FLAP_DETECTION
// Description:
//  Enables host and service flap detection on a program-wide basis.
func (c *Command) EnableFlapDetection() {
	c.Raw("ENABLE_FLAP_DETECTION")
}

// EnableHostgroupHostChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_HOST_CHECKS;<hostgroup_name>
// Description:
//  Enables active checks for all hosts in a particular hostgroup.
func (c *Command) EnableHostgroupHostChecks(hostgroup_name string) {
	c.Raw("ENABLE_HOSTGROUP_HOST_CHECKS")
	c.HostGroupName(hostgroup_name)
}

// EnableHostgroupHostNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_HOST_NOTIFICATIONS;<hostgroup_name>
// Description:
//  Enables notifications for all hosts in a particular hostgroup.  This does not enable notifications for the services associated with the hosts in the hostgroup - see the ENABLE_HOSTGROUP_SVC_NOTIFICATIONS command for that.  In order for notifications to be sent out for these hosts, notifications must be enabled on a program-wide basis as well.
func (c *Command) EnableHostgroupHostNotifications(hostgroup_name string) {
	c.Raw("ENABLE_HOSTGROUP_HOST_NOTIFICATIONS")
	c.HostGroupName(hostgroup_name)
}

// EnableHostgroupPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_PASSIVE_HOST_CHECKS;<hostgroup_name>
// Description:
//  Enables passive checks for all hosts in a particular hostgroup.
func (c *Command) EnableHostgroupPassiveHostChecks(hostgroup_name string) {
	c.Raw("ENABLE_HOSTGROUP_PASSIVE_HOST_CHECKS")
	c.HostGroupName(hostgroup_name)
}

// EnableHostgroupPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_PASSIVE_SVC_CHECKS;<hostgroup_name>
// Description:
//  Enables passive checks for all services associated with hosts in a particular hostgroup.
func (c *Command) EnableHostgroupPassiveSvcChecks(hostgroup_name string) {
	c.Raw("ENABLE_HOSTGROUP_PASSIVE_SVC_CHECKS")
	c.HostGroupName(hostgroup_name)
}

// EnableHostgroupSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_SVC_CHECKS;<hostgroup_name>
// Description:
//  Enables active checks for all services associated with hosts in a particular hostgroup.
func (c *Command) EnableHostgroupSvcChecks(hostgroup_name string) {
	c.Raw("ENABLE_HOSTGROUP_SVC_CHECKS")
	c.HostGroupName(hostgroup_name)
}

// EnableHostgroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_HOSTGROUP_SVC_NOTIFICATIONS;<hostgroup_name>
// Description:
//  Enables notifications for all services that are associated with hosts in a particular hostgroup.  This does not enable notifications for the hosts in the hostgroup - see the ENABLE_HOSTGROUP_HOST_NOTIFICATIONS command for that.  In order for notifications to be sent out for these services, notifications must be enabled on a program-wide basis as well.
func (c *Command) EnableHostgroupSvcNotifications(hostgroup_name string) {
	c.Raw("ENABLE_HOSTGROUP_SVC_NOTIFICATIONS")
	c.HostGroupName(hostgroup_name)
}

// EnableHostAndChildNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_AND_CHILD_NOTIFICATIONS;<host_name>
// Description:
//  Enables notifications for the specified host, as well as all hosts "beyond" (e.g. on all child hosts of) the specified host.  Notifications will only be sent out for these hosts if notifications are also enabled on a program-wide basis.
func (c *Command) EnableHostAndChildNotifications(host_name string) {
	c.Raw("ENABLE_HOST_AND_CHILD_NOTIFICATIONS")
	c.Hostname(host_name)
}

// EnableHostCheck is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_CHECK;<host_name>
// Description:
//  Enables (regularly scheduled and on-demand) active checks of the specified host.
func (c *Command) EnableHostCheck(host_name string) {
	c.Raw("ENABLE_HOST_CHECK")
	c.Hostname(host_name)
}

// EnableHostEventHandler is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_EVENT_HANDLER;<host_name>
// Description:
//  Enables the event handler for the specified host.
func (c *Command) EnableHostEventHandler(host_name string) {
	c.Raw("ENABLE_HOST_EVENT_HANDLER")
	c.Hostname(host_name)
}

// EnableHostFlapDetection is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_FLAP_DETECTION;<host_name>
// Description:
//  Enables flap detection for the specified host.  In order for the flap detection algorithms to be run for the host, flap detection must be enabled on a program-wide basis as well.
func (c *Command) EnableHostFlapDetection(host_name string) {
	c.Raw("ENABLE_HOST_FLAP_DETECTION")
	c.Hostname(host_name)
}

// EnableHostFreshnessChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_FRESHNESS_CHECKS
// Description:
//  Enables freshness checks of all hosts on a program-wide basis.  Individual hosts that have freshness checks disabled will not be checked for freshness.
func (c *Command) EnableHostFreshnessChecks() {
	c.Raw("ENABLE_HOST_FRESHNESS_CHECKS")
}

// EnableHostNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_NOTIFICATIONS;<host_name>
// Description:
//  Enables notifications for a particular host.  Notifications will be sent out for the host only if notifications are enabled on a program-wide basis as well.
func (c *Command) EnableHostNotifications(host_name string) {
	c.Raw("ENABLE_HOST_NOTIFICATIONS")
	c.Hostname(host_name)
}

// EnableHostSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_SVC_CHECKS;<host_name>
// Description:
//  Enables active checks of all services on the specified host.
func (c *Command) EnableHostSvcChecks(host_name string) {
	c.Raw("ENABLE_HOST_SVC_CHECKS")
	c.Hostname(host_name)
}

// EnableHostSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_HOST_SVC_NOTIFICATIONS;<host_name>
// Description:
//  Enables notifications for all services on the specified host.  Note that notifications will not be sent out if notifications are disabled on a program-wide basis.
func (c *Command) EnableHostSvcNotifications(host_name string) {
	c.Raw("ENABLE_HOST_SVC_NOTIFICATIONS")
	c.Hostname(host_name)
}

// EnableNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_NOTIFICATIONS
// Description:
//  Enables host and service notifications on a program-wide basis.
func (c *Command) EnableNotifications() {
	c.Raw("ENABLE_NOTIFICATIONS")
}

// EnablePassiveHostChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_PASSIVE_HOST_CHECKS;<host_name>
// Description:
//  Enables acceptance and processing of passive host checks for the specified host.
func (c *Command) EnablePassiveHostChecks(host_name string) {
	c.Raw("ENABLE_PASSIVE_HOST_CHECKS")
	c.Hostname(host_name)
}

// EnablePassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_PASSIVE_SVC_CHECKS;<host_name>;<service_description>
// Description:
//  Enables passive checks for the specified service.
func (c *Command) EnablePassiveSvcChecks(host_name string, service_description string) {
	c.Raw("ENABLE_PASSIVE_SVC_CHECKS")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// EnablePerformanceData is generated from the nagios external command definition:
// Desfinition: ENABLE_PERFORMANCE_DATA
// Description:
//  Enables the processing of host and service performance data on a program-wide basis.
func (c *Command) EnablePerformanceData() {
	c.Raw("ENABLE_PERFORMANCE_DATA")
}

// EnableServicegroupHostChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_HOST_CHECKS;<servicegroup_name>
// Description:
//  Enables active checks for all hosts that have services that are members of a particular hostgroup.
func (c *Command) EnableServicegroupHostChecks(servicegroup_name string) {
	c.Raw("ENABLE_SERVICEGROUP_HOST_CHECKS")
	c.ServiceGroupName(servicegroup_name)
}

// EnableServicegroupHostNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_HOST_NOTIFICATIONS;<servicegroup_name>
// Description:
//  Enables notifications for all hosts that have services that are members of a particular servicegroup.  In order for notifications to be sent out for these hosts, notifications must also be enabled on a program-wide basis.
func (c *Command) EnableServicegroupHostNotifications(servicegroup_name string) {
	c.Raw("ENABLE_SERVICEGROUP_HOST_NOTIFICATIONS")
	c.ServiceGroupName(servicegroup_name)
}

// EnableServicegroupPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_PASSIVE_HOST_CHECKS;<servicegroup_name>
// Description:
//  Enables the acceptance and processing of passive checks for all hosts that have services that are members of a particular service group.
func (c *Command) EnableServicegroupPassiveHostChecks(servicegroup_name string) {
	c.Raw("ENABLE_SERVICEGROUP_PASSIVE_HOST_CHECKS")
	c.ServiceGroupName(servicegroup_name)
}

// EnableServicegroupPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_PASSIVE_SVC_CHECKS;<servicegroup_name>
// Description:
//  Enables the acceptance and processing of passive checks for all services in a particular servicegroup.
func (c *Command) EnableServicegroupPassiveSvcChecks(servicegroup_name string) {
	c.Raw("ENABLE_SERVICEGROUP_PASSIVE_SVC_CHECKS")
	c.ServiceGroupName(servicegroup_name)
}

// EnableServicegroupSvcChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_SVC_CHECKS;<servicegroup_name>
// Description:
//  Enables active checks for all services in a particular servicegroup.
func (c *Command) EnableServicegroupSvcChecks(servicegroup_name string) {
	c.Raw("ENABLE_SERVICEGROUP_SVC_CHECKS")
	c.ServiceGroupName(servicegroup_name)
}

// EnableServicegroupSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICEGROUP_SVC_NOTIFICATIONS;<servicegroup_name>
// Description:
//  Enables notifications for all services that are members of a particular servicegroup.  In order for notifications to be sent out for these services, notifications must also be enabled on a program-wide basis.
func (c *Command) EnableServicegroupSvcNotifications(servicegroup_name string) {
	c.Raw("ENABLE_SERVICEGROUP_SVC_NOTIFICATIONS")
	c.ServiceGroupName(servicegroup_name)
}

// EnableServiceFreshnessChecks is generated from the nagios external command definition:
// Desfinition: ENABLE_SERVICE_FRESHNESS_CHECKS
// Description:
//  Enables freshness checks of all services on a program-wide basis.  Individual services that have freshness checks disabled will not be checked for freshness.
func (c *Command) EnableServiceFreshnessChecks() {
	c.Raw("ENABLE_SERVICE_FRESHNESS_CHECKS")
}

// EnableSvcCheck is generated from the nagios external command definition:
// Desfinition: ENABLE_SVC_CHECK;<host_name>;<service_description>
// Description:
//  Enables active checks for a particular service.
func (c *Command) EnableSvcCheck(host_name string, service_description string) {
	c.Raw("ENABLE_SVC_CHECK")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// EnableSvcEventHandler is generated from the nagios external command definition:
// Desfinition: ENABLE_SVC_EVENT_HANDLER;<host_name>;<service_description>
// Description:
//  Enables the event handler for the specified service.
func (c *Command) EnableSvcEventHandler(host_name string, service_description string) {
	c.Raw("ENABLE_SVC_EVENT_HANDLER")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// EnableSvcFlapDetection is generated from the nagios external command definition:
// Desfinition: ENABLE_SVC_FLAP_DETECTION;<host_name>;<service_description>
// Description:
//  Enables flap detection for the specified service.  In order for the flap detection algorithms to be run for the service, flap detection must be enabled on a program-wide basis as well.
func (c *Command) EnableSvcFlapDetection(host_name string, service_description string) {
	c.Raw("ENABLE_SVC_FLAP_DETECTION")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// EnableSvcNotifications is generated from the nagios external command definition:
// Desfinition: ENABLE_SVC_NOTIFICATIONS;<host_name>;<service_description>
// Description:
//  Enables notifications for a particular service.  Notifications will be sent out for the service only if notifications are enabled on a program-wide basis as well.
func (c *Command) EnableSvcNotifications(host_name string, service_description string) {
	c.Raw("ENABLE_SVC_NOTIFICATIONS")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// ProcessFile is generated from the nagios external command definition:
// Desfinition: PROCESS_FILE;<file_name>;<delete>
// Description:
//  Directs Nagios to process all external commands that are found in the file specified by the <file_name> argument.  If the <delete> option is non-zero, the file will be deleted once it has been processes.  If the <delete> option is set to zero, the file is left untouched.
func (c *Command) ProcessFile(file_name string, delete bool) {
	c.Raw("PROCESS_FILE")
	c.FileName(file_name)
	c.Delete(delete)
}

// ProcessHostCheckResult is generated from the nagios external command definition:
// Desfinition: PROCESS_HOST_CHECK_RESULT;<host_name>;<status_code>;<plugin_output>
// Description:
//  This is used to submit a passive check result for a particular host.  The "status_code" indicates the state of the host check and should be one of the following: 0=UP, 1=DOWN, 2=UNREACHABLE.  The "plugin_output" argument contains the text returned from the host check, along with optional performance data.
func (c *Command) ProcessHostCheckResult(host_name string, status_code int, plugin_output string) {
	c.Raw("PROCESS_HOST_CHECK_RESULT")
	c.Hostname(host_name)
	c.StatusCode(status_code)
	c.PluginOutput(plugin_output)
}

// ProcessServiceCheckResult is generated from the nagios external command definition:
// Desfinition: PROCESS_SERVICE_CHECK_RESULT;<host_name>;<service_description>;<return_code>;<plugin_output>
// Description:
//  This is used to submit a passive check result for a particular service.  The "return_code" field should be one of the following: 0=OK, 1=WARNING, 2=CRITICAL, 3=UNKNOWN.  The "plugin_output" field contains text output from the service check, along with optional performance data.
func (c *Command) ProcessServiceCheckResult(host_name string, service_description string, return_code int, plugin_output string) {
	c.Raw("PROCESS_SERVICE_CHECK_RESULT")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.ReturnCode(return_code)
	c.PluginOutput(plugin_output)
}

// ReadStateInformation is generated from the nagios external command definition:
// Desfinition: READ_STATE_INFORMATION
// Description:
//  Causes Nagios to load all current monitoring status information from the state retention file.  Normally, state retention information is loaded when the Nagios process starts up and before it starts monitoring.  WARNING: This command will cause Nagios to discard all current monitoring status information and use the information stored in state retention file!  Use with care.
func (c *Command) ReadStateInformation() {
	c.Raw("READ_STATE_INFORMATION")
}

// RemoveHostAcknowledgement is generated from the nagios external command definition:
// Desfinition: REMOVE_HOST_ACKNOWLEDGEMENT;<host_name>
// Description:
//  This removes the problem acknowledgement for a particular host.  Once the acknowledgement has been removed, notifications can once again be sent out for the given host.
func (c *Command) RemoveHostAcknowledgement(host_name string) {
	c.Raw("REMOVE_HOST_ACKNOWLEDGEMENT")
	c.Hostname(host_name)
}

// RemoveSvcAcknowledgement is generated from the nagios external command definition:
// Desfinition: REMOVE_SVC_ACKNOWLEDGEMENT;<host_name>;<service_description>
// Description:
//  This removes the problem acknowledgement for a particular service.  Once the acknowledgement has been removed, notifications can once again be sent out for the given service.
func (c *Command) RemoveSvcAcknowledgement(host_name string, service_description string) {
	c.Raw("REMOVE_SVC_ACKNOWLEDGEMENT")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// RestartProgram is generated from the nagios external command definition:
// Desfinition: RESTART_PROGRAM
// Description:
//  Restarts the Nagios process.
func (c *Command) RestartProgram() {
	c.Raw("RESTART_PROGRAM")
}

// SaveStateInformation is generated from the nagios external command definition:
// Desfinition: SAVE_STATE_INFORMATION
// Description:
//  Causes Nagios to save all current monitoring status information to the state retention file.  Normally, state retention information is saved before the Nagios process shuts down and (potentially) at regularly scheduled intervals.  This command allows you to force Nagios to save this information to the state retention file immediately.  This does not affect the current status information in the Nagios process.
func (c *Command) SaveStateInformation() {
	c.Raw("SAVE_STATE_INFORMATION")
}

// ScheduleAndPropagateHostDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_AND_PROPAGATE_HOST_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for a specified host and all of its children (hosts).  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The specified (parent) host downtime can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the specified (parent) host should not be triggered by another downtime entry.
func (c *Command) ScheduleAndPropagateHostDowntime(host_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) {
	c.Raw("SCHEDULE_AND_PROPAGATE_HOST_DOWNTIME")
	c.Hostname(host_name)
	c.Start(start_time)
	c.End(end_time)
	c.Fixed(fixed)
	c.TriggerID(trigger_id)
	c.Duration(duration)
	c.Author(author)
	c.Comment(comment)
}

// ScheduleAndPropagateTriggeredHostDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_AND_PROPAGATE_TRIGGERED_HOST_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for a specified host and all of its children (hosts).  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  Downtime for child hosts are all set to be triggered by the downtime for the specified (parent) host.  The specified (parent) host downtime can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the specified (parent) host should not be triggered by another downtime entry.
func (c *Command) ScheduleAndPropagateTriggeredHostDowntime(host_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) {
	c.Raw("SCHEDULE_AND_PROPAGATE_TRIGGERED_HOST_DOWNTIME")
	c.Hostname(host_name)
	c.Start(start_time)
	c.End(end_time)
	c.Fixed(fixed)
	c.TriggerID(trigger_id)
	c.Duration(duration)
	c.Author(author)
	c.Comment(comment)
}

// ScheduleForcedHostCheck is generated from the nagios external command definition:
// Desfinition: SCHEDULE_FORCED_HOST_CHECK;<host_name>;<check_time>
// Description:
//  Schedules a forced active check of a particular host at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).   Forced checks are performed regardless of what time it is (e.g. timeperiod restrictions are ignored) and whether or not active checks are enabled on a host-specific or program-wide basis.
func (c *Command) ScheduleForcedHostCheck(host_name string, check_time time.Time) {
	c.Raw("SCHEDULE_FORCED_HOST_CHECK")
	c.Hostname(host_name)
	c.CheckTime(check_time)
}

// ScheduleForcedHostSvcChecks is generated from the nagios external command definition:
// Desfinition: SCHEDULE_FORCED_HOST_SVC_CHECKS;<host_name>;<check_time>
// Description:
//  Schedules a forced active check of all services associated with a particular host at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).   Forced checks are performed regardless of what time it is (e.g. timeperiod restrictions are ignored) and whether or not active checks are enabled on a service-specific or program-wide basis.
func (c *Command) ScheduleForcedHostSvcChecks(host_name string, check_time time.Time) {
	c.Raw("SCHEDULE_FORCED_HOST_SVC_CHECKS")
	c.Hostname(host_name)
	c.CheckTime(check_time)
}

// ScheduleForcedSvcCheck is generated from the nagios external command definition:
// Desfinition: SCHEDULE_FORCED_SVC_CHECK;<host_name>;<service_description>;<check_time>
// Description:
//  Schedules a forced active check of a particular service at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).   Forced checks are performed regardless of what time it is (e.g. timeperiod restrictions are ignored) and whether or not active checks are enabled on a service-specific or program-wide basis.
func (c *Command) ScheduleForcedSvcCheck(host_name string, service_description string, check_time time.Time) {
	c.Raw("SCHEDULE_FORCED_SVC_CHECK")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.CheckTime(check_time)
}

// ScheduleHostgroupHostDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOSTGROUP_HOST_DOWNTIME;<hostgroup_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for all hosts in a specified hostgroup.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The host downtime entries can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the hosts should not be triggered by another downtime entry.
func (c *Command) ScheduleHostgroupHostDowntime(hostgroup_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) {
	c.Raw("SCHEDULE_HOSTGROUP_HOST_DOWNTIME")
	c.HostGroupName(hostgroup_name)
	c.Start(start_time)
	c.End(end_time)
	c.Fixed(fixed)
	c.TriggerID(trigger_id)
	c.Duration(duration)
	c.Author(author)
	c.Comment(comment)
}

// ScheduleHostgroupSvcDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOSTGROUP_SVC_DOWNTIME;<hostgroup_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for all services associated with hosts in a specified servicegroup.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The service downtime entries can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the services should not be triggered by another downtime entry.
func (c *Command) ScheduleHostgroupSvcDowntime(hostgroup_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) {
	c.Raw("SCHEDULE_HOSTGROUP_SVC_DOWNTIME")
	c.HostGroupName(hostgroup_name)
	c.Start(start_time)
	c.End(end_time)
	c.Fixed(fixed)
	c.TriggerID(trigger_id)
	c.Duration(duration)
	c.Author(author)
	c.Comment(comment)
}

// ScheduleHostCheck is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOST_CHECK;<host_name>;<check_time>
// Description:
//  Schedules the next active check of a particular host at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).  Note that the host may not actually be checked at the time you specify.  This could occur for a number of reasons: active checks are disabled on a program-wide or service-specific basis, the host is already scheduled to be checked at an earlier time, etc.  If you want to force the host check to occur at the time you specify, look at the SCHEDULE_FORCED_HOST_CHECK command.
func (c *Command) ScheduleHostCheck(host_name string, check_time time.Time) {
	c.Raw("SCHEDULE_HOST_CHECK")
	c.Hostname(host_name)
	c.CheckTime(check_time)
}

// ScheduleHostDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOST_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for a specified host.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The specified host downtime can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the specified host should not be triggered by another downtime entry.
func (c *Command) ScheduleHostDowntime(host_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) {
	c.Raw("SCHEDULE_HOST_DOWNTIME")
	c.Hostname(host_name)
	c.Start(start_time)
	c.End(end_time)
	c.Fixed(fixed)
	c.TriggerID(trigger_id)
	c.Duration(duration)
	c.Author(author)
	c.Comment(comment)
}

// ScheduleHostSvcChecks is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOST_SVC_CHECKS;<host_name>;<check_time>
// Description:
//  Schedules the next active check of all services on a particular host at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).  Note that the services may not actually be checked at the time you specify.  This could occur for a number of reasons: active checks are disabled on a program-wide or service-specific basis, the services are already scheduled to be checked at an earlier time, etc.  If you want to force the service checks to occur at the time you specify, look at the SCHEDULE_FORCED_HOST_SVC_CHECKS command.
func (c *Command) ScheduleHostSvcChecks(host_name string, check_time time.Time) {
	c.Raw("SCHEDULE_HOST_SVC_CHECKS")
	c.Hostname(host_name)
	c.CheckTime(check_time)
}

// ScheduleHostSvcDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_HOST_SVC_DOWNTIME;<host_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for all services associated with a particular host.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The service downtime entries can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the services should not be triggered by another downtime entry.
func (c *Command) ScheduleHostSvcDowntime(host_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) {
	c.Raw("SCHEDULE_HOST_SVC_DOWNTIME")
	c.Hostname(host_name)
	c.Start(start_time)
	c.End(end_time)
	c.Fixed(fixed)
	c.TriggerID(trigger_id)
	c.Duration(duration)
	c.Author(author)
	c.Comment(comment)
}

// ScheduleServicegroupHostDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_SERVICEGROUP_HOST_DOWNTIME;<servicegroup_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for all hosts that have services in a specified servicegroup.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The host downtime entries can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the hosts should not be triggered by another downtime entry.
func (c *Command) ScheduleServicegroupHostDowntime(servicegroup_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) {
	c.Raw("SCHEDULE_SERVICEGROUP_HOST_DOWNTIME")
	c.ServiceGroupName(servicegroup_name)
	c.Start(start_time)
	c.End(end_time)
	c.Fixed(fixed)
	c.TriggerID(trigger_id)
	c.Duration(duration)
	c.Author(author)
	c.Comment(comment)
}

// ScheduleServicegroupSvcDowntime is generated from the nagios external command definition:
// Desfinition: SCHEDULE_SERVICEGROUP_SVC_DOWNTIME;<servicegroup_name>;<start_time>;<end_time>;<fixed>;<trigger_id>;<duration>;<author>;<comment>
// Description:
//  Schedules downtime for all services in a specified servicegroup.  If the "fixed" argument is set to one (1), downtime will start and end at the times specified by the "start" and "end" arguments.  Otherwise, downtime will begin between the "start" and "end" times and last for "duration" seconds.  The "start" and "end" arguments are specified in time_t format (seconds since the UNIX epoch).  The service downtime entries can be triggered by another downtime entry if the "trigger_id" is set to the ID of another scheduled downtime entry.  Set the "trigger_id" argument to zero (0) if the downtime for the services should not be triggered by another downtime entry.
func (c *Command) ScheduleServicegroupSvcDowntime(servicegroup_name string, start_time time.Time, end_time time.Time, fixed bool, trigger_id int, duration time.Duration, author string, comment string) {
	c.Raw("SCHEDULE_SERVICEGROUP_SVC_DOWNTIME")
	c.ServiceGroupName(servicegroup_name)
	c.Start(start_time)
	c.End(end_time)
	c.Fixed(fixed)
	c.TriggerID(trigger_id)
	c.Duration(duration)
	c.Author(author)
	c.Comment(comment)
}

// ScheduleSvcCheck is generated from the nagios external command definition:
// Desfinition: SCHEDULE_SVC_CHECK;<host_name>;<service_description>;<check_time>
// Description:
//  Schedules the next active check of a specified service at "check_time".  The "check_time" argument is specified in time_t format (seconds since the UNIX epoch).  Note that the service may not actually be checked at the time you specify.  This could occur for a number of reasons: active checks are disabled on a program-wide or service-specific basis, the service is already scheduled to be checked at an earlier time, etc.  If you want to force the service check to occur at the time you specify, look at the SCHEDULE_FORCED_SVC_CHECK command.
func (c *Command) ScheduleSvcCheck(host_name string, service_description string, check_time time.Time) {
	c.Raw("SCHEDULE_SVC_CHECK")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.CheckTime(check_time)
}

// SendCustomHostNotification is generated from the nagios external command definition:
// Desfinition: SEND_CUSTOM_HOST_NOTIFICATION;<host_name>;<options>;<author>;<comment>
// Description:
//  Allows you to send a custom host notification.  Very useful in dire situations, emergencies or to communicate with all admins that are responsible for a particular host.  When the host notification is sent out, the $NOTIFICATIONTYPE$ macro will be set to "CUSTOM".  The <options> field is a logical OR of the following integer values that affect aspects of the notification that are sent out: 0 = No option (default), 1 = Broadcast (send notification to all normal and all escalated contacts for the host), 2 = Forced (notification is sent out regardless of current time, whether or not notifications are enabled, etc.), 4 = Increment current notification # for the host (this is not done by default for custom notifications).  The comment field can be used with the $NOTIFICATIONCOMMENT$ macro in notification commands.
func (c *Command) SendCustomHostNotification(host_name string, options int, author string, comment string) {
	c.Raw("SEND_CUSTOM_HOST_NOTIFICATION")
	c.Hostname(host_name)
	c.Options(options)
	c.Author(author)
	c.Comment(comment)
}

// SendCustomSvcNotification is generated from the nagios external command definition:
// Desfinition: SEND_CUSTOM_SVC_NOTIFICATION;<host_name>;<service_description>;<options>;<author>;<comment>
// Description:
//  Allows you to send a custom service notification.  Very useful in dire situations, emergencies or to communicate with all admins that are responsible for a particular service.  When the service notification is sent out, the $NOTIFICATIONTYPE$ macro will be set to "CUSTOM".  The <options> field is a logical OR of the following integer values that affect aspects of the notification that are sent out: 0 = No option (default), 1 = Broadcast (send notification to all normal and all escalated contacts for the service), 2 = Forced (notification is sent out regardless of current time, whether or not notifications are enabled, etc.), 4 = Increment current notification # for the service(this is not done by default for custom notifications).   The comment field can be used with the $NOTIFICATIONCOMMENT$ macro in notification commands.
func (c *Command) SendCustomSvcNotification(host_name string, service_description string, options int, author string, comment string) {
	c.Raw("SEND_CUSTOM_SVC_NOTIFICATION")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.Options(options)
	c.Author(author)
	c.Comment(comment)
}

// SetHostNotificationNumber is generated from the nagios external command definition:
// Desfinition: SET_HOST_NOTIFICATION_NUMBER;<host_name>;<notification_number>
// Description:
//  Sets the current notification number for a particular host.  A value of 0 indicates that no notification has yet been sent for the current host problem.  Useful for forcing an escalation (based on notification number) or replicating notification information in redundant monitoring environments. Notification numbers greater than zero have no noticeable affect on the notification process if the host is currently in an UP state.
func (c *Command) SetHostNotificationNumber(host_name string, notification_number int) {
	c.Raw("SET_HOST_NOTIFICATION_NUMBER")
	c.Hostname(host_name)
	c.NotificationNumber(notification_number)
}

// SetSvcNotificationNumber is generated from the nagios external command definition:
// Desfinition: SET_SVC_NOTIFICATION_NUMBER;<host_name>;<service_description>;<notification_number>
// Description:
//  Sets the current notification number for a particular service.  A value of 0 indicates that no notification has yet been sent for the current service problem.  Useful for forcing an escalation (based on notification number) or replicating notification information in redundant monitoring environments. Notification numbers greater than zero have no noticeable affect on the notification process if the service is currently in an OK state.
func (c *Command) SetSvcNotificationNumber(host_name string, service_description string, notification_number int) {
	c.Raw("SET_SVC_NOTIFICATION_NUMBER")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
	c.NotificationNumber(notification_number)
}

// ShutdownProgram is generated from the nagios external command definition:
// Desfinition: SHUTDOWN_PROGRAM
// Description:
//  Shuts down the Nagios process.
func (c *Command) ShutdownProgram() {
	c.Raw("SHUTDOWN_PROGRAM")
}

// StartAcceptingPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: START_ACCEPTING_PASSIVE_HOST_CHECKS
// Description:
//  Enables acceptance and processing of passive host checks on a program-wide basis.
func (c *Command) StartAcceptingPassiveHostChecks() {
	c.Raw("START_ACCEPTING_PASSIVE_HOST_CHECKS")
}

// StartAcceptingPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: START_ACCEPTING_PASSIVE_SVC_CHECKS
// Description:
//  Enables passive service checks on a program-wide basis.
func (c *Command) StartAcceptingPassiveSvcChecks() {
	c.Raw("START_ACCEPTING_PASSIVE_SVC_CHECKS")
}

// StartExecutingHostChecks is generated from the nagios external command definition:
// Desfinition: START_EXECUTING_HOST_CHECKS
// Description:
//  Enables active host checks on a program-wide basis.
func (c *Command) StartExecutingHostChecks() {
	c.Raw("START_EXECUTING_HOST_CHECKS")
}

// StartExecutingSvcChecks is generated from the nagios external command definition:
// Desfinition: START_EXECUTING_SVC_CHECKS
// Description:
//  Enables active checks of services on a program-wide basis.
func (c *Command) StartExecutingSvcChecks() {
	c.Raw("START_EXECUTING_SVC_CHECKS")
}

// StartObsessingOverHost is generated from the nagios external command definition:
// Desfinition: START_OBSESSING_OVER_HOST;<host_name>
// Description:
//  Enables processing of host checks via the OCHP command for the specified host.
func (c *Command) StartObsessingOverHost(host_name string) {
	c.Raw("START_OBSESSING_OVER_HOST")
	c.Hostname(host_name)
}

// StartObsessingOverHostChecks is generated from the nagios external command definition:
// Desfinition: START_OBSESSING_OVER_HOST_CHECKS
// Description:
//  Enables processing of host checks via the OCHP command on a program-wide basis.
func (c *Command) StartObsessingOverHostChecks() {
	c.Raw("START_OBSESSING_OVER_HOST_CHECKS")
}

// StartObsessingOverSvc is generated from the nagios external command definition:
// Desfinition: START_OBSESSING_OVER_SVC;<host_name>;<service_description>
// Description:
//  Enables processing of service checks via the OCSP command for the specified service.
func (c *Command) StartObsessingOverSvc(host_name string, service_description string) {
	c.Raw("START_OBSESSING_OVER_SVC")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// StartObsessingOverSvcChecks is generated from the nagios external command definition:
// Desfinition: START_OBSESSING_OVER_SVC_CHECKS
// Description:
//  Enables processing of service checks via the OCSP command on a program-wide basis.
func (c *Command) StartObsessingOverSvcChecks() {
	c.Raw("START_OBSESSING_OVER_SVC_CHECKS")
}

// StopAcceptingPassiveHostChecks is generated from the nagios external command definition:
// Desfinition: STOP_ACCEPTING_PASSIVE_HOST_CHECKS
// Description:
//  Disables acceptance and processing of passive host checks on a program-wide basis.
func (c *Command) StopAcceptingPassiveHostChecks() {
	c.Raw("STOP_ACCEPTING_PASSIVE_HOST_CHECKS")
}

// StopAcceptingPassiveSvcChecks is generated from the nagios external command definition:
// Desfinition: STOP_ACCEPTING_PASSIVE_SVC_CHECKS
// Description:
//  Disables passive service checks on a program-wide basis.
func (c *Command) StopAcceptingPassiveSvcChecks() {
	c.Raw("STOP_ACCEPTING_PASSIVE_SVC_CHECKS")
}

// StopExecutingHostChecks is generated from the nagios external command definition:
// Desfinition: STOP_EXECUTING_HOST_CHECKS
// Description:
//  Disables active host checks on a program-wide basis.
func (c *Command) StopExecutingHostChecks() {
	c.Raw("STOP_EXECUTING_HOST_CHECKS")
}

// StopExecutingSvcChecks is generated from the nagios external command definition:
// Desfinition: STOP_EXECUTING_SVC_CHECKS
// Description:
//  Disables active checks of services on a program-wide basis.
func (c *Command) StopExecutingSvcChecks() {
	c.Raw("STOP_EXECUTING_SVC_CHECKS")
}

// StopObsessingOverHost is generated from the nagios external command definition:
// Desfinition: STOP_OBSESSING_OVER_HOST;<host_name>
// Description:
//  Disables processing of host checks via the OCHP command for the specified host.
func (c *Command) StopObsessingOverHost(host_name string) {
	c.Raw("STOP_OBSESSING_OVER_HOST")
	c.Hostname(host_name)
}

// StopObsessingOverHostChecks is generated from the nagios external command definition:
// Desfinition: STOP_OBSESSING_OVER_HOST_CHECKS
// Description:
//  Disables processing of host checks via the OCHP command on a program-wide basis.
func (c *Command) StopObsessingOverHostChecks() {
	c.Raw("STOP_OBSESSING_OVER_HOST_CHECKS")
}

// StopObsessingOverSvc is generated from the nagios external command definition:
// Desfinition: STOP_OBSESSING_OVER_SVC;<host_name>;<service_description>
// Description:
//  Disables processing of service checks via the OCSP command for the specified service.
func (c *Command) StopObsessingOverSvc(host_name string, service_description string) {
	c.Raw("STOP_OBSESSING_OVER_SVC")
	c.Hostname(host_name)
	c.ServiceDescription(service_description)
}

// StopObsessingOverSvcChecks is generated from the nagios external command definition:
// Desfinition: STOP_OBSESSING_OVER_SVC_CHECKS
// Description:
//  Disables processing of service checks via the OCSP command on a program-wide basis.
func (c *Command) StopObsessingOverSvcChecks() {
	c.Raw("STOP_OBSESSING_OVER_SVC_CHECKS")
}
