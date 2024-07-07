package user

import (
	contact "contactapp/contact"
	"contactapp/contactinfo"
	"cotactapp/contactinfo"
	"errors"
)

type User struct {
	UserID           int
	Name             string
	IsAdmin          bool
	Contacts         []*contact.Contact
	currentContactID int
}

var allUsers = []User{}
var userId = 0

func NewAdmin(name string) (*User, error) {

	if name == " " {
		return nil, errors.New("invalid name!!")
	}
	return &User{
		UserID:  userId,
		Name:    name,
		IsAdmin: true,
	}, nil
}

func (admin *User) CreateUser(name string) (*User, error) {

	if !admin.IsAdmin {
		return nil, errors.New("Not Authorized")
	}
	if name == " " {
		return nil, errors.New("Please give valid user")
	}
	userId++
	user := &User{
		UserID:  userId,
		Name:    name,
		IsAdmin: false,
	}

	allUsers = append(allUsers, *user)
	return user, nil

}

func (admin *User) GetUserByID(ID int) (*User, error) {

	for _, user := range allUsers {
		if user.UserID == ID {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (admin *User) UpdateUser(ID int, parameters string, values interface{}) error {

	if !admin.IsAdmin {
		return errors.New("Not Authorized")
	}
	if ID < 0 {
		return errors.New("invalid id")
	}
	var usertoupdate User
	for _, user := range allUsers {
		if user.UserID == ID {
			usertoupdate = user
		}
	}
	switch parameters {
	case "name":
		value, check := values.(string)
		if !check || value == " " {
			return errors.New("invalid updated value passed")
		}
		usertoupdate.Name = value
	case "isadmin":
		value, check := values.(bool)
		if !check {
			return errors.New("ivalid value passed")
		}
		usertoupdate.IsAdmin = value
	}
	return errors.New("Id not found")
}

func (admin *User) DeleteUser(ID int) error {
	if !admin.IsAdmin {
		return errors.New("not authorized")
	}

	for i, user := range allUsers {
		if user.UserID == ID {
			allUsers = append(allUsers[:i], allUsers[i+1:]...)
			return nil
		}
	}

	return errors.New("User not found")
}

func (u *User) CreateUserContact(name string) (*contact.Contact, error) {

	u.currentContactID++
	newContact, err := contact.NewContact(u.currentContactID, name)
	if err != nil {
		return nil, err
	}
	u.Contacts = append(u.Contacts, newContact)
	return newContact, nil

}

func (u *User) GetContactByID(ID int) (*contact.Contact, error) {

	if ID < 0 {
		return nil, errors.New("invalid input !!! ")
	}

	for _, foundContact := range u.Contacts {
		if foundContact.ContactId == ID {
			return foundContact, nil

		}
	}

	return nil, errors.New("!!! contact not found ")
}

func (u *User) UpdateUserContact(ID int, typeOfContact string, values string) (*contact.Contact, error) {

	if ID < 0 {
		return errors.New("invalid id")
	}
	var contact, err = u.GetContactByID(ID)
	if err != nil {
		return nil, err
	}
	updatedcontact, err := contact.UpdateContact(typeOfContact, values)
	if err != nil {
		return nil, err
	}

	return updatedcontact, err
}

func (u *User) DeleteUserContactByID(ID int) error {

	if ID < 0 {
		return errors.New("invalid Id")
	}

	for i, contact := range u.Contacts {
		if contact.ContactId == ID {
			u.Contacts = append(u.Contacts[:i], u.Contacts[i+1:]...)
			return nil
		}
	}
	return errors.New("id not found")
}

func (u *User) CreateUserContactInfo(contactID int, typeOfContact string, valueofcontact string) error {

	var usercontact, err = u.GetContactByID(contactID)
	if err != nil {
		return err
	}
	err = usercontact.CreateContactInfo(typeOfContact, valueofcontact)
	if err != nil {
		return err
	}
	return nil

}

func (u *User) GetUserContactInfo(contactID int, contactInfoID int) (*contactinfo.ContactInfo, error) {

	var usercontact, err = u.GetContactByID(contactID)
	if err != nil {
		return err, nil
	}
	foundinfo, err := usercontact.GetContactInfo(contactInfoID)
	if err != nil {
		return err, nil
	}
	return foundinfo, err

}

func (u *User) UpdateUserContactInfo(contactID int, contactInfoID int, parameters string, values interface{}) (*contactinfo.ContactInfo, error) {

	var usercontact, err = u.GetContactByID(contactID)
	if err != nil {
		return err, nil
	}
	updatedcontctinfo, err := usercontact.UpdateContactInfo(contactInfoID, parameters, values)
	if err != nil {
		return nil, err
	}
	return updatedcontctinfo, err

}

func (u *User) DeleteContactInfo(contactID int, contactInfoID int) error {

	var usercontact, err = u.GetContactByID(contactID)
	if err != nil {
		return err
	}
	err = usercontact.DeleteUserContactInfo(contactInfoID)
	if err != nil {
		return err
	}
	return nil

}
