/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package helm_test

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"helm.sh/helm/v3/pkg/action"

	"github.com/kubernetes-sigs/minibroker/pkg/helm"
	"github.com/kubernetes-sigs/minibroker/pkg/log"
)

var _ = Describe("Config", func() {
	Describe("Config", func() {
		Describe("NewDefaultConfigProvider", func() {
			It("should satisfy the ConfigProvider interface", func() {
				var config helm.ConfigProvider = helm.NewDefaultConfigProvider()
				Expect(config).NotTo(BeNil())
			})
		})

		Describe("ConfigProviderProvider", func() {
			It("should fail when actionConfig.Init fails", func() {
				actionConfig := &action.Configuration{}
				configInitializerProvider := func() (*action.Configuration, helm.ConfigInitializer) {
					initializer := func(genericclioptions.RESTClientGetter, string, string, action.DebugLog) error {
						return fmt.Errorf("some error")
					}
					return actionConfig, initializer
				}

				configProvider := helm.ConfigProviderProvider(
					log.NewNoop(),
					configInitializerProvider,
					"",
					"",
				)
				cfg, err := configProvider("my-namespace")
				Expect(err).To(Equal(fmt.Errorf("failed to provide action configuration: some error")))
				Expect(cfg).To(BeNil())
			})

			It("should succeed", func() {
				actionConfig := &action.Configuration{}
				initCount := 0
				initializer := func(
					_ genericclioptions.RESTClientGetter,
					_ string,
					_ string,
					log action.DebugLog,
				) error {
					log("whatever")
					initCount += 1
					return nil
				}
				configInitializerProvider := func() (*action.Configuration, helm.ConfigInitializer) {
					return actionConfig, initializer
				}

				configProvider := helm.ConfigProviderProvider(
					log.NewNoop(),
					configInitializerProvider,
					"",
					"",
				)
				cfg, err := configProvider("my-namespace")
				Expect(err).NotTo(HaveOccurred())
				Expect(cfg).To(Equal(actionConfig))
				Expect(initCount).To(Equal(1))
			})
		})
	})

	Describe("DefaultConfigInitializerProvider", func() {
		It("should return action.Configuration.Init", func() {
			config, initializer := helm.DefaultConfigInitializerProvider()
			Expect(reflect.ValueOf(initializer).Pointer()).
				To(Equal(reflect.ValueOf(config.Init).Pointer()))
		})
	})
})
