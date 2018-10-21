/**
 *  冒泡排序是一种极其简单的排序算法，也是我所学的第一个排序算法。它重复地走访过要排序的元素，依次比较相邻两个元素，如果他们的顺序错误就把他们调换过来，直到没有元素再需要交换，排序完成。这个算法的名字由来 是因为越小(或越大)的元素会经由交换慢慢“浮”到数列的顶端。
 *
 *  　　冒泡排序算法的运作如下：
 *
 *  比较相邻的元素，如果前一个比后一个大，就把它们两个调换位置。
 *  对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
 *  针对所有的元素重复以上的步骤，除了最后一个。
 *  持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
 *
 *  Average case = O(n^2) Worst case = O(n^2) Best case = O(n)
 * */

public class BubbleSort implements Sort {
	
	@Override
	public void sort(int[] sortArray) {
		int n = sortArray.length;
		for (int i = 0; i < n -1; i++) {
			for (int j = 0; j < n - i -1; j++) {
				if (sortArray[j] > sortArray[j + 1]) {
					int tmp = sortArray[j];
					sortArray[j] = sortArray[j + 1];
					sortArray[j + 1] = tmp;
				}
			}
		}
	}
}