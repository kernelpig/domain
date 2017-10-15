var FootbarTemplate = '<footer data-am-widget="footer"\n' +
    '        class="am-footer am-footer-default am-footer am-topbar-fixed-bottom">\n' +
    '    <hr class="am-divider am-divider-default" />\n' +
    '    <div class="am-footer-miscs">\n' +
    '        <span class="Copyright"></span>\n' +
    '        <span class="RecordNumber"></span>\n' +
    '    </div>\n' +
    '</footer>';

var FootbarEnv = {
    copyright: "CopyRight©2017 <a href='/about.html'>王大脸</a> Inc.",
    record_number: "京ICP备14040736号"
};

function FootbarRender() {
    var footbar = $(FootbarTemplate);
    footbar.find(".Copyright").html(FootbarEnv.copyright);
    footbar.find(".RecordNumber").text(FootbarEnv.record_number);
    $('body').append(footbar);
}