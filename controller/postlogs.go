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
	uid, err := auth.VerifyToken(token)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	r.ParseForm()
	reqParams, err := newPostDrinkLogsRequest(uid, r.Form)

	// write error response if request is invalid
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	authorized := auth.VerifyWriteAccess(uid, reqParams.uid)

	// authorization failed
	if !authorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = data.InsertDrinkLog(reqParams.uid, reqParams.drinklog)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err) // likely an internal error
		return
	}
}

type postDrinkLogsRequest struct {
	uid      uint64
	drinklog model.DrinkLog
}

func newPostDrinkLogsRequest(uid uint64, form url.Values) (postDrinkLogsRequest, error) {
	var request postDrinkLogsRequest

	// set request uid to value in form if specified
	if formUIDStr := form.Get("uid"); len(formUIDStr) > 0 {
		if formUID, err := strconv.ParseUint(formUIDStr, 10, 64); err == nil {
			request.uid = formUID
		} else {
			// return error if argument is not an unsigned integer
			return postDrinkLogsRequest{}, fmt.Errorf("invalid uid argument: %s", formUIDStr)
		}
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
