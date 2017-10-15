package model

import (
	"fmt"
	"strings"
	"sync"

	"wangqingang/domain_lib/domain"
	"wangqingang/domain_lib/pb"
)

const (
	SubOrderActivePeriod = domain.SubOrderPeriodMin // 12月，1年最小单位
)

const (
	StatusInit    = 0
	StatusStarted = 1
	StatusStopped = 2
)

type RegLists struct {
	sync.RWMutex
	Status        int
	AppIdKey      string
	AppSecret     string
	TemplateID    string
	DomainList    []string
	DomainReqList chan string
}

var RegList *RegLists

func init() {
	RegList = new(RegLists)
	RegList.Init()
}

func (r *RegLists) Set(appId, appSecret, templateId, list string) {
	r.Lock()
	defer r.Unlock()

	r.AppIdKey = appId
	r.AppSecret = appSecret
	r.TemplateID = templateId

	ds := strings.Split(list, "\n")
	for _, d := range ds {
		r.DomainList = append(r.DomainList, strings.TrimSpace(d))
	}
	r.Status = StatusInit
}

func (r *RegLists) Start(appId, appSecret, templateId, list string) {
	r.Init()
	r.Set(appId, appSecret, templateId, list)

	r.Lock()
	r.Status = StatusStarted
	defer r.Unlock()

	go r.GenReq()
	go r.ProcReq()
}

func (r *RegLists) Stop() {
	r.Lock()
	defer r.Unlock()

	close(r.DomainReqList)
	r.Status = StatusStopped
}

func (r *RegLists) Reset() {
	r.Init()
}

func (r *RegLists) Init() {
	r.Status = StatusInit
	r.AppSecret = ""
	r.AppIdKey = ""
	r.DomainList = make([]string, 0)
	r.DomainReqList = make(chan string, 1024)
}

func (r *RegLists) GenReq() {
	for {
		if r.Status != StatusStarted {
			return
		}
		for _, s := range r.DomainList {
			subReq := &pb.SubOrderParam{
				RelatedName:      s,
				Action:           domain.SubOrderActionActivate,
				Period:           SubOrderActivePeriod,
				DomainTemplateID: r.TemplateID,
			}
			req := &pb.CreateOrderRequest{
				Action:        domain.CreateOrderAction,
				SubOrderParam: []*pb.SubOrderParam{subReq},
			}
			reqURI := domain.CreateOrderWithAccess(r.AppIdKey, r.AppSecret, req)
			r.DomainReqList <- reqURI
		}
	}
}

func (r *RegLists) ProcReq() {
	for {
		if r.Status != StatusStarted {
			return
		}
		req, ok := <-r.DomainReqList
		if !ok {
			return
		}
		fmt.Println(req)
	}
}
