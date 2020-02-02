package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/dgravesa/drinklogs-service/auth"
	"github.com/dgravesa/drinklogs-service/data"
)

func getLogs(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	uid, err := auth.VerifyToken(r.Context(), token)

	// token authentication failed
	if err != nil {
		log.Printf("[getLogs] token validation failed: %s {token:\"%s\"}\n", err, token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	reqParams, err := newGetDrinkLogsRequest(uid, r.URL.Query())

	// write error response if request is invalid
	if err != nil {
		log.Printf("[getLogs] invalid request {uid:%s, err:\"%s\"}\n", uid, err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	authorized := auth.VerifyReadAccess(uid, reqParams.uid)

	// authorization failed
	if !authorized {
		log.Printf("[getLogs] authorization failed {uid:%s, owner:%s}\n", uid, reqParams.uid)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	drinklogs, err := data.DrinkLogsInRange(reqParams.uid, reqParams.ti, reqParams.tf)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err) // likely an internal error
		return
	}

	// TODO modify these log statements to use a common scheme
	log.Printf("[getLogs] successful request {uid:%s, req:%v, reslen:%d}\n",
		uid, reqParams, len(drinklogs))
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(drinklogs)
}

type getDrinkLogsRequestParams struct {
	uid string
	ti  time.Time
	tf  time.Time
}

func newGetDrinkLogsRequest(requesterID string, query url.Values) (getDrinkLogsRequestParams, error) {
	// initialize to defaults
	currentTime := time.Now()
	drinkLogsReqParams := getDrinkLogsRequestParams{
		uid: requesterID,
		tf:  currentTime,
		ti:  currentTime.Truncate(24 * time.Hour),
	}

	// get uid from query if specified
	if queryUID := query.Get("uid"); len(queryUID) > 0 {
		drinkLogsReqParams.uid = queryUID
	}

	// get begin time from query if specified
	if beginTimeStr := query.Get("t1"); len(beginTimeStr) > 0 {
		if beginTime, err := time.Parse(time.RFC822Z, beginTimeStr); err == nil {
			drinkLogsReqParams.ti = beginTime
		} else {
			return getDrinkLogsRequestParams{}, fmt.Errorf("invalid t1 format: %s", beginTimeStr)
		}
	}

	// get end time from query if specified
	if endTimeStr := query.Get("t2"); len(endTimeStr) > 0 {
		if endTime, err := time.Parse(time.RFC822Z, endTimeStr); err == nil {
			drinkLogsReqParams.ti = endTime
		} else {
			return getDrinkLogsRequestParams{}, fmt.Errorf("invalid t2 format: %s", endTimeStr)
		}
	}

	return drinkLogsReqParams, nil
}
