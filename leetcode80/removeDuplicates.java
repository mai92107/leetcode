


// class Solution {
//     public int removeDuplicates(int[] nums) {
//         int k = 0;
//         boolean kill = false;

//         for (int i = 0; i < nums.length; i++) {
//             if (i == 0) {
//                 nums[k] = nums[i];
//                 k++;
//                 continue;
//             }
//             if (nums[i - 1] != nums[i]) {
//                 nums[k] = nums[i];
//                 k++;
//                 kill = false;
//                 continue;
//             }
//             if (nums[i - 1] == nums[i] && !kill) {
//                 nums[k] = nums[i];
//                 kill = true;
//                 k++;
//             }
//         }
//         return k;
//     }
// }