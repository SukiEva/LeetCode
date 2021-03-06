/*
 * @Author: SukiEva
 * @Date: 2022-03-03 21:38:32
 * @Description: https://github.com/SukiEva
 * わたし、気になります！
 */

#include <bits/stdc++.h>

#include "include/ListNode.h"
#include "include/TreeNode.h"

using namespace std;

typedef unsigned long long ull;

class Solution {
   public:
    int numDistinct(string s, string t) {
        int m = s.size(), n = t.size();
        vector<vector<ull>> dp(m + 1, vector<ull>(n + 1));
        for (int i = 0; i < m; i++) dp[i][0] = 1;
        // for (int j=1; j<n; j++) dp[0][j] = 0;
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (s[i - 1] == t[j - 1])
                    dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j];
                else
                    dp[i][j] = dp[i - 1][j];
            }
        }
        return dp[m][n];
    }
};