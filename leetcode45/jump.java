

// class Solution {
//     public int jump(int[] nums) {
        
//         int reachableMax = 0;
//         int curEnd = 0;
//         int count = 0;

//         if (nums.length <= 1)
//             return count;

//         for (int i = 0; i < nums.length; i++) {
//             reachableMax = Integer.max(reachableMax, i + nums[i]);

//             if (i == curEnd) {
//                 count++;
//                 curEnd = reachableMax;

//                 if (curEnd >= nums.length - 1)
//                     break;
//             }
//         }
//         return count;
//     }
// }