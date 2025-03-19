





// class Solution {
//     public String reverseWords(String s) {
//         String[] arr = s.trim().split("\\s+");
//         StringBuilder word = new StringBuilder();

//         for (int idx = arr.length-1; idx >=0; idx--) {
//             word.append(arr[idx].trim());
//             if (idx == 0) {
//                 continue;
//             }
//             word.append(" ");            
//         }

//         return word.toString();
//     }
// }