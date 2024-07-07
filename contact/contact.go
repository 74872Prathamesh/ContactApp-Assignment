package contact

import (
	contactInfo "contactapp/contactinfo"
	"errors"
)

type Contact struct {
	ContactId     int
	Name          string
	Contactinfo   []contactInfo.Contactinfo
	ContactinfoID int
}

func NewContact(contactID int, name string) (*Contact, error) {
	if name == "" {
		return nil, errors.New("invalid contact name")
	}
	contact := &Contact{
		ContactId: contactID,
		Name:      name,
	}
	return contact, nil
}

func (contact *Contact) UpdateContact(typeOfContact string, values string) (*Contact, error) {

	if ID < 0 {
		return errors.New("invalid id")
	}

	switch typeOfContact {
	case "name":
		value, check := values.(string)
		if !check || value == " " {
			return errors.New("invalid updated value passed")
		}
		cotact.Name = value
	}
	return contact, nil
}

func (c *Contact) CreateContactInfo(typeOfcontact string, valueofcontact string) error {

	newcontactinfo, err := Contactinfo.CreateNewContactInfo(typeOfContact, valueofcontact)
	if err != nil {
		return err
	}
	c.Contactinfo = append(c.Contactinfo, newcontactinfo)

	c.ContactinfoID++

	return nil

}

func (c *Contact) GetContactInfo(contactInfoID int) (*contactInfo.Contactinfo, error) {

	for _, info := range c.Contactinfo {

		if c.info.ContactinfoID == contactInfoID {
			return info, nil
		}

	}

	return nil, error.New("contactinfoID not found")

}

func (c *Contact) UpdateContactInfo(contactinfoID int, parameters string, values interface{}) (*contactInfo.Contactinfo, error) {

	if contactinfoID < 0 {
		return nil, errors.New("invalid id")
	}

	var contactinfo, err = c.GetContactInfo(contactinfoID)
	if err != nil {
		return nil, err
	}
	updatedcontactinfo, err := contactinfo.UpdateUserContactInfo(parameters, values)
	if err != nil {
		return nil, err
	}
	return updatedcontactinfo, nil
}

func (c *Contact) DeleteUserContactInfo(infoID int) error {

	if contactinfoID < 0 {
		return error.New("invalid contactinfoid ")
	}

	for i, info := range c.Contactinfo {
		if info.contactinfoID == infoID {
			c.Contactinfo = append(u.Contactinfo[:i], u.Contactinfo[i+1:]...)
			return nil
		}

	}
	return errors.New("id not found")
}
