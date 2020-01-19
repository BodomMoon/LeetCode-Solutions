// Time:  O(logn)
// Space: O(1)

class Solution {
public:
    vector<int> getNoZeroIntegers(int n) {
        int base = 1, a = 0;
        for (int curr = n; curr; curr /= 10) { 
            if (curr % 10 == 0 || (curr % 10 == 1 && curr != 1)) {
                a += 1 * base;
                curr -= 10;  // carry
            }
            a += 1 * base;
            base *= 10;
        }
        return {a, n - a};
    }
};

// Time:  O(nlogn)
// Space: O(logn)
class Solution2 {
public:
    vector<int> getNoZeroIntegers(int n) {
        vector<int> result;
        for (int a = 1; a < n; ++a) {
            if (to_string(a).find('0') == string::npos &&
                to_string(n - a).find('0') == string::npos) {
                result = {a, n - a};
                break;
            }
        }
        return result;
    }
};
