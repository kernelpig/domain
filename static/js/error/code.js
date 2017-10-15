var serverSubErrorMask = 0x0000ff00;
var SubModuleTokenErr = 0x06;

function IsTokenErr(error) {
    return ((error & serverSubErrorMask) >> 0x08) === SubModuleTokenErr;
}