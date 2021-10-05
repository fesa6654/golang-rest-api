package uuid

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/nu7hatch/gouuid"
)

func GenerateUUID(w http.ResponseWriter, r *http.Request) {

	u, err := uuid.NewV4()

	if err != nil {
		fmt.Println("UUID not Created !")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"uuid": u.String(),
	})
}
