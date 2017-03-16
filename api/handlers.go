package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/merryChris/gameLord/types"
	"github.com/merryChris/gameLord/utils"
)

type Handler struct {
	Initialized bool
}

func (this *Handler) SelfCheck() (string, bool) {
	if !this.Initialized {
		log.Println("Uninitializing UserHandler Error.")
		return types.Error101(), false
	}
	return "", true
}

func HandlerWrapper(inner http.HandlerFunc, name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctime := time.Now()
		w.Header().Set("Content-Type", "application/json")
		inner(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(ctime),
		)
	}
}

func ParseRequestJsonData(r io.ReadCloser, reqObj interface{}) (string, bool) {
	d, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println(fmt.Sprintf("Reading Request Body Error: `%s`.", err.Error()))
		return types.Error101(), false
	}
	if err = json.Unmarshal(d, reqObj); err != nil {
		log.Println(fmt.Sprintf("Unmarshalling Request Body Error: `%s`." + err.Error()))
		return types.Error101(), false
	}
	return "", true
}

func CheckRequestJsonData(reqObj interface{}) (string, bool) {
	rot := reflect.TypeOf(reqObj)
	rov := reflect.ValueOf(reqObj)
	for i := 0; i < rot.NumField(); i++ {
		switch rov.Field(i).Kind() {
		case reflect.Struct:
			if resp, ok := CheckRequestJsonData(rov.Field(i).Interface()); !ok {
				return resp, ok
			}
		case reflect.Uint8:
			if rot.Field(i).Name == "CurrentLetter" && !utils.IsLetter(uint8(rov.Field(i).Uint())) {
				return types.Error103(), false
			}
		case reflect.Int64:
			if (rot.Field(i).Name == "UserId" || rot.Field(i).Name == "GameId" || rot.Field(i).Name == "CategoryId") &&
				rov.Field(i).Int() == 0 {
				return types.Error102(), false
			}
		case reflect.String:
			if "" == rov.Field(i).String() {
				return types.Error102(), false
			}
		}
	}
	return "", true
}
