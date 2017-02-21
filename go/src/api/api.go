package main

import (
	"net/http"
	"github.com/labstack/echo"
	"gopkg.in/rana/ora.v3"
	"github.com/joho/godotenv"
  "log"
)

var DB_HOST string
var DB_PORT string
var DB_DATABASE string
var DB_USERNAME string
var DB_PASSWORD string

func conn() {
  var myEnv map[string]string
  myEnv, err := godotenv.Read()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  DB_HOST = myEnv["DB_HOST"]
  DB_PORT = myEnv["DB_PORT"]
  DB_DATABASE = myEnv["DB_DATABASE"]
  DB_USERNAME = myEnv["DB_USERNAME"]
  DB_PASSWORD = myEnv["DB_PASSWORD"]
}

func main() {
	e := echo.New()
	e.GET("/", hresh1x)
	e.Logger.Fatal(e.StartTLS(":443", "cert.pem", "key.pem"))
}

func hresh1x(c echo.Context) error {
  conn()
	var tns string = DB_USERNAME + "/" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_DATABASE
  env, srv, ses, err := ora.NewEnvSrvSes(tns, nil)
  if err != nil {
    return c.String(http.StatusServiceUnavailable, "Error connecting database")
  }
  defer env.Close()
  defer srv.Close()
  defer ses.Close()

  var resp string
  if _, err = ses.PrepAndExe("BEGIN :1 := WEB_SERVICE_HRESH1X.GEN_DATA('{\"p_coduser\": \"tjs00001\", \"p_codpswd\": \"506020187061205630914059305510426070801150087061706060721070\", \"p_lrunning\": \"136\", \"login_status\": \"1\", \"p_lang\": \"102\", \"p_page\": \"1\", \"p_limit\": \"10\", \"p_codempid\": \"53100\", \"p_dtest\": \"01012005\", \"p_dteen\": \"12122016\"}'); END;", &resp); err != nil {
    return c.String(http.StatusOK, "Error call store function")
  }
  return c.String(http.StatusOK, resp)
}