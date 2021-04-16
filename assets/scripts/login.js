// mdc ripple effect
mdc.textField.MDCTextField.attachTo(document.querySelector('.mdc-text-field-username'));
mdc.textField.MDCTextField.attachTo(document.querySelector('.mdc-text-field-password'));
mdc.ripple.MDCRipple.attachTo(document.querySelector('.cancel'));
mdc.ripple.MDCRipple.attachTo(document.querySelector('.next'));
const snackbar = mdc.snackbar.MDCSnackbar.attachTo(document.querySelector('.mdc-snackbar'));

$('#loginForm').submit(function (e) {
    e.preventDefault();

    // checking if any field is empty(only space) or not
    let flag = 1;
    if ($('#username').val().trim() == "") {
        $('#username').val("");
        notify("Please give a username!");
        flag = 0;
    }

    if (flag) {
        $('#btnLogin').prop('disabled', true);
        $('#btnLogin .mdc-button__label').text("Please wait...");

        //sending ajax post request
        let request = $.ajax({
            async: true,
            type: "POST",
            url: "/login",
            data: $('#loginForm').serialize(),
        });

        request.done(function (response) {
            if (response.trim() == "Login Done") {
                notify("Login successful!");
                window.location.href = "/dashboard";
            } else {
                //console.log(response)
                notify(response);
            }
        });

        request.fail(function (response) {
            notify(response);
        });

        request.always(function () {
            $('#btnLogin').prop('disabled', false);
            $('#btnLogin .mdc-button__label').text("Login");
        });
    }
    return false;
});

function notify(msg) {
    snackbar.timeoutMs = 5000;
    snackbar.labelText = msg;
    snackbar.actionButtonText = "OKAY";
    snackbar.open();
}