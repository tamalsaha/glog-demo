# glog-demo

## `glog` Internals

**Severity Levels**

ref: https://github.com/golang/glog/blob/23def4e6c14b4da8ac2ed8007337bc5eb5007998/glog.go#L100

```
"INFO"    : 0
"WARNING" : 1
"ERROR"   : 2
"FATAL"   : 3
```

**Flags**

ref: https://github.com/golang/glog/blob/23def4e6c14b4da8ac2ed8007337bc5eb5007998/glog.go#L398

```
flag.BoolVar(&logging.toStderr, "logtostderr", false, "log to standard error instead of files")
flag.BoolVar(&logging.alsoToStderr, "alsologtostderr", false, "log to standard error as well as files")
flag.Var(&logging.verbosity, "v", "log level for V logs")
flag.Var(&logging.stderrThreshold, "stderrthreshold", "logs at or above this threshold go to stderr")
```

**output**

ref: https://github.com/golang/glog/blob/23def4e6c14b4da8ac2ed8007337bc5eb5007998/glog.go#L671
```
	} else if l.toStderr {
		os.Stderr.Write(data)
	} else {
		if alsoToStderr || l.alsoToStderr || s >= l.stderrThreshold.get() {
			os.Stderr.Write(data)
		}
```

**V(x)**

ref:
- https://github.com/golang/glog/blob/23def4e6c14b4da8ac2ed8007337bc5eb5007998/glog.go#L1004
- https://github.com/golang/glog/blob/23def4e6c14b4da8ac2ed8007337bc5eb5007998/glog.go#L1040

```
	if logging.verbosity.get() >= level {
		return Verbose(true)
	}
```

**Notes:**

- V(x) are always output in `INFO` severity.
- If `--logtostderr` is true, severity levels are not consulted. (this is default true with appscode/go/log)
- `--stderrthreshold` is 2, ie, ERROR. So, glog.Error() logs are alway written.

To control logging using severity (Not V levels), you need to set:

```console
# only log ERROR / FATAL
glog-demo check --logtostderr=false --stderrthreshold=ERROR

# only log INFO at V=0 / WARN / ERROR / FATAL
glog-demo check --logtostderr=false --stderrthreshold=INFO

# only log INFO at V=0 / WARN / ERROR / FATAL
glog-demo check --logtostderr=true

# only log INFO at V=0,1,2 / WARN / ERROR / FATAL
glog-demo check --logtostderr=false --stderrthreshold=INFO --v=2

```

| func call               | severity, V |
|-------------------------|-------------|
| appscode/go/log.Debug   | INFO, V(4)  |
| appscode/go/log.Info    | INFO, V(3)  |
| appscode/go/log.Warning | INFO, V(2)  |
| appscode/go/log.Error   | INFO, V(1)  |
| appscode/go/log.Fatal   | INFO, V(0)  |
| glog.Info               | INFO, V(0)  |
| glog.Warning            | WARN        |
| glog.Error              | ERROR       |
| glog.Fatal              | FATAL       |


## Demo Projects

- https://github.com/tamalsaha/std-log-demo
- https://github.com/tamalsaha/glog-demo
- https://github.com/tamalsaha/kube-log-demo
- https://github.com/tamalsaha/ac-log-demo

```console

$ glide up -v
$ go install -v

$ glog-demo check -h
Check restic backup

Usage:
   check [flags]

Flags:
  -h, --help   help for check

Global Flags:
      --alsologtostderr                  log to standard error as well as files
      --analytics                        Send analytical events to Google Analytics (default true)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging

# default settings
$ glog-demo check --stderrthreshold=ERROR
2018/01/09 07:58:14 FLAG: --alsologtostderr="false"
2018/01/09 07:58:14 FLAG: --analytics="true"
2018/01/09 07:58:14 FLAG: --help="false"
2018/01/09 07:58:14 FLAG: --log_backtrace_at=":0"
2018/01/09 07:58:14 FLAG: --log_dir=""
2018/01/09 07:58:14 FLAG: --logtostderr="false"
2018/01/09 07:58:14 FLAG: --stderrthreshold="2"
2018/01/09 07:58:14 FLAG: --v="0"
2018/01/09 07:58:14 FLAG: --vmodule=""
log.Println_____
2018/01/09 07:58:14 node.Name
glog.Infoln_____
glog.Warningln_____
glog.Errorln_____
E0109 07:58:14.796966   22837 main.go:53] node.Name
glog.V(0).Infoln_____
glog.V(1).Infoln_____
glog.V(2).Infoln_____
glog.V(3).Infoln_____
glog.V(4).Infoln_____

$ glog-demo check --stderrthreshold=INFO
2018/01/09 07:59:19 FLAG: --alsologtostderr="false"
2018/01/09 07:59:19 FLAG: --analytics="true"
2018/01/09 07:59:19 FLAG: --help="false"
2018/01/09 07:59:19 FLAG: --log_backtrace_at=":0"
2018/01/09 07:59:19 FLAG: --log_dir=""
2018/01/09 07:59:19 FLAG: --logtostderr="false"
2018/01/09 07:59:19 FLAG: --stderrthreshold="0"
2018/01/09 07:59:19 FLAG: --v="0"
2018/01/09 07:59:19 FLAG: --vmodule=""
log.Println_____
2018/01/09 07:59:19 node.Name
glog.Infoln_____
I0109 07:59:19.566625   23000 main.go:47] node.Name
glog.Warningln_____
W0109 07:59:19.566811   23000 main.go:50] node.Name
glog.Errorln_____
E0109 07:59:19.566928   23000 main.go:53] node.Name
glog.V(0).Infoln_____
I0109 07:59:19.567033   23000 main.go:58] node.Name
glog.V(1).Infoln_____
glog.V(2).Infoln_____
glog.V(3).Infoln_____
glog.V(4).Infoln_____

$ glog-demo check --stderrthreshold=INFO --v=2
2018/01/09 07:59:45 FLAG: --alsologtostderr="false"
2018/01/09 07:59:45 FLAG: --analytics="true"
2018/01/09 07:59:45 FLAG: --help="false"
2018/01/09 07:59:45 FLAG: --log_backtrace_at=":0"
2018/01/09 07:59:45 FLAG: --log_dir=""
2018/01/09 07:59:45 FLAG: --logtostderr="false"
2018/01/09 07:59:45 FLAG: --stderrthreshold="0"
2018/01/09 07:59:45 FLAG: --v="2"
2018/01/09 07:59:45 FLAG: --vmodule=""
log.Println_____
2018/01/09 07:59:45 node.Name
glog.Infoln_____
I0109 07:59:45.421904   23052 main.go:47] node.Name
glog.Warningln_____
W0109 07:59:45.422051   23052 main.go:50] node.Name
glog.Errorln_____
E0109 07:59:45.422123   23052 main.go:53] node.Name
glog.V(0).Infoln_____
I0109 07:59:45.422198   23052 main.go:58] node.Name
glog.V(1).Infoln_____
I0109 07:59:45.422211   23052 main.go:61] node.Name
glog.V(2).Infoln_____
I0109 07:59:45.422220   23052 main.go:64] node.Name
glog.V(3).Infoln_____
glog.V(4).Infoln_____

$ glog-demo check --logtostderr=true --v=-1
2018/01/08 19:48:53 FLAG: --alsologtostderr="false"
2018/01/08 19:48:53 FLAG: --analytics="true"
2018/01/08 19:48:53 FLAG: --help="false"
2018/01/08 19:48:53 FLAG: --log_backtrace_at=":0"
2018/01/08 19:48:53 FLAG: --log_dir=""
2018/01/08 19:48:53 FLAG: --logtostderr="true"
2018/01/08 19:48:53 FLAG: --stderrthreshold="2"
2018/01/08 19:48:53 FLAG: --v="-1"
2018/01/08 19:48:53 FLAG: --vmodule=""
log.Println_____
2018/01/08 19:48:53 node.Name
glog.Infoln_____
I0108 19:48:53.577603   29286 main.go:47] node.Name
glog.Warningln_____
W0108 19:48:53.577620   29286 main.go:50] node.Name
glog.Errorln_____
E0108 19:48:53.577631   29286 main.go:53] node.Name
glog.V(0).Infoln_____
glog.V(1).Infoln_____
glog.V(2).Infoln_____
glog.V(3).Infoln_____
glog.V(4).Infoln_____
```


## klog flags

  -add_dir_header
    	If true, adds the file directory to the header of the log messages
  -alsologtostderr
    	log to standard error as well as files
  -log_backtrace_at value
    	when logging hits line file:N, emit a stack trace
  -log_dir string
    	If non-empty, write log files in this directory
  -log_file string
    	If non-empty, use this log file
  -log_file_max_size uint
    	Defines the maximum size a log file can grow to. Unit is megabytes. If the value is 0, the maximum file size is unlimited. (default 1800)
  -logtostderr
    	log to standard error instead of files (default true)
  -one_output
    	If true, only write logs to their native severity level (vs also writing to each lower severity level)
  -skip_headers
    	If true, avoid header prefixes in the log messages
  -skip_log_headers
    	If true, avoid headers when opening log files
  -stderrthreshold value
    	logs at or above this threshold go to stderr (default 2)
  -v value
    	number for the log level verbosity
  -vmodule value
    	comma-separated list of pattern=N settings for file-filtered logging


## glog flags

  -alsologtostderr
      log to standard error as well as files
  -log_backtrace_at value
      when logging hits line file:N, emit a stack trace
  -log_dir string
      If non-empty, write log files in this directory
  -logtostderr
      log to standard error instead of files
  -stderrthreshold value
      logs at or above this threshold go to stderr
  -v value
      log level for V logs
  -vmodule value
      comma-separated list of pattern=N settings for file-filtered logging
