package leetcode

func removeDuplicates(nums []int) int {
	slow := 0
	for fast, v := range nums {
		if fast < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}
	return slow
}







/*

//正常思路的开一个count计数器.
   int removeDuplicates(vector<int>& nums) {
        int i = 0, cnt = 1;
        for (int j = 1; j < nums.size(); j++) {
            if (nums[j] == nums[i] && cnt>=2) {
			cnt++;
			continue;
			};
			 if (nums[j] == nums[i] && cnt<2) {
			 cnt++;
			i++;
			nums[i] = nums[j];
			};
			if (nums[j] != nums[i] ) {cnt=1;
			i++;
nums[i] = nums[j];
}


        }
        return i + 1;
		}









*/

