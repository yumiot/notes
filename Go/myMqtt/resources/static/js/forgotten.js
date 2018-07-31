function forgotten(loginmail) {
    var xmlhttp;
    //var loginmail;
    var loginpassword;
    var repassword;
    var usertishi;

    if (window.XMLHttpRequest) {
        xmlhttp = new XMLHttpRequest();
    }
    else {
        xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
    }
    xmlhttp.onreadystatechange = function () {
        if (xmlhttp.readyState === 4 && xmlhttp.status === 200) {
            usertishi = xmlhttp.responseText;
            if (usertishi === "修改成功") {
                window.location = "/m/login";
            }
            else {
                document.getElementById("myDiv1").innerHTML = usertishi;
            }
        }
    };

    loginpassword = document.getElementById("password").value;
    repassword = document.getElementById("repassword").value;

    xmlhttp.open("POST", "http://localhost:8080/u/xgpass", true);
    xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xmlhttp.send("mailbox=" + loginmail + "&password=" + loginpassword + "&repassword=" + repassword);
}
