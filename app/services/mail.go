// PASWD = mailrobot
// ___________________________________________________________
// |  SPARKPOST KEY: 672c40cdb9bb75b6ccc81a9a080624877b516ca3 |
// ------------------------------------------------------------

package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var client = &http.Client{}

type contentAPI struct {
	TemplateID string `json:"template_id"`
}

type recipientAPI struct {
	Address string `json:"address"`
}

var substitution map[string]string

type transmissionApI struct {
	Recipients       []recipientAPI    `json:"recipients"`
	Content          contentAPI        `json:"content"`
	SubstitutionData map[string]string `json:"substitution_data"`
}

func betaMail(mail, name string) (result []byte, err error) {
	data := make(map[string]string)
	recipients := []recipientAPI{
		recipientAPI{
			Address: mail,
		},
	}
	data["name"] = name
	request := transmissionApI{
		Recipients: recipients,
		Content: contentAPI{
			TemplateID: "new-user-beta",
		},
		SubstitutionData: data,
	}
	result, err = json.Marshal(request)
	return
}

func SendMail(mail, name string) (err error) {
	result, err := betaMail(mail, name)
	if err != nil {
		return
	}
	// TODO - Revisar cuando la conexion con sparkpost falla
	req, err := http.NewRequest("POST", "https://api.sparkpost.com/api/v1/transmissions?num_rcpt_errors=3", bytes.NewReader(result))
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "672c40cdb9bb75b6ccc81a9a080624877b516ca3")

	_, err = client.Do(req)
	if err != nil {
		return
	}
	return
}
