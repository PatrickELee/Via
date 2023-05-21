package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PatrickELee/Via/internal/grpc"
)

type GetDangerProbabilityJsonReq struct {
	Location string `json:"location"`
	Time     string `json:"time"`
}

type GetDangerProbabilityJsonResp struct {
	Danger string `json:"danger"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func DangerProbHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "POST":
		enableCors(&w)

		log.Println("Get Danger Probabilty Request")

		d := json.NewDecoder((r.Body))

		dangerProbReq := &GetDangerProbabilityJsonReq{}

		err := d.Decode(dangerProbReq)
		if err != nil {
			fmt.Printf("Error occured, %v", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		danger := grpc.MakeDangerProbRPCRequest(dangerProbReq.Location, dangerProbReq.Time)

		dangerProbResp := &GetDangerProbabilityJsonResp{
			Danger: danger,
		}
		resp, _ := json.Marshal(dangerProbResp)
		w.Write(resp)

	default:
		fmt.Println("Error with request method")
	}
}
