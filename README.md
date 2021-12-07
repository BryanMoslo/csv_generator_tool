## CSV files generator.

The goal of this tool is to provide the Device QA team with a simple way to generate a csv file with the amount of records they want, keeping the confidentiality of real users, without wasting time on details.

## Getting Started
Please follow the instructions below to run the tool on your machine.

### Prerequisites
First off, there are a few items you'll need to have installed to run the app.

##### Go
First, go to [this page](https://golang.org/doc/install) to obtain the installer, download Go 1.16.x or higher preferably.
Check your installation result by running the following command:
```
$ go version
```

It should respond with the version that you just installed.
If you run into any problems when installing it, please reach out to your teammates for assistance.

Now that you've gone through all those steps above, you're ready to run the tool. To do that just open the terminal and go to the project's root directory and run:

```
$ go run main.go
```

![Init](https://i.ibb.co/zJNYnFY/Screen-Shot-2021-12-07-at-3-12-49-PM.png)
This starts the tool and enables the commands to generate files.

### How do I generate a user file?
```
>>> generate
```
![Generator](https://i.ibb.co/hcsHFnB/users.png)
This command opens the file types that we have available to generate, you can move and select the desired file with the keyboard. After that, a form will be displayed with the necessary data to generate the file.


### Custom Variables
![Custom Variables](https://i.ibb.co/Q9CvLpb/Screen-Shot-2021-12-07-at-3-21-52-PM.png)
After filling in the required data, you can generate as many custom variables as you want, first the number of custom variables that the file should have is filled and then the names (the data is automatically filled with random values).

### Built With
* [Go](https://golang.org/)

### Acknowledgments
Created with love by [Wawandco](https://wawand.co/)