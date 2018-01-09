# glog-demo

## Demo Projects

- https://github.com/tamalsaha/std-log-demo
- https://github.com/tamalsaha/glog-demo
- https://github.com/tamalsaha/kube-log-demo
- https://github.com/tamalsaha/ac-log-demo

```console

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

$ glog-demo check
2018/01/08 18:06:39 FLAG: --alsologtostderr="false"
2018/01/08 18:06:39 FLAG: --analytics="true"
2018/01/08 18:06:39 FLAG: --help="false"
2018/01/08 18:06:39 FLAG: --log_backtrace_at=":0"
2018/01/08 18:06:39 FLAG: --log_dir=""
2018/01/08 18:06:39 FLAG: --logtostderr="false"
2018/01/08 18:06:39 FLAG: --stderrthreshold="2"
2018/01/08 18:06:39 FLAG: --v="0"
2018/01/08 18:06:39 FLAG: --vmodule=""
log.Println_____
2018/01/08 18:06:39 node.Name
glog.Infoln_____
glog.Warningln_____
glog.Errorln_____
E0108 18:06:39.191185   10907 main.go:53] node.Name
glog.V(0).Infoln_____
glog.V(1).Infoln_____
glog.V(2).Infoln_____
glog.V(3).Infoln_____
glog.V(4).Infoln_____

$ glog-demo check --v=2
2018/01/08 18:06:53 FLAG: --alsologtostderr="false"
2018/01/08 18:06:53 FLAG: --analytics="true"
2018/01/08 18:06:53 FLAG: --help="false"
2018/01/08 18:06:53 FLAG: --log_backtrace_at=":0"
2018/01/08 18:06:53 FLAG: --log_dir=""
2018/01/08 18:06:53 FLAG: --logtostderr="false"
2018/01/08 18:06:53 FLAG: --stderrthreshold="2"
2018/01/08 18:06:53 FLAG: --v="2"
2018/01/08 18:06:53 FLAG: --vmodule=""
log.Println_____
2018/01/08 18:06:53 node.Name
glog.Infoln_____
glog.Warningln_____
glog.Errorln_____
E0108 18:06:53.404893   10958 main.go:53] node.Name
glog.V(0).Infoln_____
glog.V(1).Infoln_____
glog.V(2).Infoln_____
glog.V(3).Infoln_____
glog.V(4).Infoln_____

$ glog-demo check --v=2 --logtostderr=true
2018/01/08 18:07:40 FLAG: --alsologtostderr="false"
2018/01/08 18:07:40 FLAG: --analytics="true"
2018/01/08 18:07:40 FLAG: --help="false"
2018/01/08 18:07:40 FLAG: --log_backtrace_at=":0"
2018/01/08 18:07:40 FLAG: --log_dir=""
2018/01/08 18:07:40 FLAG: --logtostderr="true"
2018/01/08 18:07:40 FLAG: --stderrthreshold="2"
2018/01/08 18:07:40 FLAG: --v="2"
2018/01/08 18:07:40 FLAG: --vmodule=""
log.Println_____
2018/01/08 18:07:40 node.Name
glog.Infoln_____
I0108 18:07:40.711602   11070 main.go:47] node.Name
glog.Warningln_____
W0108 18:07:40.711629   11070 main.go:50] node.Name
glog.Errorln_____
E0108 18:07:40.711645   11070 main.go:53] node.Name
glog.V(0).Infoln_____
I0108 18:07:40.711662   11070 main.go:58] node.Name
glog.V(1).Infoln_____
I0108 18:07:40.711678   11070 main.go:61] node.Name
glog.V(2).Infoln_____
I0108 18:07:40.711694   11070 main.go:64] node.Name
glog.V(3).Infoln_____
glog.V(4).Infoln_____

$ glog-demo check --logtostderr=true
2018/01/08 18:07:57 FLAG: --alsologtostderr="false"
2018/01/08 18:07:57 FLAG: --analytics="true"
2018/01/08 18:07:57 FLAG: --help="false"
2018/01/08 18:07:57 FLAG: --log_backtrace_at=":0"
2018/01/08 18:07:57 FLAG: --log_dir=""
2018/01/08 18:07:57 FLAG: --logtostderr="true"
2018/01/08 18:07:57 FLAG: --stderrthreshold="2"
2018/01/08 18:07:57 FLAG: --v="0"
2018/01/08 18:07:57 FLAG: --vmodule=""
log.Println_____
2018/01/08 18:07:57 node.Name
glog.Infoln_____
I0108 18:07:57.502661   11098 main.go:47] node.Name
glog.Warningln_____
W0108 18:07:57.502676   11098 main.go:50] node.Name
glog.Errorln_____
E0108 18:07:57.502684   11098 main.go:53] node.Name
glog.V(0).Infoln_____
I0108 18:07:57.502692   11098 main.go:58] node.Name
glog.V(1).Infoln_____
glog.V(2).Infoln_____
glog.V(3).Infoln_____
glog.V(4).Infoln_____

```
