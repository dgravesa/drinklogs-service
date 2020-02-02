package controller

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/dgravesa/drinklogs-service/auth"
	"github.com/dgravesa/drinklogs-service/data"
	"github.com/dgravesa/drinklogs-service/model"
)

func postLogs(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	uid, err := auth.VerifyToken(r.Context(), token)

	if err != nil {
		log.Printf("[postLogs] token validation failed {token:\"%s\"}\n", token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	r.ParseForm()
	reqParams, err := newPostDrinkLogsRequest(uid, r.Form)

	// write error response if request is invalid
	if err != nil {
		log.Printf("[postLogs] invalid request {uid:%d, err:\"%s\"}\n", uid, err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	authorized := auth.VerifyWriteAccess(uid, reqParams.uid)

	// authorization failed
	if !authorized {
		log.Printf("[postLogs] authorization failed {uid:%d, owner:%d}\n", uid, reqParams.uid)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = data.InsertDrinkLog(reqParams.uid, reqParams.drinklog)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err) // likely an internal error
		return
	}

	// TODO modify these log statements to use a common scheme
	log.Printf("[postLogs] successful request {uid:%d, req:%v}\n", uid, reqParams)
}

type postDrinkLogsRequest struct {
	uid      string
	drinklog model.DrinkLog
}

func newPostDrinkLogsRequest(uid string, form url.Values) (postDrinkLogsRequest, error) {
	var request postDrinkLogsRequest

	// set request uid to value in form if specified
	if formUID := form.Get("uid"); len(formUID) > 0 {
		request.uid = formUID
	} else {
		request.uid = uid
	}

	// set amount from required value in form
	if amountStr := form.Get("amt"); len(amountStr) > 0 {
		if amount, err := strconv.ParseFloat(amountStr, 64); err == nil {
			request.drinklog.Amount = amount
		} else {
			return postDrinkLogsRequest{}, fmt.Errorf("invalid amount argument: %s", amountStr)
		}
	} else {
		return postDrinkLogsRequest{}, fmt.Errorf("no amount argument provided")
	}

	// set time from value in form if specified; default to now
	if timeStr := form.Get("t"); len(timeStr) > 0 {
		logtime, err := time.Parse(time.RFC822Z, timeStr)
		if err == nil {
			request.drinklog.Time = logtime
		} else {
			// return error if argument is not RFC-822 with time zone
			return postDrinkLogsRequest{}, fmt.Errorf("unable to parse time format: %s", timeStr)
		}
	} else {
		request.drinklog.Time = time.Now()
	}

	return request, nil
}
