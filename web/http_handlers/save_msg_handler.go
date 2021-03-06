package http_handlers

import (
	"github.com/rkbalgi/isosim/data"
	"github.com/rkbalgi/isosim/web/spec"
	"log"
	"net/http"
	"strconv"
)

func saveMsgHandler() {

	http.HandleFunc(SaveMsgUrl, func(rw http.ResponseWriter, req *http.Request) {

		log.Print("Handling - " + SaveMsgUrl)

		err := req.ParseForm()
		if err != nil {

			sendError(rw, err.Error())
			return
		}

		log.Print(req.PostForm)
		//log.Print("?" + req.PostForm.Get("specId") + "?")
		///log.Print(req.PostForm.Get("msgId"))
		//log.Print(strconv.Atoi(req.PostForm.Get("specId")))
		//log.Print(req.PostForm.Get("msg"))

		if specId, err := strconv.Atoi(req.PostForm.Get("specId")); err == nil {
			log.Print("Spec Id =" + strconv.Itoa(specId))
			isoSpec := spec.GetSpec(specId)
			if isoSpec == nil {
				sendError(rw, InvalidSpecIdError.Error())
				return
			}
			log.Print("Spec = " + isoSpec.Name)
			if msgId, err := strconv.Atoi(req.PostForm.Get("msgId")); err == nil {
				msg := isoSpec.GetMessageById(msgId)
				if msg == nil {
					sendError(rw, InvalidMsgIdError.Error())
					return
				}
				//log.Print("Spec Msg = " + msg.Name)

				if req.Form.Get("updateMsg") == "true" {
					err = data.DataSetManager().Update(req.PostForm.Get("specId"),
						req.PostForm.Get("msgId"),
						req.PostForm.Get("dataSetName"), req.PostForm.Get("msg"))
				} else {

					err = data.DataSetManager().Add(req.PostForm.Get("specId"),
						req.PostForm.Get("msgId"),
						req.PostForm.Get("dataSetName"), req.PostForm.Get("msg"))
				}
				if err != nil {
					if err == data.DataSetExistsError {
						sendError(rw, "Data set exists. Please choose a different name.")
						return
					}

					sendError(rw, "Failed to add data set. Error ="+err.Error())
					return

				}

			} else {
				sendError(rw, InvalidMsgIdError.Error())
				return
			}

		} else {
			sendError(rw, InvalidSpecIdError.Error())
			return
		}

	})

}
