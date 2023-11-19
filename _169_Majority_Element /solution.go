package _169_Majority_Element_

func majorityElement(nums []int) int {
	cand := nums[0]
	cnt := 0
	// first pass
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			cand = nums[i]
		}
		if cand == nums[i] {
			cnt++
		} else {
			cnt--
		}
	}
	cnt = 0
	// second pass
	for i := 0; i < len(nums); i++ {
		if cand == nums[i] {
			cnt++
		}
	}

	if cnt > len(nums)/2 {
		return cand
	}
	return -1
}

/*
  https://chat.openai.com/share/959093a7-df1b-48bc-b95a-0768c2be3be4 참고
*/
