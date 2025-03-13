
// import java.util.Arrays;

// class Solution {
//     public int trap(int[] height) {
        
//         int count = 0;
//         int[] water = new int[height.length];
//         int high = 0;
//         for (int i = 0; i < height.length; i++) {
//             if (height[i] > high) {
//                 high = height[i];
//             }
//             water[i] = high;
//         }
//         for (int i = height.length - 1; i >= 0; i--) {
//             if (i == height.length - 1)
//                 high = height[i];
//             if (height[i] > high) {
//                 high = height[i];
//             }
//             water[i] = Math.min(water[i], high);
//             count += water[i] - height[i];
//         }

//         System.out.printf("%s",Arrays.toString(water));
//         return count;
//     }
// }