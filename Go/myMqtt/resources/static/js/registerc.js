function userReg() {
    let xmlhttp;
    let mail = document.getElementById("zcMailbox").value;
    let passwordId = document.getElementById("zcPassword").value;
    let rpasswordId = document.getElementById("repassword").value;
    if (!mail) {
        document.getElementById("myDiv").innerHTML = "邮箱格式不正确";
    } else if (!passwordId) {
        document.getElementById("myDiv").innerHTML = "密码不能为空";
    } else if (!rpasswordId) {
        document.getElementById("myDiv").innerHTML = "确认密码不能为空";
    } else if (passwordId !== passwordId) {
        document.getElementById("myDiv").innerHTML = "两次密码不相同";
    } else {
        if (window.XMLHttpRequest) {
            xmlhttp = new XMLHttpRequest();
        }
        else {
            xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
        }
        xmlhttp.onreadystatechange = function () {
            if (xmlhttp.readyState === 4 && xmlhttp.status === 200) {
                document.getElementById("myDiv").innerHTML = xmlhttp.responseText;
            }
            // $('#regiter').attr("disabled",true);
        };
        xmlhttp.open("POST", "http://localhost:8080/registered", true);
        xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        xmlhttp.send("mailbox=" + mail + "&password=" + passwordId + "&repassword=" + rpasswordId);
    }
}