// mdc ripple effect
const snackbar = mdc.snackbar.MDCSnackbar.attachTo(document.querySelector('.mdc-snackbar'));

const textFieldElements = [].slice.call(document.querySelectorAll('.mdc-text-field'));
textFieldElements.forEach((textFieldEl) => {
    mdc.textField.MDCTextField.attachTo(textFieldEl);
});

$('#regForm').submit(function (e) {
    e.preventDefault();

    // checking if any field is empty(only space) or not
    let flag = 1;
    if ($('#username').val().trim() == "") {
        $('#username').val("");
        notify("Please give a username!");
        flag = 0;
    }
    if (flag == 1 && ($('#password').val() != $('#confirmPassword').val())) {
        notify("Password didn't match!");
        flag = 0;
    }

    if (flag) {
        console.log($('#regForm').serialize());
        $('#btnRegister').prop('disabled', true);
        $('#btnRegister .mdc-button__label').text("Please wait...");

        //sending ajax post request
        let request = $.ajax({
            async: true,
            type: "POST",
            url: "/register",
            data: $('#regForm').serialize(),
        });

        request.done(function (response) {
            if (response.trim() == "Registration Done") {
                notify("Registration successful. Email verification link was sent to your provided email.");

                setTimeout(function () {
                    window.location.href = "/login";
                }, 2000);
            } else if (response.trim() == "username") {
                notify("Username already taken. Please choose a different one.");
            } else if (response.trim() == "email") {
                notify("Email already exsist. Please choose a different one.");
            } else {
                notify("Registration unsuccessful. Something went wrong. Please try again!");
            }
        });

        request.fail(function (response) {
            notify(response);
        });

        request.always(function () {
            $('#btnRegister').prop('disabled', false);
            $('#btnRegister .mdc-button__label').text("Register");
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