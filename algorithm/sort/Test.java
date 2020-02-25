import java.util.*;

public class Test {
	
	public static void main(String args[]) {
		Scanner scan = new Scanner(System.in);
        System.out.println("请选择排序方式: 0(冒泡排序) 1（选择排序）2(插入排序) 3（快速排序）");
        if (scan.hasNext()) {
            String str = scan.next();
            Sort sort =  null;
            if (str.equals("0")) {
                sort = new BubbleSort();
            } else if(str.equals("1")) {
                sort = new SelectionSort();
            } else if(str.equals("2")) {
                sort = new InsertSort();
            } else if(str.equals("3")) {
                sort = new QuickSort();
            }
            int[] test = {2, 4, 1, 0, 3, 7, 8, 5, 6, 9};
            sort.sort(test);
            printArray(test);
        }
        scan.close();
	}

	public static void printArray(int[] arr) {
		for(int i = 0; i < arr.length; i++) {
			System.out.print(" " + i);
		}
		System.out.println();
	}	
}

