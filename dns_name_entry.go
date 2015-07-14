package crowbar

import (
	"errors"
	"log"
	"strconv"
)

type DnsNameEntry struct {
	ID                  int64  `json:"id,omitempty"`
	NetworkAllocationID int64  `json:"network_allocation_id,omitempty"`
	DnsNameFilterID     int64  `json:"dns_name_filter_id,omitempty"`
	Name                string `json:"name,omitempty"`
	RRType              string `json:"rr_type,omitempty"`
	CreatedAt           string `json:"created_at,omitempty"`
	UpdatedAt           string `json:"updated_at,omitempty"`
}

func (o *DnsNameEntry) Id() string {
	if o.ID != 0 {
		return strconv.FormatInt(o.ID, 10)
	} else {
		log.Panic("DnsNameEntry has no ID or name")
		return ""
	}
}

func (o *DnsNameEntry) SetId(s string) error {
	if o.ID != 0 || o.Name != "" {
		return errors.New("SetId can only be used on an un-IDed object")
	}
	if id, err := strconv.ParseInt(s, 10, 64); err == nil {
		o.ID = id
	} else {
		return errors.New("Cannot create ID")
	}
	return nil
}

func (o *DnsNameEntry) ApiName() string {
	return "dns_name_entries"
}

func (o *DnsNameEntry) Match() (res []*DnsNameEntry, err error) {
	res = make([]*DnsNameEntry, 0)
	return res, session.match(o, &res, o.ApiName(), "match")
}

func DnsNameEntrys() (res []*DnsNameEntry, err error) {
	res = make([]*DnsNameEntry, 0)
	return res, session.list(&res, "dns_name_entries")
}