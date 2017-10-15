package error

var ServiceErrs [_ServiceErrMax]string
var InterfaceErrs [_InterfaceErrMax]string
var SubModuleErrs [_SubModuleErrMax]string
var SubErrors [_SubModuleErrMax][]string

func init() {
	initServiceErrs()
	initInterfaceErr()
	initSubModuleErrs()
	initSubErrors()
}

func initServiceErrs() {
	ServiceErrs = [_ServiceErrMax]string{
		SDomain: "cunxun",
	}
}

func initSubErrors() {
	SubErrors = [_SubModuleErrMax][]string{
		_SubModuleErrMin: {
			_SubModuleErrMin: "Basic sub modules error information.",
		},
		MConfigErr: {
			_ConfigErrMin:       "Basic configuration error information.",
			ConfigLoadErr:       "Failed to load profile.",
			ConfigParseErr:      "Parsing configuration file failed.",
			ConfigParseTimeErr:  "Parse time configuration field failed.",
			ConfigLoadAvatarErr: "Load avatar image file failed.",
		},
		MLogErr: {
			_LogErrMin:        "Basic log error information.",
			LogDumpRequestErr: "Dump request parameter failed.",
		},
		MRegErr: {
			_RegErrMin:  "Basic reg error information.",
			RegStartErr: "Failed to obtain article information.",
		},
	}
}

func initSubModuleErrs() {
	SubModuleErrs = [_SubModuleErrMax]string{
		_SubModuleErrMin: "Invalid sub module", // 占位
		MConfigErr:       "config",             // 配置错误
		MLogErr:          "log",                // 日志错误
	}
}

func initInterfaceErr() {
	InterfaceErrs = [_InterfaceErrMax]string{
		_InterfaceErrMin: "Same with the top stack.",
		IRegStart:        "get /reg/start",
		IRegStop:         "get /reg/stop",
		IRegKill:         "get /reg/kill",
	}
}
