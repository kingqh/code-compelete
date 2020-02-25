/**
 *
* */

public class InsertSort implements Sort {
	@Override
	public void sort(int[] attr) {
		int n = attr.length;
		for (int i = 1; i < n; i++) {
			int index = i;
            int current = attr[i];
			while(index > 0 && current < attr[index--]) {
                attr[index + 1] = attr[index]; 
            }
            attr[index] = current;
		}		
	}
}
