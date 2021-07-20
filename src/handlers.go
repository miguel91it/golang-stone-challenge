package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/gorilla/mux"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	accounts := db.FindAccounts()

	if err := json.NewEncoder(w).Encode(accounts); err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "error to encode accounts list to return to the API caller: %s", err.Error())

		return
	}

}

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	accounts := db.FindAccounts()

	for _, account := range accounts {

		idToFind, err := strconv.Atoi(params["id"])

		if err != nil {
			w.WriteHeader(http.StatusNotFound)

			fmt.Fprintf(w, "Error trying to find an account with the taken account_id '%s'. Error: %s", params["id"], err.Error())

			return
		}

		if account.Id == idToFind {

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(struct{ Balance float64 }{account.Balance})

			return
		}
	}
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintf(w, "Account not found")

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	newAccount, err := NewAccountFromJson(json.NewDecoder(r.Body))

	if err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Error to create new Account: %s", err.Error())

		return
	}

	if err := db.SaveAccount(Accounts{*newAccount}...); err != nil {

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintf(w, "Error to create the new account: %s", err.Error())

		return
	}

	w.WriteHeader(http.StatusOK)

	fmt.Printf("\nNew account created succesfully\n")

	fmt.Fprintf(w, "New account created succesfully")

}

func GetTransfers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	// TODO: mudar isso depois ara peagr o id da conta logada por mei od token
	loggedAccount := 1

	transfers := db.FindTransfers(loggedAccount)

	if err := json.NewEncoder(w).Encode(transfers); err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "error to encode transfers list to return to the API caller: %s", err.Error())

		return
	}

}

func MakeTransfer(w http.ResponseWriter, r *http.Request) {

	transfer, err := NewTransferFromJson(json.NewDecoder(r.Body))

	if err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Error to validate the Transfer data: %s", err.Error())

		return
	}

	if err := transfer.MakeTransfer(); err != nil {

		w.WriteHeader(http.StatusNotAcceptable)

		fmt.Fprintf(w, "Error to perform the Transfer: %s", err.Error())

		return
	}

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Transfer performed succesfully")
}