package generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/jackc/fake"
)

func GenerateUsersFile(total, groupID int, domain string, customVariables []string) error {
	headers := "Client Employee ID,First Name,Last Name,Email Address,CPD Group ID,Term Date,Supervisor Email"
	for _, v := range customVariables {
		headers += fmt.Sprintf(",%s", v)
	}

	headers += "\n"
	bf := bytes.NewBufferString(headers)

	managers := []string{}
	var rnd = rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < total; i++ {
		format := `"%x","%v","%v","%v","%v","%v","%v"`

		firstname := fake.FirstName()
		lastname := fake.LastName()
		email := fmt.Sprintf("%v.%v-%x@%v", strings.ToLower(firstname), strings.ToLower(lastname), i, domain)

		isManager := rand.Intn(2) == 1
		if isManager {
			managers = append(managers, email)
		}

		manager := ""
		if !isManager && len(managers) > 0 {
			manager = managers[rand.Intn(len(managers))]
		}

		line := fmt.Sprintf(format, rnd.Uint32(), firstname, lastname, email, groupID, "", manager)

		for range customVariables {
			line += fmt.Sprintf(",\"%s\"", fake.Word())
		}

		line += "\n"

		bf.Write([]byte(line))
	}

	err := os.MkdirAll("tmp", 0777)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./tmp/users_data.csv", bf.Bytes(), 0777)
	if err != nil {
		return err
	}

	return nil
}
