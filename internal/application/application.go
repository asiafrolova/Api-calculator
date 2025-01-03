package application

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/asiafrolova/Api-calculator/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}
func (a *Application) Run() error {
	for {

		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}

		text = strings.TrimSpace(text)

		if text == "exit" {
			log.Println("aplication was successfully closed")
			return nil
		}

		result, err := calculation.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed wit error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := calculation.Calc(request.Expression)

	if err != nil {
		if errors.Is(err, calculation.ErrInvalidExpression) || errors.Is(err, calculation.ErrDivisionByZero) || errors.Is(err, calculation.ErrEmptyExp) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": %s}`, err.Error())

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Fprintf(w, `{"error": %s}`, err.Error())

		}

	} else {

		fmt.Fprintf(w, `{"result": %f}`, result)
	}

}
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (a *Application) RunServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", CalcHandler)
	handler := LoggingMiddleware(mux)

	return http.ListenAndServe(":"+a.config.Addr, handler)
}
