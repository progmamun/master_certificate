package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	models "github.com/mateors/master-erp"
	"github.com/mateors/mcb"
	"github.com/mateors/mtool"
	"golang.org/x/crypto/bcrypt"
)

var db *mcb.DB

const (
	BucketName = "master_erp" //

	CompanyTable            = "company"              //
	AccessTable             = "access"               //
	LoginTable              = "login"                //
	AccountTable            = "account"              //
	ContactTable            = "contact"              //
	LoginSessionTable       = "login_session"        //
	DeviceLogTable          = "device_log"           //
	VisitorSessionTable     = "visitor_session"      //
	VerificationTable       = "verification"         //
	MessageTable            = "message"              //
	ActivityLogTable        = "activity_log"         //
	SettingsTable           = "settings"             //
	CategoryTable           = "category"             //**ItemCategory
	ItemTable               = "item"                 //**
	WarehouseTable          = "warehouse"            //
	DocKeeperTable          = "doc_keeper"           //
	TransactionRecordTable  = "transaction_record"   //sales|purchase
	StockMovementTable      = "stock_movement"       //
	LedgerTransactionTable  = "ledger_transaction"   //
	FileStoreTable          = "file_store"           //
	DepartmentTable         = "department"           //owner = item | item_line | office
	ItemRewardTable         = "item_reward"          // item purchase point for loyalty
	ItemAttributeTable      = "item_attribute"       //
	ItemAttributeValueTable = "item_attribute_value" //
	UOMTable                = "uom"                  //
	TaxTable                = "tax"                  //tax | vat table
	AccountHeadTable        = "achead"               //account_head ledger head
	Rateplan                = "rateplan"             //
	AddressTable            = "address"              // contact address
	ShippingAddressTable    = "shipping_address"     //*
	PaymentOptionTable      = "payment_option"       //
	ShippingOptionTable     = "shipping_option"      //
	DocPayShipTable         = "docpayship"           //doc_payship_info
	//AccountGroupTable       = "account_group"        //removed account group
)

func init() {
	//couchbase connection block
	db = mcb.Connect("localhost", "root", "bootcamp")
	res, err := db.Ping()
	if err != nil {
		fmt.Println(res)
		os.Exit(1)
	}
}

func doRegistration(formData map[string]string, w http.ResponseWriter, r *http.Request) {
	//** process starts for checking into DB **//
	check := isExsist(formData) //checking if username or email already exsist or not
	if check != "" {
		fmt.Fprintln(w, check)
		return
	}
	//** process ends for checking into DB **//

	//** process starts for AccountTable **//
	cid := "company::1"
	accID, accSerial := findAailabledocID(AccountTable)
	ledgerCode := getLedgerCode(accSerial)
	fullName := strings.TrimSpace(formData["firstName"]) + " " + strings.TrimSpace(formData["lastName"])
	createDate := time.Now().String()[:19]

	r.Form.Set("bucket", BucketName)
	r.Form.Set("aid", accID)
	r.Form.Set("type", AccountTable)
	r.Form.Set("cid", cid)
	r.Form.Set("serial", strconv.FormatInt(accSerial, 10))
	r.Form.Set("photo", "")
	r.Form.Set("account_type", "STUDENT")
	r.Form.Set("account_name", fullName)
	r.Form.Set("customid", "")
	r.Form.Set("code", ledgerCode) //ledgercode
	r.Form.Set("login_id", "")
	r.Form.Set("first_name", strings.TrimSpace(formData["firstName"]))
	r.Form.Set("last_name", strings.TrimSpace(formData["lastName"]))
	r.Form.Set("dob", "")
	r.Form.Set("gender", "")
	r.Form.Set("mobile", "")
	r.Form.Set("email", formData["email"])
	r.Form.Set("remarks", "Added by Form")
	r.Form.Set("create_date", createDate)
	r.Form.Set("update_date", createDate)
	r.Form.Set("status", "1")

	var account models.Account
	insert1 := db.Insert(r.Form, &account) //insert query
	//** process ends for AccountTable **//

	//** process starts for LoginTable **//
	loginID, loginSerial := findAailabledocID(LoginTable)

	r.Form.Set("aid", loginID)
	r.Form.Set("type", LoginTable)
	r.Form.Set("cid", cid)
	r.Form.Set("serial", strconv.FormatInt(loginSerial, 10))
	r.Form.Set("account_id", accID)
	r.Form.Set("access_name", "STUDENT")
	r.Form.Set("username", formData["username"])
	r.Form.Set("passw", makeHash(formData["password"]))
	r.Form.Set("create_date", createDate)
	r.Form.Set("last_login", "")
	r.Form.Set("status", "1")

	var login models.Login
	insert2 := db.Insert(r.Form, &login)
	//** process ends for LoginTable **//

	//** process starts for AddressTable **//
	addID, addSerial := findAailabledocID(AddressTable)

	r.Form.Set("aid", addID)
	r.Form.Set("type", AddressTable)
	r.Form.Set("cid", cid)
	r.Form.Set("serial", strconv.FormatInt(addSerial, 10))
	r.Form.Set("account_id", accID)
	r.Form.Set("address_type", "Billing")
	r.Form.Set("country", "")
	r.Form.Set("state", "")
	r.Form.Set("city", "")
	r.Form.Set("address1", "")
	r.Form.Set("address2", "")
	r.Form.Set("zip", "")
	r.Form.Set("status", "1")

	var address models.Address
	insert3 := db.Insert(r.Form, &address)
	//** process ends for AddressTable **//

	//** process starts for AccountHeadTable **//
	fVal := getFieldVals("aid,account_type", AccountHeadTable, fmt.Sprintf(`cid="%s" AND name="%s" AND account_group="group"`, cid, "Account Receivable"))
	parentID := fVal["aid"]
	accHeadID, accHeadSerial := findAailabledocID(AccountHeadTable)

	r.Form.Set("aid", accHeadID)
	r.Form.Set("type", AccountHeadTable)
	r.Form.Set("cid", cid)
	r.Form.Set("serial", strconv.FormatInt(accHeadSerial, 10))
	r.Form.Set("account_group", fVal["account_group"])
	r.Form.Set("account_type", "STUDENT")
	r.Form.Set("name", fullName+"(STUDENT)")
	r.Form.Set("description", "")
	r.Form.Set("identifier", "receivable")
	r.Form.Set("code", ledgerCode) //ledgercode
	r.Form.Set("parent_id", parentID)
	r.Form.Set("balance", fmt.Sprintf("%f", 0.0))
	r.Form.Set("baltype", "Eq")
	r.Form.Set("restricted", "0")
	r.Form.Set("cost_center", "0")
	r.Form.Set("remarks", "")
	r.Form.Set("create_date", createDate)
	r.Form.Set("status", "1")

	var accountHead models.AccountHead
	insert4 := db.Insert(r.Form, &accountHead)
	//** process ends for AccountHeadTable **//

	//** process starts for updating AccountTable with login ID**//
	qs := `UPDATE master_erp SET login_id = "%s" WHERE type = "account" AND aid = "%s" RETURNING login_id`
	sql := fmt.Sprintf(qs, loginID, accID)

	pRes := db.Query(sql)
	rows := pRes.GetRows()
	updateRes := rows[0]["login_id"].(string)
	//** process ends for updating AccountTable **//

	if insert1.Status == "success" && insert2.Status == "success" && insert3.Status == "success" && insert4.Status == "success" && updateRes == loginID {
		//sending mail verification link to the user mail
		//sendMail(r.FormValue("email"), r.FormValue("username"), "http://localhost:9000/verifyemail/"+token)

		fmt.Fprintln(w, "Registration Done")
	} else {
		fmt.Fprintln(w, "Registration Error")
	}
}
func isExsist(formData map[string]string) string {
	//checking for username
	query := `SELECT username FROM master_erp WHERE type="login" AND username="` + formData["username"] + `"`
	queryRes := db.Query(query)
	rows := queryRes.GetRows()
	if len(rows) != 0 {
		return "username"
	}

	//checking for email
	query = `SELECT email FROM master_erp WHERE type="account" AND email="` + formData["email"] + `"`
	queryRes = db.Query(query)
	rows = queryRes.GetRows()
	if len(rows) != 0 {
		return "email"
	}
	return ""
}
func getLedgerCode(accSerial int64) string {
	serial := strconv.FormatInt(accSerial, 10)
	ledgerCode := "12"
	for i := 0; i < 8-len(serial); i++ {
		ledgerCode += "0"
	}
	ledgerCode += serial
	return ledgerCode
}
func makeHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
func findAailabledocID(table string) (string, int64) {
	var docID string
	var scount int64 = int64(maxDoc(BucketName, table, db))
	i := scount + 1

	for {
		count := i
		docID = fmt.Sprintf(`%s::%v`, table, count)
		//sql := fmt.Sprintf(`SELECT COUNT(*)as cnt FROM %s WHERE META().id="%s";`, BucketName, docID)
		sql := fmt.Sprintf(`SELECT META(d).id, true AS docexist FROM %s AS d USE KEYS ["%s"];`, BucketName, docID)
		//fmt.Println(sql)
		pRes := db.Query(sql)
		rows := pRes.GetRows()

		if len(rows) == 0 {
			return docID, count
		}
		exist := rows[0]["docexist"].(bool)
		fmt.Println("ALREADY EXIST", i, exist, sql)
		i++
	}
}
func maxDoc(bucketName, tableName string, db *mcb.DB) (count float64) {

	sql := fmt.Sprintf(`SELECT NVL(MAX(serial),0)as cnt FROM %s WHERE type="%s";`, bucketName, tableName)
	pResponse := db.Query(sql)
	//fmt.Println("CountDoc>", sql, pResponse.Result)
	for _, val := range pResponse.Result {
		//fmt.Printf("%v %T\n", val, val)
		//rMap:=val.(map)
		cMap := val.(map[string]interface{})
		if c, ok := cMap["cnt"]; ok {
			count = c.(float64)
		}
	}
	//fmt.Println("CountTable>", pResponse.Result)
	return
}
func getFieldVals(fieldNames, table, where string) (fVals map[string]string) {
	fields := strings.Split(fieldNames, ",")
	fVals = make(map[string]string)

	sql := fmt.Sprintf(`SELECT %v FROM %v WHERE type="%s" AND %v;`, fieldNames, BucketName, table, where)
	pRes := db.Query(sql)
	//fmt.Println("Field >>", sql)

	rows := pRes.GetRows()
	for _, row := range rows {

		for key := range row {

			if mtool.ArrayValueExist(fields, key) {
				fieldValue := fmt.Sprintf("%v", row[key])
				fVals[key] = fieldValue
			}
		}
	}
	return
}
func doLogin(username, password string) string {
	qs := `SELECT passw FROM master_erp WHERE type="login" AND username="%s"`
	sql := fmt.Sprintf(qs, username)

	pRes := db.Query(sql)
	rows := pRes.GetRows()
	if len(rows) == 0 {
		//fmt.Fprintln(w, "Username not found!")
		return "username"
	}
	originalPassword := rows[0]["passw"].(string)
	if checkHash(password, originalPassword) {
		return "Done"
	}
	return ""
}
func checkHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func getStudentList(cid string) []map[string]interface{} {
	qs := `SELECT ac.first_name, ac.last_name, ac.email, ac.mobile, ac.create_date, ac.status, l.username as username FROM master_erp ac
	LEFT JOIN master_erp as a ON a.account_id=META(ac).id AND a.type="address"
	LEFT JOIN master_erp as l ON l.account_id=META(ac).id AND l.type="login"
	WHERE ac.cid = "%s" AND ac.type = "account" AND ac.status IN [0,1] ORDER BY ac.create_date DESC;`
	sql := fmt.Sprintf(qs, cid)

	pRes := db.Query(sql)
	rows := pRes.GetRows()
	//fmt.Println(rows)

	return rows
}
func getSingleStudent(username string) []map[string]interface{} {
	qs := `SELECT ac.first_name, ac.last_name, ac.email, ac.mobile, ac.status, l.username as username, ad.city as city FROM master_erp ac
	LEFT JOIN master_erp as a ON a.account_id=META(ac).id AND a.type="address"
	LEFT JOIN master_erp as l ON l.account_id=META(ac).id AND l.type="login"
	LEFT JOIN master_erp as ad ON ad.account_id=META(ac).id AND ad.type="address"
	WHERE l.username = "%s" AND ac.type = "account" AND ac.status IN [0,1];`
	sql := fmt.Sprintf(qs, username)

	pRes := db.Query(sql)
	rows := pRes.GetRows()

	return rows
}
func updateStudentInfo(firstName, lastName, username, mobile, city string, status int) string {
	qs := `SELECT account_id FROM master_erp WHERE type = "login" AND username = "%s"`
	sql := fmt.Sprintf(qs, username)

	pRes := db.Query(sql)
	rows := pRes.GetRows()
	accID := rows[0]["account_id"].(string)

	qs = `UPDATE master_erp SET first_name = "%s", last_name = "%s", mobile = "%s", status = %d WHERE type = "account" AND aid = "%s" RETURNING aid`
	sql = fmt.Sprintf(qs, firstName, lastName, mobile, status, accID)

	pRes = db.Query(sql)
	rows = pRes.GetRows()
	update1 := rows[0]["aid"].(string)

	qs = `UPDATE master_erp SET city = "%s" WHERE type = "address" AND account_id = "%s" RETURNING account_id`
	sql = fmt.Sprintf(qs, city, accID)

	pRes = db.Query(sql)
	rows = pRes.GetRows()
	update2 := rows[0]["account_id"].(string)

	if update1 == accID && update2 == accID {
		return "OK"
	}

	return "failed"
}
