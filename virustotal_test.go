package main_test

import (
	"testing"

	main "github.com/m-mizutani/deepalert-virustotal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type VtTestConfig struct {
	VirusTotalToken string `json:"virustotal_token"`
}

func TestVirusTotalFile(t *testing.T) {
	vt := main.NewVirusTotal(testConfig.VirusTotalToken)
	report, err := vt.QueryFile("52d3df0ed60c46f336c131bf2ca454f73bafdc4b04dfa2aea80746f5ba9e6d1c")

	require.Nil(t, err)
	assert.True(t, report.Positives > 40)
	assert.NotNil(t, report.Scans["AVG"])
	assert.Equal(t, "", report.Scans["hoge"].Version)
}

func TestVirusTotalBulkFileQuery(t *testing.T) {
	vt := main.NewVirusTotal(testConfig.VirusTotalToken)
	hvList := []string{
		"52d3df0ed60c46f336c131bf2ca454f73bafdc4b04dfa2aea80746f5ba9e6d1c",
		"07b97ef6786f87a63c3acb4558a0889840c66b04732a63a462ca4f2b1b8bc141",
	}
	reports, err := vt.QueryFileBulk(hvList)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(reports))
}

func TestVirusTotalIPAddr(t *testing.T) {
	vt := main.NewVirusTotal(testConfig.VirusTotalToken)
	report, err := vt.QueryIPAddr("195.22.26.248")

	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(report.DetectedCommunicatingSamples))
	assert.NotEqual(t, 0, len(report.DetectedDownloadedSamples))
	assert.NotEqual(t, 0, len(report.DetectedReferrerSamples))
	assert.NotEqual(t, 0, len(report.DetectedURLs))
	assert.NotEqual(t, 0, len(report.Resolutions))
	assert.Equal(t, 1, report.ResponseCode)
}

func TestVirusTotalDomain(t *testing.T) {
	// 027.ru
	vt := main.NewVirusTotal(testConfig.VirusTotalToken)
	report, err := vt.QueryDomain("027.ru")

	assert.Nil(t, err)
	assert.Equal(t, 1, report.ResponseCode)
	assert.NotZero(t, len(report.Resolutions))
	assert.NotZero(t, len(report.DetectedURLs))
}
