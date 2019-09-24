package controllers

import (
	"fmt"

	"github.com/chinyart/nedum-model/models"
)

//DisplayPersonBio to diplay person bio
func DisplayPersonBio() {
	fmt.Println("hello")
	var show = models.PersonBio{
		Name: "john",
	}

	fmt.Println("hi " + show.Name)
}
