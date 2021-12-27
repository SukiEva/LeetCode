/*
 * @Author: SukiEva
 * @Date: 2021-12-27 16:44:08
 * @Description: https://github.com/SukiEva
 * わたし、気になります！
 */
#include <bits/stdc++.h>

#include "include/ListNode.h"
#include "include/TreeNode.h"

using namespace std;

class Solution {
   public:
    vector<vector<int>> ans;
    vector<int> tmp;

    void dfs(vector<int>& nums, vector<bool>& vis) {
        if (tmp.size() == nums.size()) {
            ans.push_back(tmp);
            return;
        }
        for (int i = 0; i < nums.size(); i++) {
            if (vis[i]) continue;
            if (i > 0 && nums[i] == nums[i - 1] && !vis[i - 1])
                continue;  // 排序后去重
            tmp.push_back(nums[i]);
            vis[i] = true;
            dfs(nums, vis);
            vis[i] = false;
            tmp.pop_back();
        }
    }

    vector<vector<int>> permuteUnique(vector<int>& nums) {
        vector<bool> vis(nums.size(), false);
        sort(nums.begin(), nums.end());
        dfs(nums, vis);
        return ans;
    }
};