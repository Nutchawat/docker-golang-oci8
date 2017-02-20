package main

import (
	"net/http"
	"github.com/labstack/echo"
	"gopkg.in/rana/ora.v3"
)

func main() {
	e := echo.New()
	e.GET("/", conn)
	e.Logger.Fatal(e.Start(":1323"))
}

func conn(c echo.Context) error {
    env, srv, ses, err := ora.NewEnvSrvSes("STA3/STA3@192.168.2.10:1522/HR12C", nil)
    if err != nil {
        return c.String(http.StatusOK, "error1")
    }
    defer env.Close()
    defer srv.Close()
    defer ses.Close()

    var resp string
    if _, err = ses.PrepAndExe("BEGIN :1 := WEB_SERVICE_HRESH1X.GEN_DATA('{\"p_coduser\": \"tjs00001\", \"p_codpswd\": \"506020187061205630914059305510426070801150087061706060721070\", \"p_lrunning\": \"136\", \"login_status\": \"1\", \"p_lang\": \"102\", \"p_page\": \"1\", \"p_limit\": \"10\", \"p_codempid\": \"53100\", \"p_dtest\": \"01012005\", \"p_dteen\": \"12122016\"}'); END;", &resp); err != nil {
        return c.String(http.StatusOK, "error2")
    }
    return c.String(http.StatusOK, resp)
}