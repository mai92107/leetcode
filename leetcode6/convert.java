class Solution {
    public static String convert(String s, int numRows) {

        if (numRows == 1) {
            return s;
        }
        String[] arr = new String[numRows];
        for (int i = 0; i < numRows; i++) {
            arr[i] = "";
        }
        boolean down = true;
        int idx = 0;
        for (int i = 0; i < s.length();) {
            arr[idx] += s.charAt(i);
            if (down)
                idx++;
            else
                idx--;
            i++;
            if (idx == 0 || idx == numRows - 1) {
                down = !down;
            }
        }
        StringBuilder sb = new StringBuilder();
        for (String st : arr) {
            sb.append(st);
        }
        return sb.toString();
    }
    public static void main(String[] args) {
        System.out.println(convert("HUIHHIL",5));
    }
}