package basicmath

func Prime(n int ) bool {
	if n==1 || n==2 {
		return true
	}
	for i:=3;i<n ;i++{
		if n%i ==0 {
			return false
		}
	}
	return true
}