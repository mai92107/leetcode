// import java.util.Arrays;
// import java.util.HashMap;
// import java.util.Map;
// import javax.xml.stream.events.Characters;

// class Solution {
//     public int romanToInt(String s) {
//         int i = 0;
//         String[] arr = s.split("");
//         Map<String, Integer> map = new HashMap<>();
//         int total = 0;
//         map.put("I", 1);
//         map.put("V", 5);
//         map.put("X", 10);
//         map.put("L", 50);
//         map.put("C", 100);
//         map.put("D", 500);
//         map.put("M", 1000);

//         while (i < arr.length) {
//             if (i != arr.length - 1 && map.get(arr[i]) < map.get(arr[i+1])) {
//                 total -= map.get(arr[i]);
//             } else {
//                 total += map.get(arr[i]);
//             }
//             i++;
//         }
//         return total;
//     }
// }