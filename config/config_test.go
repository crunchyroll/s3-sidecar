package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/crunchyroll/s3-sidecar/config"
)

var _ = Describe("Config", func() {
	var (
		parsedConfigs *AppConfig
		emptyConfigs  *AppConfig
		err           error
	)

	Describe("Loading Configuration from yml file", func() {
		Context("when the input is a valid config file.", func() {
			BeforeEach(func() {
				parsedConfigs, err = LoadConfiguration("sample-config.yml")
			})

			It("Can parse service port no", func() {
				Expect(parsedConfigs.Service.Listen).To(Equal(8080))
			})

			It("Can parse config for Logging", func() {
				Expect(parsedConfigs.Logging.Level).To(Equal("INFO"))
			})

			It("Can parse Newrelic license", func() {
				Expect(parsedConfigs.Newrelic.License).To(Equal("dummy"))
			})

			It("Can parse S3 Config", func() {
				Expect(parsedConfigs.S3.Region).To(Equal("us-west-2"))
				Expect(parsedConfigs.S3.Bucket).To(Equal("bucket-name"))
			})

			It("Doesn't return any error", func() {
				Expect(err).To(BeNil())
			})
		})

		Context("when the input is an invalid config file.", func() {
			BeforeEach(func() {
				emptyConfigs, err = LoadConfiguration("config/badconfig.yml")
			})

			It("Raises an exception while parsing the configs.", func() {
				Expect(err).To(Not(BeNil()))
			})

			It("Returns an empty Config after parsing.", func() {
				Expect(emptyConfigs).To(BeNil())
			})
		})
	})
})
