package Testing

func CheckError(err error) string {
	if err != nil {
		return err.Error() // Print the error message
	} else {
		return ""
	}

}
