package contactinfo

import "errors"

type Contactinfo struct {
	TypeOfContact  string
	ValueOfContact string
}

func CreateNewContactInfo(contactType string, value string) (*ContactInfo, error) {
	if contactType == " " || value == " " {
		return nil, errors.New("invalid contact info")
	}

	contactInfo := &Contactinfo{
		TypeOfContact:  contactType,
		ValueOfContact: value,
	}

	return contactInfo, nil
}

func (contactinfo *Contactinfo) UpdateUserContactInfo(parameters string, values interface{}) (*Contactinfo, error) {

	switch parameters {
	case "email":
		value, check := values.(string)
		if !check || value == " " {
			return nil, errors.New("invalid updated value passed")
		}
		contactinfo.ValueOfContact = value
	}
	return contactinfo, nil
}
