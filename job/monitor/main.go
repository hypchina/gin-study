package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"net/url"
	"strings"
	"encoding/json"
	"encoding/base64"
	"strconv"
	"crypto/md5"
)

var logFile = "./printer.log"

type Printer struct {
	printerName string
	printerUri  string
	token       string
	clientId    string
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func main() {
	logInfo("启动打印队列")
	Init()
	printer := Printer{
		printerName: Get("printer_name"),
		printerUri:  Get("printer_uri"),
		token:       Get("token"),
		clientId:    Get("client_id"),
	}
	zplPrint(printer)
}

func zplPrint(printer Printer) {
	deviceName, err := getDeviceName(printer.printerName)
	if err != nil {
		logError(err)
		return
	}
	timer := time.NewTicker(time.Second * 1)
	for range timer.C {
		runPrint(deviceName, printer)
	}
}

func runPrint(deviceName string, printer Printer) {
	respStr, err := getPrintTxt(printer)
	if err != nil {
		logError(err)
		return
	}

	lenRespStr := len(respStr)
	if respStr == "" || lenRespStr == 0 {
		printTxt("response format is error")
		return
	}

	var resp Response
	err = json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		logError(err)
		return
	}

	if resp.Code != 0 && resp.Msg != "ok" || resp.Data == "" {
		printTxt("空队列")
		return
	}

	printTxt, err := base64.StdEncoding.DecodeString(resp.Data)
	if err != nil {
		logError(err)
		return
	}

	has := md5.Sum(printTxt)
	md5str := fmt.Sprintf("%x", has)
	filename := os.TempDir() + "/" + md5str + ".zpl"
	logInfo("filename" + filename)
	err = ioutil.WriteFile(filename, printTxt, 755);
	if err != nil {
		logError(err)
		return
	}

	writeen, err := copyFile(filename, deviceName)
	if err != nil {
		logError(err)
		return
	}

	err = os.Remove(filename)
	if err != nil {
		logError(err)
		return
	}

	logInfo("writeen:" + strconv.FormatInt(writeen, 10))
}

func logInfo(msg string) {
	log(msg, "INFO")
}

func logError(err error) {
	log(err.Error(), "ERROR")
}

func printTxt(msg string) {
	date := GetDateByFormat()
	logStr := "DATE: " + date + ", LEVEL: INFO" + ", Msg: " + msg
	fmt.Println(logStr)
}

func log(msg string, level string) {

	date := GetDateByFormat()
	logStr := "DATE: " + date + ", LEVEL: " + level + ", Msg: " + msg + "\r\n"
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		printTxt("LOG: " + logStr + " ====== " + ", ERROR: " + err.Error())
		return
	}

	defer f.Close()
	_, err = f.WriteString(logStr)

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	printTxt("LOG: " + logStr + " ====== " + ", ERROR: " + errMsg)
}

func getDeviceName(printerName string) (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	device := "\\\\" + hostname + "\\" + printerName
	return device, nil
}

func getPrintTxt(printer Printer) (string, error) {
	param := url.Values{
		"token":     {printer.token},
		"client_id": {printer.clientId},
	}
	body := strings.NewReader(param.Encode())
	httpRequest, err := http.NewRequest("POST", printer.printerUri, body)
	if err != nil {
		return "", err
	}
	httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpClient := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return "", nil
	}
	return string(content), nil
}

func copyFile(dstName, srcName string) (writeen int64, err error) {

	if 1 == 1 {
		//return 0, nil
	}
	src, err := os.Open(dstName)
	if err != nil {
		return 0, err
	}

	defer src.Close()
	dst, err := os.OpenFile(srcName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}

	defer dst.Close()
	return io.Copy(dst, src)
}

const (
	StartDate       = "2006-01-02 15:04:05"
	StartDateWithMs = "2006-01-02 15:04:05.000"
)

func GetDateByFormat(timeParams ...time.Time) string {
	timeVars := time.Now()
	if len(timeParams) > 0 {
		timeVars = timeParams[0]
	}
	return timeVars.Format(StartDate)
}


var (
	flag       = false
	envContext = make(map[string]string)
)

const (
	envFn   = ".env"
	envLf   = "\n"
	envSq   = "="
	envNull = "null"
)

func Init() {
	if !flag {
		byteSet, err := ioutil.ReadFile(envFn)
		if err != nil {
			fmt.Println(err)
			return
		}
		contentSet := strings.Split(string(byteSet), envLf)
		for index := range contentSet {
			line := contentSet[index]
			if line == "" {
				continue
			}
			lineSet := strings.Split(line, envSq)
			if len(lineSet) < 2 {
				continue
			}
			envContext[lineSet[0]] = strings.Join(lineSet[1:], envSq)
		}
		flag = true
	}
}

func All() map[string]string {
	return envContext
}

func Get(key string, defaults ...string) string {
	if val, ok := envContext[key]; ok {
		return val
	}
	defaultVal := envNull
	if len(defaults) > 0 {
		defaultVal = defaults[0]
	}
	return defaultVal
}

func Set(key string, val string) {
	envContext[key] = val
}