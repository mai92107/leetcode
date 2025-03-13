
// class Solution {
//     public String intToRoman(int num) {
//         int[] in = new int[] {1000,900,500,400,100,90,50,40,10,9,5,4,1};
//         String[] s = new String[] { "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I" };
//         String ans = "";
//         for (int i = 0; i < in.length && num > 0; i++) {
//             while (num >= in[i]) {
//                 num -= in[i];
//                 ans += s[i];
//             }
//         }
//         return ans;
//     }
// }