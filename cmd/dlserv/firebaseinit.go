package main

import "github.com/dgravesa/drinklogs-service/auth"

import "log"

func createFirebaseTokenVerifier(keyname string) *auth.FirebaseTokenVerifier {
	tv, err := auth.NewFirebaseTokenVerifier(keyname)

	if err != nil {
		log.Fatalln(err)
	}

	return tv
}
