package pacttest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	bringon "github.com/antweiss/bringon"
	"github.com/gorilla/mux"
	"github.com/otomato-gh/buildinfo"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf("%s/../pacts", dir)
var logDir = fmt.Sprintf("%s/log", dir)
var port, _ = utils.GetFreePort() //34455, false //

var build3455exists = bringon.Build{
	Name:      "#3455",
	Completed: true,
	TimeStamp: time.Date(2017, time.September, 17, 16, 43, 12, 0, time.UTC),
	Info: buildinfo.BuildInfo{
		TestCoverage: 30,
		ApiVersion:   0.1,
		SwaggerLink:  "http://swagger",
		BuildTime:    230},
}

func TestPact(t *testing.T) {
	go startInstrumentedBringon()
	time.Sleep(100000 * time.Millisecond)
	pact := createPact()
	// Verify the Provider with local Pact Files
	log.Println("Start verify ", []string{filepath.ToSlash(fmt.Sprintf("%s/buildReader-bringon.json", pactDir))},
		fmt.Sprintf("http://localhost:%d/setup", port), fmt.Sprintf("http://localhost:%d", port), fmt.Sprintf("%s/buildReader-bringon.json", pactDir))
	err := pact.VerifyProvider(types.VerifyRequest{
		ProviderBaseURL:        fmt.Sprintf("http://localhost:%d", port),
		PactURLs:               []string{filepath.ToSlash(fmt.Sprintf("%s/buildReader-bringon.json", pactDir))},
		ProviderStatesSetupURL: fmt.Sprintf("http://localhost:%d/setup", port),
	})

	if err != nil {
		t.Fatal("Error:", err)
	}

}

var setupRoutes = bringon.Routes{
	bringon.Route{
		"setup",
		"GET",
		"/setup",
		Setup,
	},
}

func startInstrumentedBringon() {
	router := bringon.NewRouter()
	for _, route := range setupRoutes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = bringon.Logger(handler, route.Name)
		log.Println("adding route ", route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		// p will contain a regular expression that is compatible with regular expressions in Perl, Python, and other languages.
		// For example, the regular expression for path '/articles/{id}' will be '^/articles/(?P<v0>[^/]+)$'.
		p, err := route.GetPathRegexp()
		if err != nil {
			return err
		}
		m, err := route.GetMethods()
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(m, ","), t, p)
		return nil
	})

	//	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Println(fmt.Sprintf("[DEBUG] starting service on :%d", port))
		errs <- http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	}()
	select {}

	log.Fatal("Exit", <-errs)
}
func Setup(w http.ResponseWriter, r *http.Request) {
	log.Println("In Setup")
	session := bringon.Dbinit()
	//log.Printf("got collection %v", bCol)
	bCol := session.DB("bringon").C("builds")
	t := bCol.Insert(&build3455exists)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// Setup the Pact client.
func createPact() dsl.Pact {
	// Create Pact connecting to local Daemon
	log.Println("Creating pact")
	return dsl.Pact{
		Port:     6666,
		Consumer: "buildReader",
		Provider: "bringon",
		LogDir:   logDir,
		PactDir:  pactDir,
		LogLevel: "DEBUG",
		Host:     "localhost",
	}
}
