
// import java.util.ArrayList;
// import java.util.HashMap;
// import java.util.List;
// import java.util.Map;
// import java.util.Random;


// class RandomizedSet {
//     private List<Integer> list;
//     private Map<Integer, Integer> map;
//     private Random r;


//     public RandomizedSet() {
//         list = new ArrayList<>();
//         map = new HashMap<>();
//         r = new Random();
//     }
    
//     public boolean insert(int val) {
//         if (map.containsKey(val)) {
//             return false;
//         }
//         list.add(val);
//         map.put(val, list.size() - 1);
//         return true;
//     }
    
//     public boolean remove(int val) {
//         if (!map.containsKey(val)) {
//             return false;
//         }
//         int kindex;
//         kindex = map.get(val);
//         list.set(kindex, list.getLast());
//         map.replace(list.getLast(), list.size() - 1, kindex);

//         list.removeLast();
//         map.remove(val);
//         return true;
//     }
    
//     public int getRandom() {
//         return list.get(r.nextInt(list.size()));
//     }
// }
