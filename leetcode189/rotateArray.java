
// class Solution {
//     void reverse(int[]nums, int s, int e){
//         while (s <= e) {
//             int t = nums[s];
//             nums[s] = nums[e];
//             nums[e] = t;
//             s++;
//             e--;
//         }
//     }

//     public void rotate(int[] nums, int k) {
        
//         k = k % nums.length;
//         reverse(nums, 0, nums.length-1);
//         reverse(nums, 0, k-1);
//         reverse(nums, k, nums.length-1);
//     }
// }