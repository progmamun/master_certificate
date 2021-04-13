//js instantiate for mdc elements
const dialog = mdc.dialog.MDCDialog.attachTo(document.querySelector('.mdc-dialog'));
const dialogSearch = mdc.dialog.MDCDialog.attachTo(document.querySelector('#searchFormDialog'));
const dialogAdd = mdc.dialog.MDCDialog.attachTo(document.querySelector('#addFormDialog'));

const textFieldElements = [].slice.call(document.querySelectorAll('.mdc-text-field'));
textFieldElements.forEach((textFieldEl) => {
    mdc.textField.MDCTextField.attachTo(textFieldEl);
});

$('#btnFilter').on('click', function (e) {
    console.log("ok")
    dialogSearch.autoStackButtons = false;
    dialogSearch.open();
    dialogSearch.escapeKeyAction = "";
    dialogSearch.scrimClickAction = "";
});

$('#btnAdd').on('click', function (e) {
    console.log("ok")
    dialogAdd.autoStackButtons = false;
    dialogAdd.open();
    dialogAdd.escapeKeyAction = "";
    dialogAdd.scrimClickAction = "";

    $('#addForm h2').text('Student Add');
    $('#btnAddAction span').text("Add");
    $('#checkboxAdd').css('display', 'none');

    $('#usernameAdd').removeAttr("disabled", false);
    $('#emailAdd').prop("disabled", false);

    $('#my-label-id-password').text("Password");
    $('#passwordAdd').attr("name", "password");
    $('#passwordAdd').attr("type", "password");
    $('#passwordAdd').prop('required', true);

    $('#my-label-id-confirm-password').text("Confirm Password");
    $('#confirmPasswordAdd').attr("name", "confirmPassword");
    $('#confirmPasswordAdd').attr("type", "password");
    $('#confirmPasswordAdd').prop('required', true);

    resetForm();

    const textFieldElements = [].slice.call(document.querySelectorAll('.mdc-text-field'));
    textFieldElements.forEach((textFieldEl) => {
        mdc.textField.MDCTextField.attachTo(textFieldEl);
    });
});

//pagination
// const perPageMenu = mdc.menu.MDCMenu.attachTo(document.querySelector('#perPageMenu'));

// $('#btnRowPerPage').on('click', function (e) {
//     e.preventDefault();
//     console.log("btnRowPerPage is pressed!");
//     perPageMenu.open = !perPageMenu.open;
// });

// perPageMenu.listen("MDCMenuSurface:opened", (d) => {
//     $('#perPageMenu').removeClass("mdc-menu-surface--is-open-below");
//     $("#perPageMenu").css("transform-origin", "center bottom");
//     $("#perPageMenu").css("left", "0");
//     $("#perPageMenu").css("bottom", "0");
//     $("#perPageMenu").css("top", "");
//     $("#perPageMenu").css("min-width", "80px");
//     //$("#perPageMenu").css("max-height", "536.375px");
// });

//checkbox clicking
let chks = [];
$('.mdc-data-table__table > thead > tr').find('input[type=checkbox]').on('change', function () {
    let isChecked = $(this).is(':checked');

    $('.mdc-data-table__table > tbody  > tr').each(function (index, tr) {
        $(this).find('input[type=checkbox]').eq(0).prop('checked', isChecked);
        let chkval = $(this).find('input[type="checkbox"]').val();

        if (isChecked == true) {
            chks.push(chkval);
        } else {
            chks.pop(chkval);
        }

        if (chks.length > 0) {
            //$('#btnSForm').show();
        } else {
            //$('#btnSForm').hide();
        }
    });
    console.log(chks);
    console.log(chks.length);
});

$('.mdc-data-table__table > tbody  > tr').on("click", ".mdc-checkbox__native-control", function () {
    $curr = $(this).closest("td");

    let index = $(this).closest("tr").index();
    let chkval = $curr.find('input[type="checkbox"]').val();
    console.log(chkval);

    let isChecked = $curr.find('input[type="checkbox"]').is(':checked');
    if (isChecked == true) {
        chks.push(chkval);
    } else {
        chks.pop(chkval);
    }
    console.log(chks);
    if (chks.length > 0) {
        //$('#btnSForm').show();
    } else {
        //$('#btnSForm').hide();
    }
});

function resetForm() {
    $('#firstNameAdd').val("");
    $('#lastNameAdd').val("");
    $('#usernameAdd').val("");
    $('#emailAdd').val("");
    $('#mobileAdd').val("");
    $('#passwordAdd').val("");
    $('#confirmPasswordAdd').val("");
    $('#activeStatus').prop('checked', false);
}
function edit(tr) {
    //console.log(tr.parentElement.parentElement.cells);
    dialogAdd.autoStackButtons = false;
    dialogAdd.open();
    dialogAdd.escapeKeyAction = "";
    dialogAdd.scrimClickAction = "";

    $('#addForm h2').text('Student Update');
    $('#btnAddAction span').text("Update");
    $('#checkboxAdd').css('display', 'block');

    $('#usernameAdd').prop("disabled", true);
    $('#emailAdd').prop("disabled", true);

    $('#my-label-id-password').text("Mobile");
    $('#passwordAdd').attr("name", "mobile");
    $('#passwordAdd').attr("type", "text");
    $('#passwordAdd').removeAttr("required");

    $('#my-label-id-confirm-password').text("City");
    $('#confirmPasswordAdd').attr("name", "city");
    $('#confirmPasswordAdd').attr("type", "text");
    $('#confirmPasswordAdd').removeAttr("required");

    const textFieldElements = [].slice.call(document.querySelectorAll('.mdc-text-field'));
    textFieldElements.forEach((textFieldEl) => {
        mdc.textField.MDCTextField.attachTo(textFieldEl);
    });

    selectedRow = tr.parentElement.parentElement;
    let username = selectedRow.cells[2].innerText.trim();


    //retrieving data from backend
    //sending ajax post request
    let request = $.ajax({
        async: true,
        type: "GET",
        url: "/api/student_info/" + username,
    });
    request.done(function (response) {
        console.log(response["status"])
        $('#firstNameAdd').val(response["first_name"]);
        $('#lastNameAdd').val(response["last_name"]);
        $('#usernameAdd').val(response["username"]);
        $('#emailAdd').val(response["email"]);
        $('#passwordAdd').val(response["mobile"]);
        $('#confirmPasswordAdd').val(response["city"]);

        if (response["status"] == 1) {
            $('#activeStatus').prop('checked', true);
        } else {
            $('#activeStatus').prop('checked', false);
        }
        $('#activeStatus').val(response["status"]);
    });
    request.fail(function (response) {
        alert(response);
    });
}
$('#activeStatus').on('click', function (e) {
    let currentValue = $('#activeStatus').val().trim();
    if (currentValue == "1") {
        $('#activeStatus').val("0");
    } else {
        $('#activeStatus').val("1");
    }
});
$('#btnSearchAction').on('click', function (e) {
    console.log("search button pressed")
    console.log($('#phone').val().trim());
});
$('#btnAddAction').on('click', function (e) {
    console.log("Add button pressed")
    let btnLabel = $('#btnAddAction .mdc-button__label').text();
    console.log(btnLabel);

    $('#usernameAdd').prop("disabled", false);
    $('#emailAdd').prop("disabled", false);
    let formData = $('#addForm').serializeArray();

    if (btnLabel == "Add") {
        console.log(formData);
        //send to backend
        //sending ajax post request
        let request = $.ajax({
            async: true,
            type: "POST",
            url: "/api/register",
            data: formData,
        });
        request.done(function (response) {
            if (response.trim() == "username") {
                alert("Username already taken. Please choose a different one.");
            } else if (response.trim() == "email") {
                alert("Email already exsist. Please choose a different one.");
            } else if (response.trim() == "Registration Done") {
                alert("Registration successful. Email verification link was sent to your provided email.");
                location.reload();
            } else {
                alert("Registration unsuccessful. Something went wrong. Please try again!");
            }
        });
        request.fail(function (response) {
            alert(response);
        });
        dialogAdd.close();
    } else {
        console.log(formData);

        //send to backend
        //sending ajax post request
        let request = $.ajax({
            async: true,
            type: "POST",
            url: "/api/update",
            data: formData,
        });
        request.done(function (response) {
            console.log(response.trim())
            if (response.trim() == "OK") {
                updateFormData(formData); //update on page
                alert("Data updated successfully.");
            } else {
                alert("Update failed. Something went wrong!");
            }
        });
        request.fail(function (response) {
            alert(response);
        });
        dialogAdd.close();
    }
});
function updateFormData(formData) {
    selectedRow.cells[1].innerText = formData[0]["value"] + " " + formData[1]["value"];
    selectedRow.cells[2].innerText = formData[2]["value"];
    selectedRow.cells[3].innerText = formData[3]["value"];
    selectedRow.cells[4].innerText = formData[4]["value"];
    //selectedRow.cells[5].innerText = formData[5]["value"];    //skip for create_date
    selectedRow.cells[6].innerText = checkCheckbox();
}
function checkCheckbox() {
    if ($('#activeStatus').val().trim() == "1") {
        return "Active"
    } else {
        return "Inactive"
    }
}