package model

import (
	"strings"
	"testing"
	"time"
)

const (
	testAccessId     = "LTAIKxPeA194ogHP"
	testAccessSecret = "UtYukSQfOfM9CgVciukMPdLJJ5c4qx"
	testDomainName   = "yx.xin"
	testTemplateID   = "3834855"
)

func TestRegLists_Engine(t *testing.T) {
	req := make([]string, 0)
	for i := 0; i < 50; i++ {
		req = append(req, testDomainName)
	}
	reqs := strings.Join(req, "\n")

	RegList.Init()
	RegList.Start(testAccessId, testAccessSecret, testTemplateID, reqs)
	time.Sleep(time.Duration(1) * time.Second)
	RegList.Stop()
}
