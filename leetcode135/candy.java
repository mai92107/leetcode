// class Solution {
//     public int candy(int[] ratings) {
        
//         int[] candies = new int[ratings.length];
//         int total = 0;
//         for (int i = 0; i < ratings.length; i++) {
//             candies[i] = 1;
//             if (i != 0 && ratings[i] > ratings[i - 1]) {
//                 candies[i] = candies[i - 1] + 1;
//             }
//         }
//         for (int i = ratings.length - 1; i > 0; i--) {
//             if (ratings[i - 1] > ratings[i] && !(candies[i - 1] > candies[i])) {
//                 candies[i - 1] = candies[i] + 1;
//             }
//         }
//         for (int i = 0; i < ratings.length; i++) {
//             total += candies[i];
//         }
//         return total;
//     }
// }