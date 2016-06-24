// PASWD = mailrobot
// ___________________________________________________________
// |  SPARKPOST KEY: 672c40cdb9bb75b6ccc81a9a080624877b516ca3 |
// ------------------------------------------------------------

package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

var client = &http.Client{}

var apiKey string

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

func init() {
	apiKey = os.Getenv("CC_SPARKPOSTKEY")
	// 	apiKey = "672c40cdb9bb75b6ccc81a9a080624877b516ca3"
}

func templateMail(template, address string, data map[string]string) (result []byte, err error) {
	recipients := []recipientAPI{
		recipientAPI{
			Address: address,
		},
	}
	request := transmissionApI{
		Recipients: recipients,
		Content: contentAPI{
			TemplateID: template,
		},
		SubstitutionData: data,
	}
	result, err = json.Marshal(request)
	return
}

func SendMail(template, address string, data map[string]string) (err error) {
	result, err := templateMail(template, address, data)
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
	req.Header.Add("Authorization", apiKey)

	_, err = client.Do(req)
	if err != nil {
		return
	}
	return
}
