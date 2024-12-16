# Hi, this a calculator 
That supports simple opearations such as multiplication, division, addition and substraction.
### Types of opearations:
"/" - division
"*" - multiplication
"+" - addition
"-" - substraction
### Errors and response status
"invalid expression" - 400
"division by zero" - 400
"empty expression" -400
Other Errors - 500
### Example request (body)
(2+2)*2 (result- 8,status - 200)
2+2*2 (result- 6,status - 200)
1/0 (status - 400)
### Start
go run rpn/cmd/main.go (in projects folder)