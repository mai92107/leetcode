

// import java.util.Arrays;
// import java.util.Comparator;

// class Solution {
//     public String longestCommonPrefix(String[] strs) {

//         strs = Arrays.stream(strs).sorted(Comparator.comparingInt(String::length)).toArray(String[]::new);
// System.out.println(Arrays.toString(strs));
//         for (int i = 0; i < strs[0].length(); i++) {
//             for (String s : strs) {
//                 if (s.isEmpty()){
//                     return "";
//                 }
//                 if (strs[0].charAt(i) != s.charAt(i)){
//                     return s.substring(0,i);
//                 }
//             }
//         }
//         return strs[0];
//     }
// }