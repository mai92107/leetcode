

// class RemoveDuplicates {
//     public int removeDuplicates(int[] nums) {
        
//         int k = 0;
//         int t = 999;
//         for (int i = 0; i < nums.length; i++) {
//             if (nums[i] != t){
//                 nums[k] = nums[i];
//                 t = nums[k];
//                 k++;
//             }
//         }
//         return k;
//     }
// }