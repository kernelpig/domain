
var NavbarTemplate = '<header class="am-topbar am-topbar-inverse">\n' +
    '    <h1 class="am-topbar-brand">\n' +
    '        <a href="index.html"><img src="../i/logo.png">' + WebSiteName + '</a>\n' +
    '    </h1>\n' +
    '\n' +
    '    <button class="am-topbar-btn am-topbar-toggle am-btn am-btn-sm am-btn-success am-show-sm-only" data-am-collapse="{target: \'#doc-topbar-collapse\'}"><span class="am-sr-only">导航切换</span> <span class="am-icon-bars"></span></button>\n' +
    '\n' +
    '    <div class="am-collapse am-topbar-collapse" id="doc-topbar-collapse">\n' +
    '        <ul class="am-nav am-nav-pills am-topbar-nav">\n' +
    '        </ul>\n' +
    '    </div>\n' +
    '</header>';

// 初始化其他菜单项
function NavbarItemsRender() {
    var items = [
        {
            name: "首页",
            id: "0",
            url: "index.html?column_id=0&page_size=10&page_num=1"
        },
        {
            name: "注册",
            id: "1",
            url: "reg.html?column_id=1&page_size=10&page_num=1"
        }
    ];
    $.each(items, function (index, item) {
        var a = $("<a></a>").attr("href", item.url).text(item.name);
        var li = $("<li></li>");
        if (item.id === NavbarPageEnv.column_id) {
            $(document).attr("title", WebSiteTitle + item.name);
            li.addClass("am-active");
        }
        $('.am-nav').append(li.append(a))
    })
}

// 页面变量信息
var NavbarPageEnv = {
    creater_uid: createrUidDefault,
    column_id: columnIdDefault,
    order_by: orderByDefault,
    page_size: PageSizeDefault,
    page_num: PageStartNumberDefault,
    is_end: false
};

// 菜单栏初始化项及状态
function NavbarInit() {
    $('body').prepend($(NavbarTemplate));
    NavbarItemsRender();
}

function NavbarGetCurrentEnv(currentUrl) {
    NavbarPageEnv.column_id = GetURIParamIdValue(currentUrl, "column_id") || columnIdDefault;
    NavbarPageEnv.order_by = GetURIParamStr(currentUrl, "order_by") || orderByDefault;
    NavbarPageEnv.page_size = GetURIParamInt(currentUrl, "page_size") || PageSizeDefault;
    NavbarPageEnv.page_num = GetURIParamInt(currentUrl, "page_num") || PageStartNumberDefault;
}

function NavbarRender() {
    NavbarGetCurrentEnv(location.href);
    NavbarInit();
}
