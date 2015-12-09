package api

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("accountService", func() {
	Describe("#ValidateAction", func() {
		Context("action not set", func() {
			It("Retuns false", func() {
				client := New(Credential{}).account()
				Expect(client.ValidateAction()).To(BeFalse())
			})
		})
		Context("action set", func() {
			It("Retuns true if Insert is set", func() {
				client := New(Credential{}).account()
				client.Insert()
				Expect(client.ValidateAction()).To(BeTrue())
			})
		})
	})
	Describe("#Action", func() {
		Context("Action not set", func() {
			It("Returns a null string", func() {
				client := New(Credential{}).account()
				Expect(client.Action()).To(Equal(""))
			})
		})
		Context("Action not set", func() {
			It("returns action string", func() {
				client := New(Credential{}).account()
				client.Insert()
				Expect(client.Action()).To(Equal("insert"))
			})
		})
	})
	Describe("#Insert", func() {
		It("sets action as insert", func() {
			client := New(Credential{}).account()
			client.Insert()
			Expect(client.Action()).To(Equal("insert"))
		})
	})
	Describe("#Valid", func() {

		var accountInfo AccountDetails
		var credential Credential
		BeforeEach(func() {
			accountInfo = AccountDetails{
				Name: "sarath",
			}
			credential = Credential{UserCredentials{AccessToken: "1212121"}, ClientCredentials{ClientID: "blah", ClientSecret: "ssas"}}
		})

		Context("Action is insert", func() {
			It("Returns false if ClientSecret is not provided or is null string", func() {
				credential.ClientSecret = ""
				c := New(credential).account()
				c.Setaccount(accountInfo)
				c.Insert()
				Expect(c.Valid()).To(BeFalse())
			})
			It("Returns false if ClientID is not provided or is null string", func() {
				credential.ClientID = ""
				c := New(credential).account()
				c.Setaccount(accountInfo)
				c.Insert()
				Expect(c.Valid()).To(BeFalse())
			})
			It("Returns false if AccessToken is not provided or is null string", func() {
				credential.AccessToken = ""
				c := New(credential).account()
				c.Setaccount(accountInfo)
				c.Insert()
				Expect(c.Valid()).To(BeFalse())
			})
			It("Returns false if Name is not provided or is null string", func() {
				accountInfo.Name = ""
				c := New(credential).account()
				c.Setaccount(accountInfo)
				c.Insert()
				Expect(c.Valid()).To(BeFalse())
			})
		})

		It("Returns false if Action is not set ", func() {
			c := New(credential).account()
			c.Setaccount(accountInfo)
			Expect(c.Valid()).To(BeFalse())
		})
	})
	Describe("#Setaccount", func() {
		It("Set the Account Details", func() {
			var account = AccountDetails{Name: "Sarath"}
			client := New(Credential{}).account()
			client.Setaccount(account)
			Expect(client.account.Name).To(Equal("Sarath"))
			Expect(client.account.Sid).To(Equal(""))
		})
	})
})
