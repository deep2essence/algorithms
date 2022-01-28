- a hybrid of merge sort and insertion sort algorithm
- works on real-time data
## Steps
```
1. Divide the array into blocks known as run
2. The size of a run can either be 32 or 64
3. Sort the elements of every run using insertion sort
4. Merge the sorted runs using the merge sort algorithm
5. Double the size of the merged array after every iteration
```