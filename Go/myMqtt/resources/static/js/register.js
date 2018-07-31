function changeSucceedStyle(nameID, spanID) {
    spanID.firstChild.nodeValue = " ";
    spanID.style.fontSize = "larger";
    spanID.style.color = "green";
    nameID.style.borderColor = "limegreen";
}

function changeFailedStyle(nameID, spanID) {
    spanID.firstChild.nodeValue = " ";
    spanID.style.fontSize = "larger";
    spanID.style.color = "red";
    nameID.style.borderColor = "red";
}

function changeFailingStyle(nameID, spanID) {
    spanID.style.fontSize = "small";
    spanID.style.color = "red";
    nameID.style.borderColor = "red";
}

function spanValue(spanID, spanValue) {
    switch (spanValue) {
        case "passwordSpan":
            spanID.firstChild.nodeValue = "6-16位：由数字、字母和特殊符号组成";
            break;
        case "repasswordSpan":
            spanID.firstChild.nodeValue = "请确认密码";
            break;
        case "repasswordSpan1":
            spanID.firstChild.nodeValue = "两次密码不一致";
            break;
        case "mailboxSpan":
            spanID.firstChild.nodeValue = "请输入有效邮箱";
            break;
        case "phoneSpan":
            spanID.firstChild.nodeValue = "请输入正确的手机号";
            break;
    }
}

var mailboxnum;

function mailbox() {
    var mailbox = document.getElementById("mailbox");
    var mailboxSpan = document.getElementById("mailboxSpan");
    var pattern = /^[0-9a-zA-Z_]{5,18}@[0-9a-z]+.com$/;
    mailbox.onfocus = function () {
        if (!pattern.test(mailbox.value)) {
            spanValue(mailboxSpan, "mailboxSpan");
            changeFailingStyle(mailbox, mailboxSpan);
        }
    };
    mailbox.onkeyup = function () {
        if (pattern.test(mailbox.value)) {
            changeSucceedStyle(mailbox, mailboxSpan);
        } else {
            spanValue(mailboxSpan, "mailboxSpan");
            changeFailingStyle(mailbox, mailboxSpan);
        }
    };
    mailbox.onblur = function () {
        if (pattern.test(mailbox.value)) {
            changeSucceedStyle(mailbox, mailboxSpan);
            mailboxnum = 1;
        } else {
            changeFailedStyle(mailbox, mailboxSpan);
            mailboxnum = 0;
        }
    }
}

let passwordnum;
let repasswordnum;

function password() {
    const password = document.getElementById("password");
    const passwordSpan = document.getElementById("passwordSpan");
    const pattern = /^(?=.*?[a-zA-Z])(?=(.*[\d]){1,})(?=(.*[\W]){1,})(?!.*\s).{6,}$/,
        str = '';
    console.log(pattern.test(str));
    const repassword = document.getElementById("repassword");
    password.onfocus = function () {
        if (!pattern.test(password.value)) {
            spanValue(passwordSpan, "passwordSpan");
            changeFailingStyle(password, passwordSpan);
        }
    };
    password.onkeyup = function () {
        if (pattern.test(password.value)) {
            changeSucceedStyle(password, passwordSpan);
            if (repassword.value !== "") {
                repassword.onfocus();
            }
        } else {
            spanValue(passwordSpan, "passwordSpan");
            changeFailingStyle(password, passwordSpan);
            if (repassword.value !== "") {
                repassword.onfocus();
            }
        }
    };

    password.onblur = function () {
        if (repassword.value === "") {
            if (pattern.test(password.value)) {
                changeSucceedStyle(password, passwordSpan);
                passwordnum = 1;
            } else {
                changeFailedStyle(password, passwordSpan);
                passwordnum = 0;
            }
        } else {
            if (password.value !== repassword.value) {
                repassword.onfocus();
                repasswordnum = 0;
            }
        }
    }
}

function repassword() {
    const password = document.getElementById("password");
    const repassword = document.getElementById("repassword");
    const repasswordSpan = document.getElementById("repasswordSpan");

    repassword.onfocus = function () {
        if (!(password.value === repassword.value && password.value !== "")) {
            spanValue(repasswordSpan, "repasswordSpan");
            changeFailingStyle(repassword, repasswordSpan);
        } else {
            changeSucceedStyle(repassword, repasswordSpan);
        }
    };
    repassword.onkeyup = function () {
        if (password.value === repassword.value && password.value !== "") {
            changeSucceedStyle(repassword, repasswordSpan);
        } else {
            spanValue(repasswordSpan, "repasswordSpan1");
            changeFailingStyle(repassword, repasswordSpan);
        }
    };
    repassword.onblur = function () {
        if (password.value === repassword.value && password.value !== "") {
            changeSucceedStyle(repassword, repasswordSpan);
            repasswordnum = 1;
        } else {
            spanValue(repasswordSpan, "repasswordSpan");
            changeFailedStyle(repassword, repasswordSpan);
            repasswordnum = 0;
        }
    }
}

passwordnum = 0;
repasswordnum = 0;
mailboxnum = 0;

window.onload = function () {
    password();
    repassword();
    mailbox();
    mailbox();
};
