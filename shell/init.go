package shell

import (
	"errors"
	"strconv"
	"users_generator/generator"

	"github.com/abiosoft/ishell"
)

func Init() error {
	shell := ishell.New()
	shell.Println("CSV files generator")

	shell.AddCmd(&ishell.Cmd{
		Name: "generate",
		Help: "generate a csv file",
		Func: func(c *ishell.Context) {
			choice := c.MultiChoice([]string{
				"Users",
				"Accessories (WIP)",
				"Device Models (WIP)",
			}, "What kind of file do you want to generate?")

			switch choice {
			case 0:
				c.Println("Generating users file...")
				if err := readUsersData(c); err != nil {
					c.Println(err)
					return
				}

				c.Println("Done!, check the generated file in tmp/users_data.csv")
			case 1:
				c.Println("This feature is in progress...")
			case 2:
				c.Println("This feature is in progress...")
			}
		},
	})

	shell.Run()
	return nil
}

func readUsersData(c *ishell.Context) error {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	c.Print("Number of users: ")
	numberOfUsersStr := c.ReadLine()

	totalUsers, err := strconv.Atoi(numberOfUsersStr)
	if err != nil {
		return errors.New("invalid number of users")
	}

	c.Print("Email Domain @")
	emailDomain := c.ReadLine()

	c.Print("Program Group ID: ")
	programGroupIDStr := c.ReadLine()

	programGroupID, err := strconv.Atoi(programGroupIDStr)
	if err != nil {
		return errors.New("invalid program Group ID")
	}

	c.Print("Do you want to add custom variables? (y/n): ")
	hasVariables := c.ReadLine()

	var customVariables []string
	if hasVariables == "y" {
		c.Print("How many custom variables do you want to add?: ")
		numberOfVariablesStr := c.ReadLine()

		numberOfVariables, err := strconv.Atoi(numberOfVariablesStr)
		if err != nil {
			return errors.New("invalid number of variables")
		}

		for i := 0; i < numberOfVariables; i++ {
			c.Print("Custom variable " + strconv.Itoa(i+1) + ": ")
			customVariables = append(customVariables, c.ReadLine())
		}
	}

	generator.GenerateUsersFile(totalUsers, programGroupID, emailDomain, customVariables)

	return nil
}
