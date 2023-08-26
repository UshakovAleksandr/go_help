package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	reqID, err := GetReqIDFromCtx(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(reqID)

	w.Header().Add("msgId", reqID)
	w.Write([]byte(fmt.Sprintf("Hello, world, msgID:%s ", reqID)))
}

func GetReqIDFromCtx(r *http.Request) (string, error) {
	if m := r.Context().Value("msgId"); m != nil {
		if value, ok := m.(string); ok {
			return value, nil
		}
	}

	return "", errors.New("failed request_id")
}

func InjectReqID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msgID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "msgId", msgID)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func main() {
	//http.HandleFunc("/welcome", HelloWorld)
	http.Handle("/welcome", InjectReqID(http.HandlerFunc(HelloWorld)))
	http.ListenAndServe(":8080", nil)
}

//// create uuid
//func pseudo_uuid() (uuid string) {
//	b := make([]byte, 16)
//	_, err := rand.Read(b)
//	if err != nil {
//		fmt.Println("Error: ", err)
//		return
//	}
//	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
//	return
//}
