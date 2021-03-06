package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
)

func TestGetMiniProgram(t *testing.T) {
	log.Println(wego.GetMiniProgram().AppCode().AccessToken().GetToken())
	//2018/02/25 16:47:09 map[access_token:7_H7C27Kt3UeVSrPr75bRvjDKzWTwIrMCwzdpB5oI9ClclterJ6xlnhnc8r9GdBEjPZvnbRBbmfQZMefPVnQqnwrs9_-3fq8k7uAPHJ2AlwpVvT8Oh3SxDfA0C-xnXG06--Mgo9OxllzHYnh-5ZTGiCIAFZM expires_in:7200]

}

func TestGetAuth(t *testing.T) {
	log.Println(wego.GetMiniProgram().Auth().Session("1234"))
}

func TestGetAppCode(t *testing.T) {
	log.Println(wego.GetAppCode().Get("path", nil))
}

func TestNewAppCode(t *testing.T) {
	var v []string
	log.Println(v == nil)
}
