package algo

func insertion_sort(A []int) {
	
	for i := 1; i < len(A); i++ {
		
		key := A[i]
		
		j := i - 1
		for j >= 0 && key < A[j] {
			A[j+1] = A[j]
			j--
		}
		
		A[j+1] = key
	}
}

/*
func main() {
	A := []int{3, 1, 5, 2, 9}
	merge_sort(A, 0, len(A) - 1)
	fmt.Println(A)
}
*/

func merge_sort(A []int, p, r int) {

	if (p < r) {
        
        q := (p+r)/2
        merge_sort(A, p, q)
        merge_sort(A, q+1, r)
        merge(A, p, q, r)
    }
}

func merge(A []int, p, q, r int) {

	L := make([]int, q-p+1)
    R := make([]int, r-q)
    
    for i := 0; i < q-p+1; i++ {
        L[i] = A[i+p]
    }

    for j := 0; j < r-q; j++ {
        R[j] = A[j+q+1]
    }

    i, j := 0, 0
    for k := p; k <= r; k++ {

        if j == len(R) || (i < len(L) && L[i] <= R[j]) {
            A[k] = L[i]
            i++
        } else {
            A[k] = R[j]
            j++
        }
    }
}