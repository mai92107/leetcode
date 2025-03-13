// class Solution {
//     public int[] productExceptSelf(int[] nums) {
        
//         int n = nums.length;
//         int[] ans = new int[n];

//         //從左邊走
//         ans[0] = 1;
//         for (int i = 1; i < nums.length; i++) {
//             ans[i] = ans[i - 1] * nums[i - 1];
//         }

//         //從右邊走
//         int r = 1;
//         for (int i = n-1; i >= 0; i--){
//             ans[i] *= r;
//             r *= nums[i];
//         }

//         return ans;
//     }
// }