package chaospeddler_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi"
	. "github.com/xchapter7x/chaospeddler/service_broker"
	"github.com/xchapter7x/chaospeddler/service_broker/fake"
)

var _ = Describe("Given a ServiceBroker", func() {
	var serviceBroker = new(ServiceBroker)

	Describe("Given a Provision method", func() {
		Context("When called on a service broker with a valid persistence connection w/ valid args", func() {
			var origNewServiceInstance func(bool) InstanceProvisioner
			var f *fake.BaseService
			BeforeEach(func() {
				f = new(fake.BaseService)
				origNewServiceInstance = NewServiceInstance
				NewServiceInstance = f.NewServiceInstance
			})

			AfterEach(func() {
				NewServiceInstance = origNewServiceInstance
			})

			It("Then it should save the provisioning info", func() {
				_, err := serviceBroker.Provision("", brokerapi.ProvisionDetails{}, true)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(*f.SaveCount).Should(Equal(uint64(1)))
			})
		})
		Context("When called on a service broker not properly initialized", func() {
			var origNewServiceInstance func(bool) InstanceProvisioner
			var f *fake.BaseService
			BeforeEach(func() {
				f = new(fake.BaseService)
				f.ErrFake = errors.New("we have a connection error")
				origNewServiceInstance = NewServiceInstance
				NewServiceInstance = f.NewServiceInstance
			})

			AfterEach(func() {
				NewServiceInstance = origNewServiceInstance
			})

			It("Then it should capture and return the error", func() {
				_, err := serviceBroker.Provision("", brokerapi.ProvisionDetails{}, true)
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	Describe("Given a Deprovision method", func() {
		Context("When called on a service broker with a valid persistence connection w/ valid args", func() {
			var origNewServiceInstance func(bool) InstanceProvisioner
			var f *fake.BaseService
			BeforeEach(func() {
				f = new(fake.BaseService)
				origNewServiceInstance = NewServiceInstance
				NewServiceInstance = f.NewServiceInstance
			})

			AfterEach(func() {
				NewServiceInstance = origNewServiceInstance
			})

			It("Then it should save the provisioning info", func() {
				_, err := serviceBroker.Deprovision("", brokerapi.DeprovisionDetails{}, true)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(*f.SaveCount).Should(Equal(uint64(1)))
			})
		})
		Context("When called on a service broker not properly initialized", func() {
			var origNewServiceInstance func(bool) InstanceProvisioner
			var f *fake.BaseService
			BeforeEach(func() {
				f = new(fake.BaseService)
				f.ErrFake = errors.New("we have a connection error")
				origNewServiceInstance = NewServiceInstance
				NewServiceInstance = f.NewServiceInstance
			})

			AfterEach(func() {
				NewServiceInstance = origNewServiceInstance
			})

			It("Then it should capture and return the error", func() {
				_, err := serviceBroker.Deprovision("", brokerapi.DeprovisionDetails{}, true)
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	Describe("Given a Bind method", func() {
		Context("When called on a service broker with a valid persistence connection w/ valid args", func() {
			var origNewServiceBinding func(bool) BindingProvisioner
			var f *fake.BaseService
			BeforeEach(func() {
				f = new(fake.BaseService)
				origNewServiceBinding = NewServiceBinding
				NewServiceBinding = f.NewServiceBinding
			})

			AfterEach(func() {
				NewServiceBinding = origNewServiceBinding
			})

			It("Then it should save the provisioning info", func() {
				_, err := serviceBroker.Bind("", "", brokerapi.BindDetails{})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(*f.SaveCount).Should(Equal(uint64(1)))
			})
		})
		Context("When called on a service broker not properly initialized", func() {

			var origNewServiceBinding func(bool) BindingProvisioner
			var f *fake.BaseService
			BeforeEach(func() {
				f = new(fake.BaseService)
				f.ErrFake = errors.New("we have a connection error")
				origNewServiceBinding = NewServiceBinding
				NewServiceBinding = f.NewServiceBinding
			})

			AfterEach(func() {
				NewServiceBinding = origNewServiceBinding
			})

			It("Then it should capture and return the error", func() {
				_, err := serviceBroker.Bind("", "", brokerapi.BindDetails{})
				Ω(err).Should(HaveOccurred())
			})
		})
	})

	Describe("Given a Unbind method", func() {
		Context("When called on a service broker with a valid persistence connection w/ valid args", func() {
			var origNewServiceBinding func(bool) BindingProvisioner
			var f *fake.BaseService
			BeforeEach(func() {
				f = new(fake.BaseService)
				origNewServiceBinding = NewServiceBinding
				NewServiceBinding = f.NewServiceBinding
			})

			AfterEach(func() {
				NewServiceBinding = origNewServiceBinding
			})

			It("Then it should save the provisioning info", func() {
				err := serviceBroker.Unbind("", "", brokerapi.UnbindDetails{})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(*f.SaveCount).Should(Equal(uint64(1)))
			})
		})
		Context("When called on a service broker not properly initialized", func() {

			var origNewServiceBinding func(bool) BindingProvisioner
			var f *fake.BaseService
			BeforeEach(func() {
				f = new(fake.BaseService)
				f.ErrFake = errors.New("we have a connection error")
				origNewServiceBinding = NewServiceBinding
				NewServiceBinding = f.NewServiceBinding
			})

			AfterEach(func() {
				NewServiceBinding = origNewServiceBinding
			})

			It("Then it should capture and return the error", func() {
				err := serviceBroker.Unbind("", "", brokerapi.UnbindDetails{})
				Ω(err).Should(HaveOccurred())
			})
		})
	})
})
