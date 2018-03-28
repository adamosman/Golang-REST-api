# Golang RESTful API

A simple RESTful web service that exposes an endpoint to upload an ASCII text file and parses it for total word count, number of occurrences of each word, and removes any word that contains \'93blue\'94 (i.e. blueberry, bluegrass, bluekai, etc.)\

## Getting Started / Prerequisites

- [Go 1.10](https://golang.org/dl/) and [Node.js 8.10.0](https://nodejs.org/en/) or higher is required for the install.

- Follow instructions for setting the [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Installing

Navigate to the `client` folder and then install Node dependencies:
```
$ cd ~/Golang-REST-api/client
$ npm install
```
To launch the app in development mode execute the following:
```
$ npm start
```
This will launch a new window in your browser set to [http://localhost:3000](http://localhost:3000)

![Image of API](https://preview.ibb.co/fT2bpn/Screen_Shot_2018_03_27_at_8_30_18_PM.png)

In a new terminal window, navigate to the `Golang-REST-api` folder to build and run the `main.go` file:
```
$ cd ~/Golang-REST-api
$ go run main.go
```
You will be prompted to allow for the application to accept incoming network connections. Click 'Accept' to run the Go application.
## Running the test
On the browser, upload an arbitrary ASCII text file from your disk and select 'Upload'. A text file called `license.txt` is included in the repository for testing purposes.

Below the header a list will appear showing the number of occurences for each word in the file, as well as the total word count.

![Image of completed upload](https://image.ibb.co/mrT7Un/Screen_Shot_2018_03_27_at_8_49_15_PM.png)


## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Adam Osman** - *Initial work* - [Adamosman](https://github.com/adamosman)


## Acknowledgments

* AshikNesin for the [front-end file upload](https://gist.github.com/AshikNesin/e44b1950f6a24cfcd85330ffc1713513#file-react-file-upload-js)
* GoesToEleven for help with the back-end [file upload](https://github.com/GoesToEleven/golang-web-dev/blob/5af87d51b744cca83b945c851a7cf8aaf3e0687f/000_temp/42_class/04/main.go)

