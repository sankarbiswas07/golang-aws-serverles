package handlers

func CreateUser(){

}

func GetUser(){
	
}

func UpdateUser(){
	
}

func DeleteUser(){
	
}
func UnhandledMethod(){
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       "Method not allowed",
	}, nil
}