/*
Copyright © 2024 Colin Jacobs <colin@coljac.space>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-pdf/fpdf"
	"github.com/spf13/cobra"
)

var (
	title        string
	paper        string
	pageBreaks   bool
	pageNumbers  bool
	landscape    bool
	fontSize     float64
	output       string
	font         string
	mono         bool
	proportional bool
	openPdfFile  bool
	inputFiles   []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pdfpipe",
	Short: "Create a PDF file quickly from some text",
	Run:   run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&title, "title", "t", "", "a title string for the document")
	rootCmd.Flags().StringVarP(&paper, "paper", "g", "A4", "paper size (A4, A3, A5, Letter, Legal)")
	rootCmd.Flags().BoolVarP(&pageBreaks, "page-breaks", "b", false, "Add page break after each file's contents")
	rootCmd.Flags().BoolVarP(&pageNumbers, "page-numbers", "n", false, "Add page numbers in footer")
	rootCmd.Flags().BoolVarP(&landscape, "landscape", "l", false, "landscape orientation")
	rootCmd.Flags().Float64VarP(&fontSize, "font-size", "s", 12.0, "font size as float64 (default: 12.0)")
	rootCmd.Flags().StringVarP(&output, "output", "o", "output.pdf", "output file name")
	rootCmd.Flags().StringVarP(&font, "font", "f", "Courier", "name of font to use")
	rootCmd.Flags().BoolVar(&mono, "mono", true, "use monospace font (default)")
	rootCmd.Flags().BoolVarP(&proportional, "proportional", "p", false, "use proportional font")
	rootCmd.Flags().BoolVarP(&openPdfFile, "open-pdf-file", "x", false, "attempts to open the resulting file")
	rootCmd.Flags().StringSliceVar(&inputFiles, "input-files", []string{}, "text file input")

	// Initialize fonts
	// fpdf.SetFontLocation("") // Use built-in fonts
}

func run(cmd *cobra.Command, args []string) {
	// Override mono if proportional is set
	if proportional {
		mono = false
	}

	// Set font based on mono flag
	if mono {
		font = "Courier"
	} else if font == "Courier" {
		font = "Helvetica"
	}

	// Check if we need to read from stdin
	if len(inputFiles) == 0 {
		stdinContent, err := readStdin()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			os.Exit(1)
		}
		inputFiles = []string{"stdin"}
		createPDFFromContent([]string{stdinContent}, title, paper, fontSize, pageBreaks, pageNumbers, output, font, landscape, openPdfFile)
	} else {
		createPDF(inputFiles, title, paper, fontSize, pageBreaks, pageNumbers, output, font, landscape, openPdfFile)
	}
}

func readStdin() (string, error) {
	var content string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func createPDFFromContent(contents []string, title string, paper string, fontSize float64, pageBreaks bool, pageNumbers bool, output string, font string, landscape bool, openPdfFile bool) {
	orientation := "P"
	if landscape {
		orientation = "L"
	}

	pdf := fpdf.New(orientation, "mm", paper, "")

	// Add and set font
	err := addFont(pdf, font)
	if err != nil {
		fmt.Printf("Error setting font: %v\n", err)
		return
	}
	pdf.SetFont(font, "", fontSize)

	// Add header/footer
	pdf.SetHeaderFunc(func() {
		// Header implementation if needed
	})

	pdf.SetFooterFunc(func() {
		if pageNumbers {
			pdf.SetY(-15)
			pdf.SetFont("Arial", "I", 8)
			pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()), "", 0, "R", false, 0, "")
		}
	})

	pdf.AddPage()

	// Add title if specified
	if title != "" {
		pdf.SetFontSize(fontSize * 1.5)
		pdf.CellFormat(0, fontSize*1.5, title, "", 1, "C", false, 0, "")
		pdf.SetFontSize(fontSize)
		pdf.Ln(fontSize)
	}

	// Process input files
	for i, content := range contents {
		pdf.MultiCell(0, fontSize*0.45, content, "", "", false)

		if pageBreaks && i < len(contents)-1 {
			pdf.AddPage()
		} else {
			pdf.Ln(fontSize * 0.45 * 2)
		}
	}

	err = pdf.OutputFileAndClose(output)
	if err != nil {
		fmt.Printf("Error writing PDF: %v\n", err)
		return
	}

	fmt.Printf("PDF created: %s\n", output)

	if openPdfFile {
		openPDF(output)
	}
}

func openPDF(filename string) {
	// This is a placeholder. Implement the logic to open the PDF file based on the OS.
	fmt.Printf("Opening PDF file: %s\n", filename)
	// You might use something like exec.Command() to open the file with the default PDF viewer
}

func createPDF(inputFiles []string, title string, paper string, fontSize float64, pageBreaks bool, pageNumbers bool, output string, font string, landscape bool, openPdfFile bool) {
	orientation := "P"
	if landscape {
		orientation = "L"
	}

	pdf := fpdf.New(orientation, "mm", paper, "")

	// Add and set font
	err := addFont(pdf, font)
	if err != nil {
		fmt.Printf("Error setting font: %v\n", err)
		return
	}
	pdf.SetFont(font, "", fontSize)

	// Add header/footer
	pdf.SetHeaderFunc(func() {
		// Header implementation if needed
	})

	pdf.SetFooterFunc(func() {
		if pageNumbers {
			pdf.SetY(-15)
			pdf.SetFont("Arial", "I", 8)
			pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()), "", 0, "R", false, 0, "")
		}
	})

	pdf.AddPage()

	// Add title if specified
	if title != "" {
		pdf.SetFontSize(fontSize * 1.5)
		pdf.CellFormat(0, fontSize*1.5, title, "", 1, "C", false, 0, "")
		pdf.SetFontSize(fontSize)
		pdf.Ln(fontSize)
	}

	// Process input files
	for i, file := range inputFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		pdf.MultiCell(0, fontSize*0.45, string(content), "", "", false)

		if pageBreaks && i < len(inputFiles)-1 {
			pdf.AddPage()
		} else {
			pdf.Ln(fontSize * 0.45 * 2)
		}
	}

	err = pdf.OutputFileAndClose(output)
	if err != nil {
		fmt.Printf("Error writing PDF: %v\n", err)
		return
	}

	fmt.Printf("PDF created: %s\n", output)

	if openPdfFile {
		openPDF(output)
	}
}

func addFont(pdf *fpdf.Fpdf, font string) error {
	switch font {
	case "Courier":
		pdf.AddFont("Courier", "", "courier.json")
	case "Helvetica":
		pdf.AddFont("Helvetica", "", "helvetica_1.json")
	case "Arial":
		pdf.AddFont("Arial", "", "arial_1.json")
	default:
		return fmt.Errorf("unsupported font: %s", font)
	}
	return nil
}
