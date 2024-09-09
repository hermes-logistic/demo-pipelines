package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintHandler(t *testing.T) {
	r := Routes{}
	router := r.SetUpRouter()

	router.GET("/", r.Print)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Serve HTTP
	router.ServeHTTP(w, req)

	// Check the status code
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := `{"message":"Welcome to my API with Golang [tags]"}`

	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, expected, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateRoutes(t *testing.T) {
	// Crea una instancia de Routes
	r := Routes{}

	// Llama a la función CreateRoutes()
	r.CreateRoutes("a")

	// Verifica que la ruta "/" esté configurada correctamente
	reqRoot, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rrRoot := httptest.NewRecorder()
	r.Routes.ServeHTTP(rrRoot, reqRoot)
	if rrRoot.Code != http.StatusOK {
		t.Errorf("La ruta '/' no devolvió un código de estado OK")
	}

	// Verifica que la ruta "/tasks" esté configurada correctamente
	reqTasks, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	rrTasks := httptest.NewRecorder()
	r.Routes.ServeHTTP(rrTasks, reqTasks)
	if rrTasks.Code != http.StatusOK {
		t.Errorf("La ruta '/tasks' no devolvió un código de estado OK")
	}
}
