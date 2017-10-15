
// 服务端各模块URI
// 基础url请使用相对路径, nginx或caddy对api目录做反响代理
var serviceBaseURI      = "http://localhost:8005/api";
var regBaseURI      = serviceBaseURI + "/reg";

function APIRegStart(data, error, success) {
    var url = regBaseURI + "/start";
    AjaxNoAuth(url, "post", data, error, success);
}

function APIRegStop(error, success) {
    var url = regBaseURI + "/stop";
    AjaxNoAuth(url, "post", null, error, success);
}

function APIRegGet(error, success) {
    var url = regBaseURI + "/";
    AjaxNoAuth(url, "get", null, error, success);
}