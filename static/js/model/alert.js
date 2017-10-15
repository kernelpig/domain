var AlertTemplate = '<div class="am-modal am-modal-no-btn" tabindex="-1" id="AlertShowContainer">\n' +
    '        <div class="am-modal-dialog">\n' +
    '            <div class="am-modal-hd AlertTitleContainer">\n' +
    '                <a href="javascript: void(0)" class="am-close am-close-spin" data-am-modal-close>&times;</a>\n' +
    '            </div>\n' +
    '            <div class="am-modal-bd AlertContentContainer">\n' +
    '\n' +
    '            </div>\n' +
    '        </div>\n' +
    '    </div>';


// 显示提示信息
function AlertShow(title, content, interval) {
    if ($('#AlertShowContainer').length === 0) {
        $('body').append(AlertTemplate);
    }
    $('.AlertTitleContainer').text(title);
    $('.AlertContentContainer').html(parseSpecialChar(content));
    $('#AlertShowContainer').modal();
    if (interval !== 0) {
        setTimeout(function () {
            $('#AlertShowContainer').modal('close');
        }, interval);
    }
}

// 显示提示信息并自动关闭
function AlertShowAutoClose(title, content) {
    AlertShow(title, content, 1000);
}

// 显示提示信息不自动关闭
function AlertShowNoAutoClose(title, content) {
    AlertShow(title, content, 0);
}

// 显示错误信息
function AlertShowError(error) {
    AlertShowAutoClose("请求失败", error);
}

// 显示提示信息并自动关闭
function AlertShowAutoCloseAndGoPage(title, content, gotoPage) {
    AlertShowNoAutoClose(title, content);
    setTimeout(function () {
        GoToPage(gotoPage);
    }, 3000)
}

// 显示ajax请求后端接口错误
function AlertShowAjaxError(e) {
    if (e.responseJSON) {
        // TODO: 如果是token expired则提示需要重新登录
        if(IsTokenErr(e.responseJSON.code)) {
            resetLoginCookie();
            AlertShowAutoCloseAndGoPage("请重新登录", e.responseJSON.sub_error, "/login.html");
        } else {
            AlertShowError(e.responseJSON.sub_error);
            //带有验证码的页面需要刷新验证码
            var $captcha = $.find("#CaptchaGetImageHandler");
            if ($captcha && $captcha.length !== 0) {
                CaptchaGetImageHandler();
            }
        }
    } else if (e.responseText) {
        AlertShowError(e.responseText);
    } else {
        AlertShowError(JSON.stringify(e));
    }
}