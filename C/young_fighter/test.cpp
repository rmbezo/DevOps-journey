#include <iostream>
#include <stdio.h>

int binary_search(int arr[], int size, int target) {
    int low = 0;
    int high = size - 1;

    while (low <= high) {
        int mid = low + (high - low) / 2;

        if (arr[mid] == target) return mid;
        if (arr[mid] < target) low = mid + 1;
        else high = mid - 1;
    }
    return -1;
}

int main() {
    const int SIZE = 65;
    int array[SIZE];

    for (int i = 0; i < SIZE; i++) { 
        array[i] = i;
    }

    int result = binary_search(array, SIZE, 50);

    printf("result is: %i\n", result);
    return 0;
}
