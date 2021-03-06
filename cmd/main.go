// Define the package - this will help maintain scope in your application
package main

// Define the fmt import for formatting
import (
	"fmt"
	"io/ioutil"
	"net/http"

	logr "github.com/sirupsen/logrus"
)

func main() {
	// Print a formatted string to the terminal
	fmt.Println("Hello Fellow Gopher!")

	// Create the first route handler listening on '/'
	http.HandleFunc("/", handler)
	logr.Info("Starting up on 8080")

	http.HandleFunc("/showjoke", jokeHandler)

	// Start the sever
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Assign a variable with a string
	name := "Alex"

	// Logs a message to the terminal using the 3rd party import logrus
	logr.Info("Received request for the home page")

	// Write the response to the byte array - Sprintf formats and returns a string without printing it anywhere
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

func getJoke() (string, error) {
	logr.Infof("Getting joke from API..")

	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

	// Check the request doesnt return an error
	if err != nil {
		return "", err
	}

	// Set the request header
	req.Header.Set("Accept", "text/plain")

	// Make the HTTP request - '.Do' sends an HTTP request and returns an HTTP response
	resp, err := http.DefaultClient.Do(req)

	// Check the request doesn't return an error
	if err != nil {
		return "", err
	}

	// Closes resp.Body at the end of the function (always do this to prevent memory leaks, even if it isn't used)
	defer resp.Body.Close()

	// Read resp.Body until EOF
	body, err := ioutil.ReadAll(resp.Body)

	// Check the ReadAll doesn't return an error
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	// Write the status code 200
	w.WriteHeader(http.StatusOK)

	// Logs a message to the terminal using the 3rd party import logrus
	logr.Infof("Received request to show a joke")

	// getJoke() will return 2 values so we must assign them with x, y
	dadJoke, err := getJoke()

	// Check the request doesnt return an error
	if err != nil {
		logr.Error(err)
	}

	// Write the response to the byte array - Sprintf formats and returns a string without printing it anywhere
	w.Write([]byte(fmt.Sprintf(dadJoke)))
	logr.Info(dadJoke)
}
