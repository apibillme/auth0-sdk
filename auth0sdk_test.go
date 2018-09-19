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
			stub := stubby.StubFunc(&restlyPostJSON, nil, 0, nil)
			defer stub.Reset()
			err := New("example.auth0.com", "XVJEDJFDKLSJLDFJ", "secret")
			So(err, ShouldBeNil)
		})

		Convey("New - failure", func() {
			stub := stubby.StubFunc(&restlyPostJSON, nil, 0, errors.New("error"))
			defer stub.Reset()
			err := New("example.auth0.com", "XVJEDJFDKLSJLDFJ", "secret")
			So(err, ShouldBeError)
		})
	})

	Convey("Authenication API", t, func() {
		Convey("Signup", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := Signup("", "", "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("ChangePassword", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := ChangePassword("", "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UserInfo", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UserInfo("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("ChallengeMFA", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := ChallengeMFA("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("VerifyOTP", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := VerifyOTP("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("VerifyOOB", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := VerifyOOB("", "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("VerifyRecoveryCode", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := VerifyRecoveryCode("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("AddAuthenticator", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := AddAuthenticator([]string{}, "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("Passwordless", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := Passwordless("", "", "", "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})
	})

	Convey("Mgmt API", t, func() {
		Convey("GetClientGrants", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetClientGrants("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateClientGrant", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateClientGrant("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteClientGrant", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteClientGrant("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateClientGrant", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateClientGrant("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetClients", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetClients("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateClient", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateClient("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetClient", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetClient("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteClient", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteClient("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateClient", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateClient("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("RotateClientSecret", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := RotateClientSecret("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetConnections", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetConnections("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateConnections", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateConnection("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetConnection", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetConnection("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteConnection", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteConnection("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateConnection", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateConnection("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteConnectionUser", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteConnectionUser("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetCustomDomains", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetCustomDomains()
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateCustomDomain", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateCustomDomain("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetCustomDomain", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetCustomDomain("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteCustomDomain", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteCustomDomain("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("VerifyCustomDomain", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := VerifyCustomDomain("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetDeviceCredentials", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetDeviceCredentials("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateDeviceCredential", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateDeviceCredential("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteDeviceCredential", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteDeviceCredential("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetGrants", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetGrants("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteGrant", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteGrant("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("SearchLog", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := SearchLog("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetLog", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetLog("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetResourceServers", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetResourceServers("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateResourceServer", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateResourceServer("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetResourceServer", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetResourceServer("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteResourceServer", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteResourceServer("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateResourceServer", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateResourceServer("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetRules", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetRules("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateRule", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateRule("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetRule", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetRule("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteRule", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteRule("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateRule", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateRule("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetRulesConfigs", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetRulesConfigs()
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteRulesConfig", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteRulesConfig("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("SetRulesConfig", func() {
			stub := stubby.StubFunc(&restlyPutJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := SetRulesConfig("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetUserBlocksByIdentifier", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetUserBlocksByIdentifier("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteUserBlocksByIdentifier", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteUserBlockByIdentifier("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetUserBlock", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetUserBlock("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteUserBlock", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteUserBlock("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("SearchUsers", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := SearchUsers("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateUser", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateUser("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetUser", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetUser("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteUser", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteUser("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateUser", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateUser("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetUserGuardianEntrollments", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetUserGuardianEnrollments("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetUserLog", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetUserLog("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteUserMultifactorProvider", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteUserMultifactorProvider("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UnlinkUserIdentity", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UnlinkUserIdentity("", "", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GenerateGuardianRecoveryCode", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GenerateGuardianRecoveryCode("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("LinkUser", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := LinkUser("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("SearchUserByEmail", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := SearchUserByEmail("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetBlacklist", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetBlacklist("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateBlacklist", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateBlacklist("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetEmailTemplate", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetEmailTemplate("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("PatchEmailTemplate", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := PatchEmailTemplate("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateEmailTemplate", func() {
			stub := stubby.StubFunc(&restlyPutJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateEmailTemplate("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateEmailTemplate", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateEmailTemplate("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetEmailProvider", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetEmailProvider("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteEmailProvider", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteEmailProvider()
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateEmailProvider", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateEmailProvider("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("ConfigureEmailProvider", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := ConfigureEmailProvider("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetGuardianFactors", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetGuardianFactors()
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetGuardianEnrollment", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetGuardianEnrollment("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("DeleteGuardianEnrollment", func() {
			stub := stubby.StubFunc(&restlyDeleteJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := DeleteGuardianEnrollment("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetGuardianEnrollmentTemplates", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetGuardianEnrollmentTemplates()
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateGuardianEnrollmentTemplates", func() {
			stub := stubby.StubFunc(&restlyPutJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateGuardianEnrollmentTemplates("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetGuardianSNSConfig", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetGuardianSNSConfig()
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetGuardianTwilioConfig", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetGuardianTwilioConfig()
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateGuardianTwilioConfig", func() {
			stub := stubby.StubFunc(&restlyPutJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateGuardianTwilioConfig("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateGuardianEnrollmentTicket", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateGuardianEnrollmentTicket("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateGuardianFactor", func() {
			stub := stubby.StubFunc(&restlyPutJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateGuardianFactor("", "")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetJob", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetJob("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetFailedJob", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetFailedJob("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetJobResults", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetJobResults("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("ExportUsers", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := ExportUsers("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("ImportUsers", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := ImportUsers("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("VerifyEmail", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := VerifyEmail("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetActiveUserCount", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetActiveUserCount()
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetDailyStats", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetDailyStats("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("GetTenantSettings", func() {
			stub := stubby.StubFunc(&restlyGetJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := GetTenantSettings()
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("UpdateTenantSettings", func() {
			stub := stubby.StubFunc(&restlyPatchJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := UpdateTenantSettings("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreateEmailVerificationTicket", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreateEmailVerificationTicket("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})

		Convey("CreatePasswordChangeTicket", func() {
			stub := stubby.StubFunc(&restlyPostJSON, gjson.Parse(""), 0, nil)
			defer stub.Reset()
			res, statusCode, err := CreatePasswordChangeTicket("")
			So(res, ShouldResemble, gjson.Parse(""))
			So(statusCode, ShouldResemble, 0)
			So(err, ShouldBeNil)
		})
	})
}
