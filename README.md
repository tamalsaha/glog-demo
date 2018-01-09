# glog-demo

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
