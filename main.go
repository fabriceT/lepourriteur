package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

const (
	post_url = "http://la-lettre-de-services.info/mailtomut.php"
)

func getEmail(person *gofakeit.PersonInfo) string {
	mail := []string{"sfr.fr", "orange.fr", "nrj.fr", "hotmail.com", "outlook.fr", "wanadoo.fr"}

	return strings.ToLower(person.FirstName + "." + person.LastName + "@" + mail[rand.Intn(len(mail))])
}

func getAge() string {
	return strconv.Itoa(rand.Intn(56) + 18)
}

func getDepartement() string {
	return fmt.Sprintf("%2d", rand.Intn(97))
}

func getPhone() string {
	return fmt.Sprintf("%.2d %d %d %d %d",
		rand.Intn(6)+1,
		rand.Intn(99),
		rand.Intn(99),
		rand.Intn(99),
		rand.Intn(99),
	)
}

func main() {

	gofakeit.Seed(0)
	rand.Seed(time.Now().UnixNano())

	person := gofakeit.Person()

	data := url.Values{
		"requireddep":  {getDepartement()},
		"requirednom":  {person.FirstName + " " + person.LastName},
		"requiredmail": {getEmail(person)},
		"ageMr":        {getAge()},
		"ageMme":       {getAge()},
		"Age1":         {getAge()},
		"requiredtel":  {getPhone()},
	}

	_, err := http.PostForm(post_url, data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(data)

}
