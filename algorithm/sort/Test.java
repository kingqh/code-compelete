public class Test {
	
	public static void main(String args[]) {
		Sort sort = new BubbleSort();
		int[] test = {2, 4, 1, 0, 3, 7, 8, 5, 6, 9};
		sort.sort(test);
		printArray(test);
	}

	public static void printArray(int[] arr) {
		for(int i = 0; i < arr.length; i++) {
			System.out.print(" " + i);
		}
		System.out.println();
	}	
}

