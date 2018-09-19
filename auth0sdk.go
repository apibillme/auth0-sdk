package auth0sdk

import (
	"errors"
	"strings"

	"github.com/spf13/cast"

	"github.com/apibillme/restly"
	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
)

// for stubbing
var restlyPostJSON = restly.PostJSON
var restlyPatchJSON = restly.PatchJSON
var restlyDeleteJSON = restly.DeleteJSON
var restlyGetJSON = restly.GetJSON
var restlyPutJSON = restly.PutJSON

var req *fasthttp.Request
var clientIDGlobal string
var clientSecretGlobal string
var auth0DomainGlobal string

// New - create new Auth0 client
func New(auth0Domain string, clientID string, clientSecret string) error {
	// assign vars
	clientIDGlobal = clientID
	clientSecretGlobal = clientSecret
	auth0DomainGlobal = `https://` + auth0Domain

	req = restly.New()
	audience := auth0DomainGlobal + `/api/v2/`
	body := `{"grant_type":"client_credentials", "client_id":"` + clientIDGlobal + `", "client_secret":"` + clientSecretGlobal + `", "audience":"` + audience + `"}`
	url := auth0DomainGlobal + `/oauth/token`
	res, statusCode, err := restlyPostJSON(req, url, body, "")
	if err != nil {
		return err
	}
	if statusCode == 200 {
		accessToken := res.Get("access_token").String()
		req.Header.Add("Authorization", "Bearer "+accessToken)
	} else {
		return errors.New("could not get token - Auth0 returned a " + cast.ToString(statusCode))
	}
	return nil
}

// Signup - sign up user
func Signup(email string, password string, connection string, userMetadata string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + `/dbconnections/signup`
	body := `{"client_id":"` + clientIDGlobal + `", "email":"` + email + `", "password":"` + password + `", "connection":"` + connection + `", "user_metadata":"` + userMetadata + `"}`
	return restlyPostJSON(req, url, body, "")
}

// ChangePassword - change password for user
func ChangePassword(email string, password string, connection string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + `/dbconnections/change_password`
	body := `{"client_id":"` + clientIDGlobal + `", "email":"` + email + `", "password":"` + password + `", "connection":"` + connection + `"}`
	return restlyPostJSON(req, url, body, "")
}

// UserInfo - get user info
func UserInfo(accessToken string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + `/userinfo`
	reqDiff := restly.New()
	reqDiff.Header.Add("Authorization", "Bearer "+accessToken)
	return restlyGetJSON(reqDiff, url, "")
}

// ChallengeMFA - challenge MFA token
func ChallengeMFA(mfaToken string, challengeType string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + `/mfa/challenge`
	body := `{"client_id":"` + clientIDGlobal + `", "client_secret":"` + clientSecretGlobal + `", "mfa_token":"` + mfaToken + `", "challenge_type":"` + challengeType + `"}`
	return restlyPostJSON(req, url, body, "")
}

// VerifyOTP - verify OTP
func VerifyOTP(mfaToken string, otpCode string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + `/oauth/token`
	body := `{"client_id":"` + clientIDGlobal + `", "client_secret":"` + clientSecretGlobal + `", "mfa_token":"` + mfaToken + `", "grant_type":"http://auth0.com/oauth/grant-type/mfa-otp", "otp":"` + otpCode + `"}`
	return restlyPostJSON(req, url, body, "")
}

// VerifyOOB - verify OOB
func VerifyOOB(mfaToken string, oobCode string, bindingCode string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + `/oauth/token`
	body := `{"client_id":"` + clientIDGlobal + `", "client_secret":"` + clientSecretGlobal + `", "mfa_token":"` + mfaToken + `", "grant_type":"http://auth0.com/oauth/grant-type/mfa-oob", "oob_code":"` + oobCode + `", "binding_code":"` + bindingCode + `"}`
	return restlyPostJSON(req, url, body, "")
}

// VerifyRecoveryCode - verify recovery code
func VerifyRecoveryCode(mfaToken string, recoveryCode string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + `/oauth/token`
	body := `{"client_id":"` + clientIDGlobal + `", "client_secret":"` + clientSecretGlobal + `", "mfa_token":"` + mfaToken + `", "grant_type":"http://auth0.com/oauth/grant-type/mfa-recovery-code", "recovery_code":"` + recoveryCode + `"}`
	return restlyPostJSON(req, url, body, "")
}

// AddAuthenticator - adds an authenticator
func AddAuthenticator(authenticatorTypes []string, oobChannels string, phoneNumber string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + `/mfa/associate`
	authString := strings.Join(authenticatorTypes, ",")
	body := `{"client_id":"` + clientIDGlobal + `", "client_secret":"` + clientSecretGlobal + `", "authenticator_types":["` + authString + `]", "oob_channels":"` + oobChannels + `", "phone_number":"` + phoneNumber + `"}`
	return restlyPostJSON(req, url, body, "")
}

// Passwordless - passwordless
func Passwordless(connection string, email string, phoneNumber string, send string, authParams string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + `/passwordless/start`
	body := `{"client_id":"` + clientIDGlobal + `", "client_secret":"` + clientSecretGlobal + `", "email":"` + email + `", "phone_number":"` + phoneNumber + `", "send":"` + send + `", "authParams":"` + authParams + `"}`
	return restlyPostJSON(req, url, body, "")
}

// mgmt API version
var mgmtAPIVersion = `/api/v2/`

// mgmt api endpoints
var clientGrants = `client-grants`
var clients = `clients`
var connections = `connections`
var customDomains = `custom-domains`
var deviceCredentials = `device-credentials`
var grants = `grants`
var logs = `logs`
var resourceServers = `resource-servers`
var rules = `rules`
var rulesConfigs = `rules-configs`
var userBlocks = `user-blocks`
var users = `users`
var blackLists = `blacklists/tokens`
var emailTemplates = `email-templates`
var emailProvider = `emails/provider`
var guardianFactors = `guardian/factors`
var guardianEnrollments = `guardian/enrollments`
var jobs = `jobs`

// GetClientGrants - get all client grants
func GetClientGrants(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + clientGrants
	return restlyGetJSON(req, url, query)
}

// CreateClientGrant - create a client grant
func CreateClientGrant(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + clientGrants
	return restlyPostJSON(req, url, body, "")
}

// DeleteClientGrant - delete a client grant
func DeleteClientGrant(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + clientGrants + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// UpdateClientGrant - update client grant
func UpdateClientGrant(id string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + clientGrants + `/` + id
	return restlyPatchJSON(req, url, body, "")
}

// GetClients - get all clients
func GetClients(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + clients
	return restlyGetJSON(req, url, query)
}

// CreateClient - create a client
func CreateClient(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + clients
	return restlyPostJSON(req, url, body, "")
}

// GetClient - get a client
func GetClient(id string, query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + clients + `/` + id
	return restlyGetJSON(req, url, query)
}

// DeleteClient - delete a client
func DeleteClient(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + clients + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// UpdateClient - update a client
func UpdateClient(id string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + clients + `/` + id
	return restlyPatchJSON(req, url, body, "")
}

// RotateClientSecret - rotate the client secret
func RotateClientSecret(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + clients + `/` + id + `/rotate-secret`
	return restlyPostJSON(req, url, "", "")
}

// GetConnections - get all connections
func GetConnections(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + connections
	return restlyGetJSON(req, url, query)
}

// CreateConnection - create connection
func CreateConnection(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + connections
	return restlyPostJSON(req, url, body, "")
}

// GetConnection - get a connection
func GetConnection(id string, query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + connections + `/` + id
	return restlyGetJSON(req, url, query)
}

// DeleteConnection - delete a connection
func DeleteConnection(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + connections + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// UpdateConnection - update a connection
func UpdateConnection(id string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + connections + `/` + id
	return restlyPatchJSON(req, url, body, "")
}

// DeleteConnectionUser - delete a connection user
func DeleteConnectionUser(id string, query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + connections + `/` + id
	return restlyDeleteJSON(req, url, query)
}

// GetCustomDomains - get all custom domains
func GetCustomDomains() (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + customDomains
	return restlyGetJSON(req, url, "")
}

// CreateCustomDomain - create a custom domain
func CreateCustomDomain(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + customDomains
	return restlyPostJSON(req, url, body, "")
}

// GetCustomDomain - get a custom domain
func GetCustomDomain(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + customDomains + `/` + id
	return restlyGetJSON(req, url, "")
}

// DeleteCustomDomain - delete a custom domain
func DeleteCustomDomain(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + customDomains + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// VerifyCustomDomain - verify a custom domain
func VerifyCustomDomain(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + customDomains + `/` + id + `/verify`
	return restlyPostJSON(req, url, "", "")
}

// GetDeviceCredentials - get all device credentials
func GetDeviceCredentials(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + deviceCredentials
	return restlyGetJSON(req, url, query)
}

// CreateDeviceCredential - create a device credential
func CreateDeviceCredential(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + deviceCredentials
	return restlyPostJSON(req, url, body, "")
}

// DeleteDeviceCredential - delete a device credential
func DeleteDeviceCredential(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + deviceCredentials + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// GetGrants - get all grants
func GetGrants(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + grants
	return restlyGetJSON(req, url, query)
}

// DeleteGrant - delete a grant
func DeleteGrant(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + grants + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// SearchLog - search log
func SearchLog(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + logs
	return restlyGetJSON(req, url, query)
}

// GetLog - get log
func GetLog(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + logs + `/` + id
	return restlyGetJSON(req, url, "")
}

// GetResourceServers - get all resource servers
func GetResourceServers(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + resourceServers
	return restlyGetJSON(req, url, query)
}

// CreateResourceServer - create resource server
func CreateResourceServer(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + resourceServers
	return restlyPostJSON(req, url, body, "")
}

// GetResourceServer - get a resource server
func GetResourceServer(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + resourceServers + `/` + id
	return restlyGetJSON(req, url, "")
}

// DeleteResourceServer - delete a resource server
func DeleteResourceServer(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + resourceServers + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// UpdateResourceServer - update a resource server
func UpdateResourceServer(id string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + resourceServers + `/` + id
	return restlyPatchJSON(req, url, body, "")
}

// GetRules - get all rules
func GetRules(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + rules
	return restlyGetJSON(req, url, query)
}

// CreateRule - create a rule
func CreateRule(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + rules
	return restlyPostJSON(req, url, body, "")
}

// GetRule - get a rule
func GetRule(id string, query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + rules + `/` + id
	return restlyGetJSON(req, url, query)
}

// DeleteRule - delete a rule
func DeleteRule(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + rules + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// UpdateRule - update a rule
func UpdateRule(id string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + rules + `/` + id
	return restlyPatchJSON(req, url, body, "")
}

// GetRulesConfigs - get all rules configs
func GetRulesConfigs() (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + rulesConfigs
	return restlyGetJSON(req, url, "")
}

// DeleteRulesConfig - delete a rules config
func DeleteRulesConfig(key string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + rulesConfigs + `/` + key
	return restlyDeleteJSON(req, url, "")
}

// SetRulesConfig - set a rules config
func SetRulesConfig(key string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + rulesConfigs + `/` + key
	return restlyPutJSON(req, url, body, "")
}

// GetUserBlocksByIdentifier - get user blocks by identifier
func GetUserBlocksByIdentifier(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + userBlocks
	return restlyGetJSON(req, url, query)
}

// DeleteUserBlockByIdentifier - delete a user block by identifier
func DeleteUserBlockByIdentifier(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + userBlocks
	return restlyDeleteJSON(req, url, query)
}

// GetUserBlock - get a user block
func GetUserBlock(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + userBlocks + `/` + id
	return restlyGetJSON(req, url, "")
}

// DeleteUserBlock - delete a user block
func DeleteUserBlock(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + userBlocks + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// SearchUsers - search or list users
func SearchUsers(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users
	return restlyGetJSON(req, url, query)
}

// CreateUser - create a user
func CreateUser(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users
	return restlyPostJSON(req, url, body, "")
}

// GetUser - get a user
func GetUser(id string, query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users + `/` + id
	return restlyGetJSON(req, url, query)
}

// DeleteUser - delete a user
func DeleteUser(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// UpdateUser - update a user
func UpdateUser(id string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users + `/` + id
	return restlyPatchJSON(req, url, body, "")
}

// GetUserGuardianEnrollments - get a user's guardian enrollments
func GetUserGuardianEnrollments(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users + `/` + id + `/enrollments`
	return restlyGetJSON(req, url, "")
}

// GetUserLog - get a user's log
func GetUserLog(id string, query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users + `/` + id + `/logs`
	return restlyGetJSON(req, url, query)
}

// DeleteUserMultifactorProvider - delete a user's multifactor provider
func DeleteUserMultifactorProvider(id string, provider string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users + `/` + id + `/multifactor/` + provider
	return restlyDeleteJSON(req, url, "")
}

// UnlinkUserIdentity - unlink a user's identity
func UnlinkUserIdentity(id string, provider string, userID string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users + `/` + id + `/identities/` + provider + `/` + userID
	return restlyDeleteJSON(req, url, "")
}

// GenerateGuardianRecoveryCode - generate a new guardian recovery code for a user
func GenerateGuardianRecoveryCode(userID string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users + `/` + userID + `/recovery-code-regeneration`
	return restlyPostJSON(req, url, "", "")
}

// LinkUser - link a user
func LinkUser(id string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + users + `/` + id + `/identities`
	return restlyPostJSON(req, url, body, "")
}

// SearchUserByEmail - search for a user by email
func SearchUserByEmail(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + `users-by-email`
	return restlyGetJSON(req, url, query)
}

// GetBlacklist - get the blacklist
func GetBlacklist(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + blackLists
	return restlyGetJSON(req, url, query)
}

// UpdateBlacklist - add a token to the blacklist
func UpdateBlacklist(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + blackLists
	return restlyPostJSON(req, url, body, "")
}

// GetEmailTemplate - get an email template
func GetEmailTemplate(emailTemplate string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + emailTemplates
	return restlyGetJSON(req, url, "")
}

// PatchEmailTemplate - patch an email template
func PatchEmailTemplate(emailTemplate string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + emailTemplates + `/` + emailTemplate
	return restlyPatchJSON(req, url, body, "")
}

// UpdateEmailTemplate - update an email template
func UpdateEmailTemplate(emailTemplate string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + emailTemplates + `/` + emailTemplate
	return restlyPutJSON(req, url, body, "")
}

// CreateEmailTemplate - create an email template
func CreateEmailTemplate(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + emailTemplates
	return restlyPostJSON(req, url, body, "")
}

// GetEmailProvider - get the email provider
func GetEmailProvider(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + emailProvider
	return restlyGetJSON(req, url, query)
}

// DeleteEmailProvider - delete the email provider - use with caution
func DeleteEmailProvider() (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + emailProvider
	return restlyDeleteJSON(req, url, "")
}

// UpdateEmailProvider - update the email provider
func UpdateEmailProvider(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + emailProvider
	return restlyPatchJSON(req, url, body, "")
}

// ConfigureEmailProvider - configure the email provider
func ConfigureEmailProvider(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + emailProvider
	return restlyPostJSON(req, url, body, "")
}

// GetGuardianFactors - get a list of factors and statuses for guardian
func GetGuardianFactors() (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + guardianFactors
	return restlyGetJSON(req, url, "")
}

// GetGuardianEnrollment - get a guardian enrollment
func GetGuardianEnrollment(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + guardianEnrollments + `/` + id
	return restlyGetJSON(req, url, "")
}

// DeleteGuardianEnrollment - delete a guardian enrollment
func DeleteGuardianEnrollment(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + guardianEnrollments + `/` + id
	return restlyDeleteJSON(req, url, "")
}

// GetGuardianEnrollmentTemplates - get all guardian enrollment and verification templates
func GetGuardianEnrollmentTemplates() (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + guardianFactors + `/sms/templates`
	return restlyGetJSON(req, url, "")
}

// UpdateGuardianEnrollmentTemplates - update a guardian enrollment and verification template
func UpdateGuardianEnrollmentTemplates(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + guardianFactors + `/sms/templates`
	return restlyPutJSON(req, url, body, "")
}

// GetGuardianSNSConfig - get guardian AWS SNS factor provider configuration
func GetGuardianSNSConfig() (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + guardianFactors + `/push-notification/providers/sns`
	return restlyGetJSON(req, url, "")
}

// GetGuardianTwilioConfig - get guardian Twilio factor provider configuration
func GetGuardianTwilioConfig() (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + guardianFactors + `/push-notification/providers/twilio`
	return restlyGetJSON(req, url, "")
}

// UpdateGuardianTwilioConfig - update guardian twilio factor provider configuration
func UpdateGuardianTwilioConfig(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + guardianFactors + `/push-notification/providers/twilio`
	return restlyPutJSON(req, url, body, "")
}

// CreateGuardianEnrollmentTicket - create a guardian enrollment ticket
func CreateGuardianEnrollmentTicket(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + guardianEnrollments + `/ticket`
	return restlyPostJSON(req, url, body, "")
}

// UpdateGuardianFactor - update guardian factor
func UpdateGuardianFactor(name string, body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + guardianFactors + `/` + name
	return restlyPutJSON(req, url, body, "")
}

// GetJob - get a job
func GetJob(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + jobs + `/` + id
	return restlyGetJSON(req, url, "")
}

// GetFailedJob - get a failed job
func GetFailedJob(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + jobs + `/` + id + `/errors`
	return restlyGetJSON(req, url, "")
}

// GetJobResults - get the results of a job
func GetJobResults(id string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + jobs + `/` + id + `/results`
	return restlyGetJSON(req, url, "")
}

// ExportUsers - create a job to export users
func ExportUsers(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + jobs + `/users-exports`
	return restlyPostJSON(req, url, body, "")
}

// ImportUsers - create a job to import users
func ImportUsers(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + jobs + `/users-imports`
	return restlyPostJSON(req, url, "", query)
}

// VerifyEmail - send a `verify email address` email
func VerifyEmail(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + jobs + `/verification-email`
	return restlyPostJSON(req, url, body, "")
}

// GetActiveUserCount - get active user count
func GetActiveUserCount() (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + `active-users`
	return restlyGetJSON(req, url, "")
}

// GetDailyStats - gets the number of logins that occurred int he entered data range
func GetDailyStats(query string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + `daily`
	return restlyGetJSON(req, url, query)
}

// GetTenantSettings - get tenant settings
func GetTenantSettings() (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + `tenants/settings`
	return restlyGetJSON(req, url, "")
}

// UpdateTenantSettings - update tenant settings
func UpdateTenantSettings(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + `tenants/settings`
	return restlyPatchJSON(req, url, body, "")
}

// CreateEmailVerificationTicket - create email verification ticket
func CreateEmailVerificationTicket(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + `tickets/email-verification`
	return restlyPostJSON(req, url, body, "")
}

// CreatePasswordChangeTicket - carete a password change ticket
func CreatePasswordChangeTicket(body string) (gjson.Result, int, error) {
	url := auth0DomainGlobal + mgmtAPIVersion + `tickets/password-change`
	return restlyPostJSON(req, url, body, "")
}
