package main

import (
	"strconv"
	"time"

	logger "github.com/cihub/seelog"
)

var logConfig = `
<seelog minlevel="trace">
  <outputs formatid="common">

    <!--用来记录gin的log-->
    <filter levels="trace">
      <console formatid="access" />
      <rollingfile formatid="access" type="size" filename="logs/access_log.log" maxsize="102400000" maxrolls="7"/>
    </filter>

    <filter levels="info,debug,warn,error,critical">
      <console />
      <!--在golang中，2006代表年份，01代表月份，02代表日，因为golang是2016-01-02诞生的-->
      <rollingfile type="date" filename="logs/slog.log" datepattern="2006.01.02" maxrolls="7"/>
    </filter>
    <filter levels="error,critical">
      <!--maxsize单位是字节，102400大概是100k大小-->
      <rollingfile type="size" filename="logs/slog.error" maxsize="102400000" maxrolls="7"/>
    </filter>

  </outputs>
  <formats>
    <format id="common" format="%Date(2006-01-02 15:04:05.000) [%LEV][%File:%Line][%Func][%TID] %Msg%n" />
    <format id="access" format="[%TID] %Msg" />
  </formats>
</seelog>
`

func GetLogTraceId(param string) logger.FormatterFunc {
	return func(message string, level logger.LogLevel, context logger.LogContextInterface) interface{} {
		now := time.Now()
		return strconv.FormatInt(now.UnixNano(), 10) + "TID"
	}
}

func main() {
	// 定制化formatter, 可以在format使用%TID
	if err := logger.RegisterCustomFormatter("TID", GetLogTraceId); err != nil {
		panic(err)
	}

	log, err := logger.LoggerFromConfigAsString(logConfig)

	if err != nil {
		panic(err)
	}

	// 整个进程的logger都是用log了
	logger.ReplaceLogger(log)
	// 不要忘记flush
	defer logger.Flush()

	logger.Info("hello info")
	logger.Trace("hello trace")
	logger.Error("hello error")
}
