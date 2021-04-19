package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

//pdf requestpdf struct
type RequestPdf struct {
	body string
}

//new request to pdf function
func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

//parsing template function
func (r *RequestPdf) ParseTemplate(templateFileName string, data interface{}) error {

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

//generate pdf function
func (r *RequestPdf) GeneratePDF(pdfPath string) (bool, error) {
	time := time.Now().Unix()
	tempFileName := fmt.Sprintf("tempPDF/%s.html", strconv.FormatInt(int64(time), 10))

	// write whole the body
	err := ioutil.WriteFile(tempFileName, []byte(r.body), 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(tempFileName)
	if err != nil {
		log.Fatal(err)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		os.Remove(tempFileName)
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	//pdfg.Cover.EnableLocalFileAccess.Set(true)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	if f != nil {
		f.Close()
	}

	err = os.Remove(tempFileName)
	if err != nil {
		log.Fatal(94, err)
	}

	return true, nil
}
func pdf(ledgerCode string) bool {
	//wkhtmltopdf.SetPath(`C:\Program Files\wkhtmltopdf\bin\wkhtmltopdf.exe`)

	//preparing path for store pdf
	wDir, _ := os.Getwd()
	certificatePath := filepath.Join(wDir, "print", "certificate", ledgerCode+".pdf")

	//html template data
	templateData := struct {
		StudentsName string
		CourseName   string
		Duration     string
		QRImagePath  string
		QRText       string
	}{
		StudentsName: "Mostain Billah",
		CourseName:   "Golang Programming Language",
		Duration:     "January - March 2021",
		QRImagePath:  `https://chart.googleapis.com/chart?chs=150x150&cht=qr&chl=Hello%20world`,
		QRText:       "MA-" + ledgerCode,
	}

	//processing for PDF generate
	reqPDF := NewRequestPdf("")

	//html template path of certificate
	templateFile := "template/certificate/index.html"

	err := reqPDF.ParseTemplate(templateFile, templateData)
	checkErr(err)

	res, err := reqPDF.GeneratePDF(certificatePath)
	checkErr(err)

	return res
}
