package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/davidjulien/tesla/config"
	"github.com/davidjulien/tesla/controller"
	"github.com/davidjulien/tesla/data"
	"github.com/davidjulien/tesla/model"
	"github.com/davidjulien/tesla/server"

	"github.com/gorilla/mux"
)

func TestCreateUser(t *testing.T) {
	c := generateTestController()
	rr, _ := generateRequestJSONString(
		"/api/user/create",
		"POST",
		`{"first_name":"Bob","last_name":"Smith","phone":"6546987894"}`,
		c.CreateUser)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateUser(t *testing.T) {
	c := generateTestController()
	rr, _ := generateRequestJSONString(
		"/api/user/update",
		"POST",
		`{"user_id":"1","first_name":"FirstNameUpdated2","last_name":"LastName","phone":"1234567891"}`,
		c.UpdateUser)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetUsers(t *testing.T) {
	c := generateTestController()
	rr, _ := generateRequestJSONString(
		"/api/users",
		"GET",
		"",
		c.GetUsers)

	var users model.Users
	if err := json.NewDecoder(rr.Body).Decode(&users); err != nil {
		return
	}

	// Print all users
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateRetaurant(t *testing.T) {
	c := generateTestController()
	rr, _ := generateRequestJSONString(
		"/api/restaurant/create",
		"POST",
		`{
			"restaurant_name":"tesla",
			"category":"american",
			"address":"3500 Deer Creek Rd, Palo Alto, CA 94304, USA",
			"state": "California",
			"city": "Palo Alto",
			"zip_code": "94304"
		}`,
		c.CreateRestaurant)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestAddAddress(t *testing.T) {
	c := generateTestController()
	rr, _ := generateRequestJSONString(
		"/api/restaurant/add/address",
		"POST",
		`{
			"restaurant_id":"1",
			"address":"3501 Deer Creek Rd, Palo Alto, CA 94305, USA",
			"state": "California",
			"city": "Palo Alto",
			"zip_code": "94305"
		}`,
		c.AddAddressToRestaurant)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateRestaurant(t *testing.T) {
	c := generateTestController()
	rr, _ := generateRequestJSONString(
		"/api/restaurant/update",
		"POST",
		`{
			"restaurant_id":"1",
			"restaurant_name":"yummy restaurant",
			"category":"chinese"
		}`,
		c.UpdateRestaurant)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateRating(t *testing.T) {
	c := generateTestController()
	rr, _ := generateRequestJSONString(
		"/api/rating/create",
		"POST",
		`{
			"cost":5,
			"food":4,
			"cleanliness":3,
			"service":2,
			"address_id":"1",
			"user_id":"1"
		}`,
		c.CreateRating)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateRating(t *testing.T) {
	c := generateTestController()
	rr, _ := generateRequestJSONString(
		"/rating/update",
		"POST",
		`{
			"rating_id":"1",
			"cost":1,
			"food":1,
			"cleanliness":1,
			"service":1
		}`,
		c.UpdateRating)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// Generates a test register user request with jsonString as body
func generateRequestJSONString(
	endpoint string,
	method string,
	jsonString string,
	h func(http.ResponseWriter, *http.Request)) (rr *httptest.ResponseRecorder, r *http.Request) {

	body := strings.NewReader(jsonString)

	r, _ = http.NewRequest(method, endpoint, body)

	rr = httptest.NewRecorder()
	handler := http.HandlerFunc(h)
	handler.ServeHTTP(rr, r)

	return rr, r
}

// Generates a test controller with a server with a nil server
func generateTestController() *controller.Controller {
	cfg := config.LocalTest()
	dal := data.New(cfg)

	s := &server.Server{
		Router: mux.NewRouter().StrictSlash(true),
		Dal:    dal,
	}

	c := controller.Init(s)

	return c
}
