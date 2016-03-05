package chaospeddler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi"
	. "github.com/xchapter7x/chaospeddler/service_broker"
	"github.com/xchapter7x/chaospeddler/service_broker/fakes"
)

var _ = Describe("Given a ServiceBroker", func() {
	var serviceBroker *ServiceBroker
	var origNewServiceInstance func() InstanceProvisioner
	var origNewServiceBinding func() BindingProvisioner
	var instanceProvisioner *fakes.FakeInstanceProvisioner
	var bindingProvisioner *fakes.FakeBindingProvisioner
	var db *fakes.FakeGormDB
	var orch *fakes.FakeOrchestrator

	BeforeEach(func() {
		db = new(fakes.FakeGormDB)
		db.WhereReturns(db)
		orch = new(fakes.FakeOrchestrator)
		orch.DBReturns(db)
		serviceBroker = &ServiceBroker{
			Orchestrator: orch,
		}
		instanceProvisioner = new(fakes.FakeInstanceProvisioner)
		origNewServiceInstance = NewServiceInstance
		NewServiceInstance = func() InstanceProvisioner {
			return instanceProvisioner
		}

		bindingProvisioner = new(fakes.FakeBindingProvisioner)
		origNewServiceBinding = NewServiceBinding
		NewServiceBinding = func() BindingProvisioner {
			return bindingProvisioner
		}
	})

	AfterEach(func() {
		NewServiceInstance = origNewServiceInstance
	})

	Describe("Given a Provision method", func() {
		Context("When called on a service broker with a valid persistence connection w/ valid args", func() {
			It("Then it should capture all provisioning details", func() {
				controlID := "fakeid"
				controlProvisionDetails := brokerapi.ProvisionDetails{}
				controlAsync := true
				_, err := serviceBroker.Provision(controlID, controlProvisionDetails, controlAsync)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(instanceProvisioner.SetInstanceIDCallCount()).Should(Equal(1))
				Ω(instanceProvisioner.SetProvisionDetailsCallCount()).Should(Equal(1))
				Ω(instanceProvisioner.SetActiveCallCount()).Should(Equal(1))
				Ω(instanceProvisioner.SetInstanceIDArgsForCall(0)).Should(Equal(controlID))
				Ω(instanceProvisioner.SetProvisionDetailsArgsForCall(0)).Should(Equal(controlProvisionDetails))
				Ω(instanceProvisioner.SetActiveArgsForCall(0)).Should(Equal(controlAsync))
			})
			It("then it should save the privisoning info", func() {
				controlID := "fakeid"
				controlProvisionDetails := brokerapi.ProvisionDetails{}
				controlAsync := true
				serviceBroker.Provision(controlID, controlProvisionDetails, controlAsync)
				Ω(db.CreateCallCount()).Should(Equal(1))
			})
		})
	})

	Describe("Given a Deprovision method", func() {
		Context("When called on a service broker with a valid persistence connection w/ valid args", func() {
			It("Then it should save the de-provisioning info", func() {
				controlInstanceID := "random stuff here"
				_, err := serviceBroker.Deprovision(controlInstanceID, brokerapi.DeprovisionDetails{}, true)
				_, args := db.WhereArgsForCall(0)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(db.WhereCallCount()).Should(Equal(1))
				Ω(args[0]).Should(Equal(controlInstanceID))
				Ω(db.FirstCallCount()).Should(Equal(1))
				Ω(db.CreateCallCount()).Should(Equal(1))
			})
		})
	})

	Describe("Given a Bind method", func() {
		Context("When called on a service broker with a valid persistence connection w/ valid args", func() {
			It("Then it should save the provisioning info", func() {
				_, err := serviceBroker.Bind("", "", brokerapi.BindDetails{})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(bindingProvisioner.SetInstanceIDCallCount()).Should(Equal(1))
				Ω(bindingProvisioner.SetBindingIDCallCount()).Should(Equal(1))
				Ω(bindingProvisioner.SetBindDetailsCallCount()).Should(Equal(1))
				Ω(bindingProvisioner.SetActiveCallCount()).Should(Equal(1))
				Ω(db.CreateCallCount()).Should(Equal(1))
			})
		})
	})

	Describe("Given a Unbind method", func() {
		Context("When called on a service broker with a valid persistence connection w/ valid args", func() {
			It("Then it should save the provisioning info", func() {

				controlInstanceID := "random stuff here"
				controlBindID := "i bind stuff"
				err := serviceBroker.Unbind(controlInstanceID, controlBindID, brokerapi.UnbindDetails{})
				_, args := db.WhereArgsForCall(0)

				Ω(err).ShouldNot(HaveOccurred())
				Ω(db.WhereCallCount()).Should(Equal(1))
				Ω(args[0]).Should(Equal(controlInstanceID))
				Ω(args[1]).Should(Equal(controlBindID))
				Ω(db.FirstCallCount()).Should(Equal(1))
				Ω(db.CreateCallCount()).Should(Equal(1))
			})
		})
	})
})
