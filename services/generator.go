package services

import (
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"linkcheck/models"
)

func GeneratePDF(reportData []models.ReportData) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.Cell(0, 10, "Links Checker Report")
	pdf.Ln(36)

	pdf.Cell(20, 10, "Num")
	pdf.Cell(60, 10, "Links")
	pdf.Cell(40, 10, "Status")
	pdf.Ln(30)

	for _, data := range reportData {
		pdf.Cell(20, 10, strconv.Itoa(data.ID))
		pdf.Cell(60, 10, data.Link)
		pdf.Cell(40, 10, data.Status)
		pdf.Ln(10)
	}

	tempFile := "report.pdf"
	if err := pdf.OutputFileAndClose(tempFile); err != nil {
		return "", err
	}

	return tempFile, nil
}
