package routers

import (
	auth "BookingVehicle/controllers/auth"
	"BookingVehicle/middlewares/validator"
	"net/http"
)
/*
*	The main router for authentication
*	Author: Lee Tuan
 */
func authRouters(prefix string, r *http.ServeMux) {
	r.Handle(prefix + "/signIn", validator.SignIn(http.HandlerFunc(auth.SignIn)))
	r.Handle(prefix + "/signUp", validator.SignUp(http.HandlerFunc(auth.SignUp)))
}