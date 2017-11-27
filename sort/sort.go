package sort

/**
 * Implement various sorting algorithms
 */

import "fmt"

/**
 * Sort by insertig each item into the right place in the array
 */
func InsertionSort(A []int) []int {
    for i := 1; i < len(A); i++ {
        for j := i; j > 0; j-- {
            if A[j-1] > A[j] {
                A[j], A[j-1] = A[j-1], A[j]
            } else {
                break;
            }
        }
    }
    return A
}

/**
 * Sort by selecting each next biggest item
 */
func SelectionSort(A []int) []int {
    for i := range A {
        next := i
        for j := i + 1; j < len(A); j++ {
            if A[j] < A[next] {
                next = j
            }
        }
        A[i], A[next] = A[next], A[i]
    }
    return A
}

/**
 * Sort by splitting the list into smaller arrays and merging them
 */
func MergeSort(A []int) []int {
    return mergeSort(A, 0, len(A))
}

func mergeSort(A []int, l, r int) []int {
    if l < r - 1 {
        c := (l + r) / 2
        mergeSort(A, l, c)
        mergeSort(A, c, r)
        return merge(A, l, c, r)
    }
    return A
}

func merge(A []int, p, q, r int) []int {
    // Make a copy of the left slice of the array
    L := make([]int, q-p)
    for i := range L {
        L[i] = A[p + i]
    }

    // Make a copy of the right slice of the array
    R := make([]int, r-q)
    for j := range R {
        R[j] = A[q + j]
    }

    // Now merge
    for k, i, j := p, 0, 0; k < r; k++ {
        switch {
        // If L is depleted, grab an item from R
        case i >= len(L):
            A[k] = R[j]
            j++

        // If R is depleted, grab an item from L
        case j >= len(R):
            A[k] = L[i]
            i++

        // If L has the smaller value, grab it
        case L[i] <= R[j]:
            A[k] = L[i]
            i++

        // If R has the smaller value, grab it
        default:
            A[k] = R[j]
            j++
        }
    }

    return A
}

/**
 * Sort the input by using a max-heap
 */
func HeapSort(A []int) []int {
    buildMaxHeap(A)
    // Sort the numbers from max to min by popping from the max heap
    for i := len(A)-1; i > 0; i-- {
        A[i], A[0] = A[0], A[i]
        maxHeapify(A, 0, i-1)
    }
    return A
}

func buildMaxHeap(A []int) {
    for i := len(A) / 2; i >= 0; i-- {
        maxHeapify(A, i, len(A) - 1)
    }
}

func maxHeapify(A []int, i, boundary int) {
    // Left heap node
    l := 2 * i + 1

    // Right heap node
    r := 2 * i + 2

    // If either of the children are greater than the current node, swap
    max := i
    if l <= boundary && A[l] > A[i] {
        max = l
    }
    if r <= boundary && A[r] > A[max] {
        max = r
    }

    // If we did swap, repeat the process on the child heap
    if max != i {
        A[i], A[max] = A[max], A[i]
        maxHeapify(A, max, boundary)
    }
}

/**
 * Sort by dividing the array into 2 smaller arrays around a pivot
 */
func QuickSort(A []int) []int {
    if len(A) == 0 {
        return A
    }
    quickSort(A, 0, len(A) - 1)
    return A
}

func quickSort(A []int, l, r int) {
    if l < r {
        pivot := partition(A, l, r)
        quickSort(A, l, pivot - 1)
        quickSort(A, pivot + 1, r)
    }
}

func partition(A []int, l, r int) int {
    pivot := A[r]
    i := l
    for j := l; j < r; j++ {
        if A[j] <= pivot {
            A[i], A[j] = A[j], A[i]
            i += 1
        }
    }
    A[i], A[r] = A[r], A[i]
    return i
}
