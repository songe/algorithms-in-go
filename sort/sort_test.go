package sort

import "testing"

func TestInsertionSort(t *testing.T) {
    test(InsertionSort, t)
}

func TestSelectionSort(t *testing.T) {
    test(SelectionSort, t)
}

func TestMergeSort(t *testing.T) {
    test(MergeSort, t)
}

func TestHeapSort(t *testing.T) {
    test(HeapSort, t)
}

func TestQuickSort(t *testing.T) {
    test(QuickSort, t)
}

func TestMerge(t *testing.T) {
    cases := []struct {
        in []int; p, q, r int; expected []int
    }{
        { []int{2,1}, 0, 1, 2, []int{1,2} },
        { []int{3,7,5,9}, 0, 2, 4, []int{3,5,7,9} },
        { []int{1,2,3,4,5,6,7,8}, 0, 4, 8, []int{1,2,3,4,5,6,7,8} },
        { []int{5,6,7,8,1,2,3,4}, 0, 4, 8, []int{1,2,3,4,5,6,7,8} },
        { []int{3,5,7,9,0,1,4,8}, 0, 4, 8, []int{0,1,3,4,5,7,8,9} },
    }
    for _, c := range cases {
        actual := merge(c.in, c.p, c.q, c.r)
        if !isEquals(actual, c.expected) {
            t.Errorf("Merge(%v, %d, %d, %d) == %v, expected %v", c.in, c.p, c.q, c.r, actual, c.expected)
        }
    }
}

func test(f func([]int) []int, t *testing.T) {
    cases := []struct {
        in, expected []int
    }{
        { []int{}, []int{} },
        { []int{1}, []int{1} },
        { []int{2, 1}, []int{1, 2} },
        { []int{1, 3, 2}, []int{1, 2, 3} },
        { []int{3, 1, 4}, []int{1, 3, 4} },
    }

    for _, c := range cases {
        actual := f(c.in)
        if !isEquals(actual, c.expected) {
            t.Errorf("Sort(%v) == %v, expected %v", c.in, actual, c.expected)
        }
    }
}

func isEquals (a1, a2 []int) bool {
    if len(a1) != len(a2) {
        return false
    }

    for i := range a1 {
        if a1[i] != a2[i] {
            return false
        }
    }

    return true
}
