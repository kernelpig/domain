
var currentStatus = 0;  // 默认为初始状态

function toggleAction() {
    $(".DomainListStart").removeClass("am-btn-primary am-btn-default");
    $(".DomainListStop").removeClass("am-btn-primary am-btn-default");
    $(".DomainListReset").removeClass("am-btn-primary am-btn-default");
    if (currentStatus === 0) {  // 初始状态
        $(".DomainListStart").addClass("am-btn-primary");
        $(".DomainListStop").addClass("am-btn-default");
        $(".DomainListReset").addClass("am-btn-primary");
    } else if (currentStatus === 1) {   // 已开启状态
        $(".DomainListStart").addClass("am-btn-default");
        $(".DomainListStop").addClass("am-btn-primary");
        $(".DomainListReset").addClass("am-btn-default");
    } else {    // 已关闭状态
        $(".DomainListStart").addClass("am-btn-primary");
        $(".DomainListStop").addClass("am-btn-default");
        $(".DomainListReset").addClass("am-btn-primary");
    }
    $(".am-btn-default").attr("disabled", true);
    $(".am-btn-primary").attr("disabled", false);
}

function DomainListResetHandler() {
    $(".DomainListItems").val("");
}

function DomainListStopHandler() {
    toggleAction();
    APIRegStop(AlertShowAjaxError, function (data) {
        if (data["code"] === 0) {
            if (!data["status"]) {
                return
            }
            currentStatus = data["status"];
            toggleAction();
            return;
        } else {
            AlertShowError(data['sub_error']);
        }
    });
}

function DomainListStartHandler() {
    var req = {
        app_key_id: $(".DomainListAppIdKey").val(),
        app_key_secret: $(".DomainListAppSecret").val(),
        template_id: $(".DomainListTemplate").val(),
        domain_list: $(".DomainListItems").val(),
        thread_count: 4
    };
    APIRegStart(req, AlertShowAjaxError, function (data) {
        if (data["code"] === 0) {
            if (!data["status"]) {
                return
            }
            currentStatus = data["status"];
            toggleAction();
            return;
        } else {
            AlertShowError(data['sub_error']);
        }
    });
}

function regPageRender() {
    APIRegGet(AlertShowAjaxError, function (data) {
        if (data["code"] === 0) {
            if (!data["list"] || data["list"].length === 0) {
                return
            }
            var item = data["list"][0];
            $(".DomainListAppIdKey").val(item.app_key_id);
            $(".DomainListAppSecret").val(item.app_key_secret);
            $(".DomainListTemplate").val(item.template_id);
            $(".DomainListItems").val(item.domain_list);
        } else {
            AlertShowError(data['sub_error']);
        }
    });
}

function RegPageRender() {
    $(".DomainListStart").click(DomainListStartHandler);
    $(".DomainListStop").click(DomainListStopHandler);
    $(".DomainListReset").click(DomainListResetHandler);
    regPageRender();
}

$(document).ready(function () {
    NavbarRender();
    FootbarRender();
    RegPageRender();
    toggleAction();
});