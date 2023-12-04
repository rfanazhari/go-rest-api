package user_controller

import (
	"net/http"
	"rest-api/pkg/utils"
)

func (u *userController) GetInfoUser(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, "Success")
}
