package ldap

import (
	"fmt"

	"inwinstack/cgmh/apiserver/pkg/models"

	"gopkg.in/ldap.v3"
)

type Flag struct {
	Protocol string
	Host     string
	Username string
	Password string
	DN       string
	OU       string
}

type LDAP struct {
	flag *Flag
}

func NewServer(f *Flag) *LDAP {
	return &LDAP{flag: f}
}

func (s *LDAP) connect() (*ldap.Conn, error) {
	conn, err := ldap.Dial(s.flag.Protocol, s.flag.Host)
	if err != nil {
		return nil, err
	}

	if err = conn.Bind(s.flag.Username, s.flag.Password); err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *LDAP) Search(dn, filter string, attributes []string) ([]*ldap.Entry, error) {
	conn, err := s.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if filter == "" {
		filter = "(&(objectClass=*))"
	}

	req := ldap.NewSearchRequest(dn, 2, 0, 0, 0, false, filter, attributes, nil)
	srearch, err := conn.Search(req)
	if err != nil {
		return nil, err
	}

	if len(srearch.Entries) == 0 {
		return nil, ErrorEntryNotFound
	}
	return srearch.Entries, nil
}

func (s *LDAP) IsExist(dn string, attributes []string) bool {
	entries, err := s.Search(dn, "", attributes)
	if err == nil {
		if len(entries) > 0 {
			return true
		}
		return false
	}
	return false
}

func (s *LDAP) AddOU(name, desc string) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	dn := fmt.Sprintf("ou=%s,%s", name, s.flag.DN)
	req := ldap.NewAddRequest(dn, nil)
	req.Attribute("objectClass", []string{"organizationalUnit", "top"})
	req.Attribute("ou", []string{name})
	req.Attribute("description", []string{desc})
	return conn.Add(req)
}

func (s *LDAP) AddGroup(name, ou, desc string) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	dn := fmt.Sprintf("cn=%s,ou=%s,%s", name, ou, s.flag.DN)
	req := ldap.NewAddRequest(dn, nil)
	req.Attribute("objectClass", []string{"extensibleObject", "groupOfNames", "top"})
	req.Attribute("cn", []string{name})
	req.Attribute("description", []string{desc})
	return conn.Add(req)
}

func (s *LDAP) AddUser(user *model.User, password string) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	dn := fmt.Sprintf("uid=%s,ou=%s,%s", user.UUID, s.flag.OU, s.flag.DN)
	req := ldap.NewAddRequest(dn, nil)
	objects := []string{"extensibleObject", "uidObject", "userSecurityInformation", "account", "top"}
	req.Attribute("objectClass", objects)
	req.Attribute("uid", []string{user.UUID})
	req.Attribute("name", []string{user.Name})
	req.Attribute("email", []string{user.Email})
	req.Attribute("mobile", []string{user.Phone})
	req.Attribute("userPassword", []string{password})
	return conn.Add(req)
}

func (s *LDAP) ModifyPassword(uid, oldPasswd, newPasswd string) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	dn := fmt.Sprintf("uid=%s,ou=%s,%s", uid, s.flag.OU, s.flag.DN)
	req := ldap.NewPasswordModifyRequest(dn, oldPasswd, newPasswd)
	if _, err := conn.PasswordModify(req); err != nil {
		return err
	}
	return nil
}

func (s *LDAP) ForceModifyPassword(uid, password string) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	dn := fmt.Sprintf("uid=%s,ou=%s,%s", uid, s.flag.OU, s.flag.DN)
	req := ldap.NewModifyRequest(dn, nil)
	req.Replace("userPassword", []string{password})
	return conn.Modify(req)
}

func (s *LDAP) DelUser(uid string) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	dn := fmt.Sprintf("uid=%s,ou=%s,%s", uid, s.flag.OU, s.flag.DN)
	req := ldap.NewDelRequest(dn, nil)
	return conn.Del(req)
}
