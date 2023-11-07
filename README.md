# Golang Dynamic HTML Template to PDF Converter (api2pdf)

## What it is?

In this repository, you will see how to convert with api2pdf integration and how struct data can be integrated into the dynamic template and the pdf result.

## Prerequisite
- Let's create a data structure and set our data.
- Prepare an html template according to ourselves.
- Let's add a range loop to the parts we want to duplicate so that it gets the name struct in. For example {{ .XYZ}}
- Now we are going to parse our template, so let's write the template address, paying attention to which directory it is in. For example, "templates/dynamictemplate.html"
- It's time to send the POST request. Because we want to convert from pdf to html, api2pdf's "https://v2.api2pdf.com/chrome/pdf/html " we are sending our request to this url.
- While doing this, we need to set the necessary data.
- We are setting Authorization, HTML and FileName,
- and we write our response to our struct model and access the data returned to us.
- 
## How does the pdf look like
- When you click on the url,
  
![image](https://github.com/elcinzorlu/go-dynamic-template-with-convert-api2pdf/assets/107582166/e6d678c5-a832-4d55-aee9-3cb6248c1c90)


## Resources

- https://app.swaggerhub.com/apis-docs/api2pdf/api2pdf/2.0.0#/Headless%20Chrome/chromePdfFromHtmlPost
