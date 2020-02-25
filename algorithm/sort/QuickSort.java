
public class QuickSort implements Sort {
	@Override
	public void sort(int[] array) {
    	quickSort(array, 0, array.length - 1);
	}

    private void quickSort(int[] array, int start, int end) {
        int smallIndex = partion(array, start, end);
        if(smallIndex > start) {
            quickSort(array, start, smallIndex - 1);
        }
        if(smallIndex < end) {
            quickSort(array, smallIndex + 1, end);
        }                                    
    }

    private int partion(int[] array, int start, int end) {
        int pivot = (int)(start + Math.random()*(end - start + 1));
        int smallIndex = start - 1;
        swap(array, pivot, end);
        for(int i = start; i <= end; i++) {
            if(array[i] <= array[end]) {
                smallIndex++;
                if(i > smallIndex) {
                    swap(array, i, smallIndex);
                }
            }
        }
        return smallIndex;
   }

   private void swap(int[] array, int a1, int a2) {
       int temp = array[a1];
       array[a1] = array[a2];
       array[a2] = temp;
   }
}
