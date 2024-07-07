package main

import (
	"contactapp/user"
	"fmt"
)

func main() {

	admin, err := user.NewAdmin("prathmesh")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(admin)

	user1, err := admin.CreateUser("Girish")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user1)

	user2, _ := admin.CreateUser("Yashodip")
	fmt.Println(user2)

	foundUser, err := user1.GetUserByID(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(foundUser)

	err = admin.UpdateUser(3, "name", "Pramod")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user1)

	err = admin.DeleteUser(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User deleted:", user2)

	////For contact user operation

	contact1, err := user1.CreateUserContact("contactnew")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Contact created:", contact1)

	contact2, err := user1.GetContactByID(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(contact2)

	updateduser, err := user1.UpdateUserContact(3, "name", "Pramod")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(updateduser)

	err = user1.DeleteUserContactByID(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user1)

	//For coninfo user operation

	err = user1.CreateUserContactInfo(0, "email", "contactnewinfo@gmai.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("contactinfo updated successfully")

	contactinfo1, err := user1.GetUserContactInfo(0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("%s%v", "The contactinfo is :- ", contactinfo1)

	updatedcontctinfo, err := user1.UpdateUserContactInfo(0, 0, "email", "pratham@gmail.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(updatedcontctinfo)

	err = user1.DeleteContactInfo(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("contactinfo deleted successfully!!!!")

}
