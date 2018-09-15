package auth0sdk

import (
	"errors"
	"testing"

	"github.com/tidwall/gjson"

	"github.com/apibillme/stubby"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("New", t, func() {

		Convey("New - success", func() {
			stub := stubby.StubFunc(&restlyPostJSON, nil, nil)
			defer stub.Reset()
			err := New("example.auth0.com", "XVJEDJFDKLSJLDFJ", "secret")
			So(err, ShouldBeNil)
		})

		Convey("New - failure", func() {
			stub := stubby.StubFunc(&restlyPostJSON, nil, errors.New("error"))
			defer stub.Reset()
			err := New("example.auth0.com", "XVJEDJFDKLSJLDFJ", "secret")
			So(err, ShouldBeError)
		})
	})

	Convey("Authenication API", t, func() {
		Convey("Signup", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := Signup("", "", "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("ChangePassword", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := ChangePassword("", "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UserInfo", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UserInfo("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("ChallengeMFA", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := ChallengeMFA("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("VerifyOTP", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := VerifyOTP("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("VerifyOOB", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := VerifyOOB("", "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("VerifyRecoveryCode", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := VerifyRecoveryCode("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("AddAuthenticator", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := AddAuthenticator([]string{}, "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("Passwordless", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := Passwordless("", "", "", "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})
	})

	Convey("Mgmt API", t, func() {
		Convey("GetClientGrants", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetClientGrants("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateClientGrant", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateClientGrant("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteClientGrant", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteClientGrant("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateClientGrant", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateClientGrant("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetClients", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetClients("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateClient", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateClient("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetClient", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetClient("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteClient", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteClient("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateClient", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateClient("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("RotateClientSecret", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := RotateClientSecret("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetConnections", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetConnections("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateConnections", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateConnection("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetConnection", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetConnection("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteConnection", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteConnection("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateConnection", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateConnection("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteConnectionUser", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteConnectionUser("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetCustomDomains", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetCustomDomains()
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateCustomDomain", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateCustomDomain("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetCustomDomain", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetCustomDomain("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteCustomDomain", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteCustomDomain("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("VerifyCustomDomain", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := VerifyCustomDomain("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetDeviceCredentials", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetDeviceCredentials("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateDeviceCredential", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateDeviceCredential("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteDeviceCredential", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteDeviceCredential("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetGrants", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetGrants("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteGrant", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteGrant("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("SearchLog", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := SearchLog("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetLog", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetLog("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetResourceServers", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetResourceServers("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateResourceServer", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateResourceServer("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetResourceServer", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetResourceServer("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteResourceServer", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteResourceServer("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateResourceServer", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateResourceServer("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetRules", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetRules("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateRule", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateRule("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetRule", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetRule("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteRule", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteRule("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateRule", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateRule("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetRulesConfigs", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetRulesConfigs()
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteRulesConfig", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteRulesConfig("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("SetRulesConfig", func() {
			stub := stubby.StubFunc(&restlyPutJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := SetRulesConfig("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetUserBlocksByIdentifier", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetUserBlocksByIdentifier("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteUserBlocksByIdentifier", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteUserBlockByIdentifier("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetUserBlock", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetUserBlock("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteUserBlock", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteUserBlock("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("SearchUsers", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := SearchUsers("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateUser", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateUser("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetUser", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetUser("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteUser", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteUser("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateUser", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateUser("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetUserGuardianEntrollments", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetUserGuardianEnrollments("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetUserLog", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetUserLog("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteUserMultifactorProvider", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteUserMultifactorProvider("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UnlinkUserIdentity", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UnlinkUserIdentity("", "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GenerateGuardianRecoveryCode", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GenerateGuardianRecoveryCode("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("LinkUser", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := LinkUser("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("SearchUserByEmail", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := SearchUserByEmail("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetBlacklist", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetBlacklist("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateBlacklist", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateBlacklist("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetEmailTemplate", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetEmailTemplate("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("PatchEmailTemplate", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := PatchEmailTemplate("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateEmailTemplate", func() {
			stub := stubby.StubFunc(&restlyPutJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateEmailTemplate("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateEmailTemplate", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateEmailTemplate("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetEmailProvider", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetEmailProvider("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteEmailProvider", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteEmailProvider()
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateEmailProvider", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateEmailProvider("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("ConfigureEmailProvider", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := ConfigureEmailProvider("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetGuardianFactors", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetGuardianFactors()
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetGuardianEnrollment", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetGuardianEnrollment("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("DeleteGuardianEnrollment", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := DeleteGuardianEnrollment("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetGuardianEnrollmentTemplates", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetGuardianEnrollmentTemplates()
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateGuardianEnrollmentTemplates", func() {
			stub := stubby.StubFunc(&restlyPutJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateGuardianEnrollmentTemplates("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetGuardianSNSConfig", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetGuardianSNSConfig()
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetGuardianTwilioConfig", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetGuardianTwilioConfig()
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateGuardianTwilioConfig", func() {
			stub := stubby.StubFunc(&restlyPutJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateGuardianTwilioConfig("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateGuardianEnrollmentTicket", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateGuardianEnrollmentTicket("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateGuardianFactor", func() {
			stub := stubby.StubFunc(&restlyPutJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateGuardianFactor("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetJob", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetJob("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetFailedJob", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetFailedJob("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetJobResults", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetJobResults("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("ExportUsers", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := ExportUsers("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("ImportUsers", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := ImportUsers("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("VerifyEmail", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := VerifyEmail("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetActiveUserCount", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetActiveUserCount()
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetDailyStats", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetDailyStats("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("GetTenantSettings", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := GetTenantSettings()
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("UpdateTenantSettings", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := UpdateTenantSettings("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreateEmailVerificationTicket", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreateEmailVerificationTicket("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})

		Convey("CreatePasswordChangeTicket", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), nil)
			defer stub.Reset()
			res, err := CreatePasswordChangeTicket("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(err, ShouldBeNil)
		})
	})
}
